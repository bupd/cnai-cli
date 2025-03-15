package modelfile

import (
	"fmt"
	"os/exec"

	"github.com/CloudNativeAI/modctl/cmd/modelfile"
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
	cmd := exec.Command("cd", modelDir)
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("failed cd into: %v, %v", modelDir, err)
	}
}
