package xpath

import (
	"os/user"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// HomeToAbsolute - resolves to absolute path unix systems with tilde prefix
func HomeToAbsolute(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}
	usr, err := user.Current()
	if err != nil {
		return "", errors.Wrap(err, "get current user")
	}
	return filepath.Join(usr.HomeDir, path[2:]), nil
}
