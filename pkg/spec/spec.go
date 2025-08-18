package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/katallaxie/pkg/filex"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var validate = validator.New()

const (
	// DefaultPath is the default path for the configuration file.
	DefaultPath = ".ocictl"
	// DefaultFilename is the default filename for the configuration file.
	DefaultFilename = ".ocictl.yml"
)

// Spec is the configuration file for `csync`.
type Spec struct {
	// Version is the version of the configuration file.
	Version int `yaml:"version" validate:"required,eq=1"`
	// Folder is the folder for the configuration file.
	Folder string `yaml:".oci" validate:"required"`

	sync.Mutex `yaml:"-"`
}

// UnmarshalYAML overrides the default unmarshaler for the spec.
func (s *Spec) UnmarshalYAML(data []byte) error {
	spec := struct {
		Version int    `yaml:"version" validate:"required,eq=1"`
		Folder  string `yaml:".oci" validate:"required"`
		Stderr  bool   `yaml:"stderr,omitempty"`
		Stdout  bool   `yaml:"stdout,omitempty"`
	}{
		Folder: DefaultPath,
		Stderr: true,
		Stdout: true,
	}

	if err := yaml.Unmarshal(data, &spec); err != nil {
		return errors.WithStack(err)
	}

	s.Version = spec.Version

	return nil
}

// GetVersion returns the version of the configuration file.
func (s *Spec) GetVersion() int {
	return s.Version
}

// Default is the default configuration.
func Default() *Spec {
	return &Spec{
		Version: 1,
	}
}

// Validate is the validation function for the spec.
func (s *Spec) Validate() error {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return validate.Struct(s)
}

// Write is the write function for the spec.
func Write(s *Spec, file string, force bool) error {
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	ok, _ := filex.FileExists(filepath.Clean(file))
	if ok && !force {
		return fmt.Errorf("%s already exists, use --force to overwrite", file)
	}

	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
