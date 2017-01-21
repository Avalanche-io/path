# path

Package path implements a strongly typed path string. It enforces a forward slash ("/") separator, and and insures all paths end with a separator.

## Example 

```golang
package main

import (
    "fmt"

    "github.com/Avalanche-io/path"
)

func main() {
    p, err := path.New("/path/to/file.txt")
    if err != nil {
        panic(err)
    }
    fmt.Println(p) // "/path/to/"
    p, err = path.New("/path/to/dir/")
    if err != nil {
        panic(err)
    }
    fmt.Println(p) // "/path/to/dir/"
    p, err = path.New(".")
    if err != nil {
        panic(err)
    }
    fmt.Println(p) // "./"
}

```


# License

This package is governed by the MIT License.  See LICENSE file for details.