package builder

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

func BuildModelPackage(destination string) {
	cmd := exec.Command("modctl", "build", ".", "-t", destination)
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("failed modctl build: %v", err)
	}
}
