package main

import (
	"os/exec"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func applications() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Application to install:", []string{"BACK", "ALL", "pamac-aur", "yay-bin", "htop", "neovim", "virt-manager (also installs qemu-desktop+libvirt)", "librewolf-bin", "zed", "steam"})
		switch choice {
		case 1:
			return
		case 2:
			doAll = true
			fallthrough
		case 3:
			managePackage("-S", "pamac-aur")

			if !doAll {
				break
			}
			fallthrough
		case 4:
			managePackage("-S", "yay-bin")

			if !doAll {
				break
			}
			fallthrough
		case 5:
			managePackage("-S", "htop")

			if !doAll {
				break
			}
			fallthrough
		case 6:
			managePackage("-S", "neovim")

			if !doAll {
				break
			}
			fallthrough
		case 7:
			cmd := exec.Command("pacman", "-S", "virt-manager", "qemu-desktop", "libvirt", "--noconfirm")
			err := cmd.Run()
			if err != nil {
				other.PrintError("Failed to perform application installations for virt-manager", 1)
			}
			manageService("enable", "libvirtd.service")
			manageService("start", "libvirtd.service")
			// TODO manage user groups!

			if !doAll {
				break
			}
			fallthrough
		case 8:
			managePackage("-S", "librewolf-bin")

			if !doAll {
				break
			}
			fallthrough
		case 9:
			managePackage("-S", "zed")

			if !doAll {
				break
			}
			fallthrough
		case 10:
			managePackage("-S", "steam")
		}
	}
}
