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
		Expct *string
		Error error
	}{
		{
			In:    nil,
			Expct: addrString("<nil>"),
			Error: errors.New("Nil path."),
		},
		{
			In:    addrString(""),
			Expct: addrString("<nil>"),
			Error: errors.New("Nil path."),
		},
		{
			In:    addrString("foo.bar"),
			Expct: addrString("foo.bar/"),
			Error: nil,
		},
		{
			In:    addrString("/foo.bar"),
			Expct: addrString("/"),
			Error: nil,
		},
		{
			In:    addrString("foo.bar/"),
			Expct: addrString("foo.bar/"),
			Error: nil,
		},
		{
			In:    addrString("bat/foo.bar"),
			Expct: addrString("bat/"),
			Error: nil,
		},
		{
			In:    addrString("/bat/foo.bar"),
			Expct: addrString("/bat/"),
			Error: nil,
		},
		{
			In:    addrString("http://bat/foo.bar"),
			Expct: addrString("/bat/"),
			Error: nil,
		},
		{
			In:    addrString("/a/b/baz/bat/foo.bar"),
			Expct: addrString("/a/b/baz/bat/"),
			Error: nil,
		},
		{
			In:    addrString("/a/b/baz/bat/foo/bar"),
			Expct: addrString("/a/b/baz/bat/foo/"),
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
		is.Equal(p.String(), *tt.Expct)
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
	is.Equal(err.Error(), "Nil path.")

	is.Equal(p1.Append(nil).String(), "/foo/bar/")
	is.Equal(p1.Append(p3).String(), "/foo/bar/")
	is.Equal(p1.Append(p2).String(), "/foo/bar/some/sub/path/")
}
