// Package findcfg finds a config file.
//
//	// finder for a YAML file in os.Executable() and os.UserConfigDir()+myapp
//	finder := findcfg.New(
//	    findcfg.YAML(),
//	    findcfg.ExecutableDir(),
//	    findcfg.UserConfigDir("myapp"),
//	)
//	found := finder.Find("config") // config.yaml
//	if found != nil {
//		return found.Path
//	}
//	return finder.FallbackPath(configName)
package findcfg

import (
	"os"
	"path/filepath"
)

type DirFunc func() (path, desc string)

type Finder struct {
	Exts   []string
	Names  []string
	Dirs   []DirFunc
	Exacts []string // absolute file path
}

type Found struct {
	Path    string
	Ext     string
	DirDesc string
}

// New returns a Finder with given options.
//
// see [FinderOption].
func New(opts ...FinderOption) *Finder {
	f := &Finder{}
	for _, o := range opts {
		o(f)
	}
	return f
}

// AddExts adds extentions. (needs a period like ".yaml")
// If no exts given by AddExts and New, the Finder does not find anything.
func (f *Finder) AddExts(exts ...string) {
	f.Exts = append(f.Exts, exts...)
}

// AddNames adds base names without extentions.
func (f *Finder) AddNames(names ...string) {
	f.Names = append(f.Names, names...)
}

// AddDirs adds search paths.
// If no dirs given by AddDirs and New, the Finder does not find anything.
func (f *Finder) AddDirs(dirs ...DirFunc) {
	f.Dirs = append(f.Dirs, dirs...)
}

// AddExact adds an absolute file path.
func (f *Finder) AddExact(path string) {
	if path == "" {
		return
	}
	f.Exacts = append(f.Exacts, path)
}

// Find try to find a config file accordding to its dirs and exts.
//
// You should call AddXXX before call this method.
//
// baseName is without ".ext", like "myconfig"
func (f *Finder) Find() *Found {
	for _, p := range f.Exacts {
		if s, err := os.Stat(p); err != nil || s.IsDir() {
			continue
		}
		return &Found{
			Path:    p,
			Ext:     filepath.Ext(p),
			DirDesc: "exact",
		}
	}

	for _, getdir := range f.Dirs {
		dir, desc := getdir()
		if dir == "" {
			continue
		}

		for _, name := range f.Names {
			for _, ext := range f.Exts {
				p := filepath.Join(dir, name+ext)
				if s, err := os.Stat(p); err != nil || s.IsDir() {
					continue
				}
				return &Found{
					Path:    p,
					Ext:     ext,
					DirDesc: desc,
				}
			}
		}
	}

	return nil
}

// FallbackPath returns a path that may not be exist.
// Call this method if Find does not find anything.
//
// This method returns the firstly tried path.
func (f *Finder) FallbackPath() string {
	if len(f.Exacts) > 0 {
		return f.Exacts[0]
	}

	if len(f.Dirs) == 0 {
		return ""
	}
	if len(f.Names) == 0 {
		return ""
	}
	if len(f.Exts) == 0 {
		return ""
	}
	d, _ := f.Dirs[0]()
	return filepath.Join(d, f.Names[0]+f.Exts[0])
}
