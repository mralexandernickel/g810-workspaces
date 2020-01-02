# g810-workspaces

Script that is colouring the keyboard-leds as an indicator of available and
focused workspaces in [sway](https://github.com/swaywm/sway) or
[i3](https://github.com/i3/i3) window manager.

This script is using [g810-led](https://github.com/MatMoul/g810-led), so every
keyboard that is supported by it, should also be supported by this script.

## How to use?

Put the script into your `$PATH` and edit your sway- or i3-config to call the
script whenever you switch your workspace. An example-config could look like
this:

(...explain what is to edit and how else it could look...)

```bash
bindsym $mod+1 workspace number $ws1, exec g810-workspaces
bindsym $mod+2 workspace number $ws2, exec g810-workspaces
bindsym $mod+3 workspace number $ws3, exec g810-workspaces
bindsym $mod+4 workspace number $ws4, exec g810-workspaces
bindsym $mod+5 workspace number $ws5, exec g810-workspaces
bindsym $mod+6 workspace number $ws6, exec g810-workspaces
bindsym $mod+7 workspace number $ws7, exec g810-workspaces
bindsym $mod+8 workspace number $ws8, exec g810-workspaces
bindsym $mod+9 workspace number $ws9, exec g810-workspaces
bindsym $mod+0 workspace number $ws10, exec g810-workspaces
```

...in this config we are using the default-values. To find out which options are
available (and which ones are default) you can call `g810-workspaces --help`.

## Dependencies

Please make sure that [g810-led](https://github.com/MatMoul/g810-led) is
available on your system.