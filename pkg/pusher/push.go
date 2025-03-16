package pusher

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func PushModel(regUrl string) {
	// cmd.Execute()
	// .Execute().SetArgs([]string{"generate", "."})
	fmt.Printf("Pushing model to: %v", regUrl)
	cmd := exec.Command("modctl", "push", regUrl, "--plain-http")
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("failed modctl push: %v, %v", regUrl, err)
	}
}
