# go-qt-wasm
Realtime svg generation using golang webassembly + [quicktemplates](https://github.com/valyala/quicktemplate) at 60 fps

# FLASHING COLOR SIEZURE WARNING
I took the famous tiger svg image and changed every color to be a random color.
I tried to make it less obnoxious by making the colors greyscale but results may vary.

# building
You need go 1.11 beta3 for wasm and go module support
```
go get -u github.com/valyala/quicktemplate/qtc
make build
```

# viewing
after building `make serve` http://localhost:3000