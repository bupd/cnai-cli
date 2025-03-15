package pusher

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

func PushModel(regUrl string) {
	// cmd.Execute()
	// .Execute().SetArgs([]string{"generate", "."})
	cmd := exec.Command("modctl", "push", regUrl)
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("failed modctl push: %v, %v", regUrl, err)
	}
}
