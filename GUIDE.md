# User Guide

The folder that contains all the saved configuration files will be located at `EqualizerAPO\config\E-APO-Config-Switcher\config-files`

## Adding your own config

Let's say you have this configuration for Equalizer APO:

```
Channel: all
Preamp: -2.5 dB
Filter 1: ON LS Fc 500 Hz Gain 5 dB
```

and you want to call the configuration `My Config`

Copy the config text (shown above) and place in a file named `My Config.txt`

Move this file into the EACS configuration folder

Restart E-APO-Config-Switcher and you will see your new configuration appear in the list

## Editing a config

Find the `.txt` in the configuration folder of the config you want to edit

Open the file and edit the text how you want

Once you save the file, un-check (if it was previously checked) and check the configuration in E-APO-Config-Switcher and your new configuration will be applied

## Remove a config

To remove a config all you have to do is delete the `.txt` file associated with that configuration

Restart E-APO-Config-Switcher and you will see the configuration does not appear in the list anymore
