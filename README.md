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


## 🛈 About The Project

Guark is an open-source framework to build cross platform desktop GUI applications.

### 📢   What Guark Stands For?
Go + Quark = Guark

### 🔮   Guark Mission
Simplify cross platform desktop apps development.

### 🎸   Who It Works
Guark backend and logic part handled by native Go code, while the user interfaces built with modern web technologies (vue, react, etc...), and guark javascript API allows you to call your exported go functions and plugins.


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

// React template coming soon...
// guark new --template react --dest myreactapp
// You can create new project from git url:
// guark new --template https://github.com/username/template --dest myreactapp
``` 

### Start dev server
After creating new project change your working directory to it and run: `guark dev`

### Build your app
You can build you app with `guark build`.  


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

## TODOs Before v1:

- [ ] Add More Tests.
- [ ] Test Guark Apps on MacOS and Windows.
- [ ] Auto Reload App When Go Code Changes.
- [ ] Guark App Installer (Cross Platform).
- [ ] Add More Templates (React.js, Framework7, and more..)
- [ ] Strip Builds


