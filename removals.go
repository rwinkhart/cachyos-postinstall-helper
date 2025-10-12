package main

import (
	"github.com/rwinkhart/go-boilerplate/front"
)

// TODO remove unneeded linux-firmware packages!
func removals() {
	for {
		var choice int
		if performAllTweaks {
			choice = 1
		} else {
			choice = front.InputMenuGen("Application to remove:", []string{"BACK", "linux-cachyos-lts"})
		}
		var targets []string
		switch choice {
		case 1:
			return
		case 2:
			targets = append(targets, "linux-cachyos-lts")
		}
		managePackages("-Rcns", targets)
		if performAllTweaks {
			break
		}
	}
}
