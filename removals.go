package main

import (
	"github.com/rwinkhart/go-boilerplate/front"
)

func removals() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Application to remove:", []string{"back", "all", "linux-cachyos-lts"})
		var target string
		switch choice {
		case 1:
			return
		case 2:
			target = "linux-cachyos-lts"
		}
		if !doAll {
			managePackage("-Rcns", target)
		}
	}
}
