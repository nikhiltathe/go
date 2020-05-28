package main

import (
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	t.Skip()
	main()
}

func TestMain2(t *testing.T) {
	cmd := exec.Command("./Coverage.exe")
	//cmd := exec.Command("./test.sh")
	// out, err := cmd.Output()
	//t.Log("\n\n\n :", string(out), "\n\n")
	err := cmd.Run()
	t.Log("\n\n\n :", err, "\n\n")
	if err != nil {
		t.Error("Failed with error as", err)
	}
}
