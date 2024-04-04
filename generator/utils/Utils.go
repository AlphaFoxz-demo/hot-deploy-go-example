package utils

import (
	"github.com/AlphaFoxz/hot-deploy-go-example/generator/utils/logutils"
	"os"
)

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		logutils.LogRed("Error", err.Error())
	}
}
