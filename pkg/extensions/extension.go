package extension

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"plugin"
	"strings"
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
	// CurrentVersion is the current version of the extension
	CurrentVersion() string
	// LatestVersion is the latest version of the extension
	LatestVersion() string
	// IsPinned indicates if the extension is pinned
	IsPinned() bool
	// UpdateAvailable indicates if an update is available for the extension
	UpdateAvailable() bool
	// IsBinary indicates if the extension is a binary
	IsBinary() bool
	// IsLocal indicates if the extension is a local extension
	IsLocal() bool
	// Owner is the owner of the extension
	Owner() string
}

// ExtensionManager manages a collection of extensions.
type ExtensionManager interface {
	// ListExtensions lists all installed extensions
	ListExtensions() []Extension
	// Dispatch dispatches a command to the extension
	Dispatch(args []string, stdin io.Reader, stdout, stderr io.Writer) (bool, error)
	// EnableDryRunMode enables dry run mode
	// EnableDryRunMode enables dry run mode
	EnableDryRunMode()
}

var _ Extension = (*UnimplementedExtension)(nil)

type UnimplementedExtension struct{}

// Name implements Extension.Name
func (e *UnimplementedExtension) Name() string {
	return Unknown
}

// Path implements Extension.Path
func (e *UnimplementedExtension) Path() string {
	return Unknown
}

// CurrentVersion implements Extension.CurrentVersion
func (e *UnimplementedExtension) CurrentVersion() string {
	return Unknown
}

// LatestVersion implements Extension.LatestVersion
func (e *UnimplementedExtension) LatestVersion() string {
	return Unknown
}

// IsPinned implements Extension.IsPinned
func (e *UnimplementedExtension) IsPinned() bool {
	return false
}

// UpdateAvailable implements Extension.UpdateAvailable
func (e *UnimplementedExtension) UpdateAvailable() bool {
	return false
}

// IsBinary implements Extension.IsBinary
func (e *UnimplementedExtension) IsBinary() bool {
	return false
}

// IsLocal implements Extension.IsLocal
func (e *UnimplementedExtension) IsLocal() bool {
	return false
}

// Owner implements Extension.Owner
func (e *UnimplementedExtension) Owner() string {
	return Unknown
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
		return nil, err
	}

	return e, nil
}
