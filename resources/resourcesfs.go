package resources

import (
	"embed"
	"io/fs"
)

//go:embed all:static
var static embed.FS

func GetStaticFs() fs.FS {
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	return sub
}
