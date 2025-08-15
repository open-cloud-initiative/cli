package extension

import "io"

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

type Manager interface {
	// ListExtensions lists all installed extensions
	ListExtensions() []Extension
	// Dispatch dispatches a command to the extension
	Dispatch(args []string, stdin io.Reader, stdout, stderr io.Writer) (bool, error)
	// EnableDryRunMode enables dry run mode
	EnableDryRunMode()
}
