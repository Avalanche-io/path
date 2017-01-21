package path

import (
	"errors"
	"path/filepath"
	"strings"
)

type Path string

const Separator string = "/"
const DoubleSeparator string = "//"

// New creates a path from a string, with the following rules:
//
// Removes all URL prefixes.
// Reduces two or more separators to one.
// Appends a separator if there are no separators in the string.
// Removes all characters after the last separator.
//
// If input is empty or nil, return nil path and error.
//
func New(path string) (*Path, error) {
	path = removeUrlPrefix(path)
	path = deduplicateSlash(path)

	l := len(path)
	switch {
	case l == 0:
		return nil, errors.New("Nil path.")
	case path[l-1] == Separator[0]:
		p := Path(path)
		return &p, nil
	case strings.Index(path, Separator) == -1:
		p := Path(path + Separator)
		return &p, nil // No slashes at all. Appending slash to end of string.
	default:
		d, _ := filepath.Split(path)
		p := Path(d)
		return &p, nil // No trailing slash. Removed last item in path.
	}
}

// Appends sub-path to path managing slashes appropriately.
func (p *Path) Append(subpath *Path) *Path {
	if subpath == nil {
		return p
	}
	path, _ := New(p.String() + Separator + subpath.String())
	return path
}

// Implements Stringer interface
func (p *Path) String() string {
	if p == nil {
		return "<nil>"
	}
	return string(*p)
}

func deduplicateSlash(path string) string {
	p := path
	i := strings.Index(p, DoubleSeparator)

	for i != -1 {
		// p = p[:i] + p[i+1:]
		p = strings.Replace(p, DoubleSeparator, Separator, -1)
		i = strings.Index(p, DoubleSeparator)
	}
	return p
}

func removeUrlPrefix(path string) string {
	p := path
	if i := strings.LastIndex(path, ":"); i != -1 {
		p = path[i+1:]
	}
	return p
}
