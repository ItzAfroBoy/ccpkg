package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	var projectName string
	platforms := []string{"darwin/amd64", "darwin/arm64", "windows/386", "windows/amd64", "linux/amd64"}
	raw, err := os.ReadFile("go.mod")
	if err != nil {
		projectName = "useless"
	} else {
		projectName = path.Base(strings.Split(string(raw), "\n")[0])
	}
	for _, v := range platforms {
		var output string
		split := strings.Split(v, "/")
		_os, arch := fmt.Sprintf("GOOS=%s", split[0]), fmt.Sprintf("GOARCH=%s", split[1])
		if split[0] == "windows" {
			output = fmt.Sprintf("bin/%s-%s-%s.exe", projectName, split[0], split[1])
		} else {
			output = fmt.Sprintf("bin/%s-%s-%s", projectName, split[0], split[1])
		}
		cmd := exec.Command("go", "build", "-o", output)
		cmd.Env = append(cmd.Environ(), _os, arch)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		fmt.Printf("Project built for: %s\n", v)
	}
}
