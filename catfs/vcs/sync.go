package vcs

import (
	"fmt"
	"path"

	c "github.com/disorganizer/brig/catfs/core"
	ie "github.com/disorganizer/brig/catfs/errors"
	n "github.com/disorganizer/brig/catfs/nodes"
)

const (
	ConflictStragetyMarker = iota
	ConflictStragetyIgnore
	ConflictStragetyUnknown
)

type ConflictStrategy int

func (cs ConflictStrategy) String() string {
	switch cs {
	case ConflictStragetyMarker:
		return "marker"
	case ConflictStragetyIgnore:
		return "ignore"
	default:
		return "unknown"
	}
}

func ConflictStrategyFromString(spec string) ConflictStrategy {
	switch spec {
	case "marker":
		return ConflictStragetyMarker
	case "ignore":
		return ConflictStragetyIgnore
	default:
		return ConflictStragetyUnknown
	}
}

// SyncConfig gives you the possibility to configure the sync algorithm.
// The zero value of each option is the
type SyncConfig struct {
	ConflictStrategy ConflictStrategy
	IgnoreDeletes    bool
}

var (
	DefaultSyncConfig = &SyncConfig{}
)

type syncer struct {
	cfg    *SyncConfig
	lkrSrc *c.Linker
	lkrDst *c.Linker
}

func (sy *syncer) add(src n.ModNode, srcParent, srcName string) error {
	fmt.Println("ADD", src, srcParent, srcName)
	var newDstNode n.ModNode
	var err error

	parentDir, err := sy.lkrDst.LookupDirectory(srcParent)
	if err != nil {
		return err
	}

	switch src.Type() {
	case n.NodeTypeDirectory:
		newDstNode, err = n.NewEmptyDirectory(
			sy.lkrDst,
			parentDir,
			srcName,
			sy.lkrDst.NextInode(),
		)

		if err != nil {
			return err
		}
	case n.NodeTypeFile:
		newDstFile, err := n.NewEmptyFile(
			parentDir,
			srcName,
			sy.lkrDst.NextInode(),
		)

		if err != nil {
			return err
		}

		newDstNode = newDstFile

		srcFile, ok := src.(*n.File)
		if ok {
			newDstFile.SetContent(sy.lkrDst, srcFile.Content())
			newDstFile.SetSize(srcFile.Size())
			newDstFile.SetKey(srcFile.Key())
		}

		// TODO: This is inconsistent:
		// NewEmptyDirectory calls Add(), NewEmptyFile does not
		if err := parentDir.Add(sy.lkrDst, newDstFile); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unexpected node type in handleAdd")
	}
	fmt.Println("STAGE")

	return sy.lkrDst.StageNode(newDstNode)
}

func (sy *syncer) handleAdd(src n.ModNode) error {
	return sy.add(src, path.Dir(src.Path()), src.Name())
}

func (sy *syncer) handleRemove(dst n.ModNode) error {
	if sy.cfg.IgnoreDeletes {
		return nil
	}

	_, _, err := c.Remove(sy.lkrDst, dst, true, true)
	return err
}

func (sy *syncer) handleConflict(src, dst n.ModNode, srcMask, dstMask ChangeType) error {
	fmt.Println("CONFLICT", src, dst)
	if sy.cfg.ConflictStrategy == ConflictStragetyIgnore {
		return nil
	}

	// Find a path that we do not have yet.
	// stamp := time.Now().Format(time.RFC3339)
	conflictNameTmpl := fmt.Sprintf("%s.conflict.%%d", dst.Name())
	conflictName := ""

	// Fix the unlikely case that there is already a node at the conflict path:
	for tries := 0; tries < 100; tries++ {
		conflictName = fmt.Sprintf(conflictNameTmpl, tries)
		dstNd, err := sy.lkrDst.LookupNode(conflictName)
		if err != nil && !ie.IsNoSuchFileError(err) {
			return err
		}

		if dstNd == nil {
			break
		}
	}

	dstDirname := path.Dir(dst.Path())
	fmt.Println("Writing conflict file to ", dstDirname, conflictName)
	return sy.add(src, dstDirname, conflictName)
}

func (sy *syncer) handleMerge(src, dst n.ModNode, srcMask, dstMask ChangeType) error {
	if src.Path() != dst.Path() {
		// Only move the file if it was only moved on the remote side.
		if srcMask&ChangeTypeMove != 0 && dstMask&ChangeTypeMove == 0 {
			// TODO: Sanity check that there's nothing at src.Path(),
			//       but Mapper should already have checked that.
			if err := c.Move(sy.lkrDst, dst, src.Path()); err != nil {
				return err
			}
		}
	}

	// If src did not change, there's no need to sync the content.
	// If src has no changes, we know that dst must have changes,
	// otherwise it would have been reported as conflict.
	if srcMask&ChangeTypeModify == 0 && srcMask&ChangeTypeAdd == 0 {
		return nil
	}

	dstParent, err := n.ParentDirectory(sy.lkrDst, dst)
	if err != nil {
		return err
	}

	if err := dstParent.RemoveChild(sy.lkrSrc, dst); err != nil {
		return err
	}

	dstFile, ok := dst.(*n.File)
	if !ok {
		return ie.ErrBadNode
	}

	srcFile, ok := src.(*n.File)
	if !ok {
		return ie.ErrBadNode
	}

	dstFile.SetContent(sy.lkrDst, srcFile.Content())
	dstFile.SetSize(srcFile.Size())
	dstFile.SetKey(srcFile.Key())

	return sy.lkrDst.StageNode(dstFile)
}

func (sy *syncer) handleTypeConflict(src, dst n.ModNode) error {
	// Simply do nothing.
	return nil
}

func Sync(lkrSrc, lkrDst *c.Linker, cfg *SyncConfig) error {
	if cfg == nil {
		cfg = DefaultSyncConfig
	}

	syncer := &syncer{
		cfg:    cfg,
		lkrSrc: lkrSrc,
		lkrDst: lkrDst,
	}

	resolver, err := newResolver(lkrSrc, lkrDst, nil, nil, syncer)
	if err != nil {
		return err
	}

	if err := resolver.resolve(); err != nil {
		return err
	}

	wasModified, err := lkrDst.HaveStagedChanges()
	if err != nil {
		return err
	}

	// If something was changed, we should set the merge marker
	// and also create a new commit.
	if wasModified {
		srcOwner, err := lkrSrc.Owner()
		if err != nil {
			return err
		}

		srcHead, err := lkrSrc.Head()
		if err != nil {
			return err
		}

		// If something was changed, remember that we merged with src.
		// This avoids merging conflicting files a second time in the next resolve().
		if err := lkrDst.SetMergeMarker(srcOwner, srcHead.Hash()); err != nil {
			return err
		}

		message := fmt.Sprintf("Merge with %s", srcOwner)
		return lkrDst.MakeCommit(srcOwner, message)
	}

	return nil
}
