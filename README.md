<p align="center">
  <img src="https://raw.githubusercontent.com/Pettrus/go-face-unlock/master/mascot.png" width="40%">
</p>
Go face unlock is a script created to provide similar functionality as Windows Hello but without a IR emitters to validate you using facial recognition on linux machines.

It uses [PAM](https://en.wikipedia.org/wiki/Linux_PAM)(Pluggable Authentication Modules) so it works on login screen, sudo and su on the terminal.

### Installation
```
git clone https://github.com/Pettrus/go-face-unlock.git
cd go-face-unlock
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
