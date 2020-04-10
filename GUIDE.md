# User Guide

The folder that contains all the saved configuration files will be located at `EACS\config-files`

## Adding your own config

Let's say you have this configuration for Equalizer APO:

```
Channel: all
Preamp: -2.5 dB
Filter 1: ON LS Fc 500 Hz Gain 5 dB
```

and you want to call the configuration `My Config`

Create a file in `config-files` names `My Config.txt`

Copy the configuration text and paste it into `My Config.txt`

Restart EACS and you will see your new configuration appear in the list

## Editing a config

Simply find the `.txt` in the configuration folder of the config you want to edit, open the file and edit the text how you want

Once you save the file, you can click on it in EACS and your new configuration will be applied. If it was previously checked, just un-check and then check it again.

## Remove a config

To remove a config all you have to do is delete the `.txt` file associated with that configuration

Restart EACS and you will see the configuration does not appear in the list anymore
