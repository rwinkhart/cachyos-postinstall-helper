package main

import (
	"os"

	"github.com/rwinkhart/go-boilerplate/other"
)

func universalTweaks() {
	// Lock root account
	chrootCommandRun([]string{"passwd", "--delete", "root"})
	chrootCommandRun([]string{"usermod", "--lock", "root"})

	// Add user to uucp group for non-root serial console usage
	chrootCommandRun([]string{"usermod", "-aG", "uucp", getUsername()})

	// Force disable kernel watchdog (kernel parameters should do this as well)
	f, err := os.Create(mountPoint + "/etc/modprobe.d/blacklist.conf")
	if err != nil {
		other.PrintError("Failed to create "+mountPoint+"/etc/modprobe.d/blacklist.conf", 1)
	}
	f.WriteString("# Disable kernel watchdog\nblacklist iTCO_wdt")
	f.Close()
}
