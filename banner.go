package utils

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"runtime"
)

func GetBanner(title string, version string) string {
	if version == "" {
		version = "develop"
	}

	banner := figure.NewColorFigure(fmt.Sprintf("%s - %s", title, version), "", "blue", true)
	banner.Print()

	return fmt.Sprintf("go version: %s", runtime.Version())
}
