# E-APO-Config-Switcher

A small Windows app that sits in the tray and allows you to quickly switch between using different [Equalizer APO](https://sourceforge.net/projects/equalizerapo/) configuration files

![screenshot](screenshot.png)

## Features

- Easily enable and disable configurations
- Edit, add, and remove configurations just by changing `.txt` files
- Everything happens in the system tray, which keeps things simple

## Warnings

*Both of these will be fixed in a future release*

- This will erase any configuration you currently have (and will continue to erase any changes you make to `config.txt`)
- This will not work alongside other configuration programs, such as Peace

## Install

- Download the latest `E-APO-Config-Switcher.zip` from [releases](https://github.com/psidex/E-APO-Config-Switcher/releases/latest)
- Extract it to `<Equalizer APO install location>\EqualizerAPO\config`
- Run `E_APO_Config_Switcher.exe` that is now inside `config\E-APO-Config-Switcher`
- You should now have the icon in your system tray, you can now left/right click on it to switch configurations
- If you want it to run on system startup, create a shortcut to the exe and move it to the windows startup directory

## Edit / Add / Remove Configurations

Simply copy the configuration text you want and put it in a `.txt` file inside `E-APO-Config-Switcher\E-APO-Config-Files`. If you look in there you will see the configurations that are included with this app

## Credits

- [zserge/tray](https://github.com/zserge/tray) to handle the system tray interaction
- Inspired by [Peace](https://sourceforge.net/projects/peace-equalizer-apo-extension/)
