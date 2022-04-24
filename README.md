# NotifyXF

[![Go Reference](https://pkg.go.dev/badge/github.com/Mobilpadde/notifyxf.svg)](https://pkg.go.dev/github.com/Mobilpadde/notifyxf)

## What

This is a package to make it easier to use [notifyxf.com](https://notifyxf.com) in projects.

## Usage

If you want to control what notifications you'll get, please use the `notifyxf.Notify`-func.

```go
tkn := "...notifyxf token..."

n, err := notifyxf.NewNotifier(os.Getenv("NOTIFYXF_TOKEN"))
if err != nil {
    panic(err)
}

n.Notify("some message")
```
