package main

import (
	"os"
	"os/exec"

	"github.com/rwinkhart/go-boilerplate/other"
)

func universalTweaks() {
	// Lock root account
	cmd := exec.Command("usermod", "--delete", "root")
	cmd.Run()
	cmd = exec.Command("usermod", "--lock", "root")
	cmd.Run()

	// Add user to uucp group for non-root serial console usage
	cmd = exec.Command("usermod", "-aG", "uucp", getUsername())
	cmd.Run()

	// Force disable kernel watchdog (kernel parameters should do this as well)
	f, err := os.Create("/etc/modprobe.d/blacklist.conf")
	if err != nil {
		other.PrintError("Failed to create /etc/modprobe.d/blacklist.conf", 1)
	}
	f.WriteString("# Disable kernel watchdog\nblacklist iTCO_wdt")
	f.Close()
}
