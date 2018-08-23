---
title: "Go HEP release 0.14.0"
date: 2018-08-23T17:02:57+02:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1402652.svg)](https://doi.org/10.5281/zenodo.1402652)

Release [`v0.14.0`](https://github.com/go-hep/hep/tree/v0.14.0) is fresh from the oven.

This release is the result of some massive work in the [xrootd](https://go-hep.org/x/hep/xrootd) package thanks to [Mikhail Ivchenko (a.k.a @EgorMatirov)](https://github.com/EgorMatirov), our [Google Summer of Code 2018](https://summerofcode.withgoogle.com/) student.

While GSoC-2018 is now [over](https://opensource.googleblog.com/2018/08/thats-a-wrap-gsoc-2018.html), it's time to reflect on what Mikhail wrote:

- an almost complete `xrootd` [client](https://go-hep.org/x/hep/xrootd/client), compatible with the C++ implementation;
- [xrd-fuse](https://go-hep.org/x/hep/xrootd/cmd/xrd-fuse), a command to mount the contents of a remote XRootD server, locally;
- the beginnings of an `xrootd` [server](https://go-hep.org/x/hep/xrootd/server).

Most notably, the `client` package allowed to:

- to create the `xrd-cp` and `xrd-ls` commands that copy and list the contents of a remote XRootD server,
- to seamlessly read ROOT files over XRootD.

The `client` package handles authentication with the `unix` and `kerberos` protocols.
Unfortunately, authentication via `GSI` couldn't be implemented because there were no publicly available specifications for that protocol, see [xrootd/xrootd#757](https://github.com/xrootd/xrootd/issues/757) for more details.

Here is the final report of this year's GSoC:

- https://gist.github.com/EgorMatirov/303e4019606677d1838677a5a7145420

Thanks a lot Mikhail, hope we'll see you around :)


Another big new feature is the ability to write ROOT files, directly with `go-hep/rootio`.
This is still very much a work in progress, though, as only writing "empty" ROOT files or writing ROOT files with `TObjString`s have been explicitly tested.
Next release should see explicit support for writing histograms and graphs.

Lastly, improvements on the build and continuous integration procedure have been applied during this release cycle:

- added [AppVeyor](https://ci.appveyor.com/project/sbinet/hep) CI, for Windows: [![appveyor](https://ci.appveyor.com/api/projects/status/qnnp26vv2c71f560?svg=true)](https://ci.appveyor.com/project/sbinet/hep)
- added [code coverage](https://codecov.io/gh/go-hep/hep), with [codecov.io](https://codecov.io): [![codecov](https://codecov.io/gh/go-hep/hep/branch/master/graph/badge.svg)](https://codecov.io/gh/go-hep/hep)

## Code coverage improvements

The following packages have been updated, with additional tests, to improve their code coverage:

- `brio`, `csvutil`, `csvutil/csvdriver`,
- `fit`, `fmom`,
- `heppdt`, `hepmc`, `hepevt`,
- `hbook/ntup`, `hplot`,
- `lcio`, `lhef`,
- `rootio`, `sio`,
- `xrootd`.

Still some more work is needed to bring code coverage to a saner level (from `~55%` to `~70-80%`.)
Help more than welcome: it's "just" a matter of creating examples and tests.

## hepevt

- `AsciiEncoder` was renamed into simply `Encoder`
- `AsciiDecoder` was renamed into simply `Decoder`

## rootio

As noted above, it is now possible to create ROOT files.
The `rootio` package has a couple of examples:

- [create an empty file](https://godoc.org/go-hep.org/x/hep/rootio#example-Create--Empty)
- [create a file with a TObjString](https://godoc.org/go-hep.org/x/hep/rootio#example-Create)

Here is how you would create a ROOT file with one `TObjString` in it:

```go
package main

import (
	"fmt"
	"log"

	"go-hep.org/x/hep/rootio"
)

func main() {
	w, err := rootio.Create(fname)
	if err != nil {
	    log.Fatal(err)
	}
	defer w.Close()

	var (
	    k   = "my-objstring"
	    v   = rootio.NewObjString("Hello World from Go-HEP!")
	)

	err = w.Put(k, v)
	if err != nil {
	    log.Fatal(err)
	}

	fmt.Printf("wkeys: %d\n", len(w.Keys()))
	
	err = w.Close()
	if err != nil {
	    log.Fatalf("could not close file: %v", err)
	}
}
```

It is quite possible the API may change as we gain experience with what it ought to look like in a Go world.
_e.g._ it is possible that `rootio.Directory.Put` would get clever enough to automatically translate a Go builtin, like `string`, in the ROOT equivalent, say `TObjString`, on the fly.
Still pondering on that...

Again, still a lot of work to do on the writing side of things:

- support for `TH1x`,
- support for `TH2x`,
- support for `TGraph`, `TGraphErrors`, `TGraphAsymErrors`
- support for `TTree` (this will take some time)
- support for user-provided types.

## xrootd

The `xrootd/server` package ships with support for the following XRootD requests:

- `open`, `read`, `close`,
- `write`, `stat`, `truncate`, `sync`, `rename`.

There is no authentication support (yet) on the server: **DO NOT RUN THIS ON PUBLICLY ACCESSIBLE MACHINES** :)

The `xrootd` package gained a new sub-command:

- [xrootd/cmd/xrd-srv](https://go-hep.org/x/hep/xrootd/cmd/xrd-srv)

```
$> xrd-srv -h
xrd-srv serves data from a local filesystem over the XRootD protocol. 

Usage:

 $> xrd-srv [OPTIONS] <base-dir>

Example:

 $> xrd-srv /tmp
 $> xrd-srv -addr=0.0.0.0:1094 /tmp

Options:
  -addr string
    	listen to the provided address (default "0.0.0.0:1094")
```

## AOB

Support for Go-1.6 and Go-1.7 has been dropped.
Please upgrade to the latest and finest Go version (`1.11` is around the corner, with support for [Go modules](https://tip.golang.org/doc/go1.11#modules).)

Another interesting possible development avenue: exposing ROOT `TTree`s as [Apache Arrow Arrays](https://arrow.apache.org/).
This would allow for a better interoperability with that ecosystem and the tools it provides for data science and data analysis in general.
Help wanted!
