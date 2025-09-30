package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

const mountPoint = "/mnt/cph"

var username string
var partitionR, partitionB string

// TODO add mode to perform EVERYTHING
func main() {
	userInfo, err := user.Current()
	if err != nil {
		other.PrintError("Failed to get current user", 1)
	}
	if userInfo.Uid != "0" {
		other.PrintError("This utility must be run as root", 1)
	}

	partitionR = front.Input("Target root partition (/dev/driveparition):")
	partitionB = front.Input("Target boot (efi) partition (/dev/driveparition):")
	err = os.MkdirAll("/mnt/cph", 0777)
	if err != nil {
		other.PrintError("Failed to create mount directory (/mnt/cph)", 1)
	}

	for {
		fmt.Println()
		choice := front.InputMenuGen("Option:", []string{"pre-mount tweaks", "universal tweaks", "system tweaks", "removals", "applications", "configuration", "cosmic"})
		switch choice {
		case 1:
			if partitionIsMounted() {
				fmt.Println(partitionR + " is already mounted on " + mountPoint)
			} else {
				premount()
			}
		case 2:
			mountPartition()
			universalTweaks()
		case 3:
			mountPartition()
			system()
		case 4:
			mountPartition()
			removals()
			// TODO remove config/data
		case 5:
			mountPartition()
			applications()
		case 6:
			mountPartition()
			configurations()
		case 7:
			mountPartition()
			cosmic()
			// TODO keyboard shortcuts for screenshot/sticky window/terminal
			// TODO window decoration position (not yet implemented?)
			// TODO link CachyOS wallpapers?
		}
	}
}
