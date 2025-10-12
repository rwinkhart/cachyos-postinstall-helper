# CachyOS Post-Install Helper
This is an opinionated utility meant to customize the base CachyOS install to match my preferences.

Notable Tweaks:
- ext4 fast_commit option is enabled
- root account is locked
- user is added to uucp group for serial console access
- sudo is replaced with doas
- ipv6 privacy extensions are enabled
- SysRq reboots are enabled
- kernel security mitigations are disabled
- amdgpu tuning is fully unlocked
- amd_pstate is enabled in passive mode
- the lts kernel is removed
- many applications can be installed alongside custom configs
- custom configs for pacman/makepkg (optimal builds)
- custom Cosmic DE config
- option to select specific linux-firmware packages required for your system

## Usage
Begin by installing CachyOS via the standard GUI installer.

If you intend on using all options provided by this utility, it is best to use systemd-boot, ext4, and to ***un***check the following during installation:
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

As soon as the system is installed (prior to booting into it), run this utility as root and follow the prompts! That's it.
