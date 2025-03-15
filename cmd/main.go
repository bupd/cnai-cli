package main

import (
	"fmt"
	"os"

	"github.com/bupd/cnai-cli/pkg/builder"
	"github.com/bupd/cnai-cli/pkg/cloner"
	"github.com/bupd/cnai-cli/pkg/modelfile"
	"github.com/bupd/cnai-cli/pkg/pusher"
	"github.com/bupd/cnai-cli/pkg/storage"
	"github.com/bupd/cnai-cli/pkg/utils"
)

func main() {
	fmt.Println("I am back")

	origin := os.Args[1]
	destination := utils.RemoveHttps(os.Args[2])

	storage.GetCurrDir()
	// download stage
	modelDir := cloner.DownloadModel(origin)
  // GenerateModelFile
  modelfile.GotoModelDir(modelDir)
	modelfile.GenerateModelFile()
	// build stage
	builder.BuildModelPackage(destination)
	// push stage
	pusher.PushModel(destination)
}
