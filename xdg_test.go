package xdg

import (
	"os"
	"strings"
	"testing"
)

type expectation struct {
	env  map[string]string
	path string
}

func ConfigDirsTest(t *testing.T) {
	t.Parallel()

	os.Setenv(envConfigDirs, "/foo:/bar")
	if strings.Join(ConfigDirs(), ":") != "/foo:/bar" {
		t.Errorf("expected %q got %q", []string{"/foo", "/bar"}, ConfigDirs())
	}

	os.Setenv(envConfigDirs, "")
	if strings.Join(ConfigDirs(), ":") != "/etc/xdg" {
		t.Errorf("expected %q got %q", []string{"/etc/xdg"}, ConfigDirs())
	}
}

func DataDirsTest(t *testing.T) {
	t.Parallel()

	os.Setenv(envDataDirs, "/foo:/bar")
	if strings.Join(DataDirs(), ":") != "/foo:/bar" {
		t.Errorf("expected %q got %q", []string{"/foo", "/bar"}, DataDirs())
	}

	os.Setenv(envDataDirs, "")
	if strings.Join(DataDirs(), ":") != "/usr/local/share/:/usr/share/" {
		t.Errorf("expected %q got %q", []string{"/usr/local/share", "/usr/share"}, DataDirs())
	}
}

func CacheHomeTests(t *testing.T) {
	t.Parallel()

	var tests = []expectation{
		expectation{
			path: "/home/manfred/.cache",
			env: map[string]string{
				envCacheHome: "",
				"HOME":       "/home/manfred",
			},
		},
		expectation{
			path: "/tmp",
			env: map[string]string{
				envCacheHome: "/tmp",
				"HOME":       "/home/manfred",
			},
		},
	}
	for _, test := range tests {
		for key, val := range test.env {
			os.Setenv(key, val)
		}
		if CacheHome() != test.path {
			t.Errorf("expected %q got %q", test.path, CacheHome())
		}
	}
}

func ConfigHomeTests(t *testing.T) {
	t.Parallel()

	var tests = []expectation{
		expectation{
			path: "/home/manfred/.config",
			env: map[string]string{
				envConfigHome: "",
				"HOME":        "/home/manfred",
			},
		},
		expectation{
			path: "/home/manfred/cunfig",
			env: map[string]string{
				envConfigHome: "/home/manfred/cunfig",
				"HOME":        "/home/manfred",
			},
		},
	}
	for _, test := range tests {
		for key, val := range test.env {
			os.Setenv(key, val)
		}
		if ConfigHome() != test.path {
			t.Errorf("expected %q got %q", test.path, ConfigHome())
		}
	}
}

func DataHomeTests(t *testing.T) {
	t.Parallel()

	var tests = []expectation{
		expectation{
			path: "/home/manfred/.local/share",
			env: map[string]string{
				envDataHome: "",
				"HOME":      "/home/manfred",
			},
		},
		expectation{
			path: "/home/manfred/.custom",
			env: map[string]string{
				envDataHome: "/home/manfred/custom",
				"HOME":      "/home/manfred",
			},
		},
	}

	for _, test := range tests {
		for key, val := range test.env {
			os.Setenv(key, val)
		}
		if DataHome() != test.path {
			t.Errorf("expected %q got %q", test.path, DataHome())
		}
	}
}
