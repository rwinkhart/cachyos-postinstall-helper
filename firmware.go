package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rwinkhart/go-boilerplate/back"
	"github.com/rwinkhart/go-boilerplate/front"
)

const firm = "linux-firmware-"

func firmwareSelector() {
	var nonOptional = []string{firm + "whence", firm + "cirrus", firm + "other"}
	var optional []string
	for {
		fmt.Println(back.AnsiBold + "\nSelected firmware: " + strings.Join(optional, ", ") + back.AnsiReset + "\n")
		choice := front.InputMenuGen("Firmware to install:", []string{"BACK", "INSTALL", "intel", "amdgpu", "nvidia", "mediatek", "realtek", "atheros", "qcom", "broadcom"})
		switch choice {
		case 1:
			return
		case 2:
			allPackages := append(nonOptional, optional...)
			managePackages("-Rcns", []string{"linux-firmware"})
			managePackages("-S", allPackages)
			if performAllTweaks {
				os.Exit(0)
			}
		case 3:
			optional = append(optional, firm+"intel")
		case 4:
			optional = append(optional, firm+"amdgpu")
		case 5:
			optional = append(optional, firm+"nvidia")
		case 6:
			optional = append(optional, firm+"mediatek")
		case 7:
			optional = append(optional, firm+"realtek")
		case 8:
			optional = append(optional, firm+"atheros")
		case 9:
			optional = append(optional, firm+"qcom")
		case 10:
			optional = append(optional, firm+"broadcom")
		}
	}
}
