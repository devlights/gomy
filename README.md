# What's this?

This package is a personal library for Go language.

For more information, see [![PkgGoDev](https://pkg.go.dev/badge/github.com/devlights/gomy)](https://pkg.go.dev/github.com/devlights/gomy) .

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/devlights/gomy)

![Go Version](https://img.shields.io/badge/go-1.20-blue.svg)[![CodeFactor](https://www.codefactor.io/repository/github/devlights/gomy/badge/master)](https://www.codefactor.io/repository/github/devlights/gomy/overview/master)[![Go Report Card](https://goreportcard.com/badge/github.com/devlights/gomy)](https://goreportcard.com/report/github.com/devlights/gomy)



## Environments

```sh
$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 22.04.2 LTS
Release:        22.04
Codename:       jammy
```

```sh
$ go version
go version go1.20.5 linux/amd64
```

## Requirements

### [go-task](https://taskfile.dev/)

```sh
$ go install github.com/go-task/task/v3/cmd/task@latest
```

## Howto

### Build

```sh
$ task fmt vet
$ task build
```

### Test

```sh
$ task test
```

### Cover

```sh
$ task cover
```
