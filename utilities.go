package main

import (
	"os"
	"os/exec"
	"strings"

	mnt "github.com/moby/sys/mountinfo"
	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func manageService(action, service string) {
	chrootCommandRun([]string{"systemctl", "daemon-reload"})
	chrootCommandRun([]string{"systemctl", action, service})
}

func managePackages(action string, targetPackages []string) {
	chrootCommandRun(append([]string{"pacman", action, "--noconfirm"}, targetPackages...))
}

func writeFile(path, data, owner string, perms os.FileMode) {
	err := os.WriteFile(path, []byte(data), perms)
	if err != nil {
		other.PrintError("Failed to write to "+path+":"+err.Error(), 1)
	}
	if owner != "root" {
		err = os.Chown(path, 1000, 1000)
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

func partitionIsMounted() bool {
	if partitionR == "/" {
		return true
	}

	isMounted, err := mnt.Mounted(mountPoint)
	if err != nil {
		other.PrintError("Failed to verify whether "+partitionR+" is mounted", 1)
	}
	return isMounted
}

func mountPartition() {
	if partitionR == "/" {
		return
	}

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
	var cmd *exec.Cmd
	var errorSlice int
	if partitionR == "/" {
		errorSlice = 0
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		errorSlice = 1
		args = append([]string{mountPoint}, args...)
		cmd = exec.Command("arch-chroot", args...)
	}
	err := cmd.Run()
	if err != nil {
		other.PrintError("Failed to run \""+strings.Join(args[errorSlice:], " ")+"\" within chroot: "+err.Error(), 1)
	}
}

func mkdir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		other.PrintError("Failed to create \""+path+"/\" directory", 1)
	}
	err = os.Chown(path, 1000, 1000)
	if err != nil {
		other.PrintError("Failed to set ownership of \""+path+"\"", 1)
	}
}
