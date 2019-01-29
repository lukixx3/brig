package main

// Simple test gateway. This is not very useful on its own,
// but very useful for frontend development. Use it like this:
// $ go generate ./... && go run gateway/standalone/*.go some/test/data

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/sahib/brig/catfs"
	"github.com/sahib/brig/defaults"
	"github.com/sahib/brig/gateway"
	"github.com/sahib/brig/gateway/remotesapi"
	"github.com/sahib/config"
)

const (
	dbPath     = "/tmp/gw-runner"
	configPath = "/tmp/gw-runner/config.cfg"
)

func loadConfig(configPath string) *config.Config {
	cfg, err := defaults.OpenMigratedConfig(configPath)
	if err == nil {
		return cfg
	}

	log.Printf("failed to open config: %v", err)
	if _, err := os.Stat(configPath); err != nil && !os.IsNotExist(err) {
		os.Exit(1)
	}

	log.Printf("creating empty config at %s", configPath)
	cfg, err = config.Open(nil, defaults.Defaults, config.StrictnessPanic)
	if err != nil {
		log.Printf("failed to load defaults: %v", err)
		return cfg
	}

	saveConfig(cfg, configPath)
	return cfg
}

func saveConfig(cfg *config.Config, configPath string) {
	fd, err := os.OpenFile(configPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("failed to write to config location %s: %v", configPath, err)
	}

	defer fd.Close()

	if err := cfg.Save(config.NewYamlEncoder(fd)); err != nil {
		log.Fatalf("failed to serialize config: %v", err)
	}
}

func main() {
	log.SetLevel(log.DebugLevel)
	if err := os.MkdirAll(dbPath, 0700); err != nil {
		log.Fatalf("failed to create dir %s: %v", dbPath, err)
	}

	cfg := loadConfig(configPath)
	cfg.SetBool("gateway.enabled", true)
	cfg.SetBool("gateway.ui.enabled", true)
	cfg.SetBool("gateway.ui.debug_mode", true)
	cfg.SetInt("gateway.port", 5001)
	cfg.SetBool("gateway.cert.redirect.enabled", false)

	cfg.SetBool("gateway.auth.enabled", true)
	cfg.SetString("gateway.cert.domain", "nwzmlh4iouqikobq.myfritz.net")
	cfg.SetString("gateway.cert.certfile", "/tmp/fullchain.pem")
	cfg.SetString("gateway.cert.keyfile", "/tmp/privkey.pem")

	bk, err := NewTmpFsBackend(filepath.Join(dbPath, "backend"))
	if err != nil {
		log.Fatalf("failed to open backend: %v", err)
	}

	fsPath := filepath.Join(dbPath, "metadata")
	fs, err := catfs.NewFilesystem(bk, fsPath, "ali", false, cfg.Section("fs"))
	if err != nil {
		log.Fatalf("failed to open fs: %v", err)
	}

	for _, root := range os.Args[1:] {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}

			/* #nosec */
			fd, err := os.Open(path)
			if err != nil {
				log.Fatalf("failed to open: %v", err)
			}

			if err := fs.Stage(path[len(root):], fd); err != nil {
				log.Fatalf("failed to stage: %s: %v", path, err)
			}

			return fd.Close()
		})

		if err != nil {
			log.Fatalf("walk failed: %v", err)
		}
	}

	rmtMock := remotesapi.NewMock("ali", "<fingerprint>")

	userDbPath := filepath.Join(dbPath, "users")
	gw, err := gateway.NewGateway(fs, rmtMock, cfg.Section("gateway"), userDbPath)
	if err != nil {
		log.Fatalf("failed to open gateway: %v", err)
	}

	if err := gw.UserDatabase().Add("admin", "password", nil); err != nil {
		log.Fatalf("failed to add user: %v", err)
	}

	gw.Start()

	time.Sleep(1 * time.Second)
	saveConfig(cfg, configPath)

	defer func() {
		if err := gw.Stop(); err != nil {
			log.Warningf("failed to stop properly: %v", err)
		}
	}()

	// Block until hitting Ctrl-C
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	fmt.Println("Hit Ctrl-C to interrupt.")
	<-ch
	fmt.Println("Interrupted. Shutting down.")
}