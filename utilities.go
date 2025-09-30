package main

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	mnt "github.com/moby/sys/mountinfo"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func manageService(action, service string) {
	chrootCommandRun([]string{"systemctl", "daemon-reload"})
	chrootCommandRun([]string{"systemctl", action, service})
}

func managePackage(action, targetPackage string) {
	chrootCommandRun([]string{"pacman", action, targetPackage, "--noconfirm"})
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

func partitionIsMounted() bool {
	isMounted, err := mnt.Mounted(mountPoint)
	if err != nil {
		other.PrintError("Failed to verify whether "+partitionR+" is mounted", 1)
	}
	return isMounted
}

func mountPartition() {
	if !partitionIsMounted() {
		// root
		cmd := exec.Command("mount", partitionR, mountPoint)
		err := cmd.Run()
		if err != nil {
			other.PrintError("Failed to mount "+partitionR+" on "+mountPoint, 1)
		}

		// boot
		cmd = exec.Command("mount", partitionB, mountPoint+"/boot")
		err = cmd.Run()
		if err != nil {
			other.PrintError("Failed to mount "+partitionB+" on "+mountPoint+"/boot", 1)
		}
	}
}

func chrootCommandRun(args []string) {
	args = append([]string{mountPoint}, args...)
	cmd := exec.Command("arch-chroot", args...)
	err := cmd.Run()
	if err != nil {
		other.PrintError("Failed to run \""+strings.Join(args[1:], " ")+"\" within chroot: "+err.Error(), 1)
	}
}
