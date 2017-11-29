package client

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/disorganizer/brig/brigd/capnp"
	h "github.com/disorganizer/brig/util/hashlib"
)

type StatInfo struct {
	Path    string
	Hash    h.Hash
	Size    uint64
	Inode   uint64
	IsDir   bool
	Depth   int
	ModTime time.Time
}

func convertCapStatInfo(capInfo *capnp.StatInfo) (*StatInfo, error) {
	info := &StatInfo{}

	path, err := capInfo.Path()
	if err != nil {
		return nil, err
	}

	hashBytes, err := capInfo.Hash()
	if err != nil {
		return nil, err
	}

	hash, err := h.Cast(hashBytes)
	if err != nil {
		return nil, err
	}

	modTimeData, err := capInfo.ModTime()
	if err != nil {
		return nil, err
	}

	if err := info.ModTime.UnmarshalText([]byte(modTimeData)); err != nil {
		return nil, err
	}

	info.Path = path
	info.Hash = hash
	info.Size = capInfo.Size()
	info.Inode = capInfo.Inode()
	info.IsDir = capInfo.IsDir()
	info.Depth = int(capInfo.Depth())
	return info, nil
}

func (cl *Client) List(root string, maxDepth int) ([]*StatInfo, error) {
	call := cl.api.List(cl.ctx, func(p capnp.FS_list_Params) error {
		p.SetMaxDepth(int32(maxDepth))
		return p.SetRoot(root)
	})

	result, err := call.Struct()
	if err != nil {
		return nil, err
	}

	results := []*StatInfo{}
	statList, err := result.Entries()
	for idx := 0; idx < statList.Len(); idx++ {
		capInfo := statList.At(idx)
		info, err := convertCapStatInfo(&capInfo)
		if err != nil {
			return nil, err
		}

		results = append(results, info)
	}

	return results, err
}

func (cl *Client) Stage(localPath, repoPath string) error {
	call := cl.api.Stage(cl.ctx, func(p capnp.FS_stage_Params) error {
		if err := p.SetRepoPath(repoPath); err != nil {
			return err
		}

		return p.SetLocalPath(localPath)
	})

	_, err := call.Struct()
	return err
}

func (cl *Client) Cat(path string) (io.ReadCloser, error) {
	call := cl.api.Cat(cl.ctx, func(p capnp.FS_cat_Params) error {
		return p.SetPath(path)
	})

	result, err := call.Struct()
	if err != nil {
		return nil, err
	}

	port := result.Port()
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (cl *Client) Mkdir(path string, createParents bool) error {
	call := cl.api.Mkdir(cl.ctx, func(p capnp.FS_mkdir_Params) error {
		p.SetCreateParents(createParents)
		return p.SetPath(path)
	})

	_, err := call.Struct()
	return err
}

func (cl *Client) Remove(path string) error {
	call := cl.api.Remove(cl.ctx, func(p capnp.FS_remove_Params) error {
		return p.SetPath(path)
	})

	_, err := call.Struct()
	return err
}

func (cl *Client) Move(srcPath, dstPath string) error {
	call := cl.api.Move(cl.ctx, func(p capnp.FS_move_Params) error {
		if err := p.SetSrcPath(srcPath); err != nil {
			return err
		}

		return p.SetDstPath(dstPath)
	})

	_, err := call.Struct()
	return err
}

func (cl *Client) Pin(path string) error {
	call := cl.api.Pin(cl.ctx, func(p capnp.FS_pin_Params) error {
		return p.SetPath(path)
	})

	_, err := call.Struct()
	return err
}

func (cl *Client) Unpin(path string) error {
	call := cl.api.Unpin(cl.ctx, func(p capnp.FS_unpin_Params) error {
		return p.SetPath(path)
	})

	_, err := call.Struct()
	return err
}

func (cl *Client) IsPinned(path string) (bool, error) {
	call := cl.api.IsPinned(cl.ctx, func(p capnp.FS_isPinned_Params) error {
		return p.SetPath(path)
	})

	result, err := call.Struct()
	if err != nil {
		return false, err
	}

	return result.IsPinned(), nil
}
