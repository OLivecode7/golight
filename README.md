# GoLight
##  Firmware for lighting a WS2812 LED Strip for various Arduino chipsets using TinyGo
Author: olive@code7.dev /  2023-10-23

### This program use the WS2812 tinygo driver, a Arduino Board and a push button to switch between different light modes.

# Build and Flashing tools

Tinygo install is required and additional installs for flashing Arduino borads :

[tinygo.org/getting-started/install/](https://tinygo.org/getting-started/install/)

```bash
wget https://github.com/tinygo-org/tinygo/releases/download/v0.32.0/tinygo_0.32.0_amd64.deb

sudo dpkg -i tinygo_0.32.0_amd64.deb

Add path to bashrc:

nano ~/.bashrc

Scroll down to the end, add this line, press Ctrl+X then Y to save:

export PATH=$PATH:/usr/local/tinygo/bin

tinygo version 

tinygo version 0.32.0 linux/amd64 (using go version go1.22.0 and LLVM version 18.1.2)
```
### Install Arduino UNO Tools
```bash
sudo apt-get install gcc-avr
sudo apt-get install avr-libc
sudo apt-get install avrdude
```
### Build and flash command for Arduino Nano :
```bash
tinygo flash -target=arduino-nano -port /dev/ttyUSB0 -monitor  -scheduler=tasks -baudrate 9600 main.go
```
### Build and flash command for Arduino Uno :
```bash
tinygo flash -target arduino -monitor -port /dev/ttyACM0 main.go 
```
### For Arduino 33 BLE :

In order to flash your TinyGo programs onto the Arduino Nano33 BLE board, you will need to install the “bossac_arduino2” command line utility which is a special build of the BOSSA command line utilities.
```
Linux 

Download the bossac_arduino2 program from http://downloads.arduino.cc/tools/bossac-1.9.1-arduino2-linux64.tar.gz

Extract the downloaded file to a directory on your computer.

Make sure to add that directory into your PATH.
```

### Build and flash command for Arduino nano 33 BLE :
```bash
tinygo flash -target=nano-33-ble -port /dev/ttyACM0 main.go
```

## TODO

- Persist mode beetween reboot with Epromm storage :

https://github.com/tinygo-org/tinyfs/tree/release

- Monitor power voltage : If batteries are used need to switch card off under a certain voltage level

https://forum.arduino.cc/t/measurement-of-bandgap-voltage/38215/16

https://forum.arduino.cc/t/making-accurate-analog-measurements-and-detecting-power-fails/526108/3
