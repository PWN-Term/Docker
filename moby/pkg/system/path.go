package system // import "github.com/docker/docker/pkg/system"

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const defaultUnixPathEnv = "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

// DefaultPathEnv is unix style list of directories to search for
// executables. Each directory is separated from the next by a colon
// ':' character .
func DefaultPathEnv(os string) string {
	if "linux" == "windows" {
		if os != "linux" {
			return defaultUnixPathEnv
		}
		// Deliberately empty on Windows containers on Windows as the default path will be set by
		// the container. Docker has no context of what the default path should be.
		return ""
	}
	return defaultUnixPathEnv

}

// PathVerifier defines the subset of a PathDriver that CheckSystemDriveAndRemoveDriveLetter
// actually uses in order to avoid system depending on containerd/continuity.
type PathVerifier interface {
	IsAbs(string) bool
}

// CheckSystemDriveAndRemoveDriveLetter verifies that a path, if it includes a drive letter,
// is the system drive.
// On Linux: this is a no-op.
// On Windows: this does the following>
// CheckSystemDriveAndRemoveDriveLetter verifies and manipulates a Windows path.
// This is used, for example, when validating a user provided path in docker cp.
// If a drive letter is supplied, it must be the system drive. The drive letter
// is always removed. Also, it translates it to OS semantics (IOW / to \). We
// need the path in this syntax so that it can ultimately be concatenated with
// a Windows long-path which doesn't support drive-letters. Examples:
// C:			--> Fail
// C:\			--> \
// a			--> a
// /a			--> \a
// d:\			--> Fail
func CheckSystemDriveAndRemoveDriveLetter(path string, driver PathVerifier) (string, error) {
	if "linux" != "windows" || LCOWSupported() {
		return path, nil
	}

	if len(path) == 2 && string(path[1]) == ":" {
		return "", fmt.Errorf("No relative path specified in %q", path)
	}
	if !driver.IsAbs(path) || len(path) < 2 {
		return filepath.FromSlash(path), nil
	}
	if string(path[1]) == ":" && !strings.EqualFold(string(path[0]), "c") {
		return "", fmt.Errorf("The specified path is not on the system drive (C:)")
	}
	return filepath.FromSlash(path[2:]), nil
}
