## binimg

A command-line application for drawing single color images with transparent background.

The **color**, **scale** and **output location** can be tweaked via flags.

Call with `-h` flag to find out how. It is super simple :wink:.

### Usage
- draw with your mouse inside the terminal
    - left click draws
    - right click erases
- press `Ctrl+C` to clear
- press `Ctrl+S` when happy
- image will be saved in **PNG** format to the specified location or *image.png* by default
- press `Esc` to quit

### External dependencies

- github.com/gdamore/tcell
- github.com/nfnt/resize

### Installation

```sh
git clone https://github.com/rokf/binimg
cd binimg
make install
```

### License

This library is free software; you can redistribute it and/or modify it under the terms of the MIT license. See LICENSE for details.
