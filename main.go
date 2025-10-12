package main

import (
	"fmt"
	"os/user"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

var username string
var partitionR, partitionB string
var mountPoint = "/mnt/cph"
var performAllTweaks bool

func main() {
	userInfo, err := user.Current()
	if err != nil {
		other.PrintError("Failed to get current user", 1)
	}
	if userInfo.Uid != "0" {
		other.PrintError("This utility must be run as root", 1)
	}

	partitionR = front.Input("Target root partition (/dev/driveparition OR \"/\" if already booted):")
	if partitionR != "/" {
		partitionB = front.Input("Target boot (efi) partition (/dev/driveparition):")
		mkdir("/mnt/cph")
	} else {
		mountPoint = ""
	}

	for {
		fmt.Println()
		choice := front.InputMenuGen("Option:", []string{"ALL", "pre-mount tweaks", "universal tweaks", "system tweaks", "removals", "applications", "configuration", "cosmic"})
		switch choice {
		case 1:
			performAllTweaks = true
			fallthrough
		case 2:
			if partitionIsMounted() {
				fmt.Println("\n" + partitionR + " is already mounted on " + mountPoint)
			} else {
				premount()
			}

			if !performAllTweaks {
				break
			}
			fallthrough
		case 3:
			mountPartition()
			universalTweaks()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 4:
			mountPartition()
			system()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 5:
			mountPartition()
			removals()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 6:
			mountPartition()
			applications()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 7:
			mountPartition()
			configurations()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 8:
			mountPartition()
			cosmic()
			// TODO window decoration position (not yet implemented)
			// TODO screen blank shortcut (not yet implemented)
		}
	}
}
