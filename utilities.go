package main

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func manageService(action, service string) {
	cmd := exec.Command("systemctl", "daemon-reload")
	err := cmd.Run()
	if err != nil {
		other.PrintError("Failed to perform daemon-reload", 1)
	}
	cmd = exec.Command("systemctl", action, service)
	err = cmd.Run()
	if err != nil {
		other.PrintError("Failed to "+action+" "+service, 1)
	}
}

func managePackage(action, targetPackage string) {
	cmd := exec.Command("pacman", action, targetPackage, "--noconfirm")
	err := cmd.Run()
	if err != nil {
		other.PrintError("Failed to "+action+" "+targetPackage, 1)
	}
}

func writeFile(path, data, owner string, perms os.FileMode) {
	err := os.WriteFile(path, []byte(data), perms)
	if err != nil {
		other.PrintError("Failed to write to "+path+":"+err.Error(), 1)
	}
	if owner != "root" {
		uid := getUID()
		err = os.Chown(path, uid, uid)
		if err != nil {
			other.PrintError("Failed to set ownership of "+path, 1)
		}
	}
}

func getUsername() string {
	if username == "" {
		username = front.Input("Non-root username:")
	}
	localUsername := username
	return localUsername
}

func getUID() int {
	userInfo, err := user.Lookup(getUsername())
	if err != nil {
		other.PrintError("Failed to lookup UID for "+getUsername(), 1)
	}
	uid, err := strconv.Atoi(userInfo.Uid)
	if err != nil {
		other.PrintError("Failed to convert UID to integer", 1)
	}
	return uid
}
