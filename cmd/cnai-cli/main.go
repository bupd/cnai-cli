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

	if len(os.Args) < 3 || len(os.Args[1]) < 5 || len(os.Args[2]) < 1 {
		fmt.Println("Please provide the correct args for the command")
		fmt.Println("Example")
		fmt.Println("cnai-cli 'hugging-face_url' 'destination_oci_reg_url'")
    return
	}

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
