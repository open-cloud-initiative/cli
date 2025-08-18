package extensions

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/katallaxie/pkg/slices"
	"github.com/spf13/cobra"
)

// Unknown is returned when an extension is not found.
var Unknown = "unknown"

// ErrUnimplemented is returned when an extension is not implemented.
var ErrUnimplemented = errors.New("unimplemented")

// Extension represents a single extension.
type Extension interface {
	// Name is the name of the extension
	Name() string
	// Path is the path to the extension
	Path() string
	// Version is the version of the extension
	Version() string
	// Owner is the owner of the extension
	Owner() string
	// Cmd is the command for the extension
	Cmd() *cobra.Command
}

// Manager manages a collection of extensions.
type Manager interface {
	// Scan is scanning for extensions
	Scan(path string) error
	// ListExtensions lists all installed extensions
	ListExtensions() []Extension
	// EnableDryRunMode enables dry run mode
	EnableDryRunMode()
}

var _ Manager = (*manager)(nil)

type manager struct {
	dryRun     bool
	extensions []Extension
}

// ListExtensions implements Manager.ListExtensions.
func (m *manager) ListExtensions() []Extension {
	return m.extensions
}

// EnableDryRunMode implements Manager.EnableDryRunMode.
func (m *manager) EnableDryRunMode() {
	m.dryRun = true
}

// Scan implements Manager.Scan.
func (m *manager) Scan(path string) error {
	exts, err := Scan(path)
	if err != nil {
		return err
	}

	m.extensions = slices.Append(m.extensions, exts...)

	return nil
}

// NewManager creates a new extension manager.
func NewManager() Manager {
	return &manager{}
}

var _ Extension = (*UnimplementedExtension)(nil)

type UnimplementedExtension struct{}

// Name implements Extension.Name.
func (e *UnimplementedExtension) Name() string {
	return Unknown
}

// Path implements Extension.Path.
func (e *UnimplementedExtension) Path() string {
	return Unknown
}

// Version implements Extension.Version.
func (e *UnimplementedExtension) Version() string {
	return Unknown
}

// Owner implements Extension.Owner.
func (e *UnimplementedExtension) Owner() string {
	return Unknown
}

// Cmd implements Extension.Cmd.
func (e *UnimplementedExtension) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "unknown",
		Short: "Unknown extension",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("Unknown extension")
		},
	}
}

// Scan is scanning a folder for extensions.
// Extensions start with "oci-" for the ocictl.
func Scan(folder string) ([]Extension, error) {
	var extensions []Extension

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), "oci-") {
			ext, err := Load(path)
			if err != nil {
				return err
			}
			extensions = append(extensions, ext)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return extensions, nil
}

// Load loads an extension with the given path and returns an instance.
func Load(path string) (Extension, error) {
	ext, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symPlugin, err := ext.Lookup("Extension")
	if err != nil {
		return nil, err
	}

	e, ok := symPlugin.(Extension)
	if !ok {
		return nil, fmt.Errorf("plugin %s does not implement Extension interface", path)
	}

	return e, nil
}

// DataDir returns the data directory for the extension.
func DataDir() string {
	return filepath.Join(os.Getenv("HOME"), ".ocictl")
}
