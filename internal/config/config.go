/*
©Artefhack | 2026
config.go

Loads vulnforge settings from .vulnforge.toml or the [tool.vulnforge] section in pyproject.toml

Implements a cascading resolution order: checks for a standalone
.vulnforge.toml first, then falls back to [tool.vulnforge] inside
pyproject.toml. Returns a zero-value Config if neither source is present.

Connects to:
  update.go - calls Load() before each scan or update run
*/

package config

import (
	"os"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

// Config holds project-level vulnforge settings loaded from .vulnforge.toml
// or [tool.vulnforge] in pyproject.toml
type Config struct {
	MinSeverity string   `toml:"min-severity"`
	Ignore      []string `toml:"ignore"`
	IgnoreVulns []string `toml:"ignore-vulns"`
}

type pyprojectWrapper struct {
	Tool struct {
		VulnForge Config `toml:"vulnforge"`
	} `toml:"tool"`
}

// Load reads vulnforge configuration using a cascading resolution order:
// .vulnforge.toml in current directory, then [tool.vulnforge] in pyproject.toml
func Load(pyprojectPath string) (Config, error) {
	if cfg, err := loadFile(".vulnforge.toml"); err == nil {
		return cfg, nil
	} else if !os.IsNotExist(err) {
		return Config{}, err
	}

	if cfg, ok := loadFromPyproject(pyprojectPath); ok {
		return cfg, nil
	}

	return Config{}, nil
}

func loadFile(path string) (Config, error) {
	data, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	cfg.MinSeverity = strings.ToLower(
		strings.TrimSpace(cfg.MinSeverity),
	)
	return cfg, nil
}

func loadFromPyproject(path string) (Config, bool) {
	data, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		return Config{}, false
	}

	var wrapper pyprojectWrapper
	if err := toml.Unmarshal(data, &wrapper); err != nil {
		return Config{}, false
	}

	cfg := wrapper.Tool.VulnForge
	if cfg.MinSeverity == "" &&
		len(cfg.Ignore) == 0 &&
		len(cfg.IgnoreVulns) == 0 {
		return Config{}, false
	}

	cfg.MinSeverity = strings.ToLower(
		strings.TrimSpace(cfg.MinSeverity),
	)
	return cfg, true
}
