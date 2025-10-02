package main

import (
	"github.com/rwinkhart/go-boilerplate/front"
)

// TODO remove unneeded linux-firmware packages!
func removals() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Application to remove:", []string{"BACK", "linux-cachyos-lts"})
		var target string
		switch choice {
		case 1:
			return
		case 2:
			target = "linux-cachyos-lts"
		}
		if !doAll {
			managePackages("-Rcns", []string{target})
		}
	}
}
