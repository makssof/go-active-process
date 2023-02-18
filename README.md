# Go-Active-Process

This package allows you to get information (process filename) about the currently active (foreground) window.

> Note: works only on Windows OS!

[![GoDoc](https://godoc.org/github.com/makssof/go-active-process?status.svg)](https://godoc.org/github.com/makssof/go-active-process)
[![Release](https://img.shields.io/github/v/release/makssof/go-active-process.svg)](https://github.com/makssof/go-active-process/releases/)
[![License](https://img.shields.io/github/license/makssof/go-active-process.svg)](https://github.com/makssof/go-active-process/blob/master/LICENSE)

## Content

- [Installation](#installation)
- [Usage](#usage)
- [Contribution](#contribution)

## Installation

To install the package just run:

```bash
go get -u github.com/makssof/go-active-process
```

## Usage

```GO
package main

import (
    "fmt"

    ActiveProcess "github.com/makssof/go-active-process"
)

func main() {
    activeProcess, err := ActiveProcess.Get()

    if err != nil {
        fmt.Println(fmt.Errorf("error: %v", err))
    } else {
        fmt.Println(activeProcess)
    }
}
```

## Contribution

The package is open-sourced under the [MIT](LICENSE) license.

If you will find some error, want to add something or ask a question - feel free to create an issue and/or make a pull request.

Any contribution is welcome.