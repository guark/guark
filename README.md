<div align="center">
    <a href="https://github.com/guark/guark">
        <img src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/svg/51.svg" width="200">
    </a>
    <h1>Guark</h1>
    <p>Guark allows you to build beautiful user interfaces using modern web technologies such as Vue.js, React.js..., while your app logic handled and powered by the amazing <b>Go</b>.</p>
</div>

![Guark Framework7 Demo App](https://github.com/guark/guark/raw/master/testdata/demo.png)

<p align="center">
    <a href="#installation">Installation</a> ❘
    <a href="#getting-started">Getting Started</a> ❘
    <a href="#contributing">Contributing</a> ❘
    <a href="#license">License</a>
</p>


## 🖳 About The Project

Guark is an open-source framework to build cross platform desktop GUI applications.

### 📢   What Guark stands for?

Go + Quark = Guark

### 🔮   Guark mission

Simplify cross platform desktop apps development.

### 🎸   Who it works

Guark backend and logic part handled by native Go code, while the user interfaces built with modern web technologies (vue, react, etc...), and guark javascript API allows you to call your exported go functions and plugins.

### 📐   Important note

This is a v0 "WIP" prototype of guark, still a lot to do to make it to production v1. your feedback is very appreciated.

## 💌  Features

- Desktop applications with GO ♥
- One codebase for: Linux, MacOS, and Windows
- Hot Reload UI (and auto restart app on Go code change coming soon)
- You can use any front end framework
- Using native system webview (unlike electron we do not embed chrome in the builds)
- Cross Compile (WIP)
- App Hooks
- App Watchers
- App Plugins (You can make your own)
- Simple
- And more..


## 📜   Installation

Install guark cli tool.
```bash
go install github.com/guark/guark/cmd/guark
```

If you on Linux❤ you need to install `webkit2gtk3`:
```bash
// fedora
sudo dnf install webkit2gtk3-devel

// Ubuntu
sudo apt install libwebkit2gtk-4.0-dev
```

## Getting Started

After installing guark CLI tool, the next step is to create a new guark project based on template that you like:

### Create new project

```bash
guark new --template vue --dest myapp
``` 

React and more templates coming soon...


### Start dev server

After creating new project change your working directory to it and run: `guark dev`

### Build your app

You can build your app with `guark build`. 

### Export your first GO function to JS API

You need to create a file for your function in `lib/funcs` directory.
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

Then export it to Guark JS API in `lib/config.go` file.
```go
import "github.com/your_username/your_app/lib/funcs"

// Exposed functions to guark Javascript api.
var Funcs = app.Funcs{
	"foo": funcs.FooBar,
}

```

Then you can call your function in javascript:
```js
import g from "guark"

g.call("foo")
 .then(res => console.log(res))
 .catch(err => console.error(err))
```

See Vue template as an example: https://github.com/guark/vue


## Cross Compiling Status.

|   Platform  ⬇️  ⼁ Build Target ➡️ |  Build to linux | Build to Windows  | Build to MacOS  |
|---|---|---|---|
| Linux    |  ✔  | [1] ✔ |  ⚠  |
| MacOS    |  ⚠  |   ⚠   |  ✔  |
| Windows  |  ⚠  |   ✔   |  ⚠  |

- Work In Progress. ⚠
- Supported. ✔
- [1]: `mingw64-gcc` Required.


#### Install mingw64:
```bash
// Fedora
sudo dnf install mingw64-gcc

// Ubuntu
sudo apt install binutils-mingw-w64
```

You can use any cross compiler for example: `env CC=.. CXX=.. guark build --target darwin`.

## V1 Roadmap:

- [ ] Add App Watchers.
- [ ] Add More Templates (React.js, Framework7, and more..)
- [ ] Add More Tests.
- [ ] Test Guark Apps on MacOS and Windows.
- [ ] Auto Reload App When Go Code Changes.
- [ ] Guark App Installer (Cross Platform).
- [ ] Strip Binaries on build


