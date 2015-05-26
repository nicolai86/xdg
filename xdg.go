// Package xdg implements the freedesktop xdg specification
//
// http://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html#introduction
package xdg

import (
	"errors"
	"os"
	"strings"
)

const (
	envDataHome   = "XDG_DATA_HOME"
	envDataDirs   = "XDG_DATA_DIRS"
	envConfigHome = "XDG_CONFIG_HOME"
	envConfigDirs = "XDG_CONFIG_DIRS"
	envCacheHome  = "XDG_CACHE_HOME"
	envRuntimeDir = "XDG_RUNTIME_DIR"
	envHome       = "HOME"
)

var (
	// ErrRuntimeDirNotSet is returned by RuntimeDir() if the XDG_RUNTIME_DIR environment
	// variable is not set.
	// Applications should fall back to a replacement directory with similar
	// capabilities and print a warning message.
	// Applications should use this directory for communication and synchronization
	// purposes and should not place larger files in it, since it might reside in
	// runtime memory and cannot necessarily be swapped out to disk.
	ErrRuntimeDirNotSet = errors.New("XDG_RUNTIME_DIR not set")
)

// RuntimeDir defines the base directory relative to which user-specific
// non-essential runtime files and other file objects (such as sockets, named pipes, ...)
// should be stored.
// The directory MUST be owned by the user, and he MUST be the only one having
// read and write access to it.
//
// Its Unix access mode MUST be 0700.
// The lifetime of the directory MUST be bound to the user being logged in.
// It MUST be created when the user first logs in and if the user fully logs out
// the directory MUST be removed.
// If the user logs in more than once he should get pointed to the same directory,
// and it is mandatory that the directory continues to exist from his first login
// to his last logout on the system, and not removed in between.
// Files in the directory MUST not survive reboot or a full logout/login cycle.
//
// The directory MUST be on a local file system and not shared with any other system.
// The directory MUST by fully-featured by the standards of the operating system.
// More specifically, on Unix-like operating systems AF_UNIX sockets,
// symbolic links, hard links, proper permissions, file locking, sparse files, memory mapping,
// file change notifications, a reliable hard link count must be supported, and
// no restrictions on the file name character set should be imposed.
// Files in this directory MAY be subjected to periodic clean-up.
// To ensure that your files are not removed, they should have their access time
// timestamp modified at least once every 6 hours of monotonic time or the
// 'sticky' bit should be set on the file.
func RuntimeDir() (string, error) {
	if os.Getenv(envRuntimeDir) != "" {
		return os.Getenv(envRuntimeDir), nil
	}
	return "", ErrRuntimeDirNotSet
}

// CacheHome defines the base directory relative to which user specific
// non-essential data files should be stored
func CacheHome() string {
	if os.Getenv(envCacheHome) != "" {
		return os.Getenv(envCacheHome)
	}
	return os.Getenv(envHome) + "/.cache"
}

// DataHome defines the base directory relative to which use specific data
// should be stored.
func DataHome() string {
	if os.Getenv(envDataHome) != "" {
		return os.Getenv(envDataHome)
	}
	return os.Getenv(envHome) + "/.local/share"
}

// DataDirs defines the preference-ordered set of base directories to search for
// data files in addition to the DataHome()
//
// The order of base directories denotes their importance;
// the first directory listed is the most important. When the same information
// is defined in multiple places the information defined relative to the more
// important base directory takes precedent.
// The base directory defined by DataHome is considered more important than any
// of the base directories defined by DataDirs().
func DataDirs() []string {
	if os.Getenv(envDataDirs) != "" {
		return strings.Split(os.Getenv(envDataDirs), ":")
	}
	return []string{"/usr/local/share", "/usr/share"}
}

// ConfigHome defines the base directory relative to which user specific
// configuration files should be stored
func ConfigHome() string {
	if os.Getenv(envConfigHome) != "" {
		return os.Getenv(envConfigHome)
	}
	return os.Getenv(envHome) + "/.config"
}

// ConfigDirs defines the preference-ordered set of base directories to search
// for configuration files in addition to the ConfigHome()
//
// The order of base directories denotes their importance;
// the first directory listed is the most important. When the same information
// is defined in multiple places the information defined relative to the more
// important base directory takes precedent.
// The base directory defined by ConfigHome() is considered more important than
// any of the base directories defined by ConfigDirs().
func ConfigDirs() []string {
	if os.Getenv(envConfigDirs) != "" {
		return strings.Split(os.Getenv(envConfigDirs), ":")
	}
	return []string{"/etc/xdg"}
}
