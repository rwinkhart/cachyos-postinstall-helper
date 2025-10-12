package main

import (
	"fmt"
	"os"
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
		choice := front.InputMenuGen("Option:", []string{"ALL", "pre-mount tweaks", "universal tweaks", "system tweaks", "removals", "applications", "configuration", "cosmic", "firmware selector"})
		switch choice {
		case 1:
			performAllTweaks = true
			fallthrough
		case 2:
			if partitionIsMounted() {
				fmt.Println("\n" + partitionR + " is already mounted on " + mountPoint)
				os.Exit(1)
			} else {
				allTweaksAnnounce("pre-mount tweaks")
				premount()
			}

			if !performAllTweaks {
				break
			}
			fallthrough
		case 3:
			allTweaksAnnounce("universal tweaks")
			mountPartition()
			universalTweaks()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 4:
			allTweaksAnnounce("system tweaks")
			mountPartition()
			system()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 5:
			allTweaksAnnounce("removals")
			mountPartition()
			removals()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 6:
			allTweaksAnnounce("application installs")
			mountPartition()
			applications()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 7:
			allTweaksAnnounce("configurations")
			mountPartition()
			configurations()

			if !performAllTweaks {
				break
			}
			fallthrough
		case 8:
			allTweaksAnnounce("cosmic tweaks")
			mountPartition()
			cosmic()
			// TODO window decoration position (not yet implemented)
			// TODO screen blank shortcut (not yet implemented)

			if !performAllTweaks {
				break
			}
			fallthrough
		case 9:
			allTweaksAnnounce("firmware selection")
			firmwareSelector()
		}
		if !performAllTweaks {
			os.Exit(0)
		}
	}
}
