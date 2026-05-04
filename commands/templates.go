package commands

import (
	"embed"
)

//go:embed template/*.tmpl
var TemplatesFS embed.FS
