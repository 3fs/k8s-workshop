# [Workshop](../README.md) &raquo;Installation instructions

## Table of contents

- [Install Docker - Ubuntu](#install-docker---ubuntu)
  - [System requirements](#system-requirements)
  - [Installation steps](#installation-steps)
- [Install Docker - Windows](#install-docker---windows)
  - [System requirements](#system-requirements-1)
  - [Installation steps](#installation-steps-1)
- [Install Docker - Mac](#install-docker---mac)
  - [System requirements](#system-requirements-2)
  - [Installation steps](#installation-steps-2)
- [Official instructions](#official-instructions)

## Install Docker - Ubuntu

### System requirements

- Cosmic 18.10 or
- Bionic 18.04 (LTS) or
- Xenial 16.04 (LTS)

### Installation steps

1. Add GPG key:

```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

2. Add repository

```bash
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
```

3. Install

```bash
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

## Install Docker - Windows

- After Hyper-V is enabled, VirtualBox no longer works, but any VirtualBox VM
  images remain.

### System requirements

- Windows 10 64bit
- At least Intel i* processor
- Virtualization is enabled in BIOS
- At least 4GB of RAM.

### Installation steps

1. Download installer from
   `https://download.docker.com/win/stable/Docker%20for%20Windows%20Installer.exe`
2. Follow the install wizard to accept the license, authorize the installer, and
   proceed with the install

## Install Docker - Mac

### System requirements

- Mac hardware must be a 2010 or newer model
- macOS Sierra 10.12 and newer macOS releases are supported
- At least 4GB of RAM
- VirtualBox prior to version 4.3.30 must NOT be installed

### Installation steps

1. Download installer from:
   `https://download.docker.com/mac/stable/31259/Docker.dmg`
2. Double-click Docker.dmg to open the installer, then drag Moby the whale to
   the Applications folder.

## Official instructions

- [Linux - Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
- [Windows](https://docs.docker.com/docker-for-windows/install/)
- [Mac](https://docs.docker.com/docker-for-mac/install/)
