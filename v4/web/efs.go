package web

import "embed"

//go:embed "js"
var Files embed.FS

//go:embed "styles"
var Static embed.FS
