package flags

import (
	"github.com/labbs/mynotes/pkg/config"
	"github.com/urfave/cli/v2"
)

func DocumentFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "excalidraw-libs-path",
			Aliases:     []string{"e"},
			EnvVars:     []string{"EXCALIDRAW_LIBS_PATH"},
			Usage:       "Path to Excalidraw libraries",
			Value:       "./excalidraw_libs",
			Destination: &config.Document.ExcalidrawLibsPath,
		},
	}
}
