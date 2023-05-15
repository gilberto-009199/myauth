# Manager OTP


  Aplication for OTP in Desktop


# How Build from files

Dependencies Linux:

`Debian/Ubuntu`:
```shell
    sudo apt install libxcursor-dev libxinerama-dev libxrandr-dev libxi-dev libgl-dev libxxf86vm-dev
```
`Fedora`:
```shell
    sudo dnf install golang gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel
```
`Arch Linux`:
```shell
    sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi
```
`Solus`:
```shell
    sudo eopkg it -c system.devel golang mesalib-devel libxrandr-devel libxcursor-devel libxi-devel libxinerama-devel
```
`openSUSE`:
```shell
    sudo zypper install go gcc libXcursor-devel libXrandr-devel Mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel
```
`Void Linux`:
```shell
    sudo xbps-install -S go base-devel xorg-server-devel libXrandr-devel libXcursor-devel libXinerama-devel
```
`Alpine Linux`:
```shell
    sudo apk add go gcc libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev linux-headers mesa-dev
```
`NixOS`:
```shell
    nix-shell -p libGL pkg-config xorg.libX11.dev xorg.libXcursor xorg.libXi xorg.libXinerama xorg.libXrandr xorg.libXxf86vm
```