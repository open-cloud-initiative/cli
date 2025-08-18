package config

import (
	"os"
	"sync"

	"github.com/katallaxie/pkg/filex"
)

// Flags contains the command line flags.
type Flags struct {
	// Dry toggles the dry run mode.
	Dry bool
	// Force toggles the force mode.
	Force bool
	// Root runs the command as root.
	Root bool
	// Verbose toggles the verbosity.
	Verbose bool
	// Version toggles the version.
	Version bool
}

// NewFlags returns a new flags.
func NewFlags() *Flags {
	return &Flags{}
}

// Config contains the configuration.
type Config struct {
	// Verbose toggles the verbosity
	Verbose bool
	// File...
	File string
	// Flags ...
	Flags *Flags
	// Stdin ...
	Stdin *os.File
	// Stdout ...
	Stdout *os.File
	// Stderr ...
	Stderr *os.File

	sync.RWMutex
}

// New returns a new config.
func New() *Config {
	return &Config{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Flags:  &Flags{},
	}
}

// InitDefaultConfig initializes the default configuration.
func (c *Config) InitDefaultConfig() error {
	folder, err := filex.ExpandHomeFolder(c.File)
	if err != nil {
		return err
	}

	c.File = folder

	return nil
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
