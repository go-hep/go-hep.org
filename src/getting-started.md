+++
date = "2016-11-24T14:09:34+01:00"
title = "Getting started"
weight = 11
tags = ["sphinx", "documentation"]
+++

`go-hep` is written in [Go](https://golang.org).
You will need to have a valid Go toolchain installed.

## Install Go

Installing a Go toolchain is explained on the main Go website: https://golang.org/doc/install.

You can either:

1. download the binaries for your operating system, 
2. install Go via your O/S package manager, or
3. recompiled from sources.

Option 1) is recommanded (Go has a quick release cycle so package managers tend to lag a bit.)

It is always OK to download the latest Go version and upgrade regularly as the Go toolchain has a solid reputation of backward compatibility (since version `1.0`.)

Once Go has been installed, you should be able to run:

```sh
$> go version
$> go env
```

## Install Go-HEP

`go-hep` packages are installable with the `go get` command:

```sh
$> go get go-hep.org/x/hep/fads
```

This will:

- fetch the sources for `fads` (via [git](https://git-scm.com/)),
- inspect the sources for package dependencies,
- fetch the dependencies (recursively),
- build and install the dependencies,
- build and install `fads`.

By default, `go get` compiles the package for your operating system and architecture.
But you can easily cross-compile for other operating systems and architectures:

```sh
$> GOOS=linux   GOARCH=amd64 go get go-hep.org/x/hep/fads
$> GOOS=windows GOARCH=386   go get go-hep.org/x/hep/fads
```

The list of supported operating systems and architectures is given by:

```sh
$> go tool dist list
```

To install a package and all the packages or commands under it (and, of course, their dependencies), use the ellipsis `...`:

```sh
$> go get go-hep.org/x/hep/fads/...
```

## Updating Go-HEP

To update your local install of `go-hep`, you would use again the `go get` command, with the `-u` switch:

```sh
$> go get -u go-hep.org/x/hep/fads
```
