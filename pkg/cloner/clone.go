package cloner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func DownloadModel(repo string) string {
	if len(repo) < 1 {
		logrus.Fatalf("give the correct hugging face url as arg")
	}

	giturllong := strings.Split(repo, "/")
	lastPart := giturllong[len(giturllong)-1]

  fmt.Println("cloning model", repo)
  fmt.Println("This could take few minutes, based on your model size", repo)

	cmd := exec.Command("git", "clone", repo)
	err := cmd.Run()
	if err != nil {
		logrus.Fatalf("failed git clone: %v, %v", repo, err)
	}

	return lastPart
}
