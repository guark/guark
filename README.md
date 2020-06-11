# Guark

Guark allows you to build beautiful user interfaces using modern web technologies such as Vue.js, React.js..., while your app logic handled and powered by the amazing **Go**.


## Cross Compiling Status.

|  ⬇️ Platform ⼁ Build Target ➡️ |  Build to linux | Build to Windows  | Build to MacOS  |
|---|---|---|---|
| Linux    |  ✔  | [1] ✔ |  ⚠  |
| MacOS    |  ⚠  |   ⚠   |  ✔  |
| Windows  |  ⚠  |   ✔   |  ⚠  |

- ⚠: Work In Progress.
- ✔: Supported.
- [1]: `mingw64-gcc` Required.


#### Install mingw64:
```bash
// Fedora
sudo dnf install mingw64-gcc

// Ubuntu
sudo apt install binutils-mingw-w64
```

You can use any GCC cross compiler for example: `env CC=.. CXX=.. guark build --target darwin`.

## TODOs Before v1:

- [ ] Add More Tests.
- [ ] Test Guark Apps on MacOS and Windows.
- [ ] Auto Reload App When Go Code Changes.
- [ ] Guark App Installer (Cross Platform).
- [ ] Add More Templates (React.js, Framework7, and more..)
- [ ] Strip Builds


