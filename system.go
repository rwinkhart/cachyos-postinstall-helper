package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/rwinkhart/go-boilerplate/front"
	"github.com/rwinkhart/go-boilerplate/other"
)

func system() {
	for {
		var doAll bool
		choice := front.InputMenuGen("Tweak:", []string{"BACK", "ALL", "replace sudo with doas", "enable ipv6 privacy extensions", "enable SysRq reboots", "disable kernel security mitigations", "unlock amdgpu tuning", "enable amd_pstate in passive mode (recommended for zenver2)"})
		switch choice {
		case 1:
			return
		case 2:
			doAll = true
			fallthrough
		case 3:
			writeFile(mountPoint+"/etc/doas.conf", "permit persist keepenv :wheel as root\n", "root", 0644)
			managePackage("-S", "opendoas")
			chrootCommandRun([]string{"pacman", "-S", "--asexplicit", "--dbonly", "archlinux-keyring", "autoconf", "automake", "binutils", "bison", "debugedit", "fakeroot", "file", "findutils", "flex", "gawk", "gcc", "gettext", "grep", "groff", "gzip", "libtool", "m4", "make", "pacman", "patch", "pkgconf", "sed", "texinfo", "which", "--noconfirm"})
			chrootCommandRun([]string{"pacman", "-R", "base-devel", "sudo", "--noconfirm"})

			if !doAll {
				break
			}
			fallthrough
		case 4:
			nics, err := os.ReadDir("/sys/class/net")
			if err != nil {
				other.PrintError("Failed to retrieve NICs", 1)
			}
			var sysctlConf strings.Builder
			sysctlConf.WriteString("net.ipv6.conf.all.use_tempaddr = 2\nnet.ipv6.conf.default.use_tempaddr = 2")
			for _, nic := range nics {
				sysctlConf.WriteString("\nnet.ipv6.conf." + nic.Name() + ".use_tempaddr = 2")
			}
			writeFile(mountPoint+"/etc/sysctl.d/40-ipv6-priv-ext.conf", sysctlConf.String()+"\n", "root", 0644)

			if !doAll {
				break
			}
			fallthrough
		case 5:
			writeFile(mountPoint+"/etc/sysctl.d/35-sysrq.conf", "kernel.sysrq = 244\n", "root", 0644)

			if !doAll {
				break
			}
			fallthrough
		case 6:
			addKernelParams([]string{"mitigations=off"})

			if !doAll {
				break
			}
			fallthrough
		case 7:
			addKernelParams([]string{"amdgpu.ppfeaturemask=0xffffffff"})

			if !doAll {
				break
			}
			fallthrough
		case 8:
			addKernelParams([]string{"amd_pstate=passive"})
		}
	}
}

func addKernelParams(newParameters []string) {
	f, err := os.Open(mountPoint + "/boot/loader/entries/linux-cachyos.conf")
	if err != nil {
		other.PrintError("Failed to open systemd-boot entry", 1)
	}
	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var parameters []string
	for i := 0; fScanner.Scan(); i++ {
		if i == 1 {
			parameters = strings.Split(fScanner.Text(), " ")[2:]
			break
		}
	}
	parameters = append(parameters, newParameters...)
	writeFile(mountPoint+"/boot/loader/entries/linux-cachyos.conf", "title Linux Cachyos\noptions root=UUID=e2ff600b-0202-4620-bed3-54d7a58414f8 "+strings.Join(parameters, " ")+"\nlinux /vmlinuz-linux-cachyos\ninitrd /initramfs-linux-cachyos.img\n", "root", 0700)
}
