package path_test

import (
	"errors"
	"testing"

	"github.com/cheekybits/is"

	"github.com/Avalanche-io/path"
)

func addrString(s string) *string {
	return &s
}

func TestPath(t *testing.T) {
	// init
	is := is.New(t)
	tests := []struct {
		In    *string
		Exp   *string
		Error error
	}{
		{
			In:    nil,
			Exp:   addrString("<nil>"),
			Error: nil,
		},
		{
			In:    addrString(""),
			Exp:   addrString("<nil>"),
			Error: errors.New("empty path"),
		},
		{
			In:    addrString("foo.bar"),
			Exp:   addrString("foo.bar/"),
			Error: nil,
		},
		{
			In:    addrString("/foo.bar"),
			Exp:   addrString("/"),
			Error: nil,
		},
		{
			In:    addrString("foo.bar/"),
			Exp:   addrString("foo.bar/"),
			Error: nil,
		},
		{
			In:    addrString("bat/foo.bar"),
			Exp:   addrString("bat/"),
			Error: nil,
		},
		{
			In:    addrString("/bat/foo.bar"),
			Exp:   addrString("/bat/"),
			Error: nil,
		},
		{
			In:    addrString("http://bat/foo.bar"),
			Exp:   addrString("/bat/"),
			Error: nil,
		},
		{
			In:    addrString("/a/b/baz/bat/foo.bar"),
			Exp:   addrString("/a/b/baz/bat/"),
			Error: nil,
		},
		{
			In:    addrString("/a/b/baz/bat/foo/bar"),
			Exp:   addrString("/a/b/baz/bat/foo/"),
			Error: nil,
		},
	}

	// nothing to do

	// check
	for _, tt := range tests {
		var p *path.Path
		var err error
		if tt.In != nil {
			p, err = path.New(*tt.In)
			is.Equal(err, tt.Error)
		}
		is.Equal(p.String(), *tt.Exp)
	}
}

func TestAppend(t *testing.T) {
	// init
	is := is.New(t)
	p1, err := path.New("/foo/bar/")
	is.NoErr(err)

	p2, err := path.New("/some/sub/path/")
	is.NoErr(err)
	p3, err := path.New("")
	is.Equal(err.Error(), "empty path")

	is.Equal(p1.Append(nil).String(), "/foo/bar/")
	is.Equal(p1.Append(p3).String(), "/foo/bar/")
	is.Equal(p1.Append(p2).String(), "/foo/bar/some/sub/path/")
}

func TestEveryPath(t *testing.T) {
	// init
	is := is.New(t)
	allPaths := []string{
		"one/",
		"one/two/",
		"one/two/three/",
		"one/two/three/four/",
		"one/two/three/four/five/",
		"one/two/three/four/five/six/",
		"one/two/three/four/five/six/seven/",
		"one/two/three/four/five/six/seven/eight/",
		"one/two/three/four/five/six/seven/eight/nine/",
		"one/two/three/four/five/six/seven/eight/nine/ten/",
		"one/two/three/four/five/six/seven/eight/nine/ten/eleven/",
	}
	allPathsAbs := []string{
		"/one/",
		"/one/two/",
		"/one/two/three/",
		"/one/two/three/four/",
		"/one/two/three/four/five/",
		"/one/two/three/four/five/six/",
		"/one/two/three/four/five/six/seven/",
		"/one/two/three/four/five/six/seven/eight/",
		"/one/two/three/four/five/six/seven/eight/nine/",
		"/one/two/three/four/five/six/seven/eight/nine/ten/",
		"/one/two/three/four/five/six/seven/eight/nine/ten/eleven/",
	}

	testp := "one/two/three/four/five/six/seven/eight/nine/ten/eleven/"
	testp_abs := "/one/two/three/four/five/six/seven/eight/nine/ten/eleven/"
	// non-absolute path
	p, err := path.New(testp)
	is.NoErr(err)
	for i, f := range p.EveryPath() {
		is.Equal(allPaths[i], f)
	}
	// absolute path
	p, err = path.New(testp_abs)
	is.NoErr(err)
	for i, f := range p.EveryPath() {
		is.Equal(allPathsAbs[i], f)
	}
}

func TestSplit(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		In  string
		Exp []string
	}{
		{
			In:  "foo.bar/",
			Exp: []string{"", "foo.bar/"},
		},
		{
			In:  "/",
			Exp: []string{"", "/"},
		},
		{
			In:  "bat/",
			Exp: []string{"", "bat/"},
		},
		{
			In:  "/bat/",
			Exp: []string{"/", "bat/"},
		},
		{
			In:  "/a/b/baz/bat/",
			Exp: []string{"/a/b/baz/", "bat/"},
		},
		{
			In:  "/a/b/baz/bat/foo/",
			Exp: []string{"/a/b/baz/bat/", "foo/"},
		},
	}
	for _, test := range tests {
		p, err := path.New(test.In)
		is.NoErr(err)
		dir, name := p.Split()
		is.Equal(dir, test.Exp[0])
		is.Equal(name, test.Exp[1])
	}
}

func TestIsDir(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		In  string
		Exp bool
	}{
		{
			In:  "foo.bar/",
			Exp: true,
		},
	}
	for _, test := range tests {
		is.Equal(path.IsDir(test.In), test.Exp)
	}
}
