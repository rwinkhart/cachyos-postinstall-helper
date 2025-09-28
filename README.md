# CachyOS Post-Install Helper
This is a simple utility meant to address what I personally dislike about the base CachyOS installation.

## Usage
Begin by installing CachyOS via the standard GUI installer.

If you intend on using all options provided by this utility, it is best to use systemd-boot, ext4, and to **_un_**check the following during installation:
- CachyOS Packages
    - cachyos-micro-settings
    - cachyos-wallpapers
- CachyOS shell configuration
- Base-devel + Common packages
    - X-system
    - Network
        - modemmanager
    - firewall
    - packages management
        - paru
        - octopi
    - audio
        - pavucontrol
    - Some applications selection
        - alacritty
        - btop
        - nano-syntax-highlighting
        - vi
        - ripgrep
        - micro
        - nano
        - vim
- CPU specific Microcode update packages
    - whichever one you don't need
- Firefox and language package

As soon as the system is installed and booted, run this utility as root and follow the prompts! That's it.

If using to configure Cosmic DE, I recommend running this utility from a raw TTY without Cosmic running.

A reboot is recommended if applying system tweaks.
