package builder

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func BuildModelPackage(destination string) {
  fmt.Println("building model")
	cmd := exec.Command("modctl", "build", ".", "-t", destination)
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("failed modctl build: %v", err)
	}
}
