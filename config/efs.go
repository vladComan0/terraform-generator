package config

import "embed"

//go:embed base.tmpl files/*.tmpl
var Files embed.FS
