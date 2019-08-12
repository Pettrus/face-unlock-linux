<p align="center">
  <img src="https://raw.githubusercontent.com/Pettrus/go-face-unlock/master/mascot.png" width="40%">
</p>
Go face unlock is a script created to provide similar functionality as Windows Hello but without a IR emitters to validate you using facial recognition on linux machines.

It uses [PAM](https://en.wikipedia.org/wiki/Linux_PAM)(Pluggable Authentication Modules) so it works on login screen, sudo and su on the terminal.

### Pre-Install
Before we continue first we need to get dlib and go face going

#### Ubuntu 16.04, Ubuntu 18.04

You may use [dlib PPA](https://launchpad.net/~kagamih/+archive/ubuntu/dlib)
which contains latest dlib package compiled with Intel MKL support:

```bash
sudo add-apt-repository ppa:kagamih/dlib
sudo apt-get update
sudo apt-get install libdlib-dev libjpeg-turbo8-dev
```

#### Ubuntu 18.10+, Debian sid

Latest versions of Ubuntu and Debian provide suitable dlib package so just run:

```bash
# Ubuntu
sudo apt-get install libdlib-dev libopenblas-dev libjpeg-turbo8-dev
# Debian
sudo apt-get install libdlib-dev libopenblas-dev libjpeg62-turbo-dev
```

**ONLY FOR UBUNTU 18.10+ AND DEBIAN SID:**  
It won't install pkgconfig metadata file so create one in
`/usr/local/lib/pkgconfig/dlib-1.pc` with the following content:

```
libdir=/usr/lib/x86_64-linux-gnu
includedir=/usr/include

Name: dlib
Description: Numerical and networking C++ library
Version: 19.10.0
Libs: -L${libdir} -ldlib -lblas -llapack
Cflags: -I${includedir}
Requires:
```

### Installation

## From binary
Download the lastest compiled version [here](https://github.com/Pettrus/face-unlock-linux/releases)

Then run:
```
sudo ./main install
```

## From source
```
git clone https://github.com/Pettrus/go-face-unlock.git
cd go-face-unlock
go get github.com/Kagami/go-face
go build main.go camera.go facialRecognition.go file.go
```

Finally run the installer:
```
sudo ./main install
```

The script will take a picture of you to use as base for the facial recognition, so stand still and be on a well lit room for this first photo.

After installation everything should be working, open a new terminal and type some sudo command to see it working.

### Add new pictures

You can add more pictures to improve the recognition, for that use the command:
```
sudo ./main add
```

### Uninstall

```
sudo ./main uninstall
```

### Security

This script is by no mean a secure way to unlock your computer, just a convenient way to do it, if someone looks limilar to you or just a printed photo of you could fool the script.

### Dependencies
- [go-face](https://github.com/Kagami/go-face) - Face recognition with Go
- [go-webcam](https://github.com/blackjack/webcam) - Golang webcam library for Linux


### Similar projects

- [Howdy](https://github.com/boltgolt/howdy) - Windows Hello™ style facial authentication for Linux
