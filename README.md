<div align="center">
    <h1>Guark</h1>
    <p>Guark allows you to build beautiful user interfaces using modern web technologies such as Vue.js, React.js..., while your app logic handled and powered by the amazing <b>Go</b>.</p>
    <p align="center">
        <a href="#installation">Installation</a> ❘
        <a href="#getting-started">Getting Started</a> ❘
        <a href="#contributing">Contributing</a> ❘
        <a href="#license">License</a>
    </p>
</div>

![Guark Vue Template](https://i.imgur.com/RhU6bh7.png)



## 🖳  About The Project

Guark is an open-source framework to build cross platform desktop GUI applications.

### 📢 What Guark stands for?

Go + Quark = Guark

### 🔮 Guark mission

Simplify cross platform desktop apps development.

### 🎸 How it works

Demo Video: https://youtu.be/_k_tq9sj-do

Guark backend and logic part handled by native Go code, while the user interfaces built with modern web technologies (Vue.js, React.js, etc...), and with Guark javascript API you can communicate between your Go and UI framework, and calling your exported Go functions and plugin(s) methods.

## 💌  Main Features

- Desktop applications with GO ♥
- One codebase for Gnu/Linux, macOS, and Windows.
- UI Hot Reload.
- You can use any front end framework.
- Supports Chrome, and native webview engine or both.
- Windows MSI bundler builtin.


## 📜  Installation

#### 1. Install guark CLI tool:
```bash
go install github.com/guark/guark/cmd/guark@latest
```

#### 2. Some Requirements:

```bash
// fedora
sudo dnf install gtk3-devel webkit2gtk3-devel gcc-c++ pkgconf-pkg-config

// Ubuntu
sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev build-essential

// Windows
// https://jmeubank.github.io/tdm-gcc/download/
```
Open a console and make sure the tdm-gcc tools chain are in the PATH:

```
gcc --version
gcc (tdm64-1) 10.3.0
Copyright (C) 2020 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```

## Getting Started

After installing guark CLI tool, the next step is to create a new guark project based on the template that you like:

### Create a new project

```bash
# cd to empty directory and run:
guark init --template vue --mod github.com/username/appname
```

### Start Dev Server

After creating new project run:
```bash
guark run
```

### Export your first GO function to JS API

##### 1. Create a new file in `lib/funcs` directory:
```go
// lib/funcs/foo_bar.go

import (
    "github.com/guark/guark/app"
)

func FooBar(c app.Context) (interface{}, error) {

   c.App.Log.Debug("Foo bar called")

   return "This is a my return value to javasript", nil
}
```

##### 2. Then export it to Guark JS API in `lib/config.go` file:

```go
// Exposed functions to guark Javascript api.
var Funcs = app.Funcs{
    "foo": funcs.FooBar,
}

```

##### 3. Then you can call your function in javascript:

```js
import g from "guark"

g.call("foo")
    .then(res => console.log(res))
    .catch(err => console.error(err))

```

See Vue template as an example: https://github.com/guark/vue

## Guark Engines:

You can change the engine in `guark.yaml` file.

- `webview`: Uses native system webview.
- `chrome`: Uses and requires google chrome.
- `hybrid`: Uses chrome if exists in the system, if chrome not available guark will switch to native webview by default.


## Build Your App

### Configure the build

guark-build.yaml contains all configuration and path required for the build, like the compilers or default platform.

The C/C++ compilers are not using the PATH environment and need an absolute path to them, on linux it will use the default paths for Linux builds.

For cross platform build, it will look like:

```
linux:
  ldflags: ""

darwin:
  ldflags: ""

windows:
  cc: /usr/bin/x86_64-w64-mingw32-gcc
  cxx: /usr/bin/x86_64-w64-mingw32-g++
  ldflags: "-H windowsgui"
  windres: /usr/bin/x86_64-w64-mingw32-windres
```

Update these paths accordingly (check the GCC/distributions documentations for the right paths)

On Windows, make sure to have installed tdm-gcc, once done, update the guark-build.yaml:

```
# Guark build config file.

setup:
  - cmd: yarn install
    dir: ui
  - cmd: go mod download
  - cmd: go mod tidy
  - cmd: go mod verify

linux:
  ldflags: ""

darwin:
  ldflags: ""

windows:
  cc: C:\apps\tdm-gcc1030\bin\gcc
  cxx: C:\apps\tdm-gcc1030\bin\g++
  ldflags: "-H windowsgui"
  windres: C:\apps\tdm-gcc1030\bin\windres.exe
```

### Build

You can build your app with
```bash
guark build
```

## Bundle Windows App

After building your app you can bundle your windows app into msi using WIX.
```bash
guark bundle
```

#### Wix required!
Install it from: https://wixtoolset.org/

## Cross Compiling To Windows From Gnu/Linux:

You can build windows app from your linux based system, using `mingw64`

#### 1. Install mingw64:
```bash
// Fedora
sudo dnf install mingw64-gcc

// Ubuntu
sudo apt install binutils-mingw-w64
```

#### 2. Configure `guark-build.yaml` File:

Double check the binary paths in `guark-build.yaml`.

#### 3. Build The App:

```bash
# this command will build and compile linux, and windows app. you can find your compiled apps in `dist/` directory.
guark build --target linux --target windows
```

You can use any cross compiler for example: `guark build --target darwin`. just change the options in `guark-build.yaml` file.

#### Note
You can also bundle windows app into MSI file from your linux based system via `guark bundle`, but you need to install wix tools:

```bash
# fedora
dnf install msitools

# Ubuntu
sudo apt-get install wixl
```

## Contributing

PRs, issues, and feedback from ninja gophers are very welcomed.

## License

Guark is provided under the [MIT License](https://github.com/guark/guark/blob/master/LICENSE).

