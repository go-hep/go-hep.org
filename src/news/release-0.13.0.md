---
title: "Go HEP release 0.13.0"
date: 2018-06-19T17:49:25+02:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1292867.svg)](https://doi.org/10.5281/zenodo.1292867)

Release [`v0.13.0`](https://github.com/go-hep/hep/tree/v0.13.0) is fresh from the oven.

This release ships with major improvements in the [xrootd](https://godoc.org/go-hep.org/x/hep/xrootd) implementation and a few fixes in `rootio`.

## rootio

- leveraging the work that happened in `xrootd`, `rootio` is now able to read files over `[x]root`:

```go
import (
    "go-hep.org/x/hep/rootio"
)

func foo() {
    f, err := rootio.Open("root://ccxrootdgotest.in2p3.fr:9001/tmp/rootio/testdata/small-flat-tree.root")
    if err != nil { ... }
    defer f.Close()
}
```

- all the `root-xyz` commands can now also leverage `xrootd`:
```
$> root-ls -t root://ccxrootdgotest.in2p3.fr:9001/tmp/rootio/testdata/small-flat-tree.root
=== [root://ccxrootdgotest.in2p3.fr:9001/tmp/rootio/testdata/small-flat-tree.root] ===
version: 60806
TTree          tree                 my tree title (entries=100)
│ Int32        "Int32/I"            TBranch
│ Int64        "Int64/L"            TBranch
│ UInt32       "UInt32/i"           TBranch
│ UInt64       "UInt64/l"           TBranch
│ Float32      "Float32/F"          TBranch
│ Float64      "Float64/D"          TBranch
│ Str          "Str/C"              TBranch
│ ArrayInt32   "ArrayInt32[10]/I"   TBranch
│ ArrayInt64   "ArrayInt64[10]/L"   TBranch
│ ArrayUInt32  "ArrayInt32[10]/i"   TBranch
│ ArrayUInt64  "ArrayInt64[10]/l"   TBranch
│ ArrayFloat32 "ArrayFloat32[10]/F" TBranch
│ ArrayFloat64 "ArrayFloat64[10]/D" TBranch
│ N            "N/I"                TBranch
│ SliceInt32   "SliceInt32[N]/I"    TBranch
│ SliceInt64   "SliceInt64[N]/L"    TBranch
│ SliceUInt32  "SliceInt32[N]/i"    TBranch
│ SliceUInt64  "SliceInt64[N]/l"    TBranch
│ SliceFloat32 "SliceFloat32[N]/F"  TBranch
│ SliceFloat64 "SliceFloat64[N]/D"  TBranch
```

- support for seeking (_i.e.:_ event random access) has been added to scanners connected to chains of `rootio.Tree`s,

- `rootio` can now automatically generate streamers for `std::vector<T>` when a streamer for `T` exists,

- `rootio` has been updated to `v2` of `pierrec/lz4` library to decode `LZ4` compressed `ROOT` files.

## xrootd

- support for `ping` and `protocol` requests
- support for `dirlist`, `open`, `close` and `sync` requests
- support for `read` and `write` requests
- support for `rm`, `rmdir` and `truncate` requests
- support for `stat`, `vstat`, `statx`, `mkdir`, `mv` and `chmod` requests
- support for `signing` requests
- support for `auth+unix` request
- introduction of the `xrd-cp` command to copy files from a remote `xrootd` server:

```
$> go doc go-hep.org/x/hep/xrootd/cmd/xrd-cp
Command xrd-cp copies files and directories from a remote xrootd server to
local storage.

Usage:

    $> xrd-cp [OPTIONS] <src-1> [<src-2> [...]] <dst>

Example:

    $> xrd-cp root://server.example.com/some/file1.txt .
    $> xrd-cp root://gopher@server.example.com/some/file1.txt .
    $> xrd-cp root://server.example.com/some/file1.txt foo.txt
    $> xrd-cp root://server.example.com/some/file1.txt - > foo.txt
    $> xrd-cp -r root://server.example.com/some/dir .
    $> xrd-cp -r root://server.example.com/some/dir outdir

Options:

    -r	copy directories recursively
    -v	enable verbose mode
```

- introduction of the `xrd-ls` command to list the contents of directories on a remote `xrootd` server:

```
$> go doc go-hep.org/x/hep/xrootd/cmd/xrd-ls
Command xrd-ls lists directory contents on a remote xrootd server.

Usage:

    $> xrd-ls [OPTIONS] <dir-1> [<dir-2> [...]]

Example:

    $> xrd-ls root://server.example.com/some/dir
    $> xrd-ls -l root://server.example.com/some/dir
    $> xrd-ls -R root://server.example.com/some/dir
    $> xrd-ls -l -R root://server.example.com/some/dir

Options:

    -R	list subdirectories recursively
    -l	use a long listing format
```

- a convenience `xrootd/xrdfs.FileSystem` interface has been introduced to model interacting with the remote `xrootd` server's filesystem, following the [os](https://godoc.org/os) package API
- a convenience `xrootd/xrdfs.File` interface has been introduced to model interacting with the remote `xrootd` file, following the [os.File](https://godoc.org/os#File) API
- a convenience `xrootd/xrdio.File` type, implementing various `io.Xyz` interfaces has been introduced as well.
