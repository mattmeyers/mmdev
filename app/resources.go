package app

import (
	"embed"
	"io/fs"
)

//go:embed templates style.css
var Resources embed.FS

//go:embed favicons
var favicons embed.FS

var Favicons = sub(favicons, "favicons")

func sub(f fs.FS, dir string) fs.FS {
	subFS, err := fs.Sub(f, dir)
	if err != nil {
		panic(err)
	}
	return subFS
}
