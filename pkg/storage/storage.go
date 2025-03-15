package storage

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func GetCurrDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("failed to get current directory: %v", err)
	}
	fmt.Println("current dir: ", dir)
	return dir
}
