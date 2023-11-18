package main

import (
	"os"

	"github.com/toj900/grab/cmd"
	"k8s.io/klog/v2"
)

func main() {
	if err := cmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}
