# What is it?
This project is a CLI appliction to control Quarta Radex One personal dosimeter.

At now, application is in early alpha and can only get measures from Radex One directly. But in future it can be drop-in replacement for official Radex Data Center software
# Installing drivers
Radex One is using Silicon Labs CP210x USB to UART drivers
## Windows
Drivers will be installed simultaneously with Radex Data Center (official application).
You can found it here - https://www.quarta-rad.ru/en/catalog/dozimetr-radiometr-radon/dozimetr-radex-one/

Tested on Windows 10
## Linux
From kernel version 5.0 driver is already prebuild with it.
But due to VID and PID are custom, device doesn't detects by default.

To fix this behavior you can run these commands:

    # Load silabs driver
    sudo modprobe cp210x
    # Forcing driver to device match
    sudo sh -c 'echo abba a011 > /sys/bus/usb-serial/drivers/cp210x/new_id'
Tested on Ubuntu 20.04 LTS

# Quick start
1. Install drivers
1. Download release from https://github.com/iamtio/goradex/releases

## Run on windows

    > ./goradex.exe -s COM9 measure
    CPM: 16, Ambient: 13, Accumulated: 208

## Run on linux based systems

    $ ./goradex -s /dev/ttyUSB0 measure
    CPM: 17, Ambient: 9, Accumulated: 208

# Contributing
This project is written in Golang. If you want to contribute code:

1. Ensure you are running golang version 1.15 or greater for go module support
1. Check-out the project: `git clone https://github.com/iamtio/goradex && cd goradex`
1. Make changes to the code
1. Build the project, e.g. via `go build -o goradex[.exe]`
1. Evaluate and test your changes `./goradex [SOME_COMMAND]`
1. Make a pull request
