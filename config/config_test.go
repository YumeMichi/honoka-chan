package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadCreatesDefaultConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")

	conf, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if conf.Settings.ListenPort != "8080" {
		t.Fatalf("ListenPort = %q, want 8080", conf.Settings.ListenPort)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("default config was not written: %v", err)
	}
}

func TestLoadBacksUpInvalidConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, []byte("{"), 0644); err != nil {
		t.Fatalf("write invalid config: %v", err)
	}

	conf, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if conf.Settings.ListenPort != "8080" {
		t.Fatalf("ListenPort = %q, want 8080", conf.Settings.ListenPort)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "config.json.backup") {
			return
		}
	}
	t.Fatal("invalid config backup was not created")
}

func TestValidateRequiresSessionKey(t *testing.T) {
	conf := DefaultConfigs()
	if err := conf.Validate(); err == nil {
		t.Fatal("Validate accepted a missing session key")
	}

	conf.Settings.SessionKey = "01234567890123456789012345678901"
	if err := conf.Validate(); err != nil {
		t.Fatalf("Validate rejected a sufficiently long session key: %v", err)
	}
}

func TestValidateRejectsWhitespaceAroundSessionKey(t *testing.T) {
	conf := DefaultConfigs()
	conf.Settings.SessionKey = " 01234567890123456789012345678901"
	if err := conf.Validate(); err == nil {
		t.Fatal("Validate accepted a session key with leading whitespace")
	}
}
