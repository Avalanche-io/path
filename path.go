package path

import (
	"errors"
	slash "path"
	"strings"
)

type Path string

// New creates a path from a string, with the following rules:
//
// Removes all URL prefixes.
// Reduces two or more "/" to one.
// Appends a "/" if there are no "/" in the string.
// Removes all characters after the last "/".
//
// If input is empty or nil, return nil path and error.
//
func New(path string) (*Path, error) {
	if len(path) == 0 {
		return nil, errors.New("empty path")
	}
	file := true
	if path[len(path)-1:] == "/" {
		file = false
	}
	path = removeUrlPrefix(path)
	path = slash.Clean(path)
	if strings.Index(path, "/") == -1 {
		file = false
	}
	if file {
		path = slash.Dir(path)
	}
	if len(path) > 1 {
		path += "/"
	}
	p := Path(path)
	return &p, nil
}

// Appends sub-path to path managing slashes appropriately.
func (p *Path) Append(subpath *Path) *Path {
	if subpath == nil {
		return p
	}
	path, _ := New(p.String() + "/" + subpath.String())
	return path
}

// Implements Stringer interface
func (p *Path) String() string {
	if p == nil {
		return "<nil>"
	}
	return string(*p)
}

// Returns true of path starts with "/", false otherwise.
func (p *Path) IsAbsolute() bool {
	if len(*p) == 0 {
		return false
	}
	return string(*p)[0:1] == "/"
}

func (p *Path) EveryPath() []string {
	names := strings.Split(string(*p), "/")
	var paths []string
	var working string
	if p.IsAbsolute() {
		working = "/"
		names = names[1:]
	}
	names = names[:len(names)-1]
	for _, n := range names {
		working += n + "/"
		paths = append(paths, working)
	}
	return paths
}

func (p *Path) Split() (string, string) {
	if p == nil || len(string(*p)) == 0 {
		return "", ""
	}
	path := string(*p)
	if path[:1] != "/" {
		path = "./" + path
	}
	d, n := slash.Split(path[:len(path)-1])
	if d == "./" {
		d = ""
	}
	return d, n + "/"
}

func removeUrlPrefix(path string) string {
	p := path
	if i := strings.LastIndex(path, ":"); i != -1 {
		p = path[i+1:]
	}
	return p
}
