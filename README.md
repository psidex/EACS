# E-APO-Config-Switcher

[![Build Status](https://travis-ci.com/psidex/E-APO-Config-Switcher.svg?branch=master)](https://travis-ci.com/psidex/E-APO-Config-Switcher)
[![license](https://img.shields.io/github/license/psidex/E-APO-Config-Switcher.svg)](LICENSE)

A small Windows tray app that allows you to quickly switch between using different [Equalizer APO](https://sourceforge.net/projects/equalizerapo/) configuration files

![screenshot](screenshot.png)

## Features

- Easily enable and disable configurations
- Edit, add, and remove configurations just by changing `.txt` files
- Everything happens in the system tray, which keeps things simple

## Warnings

- This won't work alongside other configuration programs such as Peace or the default configuration program that comes with Equalizer APO
- This will overwrite Equalizer APO's `config.txt`. Make a backup if you need to!

## Install

- Download the latest `E-APO-Config-Switcher.zip` from [releases](https://github.com/psidex/E-APO-Config-Switcher/releases/latest)
- Extract it to `<Equalizer APO install location>\EqualizerAPO\config`
- Run `E_APO_Config_Switcher.exe` that is now inside `config\E-APO-Config-Switcher`
- You should now have the icon in your system tray, you can now left/right click on it to switch configurations
- If you want it to run on system startup, create a shortcut to the exe and move it to the windows startup directory

## Edit / Add / Remove Configurations

The configuration folder (the folder that contains all the configuration files) is located at `EqualizerAPO\config\E-APO-Config-Switcher\config-files`

### Adding your own config

Let's say you have this configuration for Equalizer APO:

```
Channel: all
Preamp: -2.5 dB
Filter 1: ON LS Fc 500 Hz Gain 5 dB
```

and you want to call the configuration `My Config`

Copy the config text (shown above) and place in a file named `My Config.txt`

Move this file into the configuration folder

Restart E-APO-Config-Switcher and you will see your new configuration appear in the list

### Editing a config

Find the `.txt` in the configuration folder of the config you want to edit

Open the file and edit the text how you want

Once you save the file, un-check (if it was previously checked) and check the configuration in E-APO-Config-Switcher and your new configuration will be applied

### Remove a config

To remove a config all you have to do is delete the `.txt` file associated with that configuration

Restart E-APO-Config-Switcher and you will see the configuration does not appear in the list anymore

## Credits

- Inspired by [Peace](https://sourceforge.net/projects/peace-equalizer-apo-extension/)
