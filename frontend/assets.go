//go:build !dev
// +build !dev

package frontend

import "embed"

//go:embed dist/assets/*
//go:embed dist/img/*
//go:embed dist/public/index.html
var assets embed.FS

func Assets() embed.FS {
	return assets
}
