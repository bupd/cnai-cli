package modelfile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/CloudNativeAI/modctl/cmd/modelfile"
	"github.com/bupd/cnai-cli/pkg/storage"
	"github.com/sirupsen/logrus"
)

func GenerateModelFile() {
	fmt.Println("generating model file")
	// get the modelfile pkg from modctl
	// Call the Cobra root command to start the CLI
	modelfile.RootCmd.SetArgs([]string{"generate", "."})
	if err := modelfile.RootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func GotoModelDir(modelDir string) {
	currDir := storage.GetCurrDir()
	err := os.Chdir(filepath.Join(currDir, modelDir))
	if err != nil {
		logrus.Fatalf("failed cd into model dir: %v/%v, %v", currDir, modelDir, err)
	}
}
