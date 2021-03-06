package static

import (
	"github.com/pkg/errors"
	"github.com/spiral/roadrunner/service"
	"os"
	"path"
	"strings"
)

// Config describes file location and controls access to them.
type Config struct {
	// Enables StaticFile service.
	Enable bool

	// Dir contains name of directory to control access to.
	Dir string

	// Forbid specifies list of file extensions which are forbidden for access.
	// Example: .php, .exe, .bat, .htaccess and etc.
	Forbid []string
}

// Hydrate must populate Config values using given Config source. Must return error if Config is not valid.
func (c *Config) Hydrate(cfg service.Config) error {
	if err := cfg.Unmarshal(c); err != nil {
		return err
	}

	return c.Valid()
}

// Valid returns nil if config is valid.
func (c *Config) Valid() error {
	st, err := os.Stat(c.Dir)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("root directory does not exists")
		}

		return err
	}

	if !st.IsDir() {
		return errors.New("invalid root directory")
	}

	return nil
}

// Forbids must return true if file extension is not allowed for the upload.
func (c *Config) Forbids(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))

	for _, v := range c.Forbid {
		if ext == v {
			return true
		}
	}

	return false
}
