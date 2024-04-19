package findcfg

import (
	"os"
	"path/filepath"
	"strings"
)

type FinderOption func(*Finder)

// .toml
func TOML() FinderOption {
	return func(f *Finder) {
		f.AddExts(".toml")
	}
}

// .ini
func INI() FinderOption {
	return func(f *Finder) {
		f.AddExts(".ini")
	}
}

// .json
func JSON() FinderOption {
	return func(f *Finder) {
		f.AddExts(".json")
	}
}

// .yaml, .yml
func YAML() FinderOption {
	return func(f *Finder) {
		f.AddExts(".yaml", ".yml")
	}
}

// arbitrary exts
func Ext(exts ...string) FinderOption {
	return func(f *Finder) {
		for _, e := range exts {
			if !strings.HasPrefix(e, ".") {
				e = "." + e
			}
			f.AddExts(e)
		}
	}
}

// arbitrary dirs
func Dir(sub ...string) FinderOption {
	return func(f *Finder) {
		f.AddDirs(func() (string, string) {
			const desc = "const"
			return filepath.Join(sub...), desc
		})
	}
}

// cwd
func CurrentDir(sub ...string) FinderOption {
	return func(f *Finder) {
		f.AddDirs(func() (string, string) {
			const desc = "cwd"
			wd, err := os.Getwd()
			if err != nil {
				return "", desc
			}
			p := append([]string{wd}, sub...)
			return filepath.Join(p...), desc
		})
	}
}

// os.UserHomeDir()
func HomeDir(sub ...string) FinderOption {
	return func(f *Finder) {
		f.AddDirs(func() (string, string) {
			const desc = "home"
			ud, err := os.UserHomeDir()
			if err != nil {
				return "", desc
			}
			p := append([]string{ud}, sub...)
			return filepath.Join(p...), desc
		})
	}
}

// os.UserConfigDir()
func UserConfigDir(sub ...string) FinderOption {
	return func(f *Finder) {
		f.AddDirs(func() (string, string) {
			const desc = "userconfig"
			ud, err := os.UserConfigDir()
			if err != nil {
				return "", desc
			}
			p := append([]string{ud}, sub...)
			return filepath.Join(p...), desc
		})
	}
}

// os.Executable()
func ExecutableDir(sub ...string) FinderOption {
	return func(f *Finder) {
		f.AddDirs(func() (string, string) {
			const desc = "exe"
			ed, err := os.Executable()
			if err != nil {
				return "", desc
			}
			p := append([]string{filepath.Dir(ed)}, sub...)
			return filepath.Join(p...), desc
		})
	}
}

func ExactPath(exacts ...string) FinderOption {
	return func(f *Finder) {
		f.AddExacts(exacts...)
	}
}
