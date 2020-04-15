package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	baseDir := filepath.Dir(os.Args[0])
	cmd := exec.Command(fmt.Sprintf("%s\\__app.exe", baseDir))

	var out bytes.Buffer
	cmd.Stdout = &out

	path := os.Getenv("PATH")
	path = fmt.Sprintf("%s;%s\\lib;%s", baseDir, baseDir, path)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PATH=%s", path))

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		fmt.Println("OUT: ", out)
	}
}
