---
title: "Go HEP release 0.16.0"
date: 2018-12-03T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1886496.svg)](https://doi.org/10.5281/zenodo.1886496)

Release [`v0.16.0`](https://github.com/go-hep/hep/tree/v0.16.0) is fresh from the oven.

This new release introduces a new package, [groot](https://go-hep.org/x/hep/groot), that replaces `rootio`.

## groot & rootio

![groot](/images/news/release-0.16.0/groot.jpg)

`groot` is a new package that is meant to replace `rootio`.
For backward compatibility, `rootio` will remain in the repository for a couple of releases (presumably until `v1.0.0`.)

`groot` is the result of refactoring `rootio` in a couple of more focused packages with clear API boundaries:

- [groot/rbase](https://go-hep.org/x/hep/groot/rbase): definitions of basic ROOT classes (`Object`, `Named`, `ObjString`, ...)
- [groot/rbytes](https://go-hep.org/x/hep/groot/rbytes): definitions of types useful for serializing and deserializing ROOT data buffers, interfaces to interact with ROOT's metadata classes such as `StreamerInfo` and `StreamerElement`s.
- [groot/rcont](https://go-hep.org/x/hep/groot/rcont): definitions of ROOT container types (`TList`, `THashList`, `TObjArray`, `TArrayX`)
- [groot/rdict](https://go-hep.org/x/hep/groot/rdict): definitions of ROOT streamers (`TStreamerArtificial`, `TStreamerLoop`, ...)
- [groot/rhist](https://go-hep.org/x/hep/groot/rhist): definitions of ROOT types related to histograms and graphs (`TH1x`, `TH2x`, `TGraph`, `TGraphErrors`, `TGraphAsymmErrors`)
- [groot/riofs](https://go-hep.org/x/hep/groot/riofs): low-level types and functions to deal with opening and creating ROOT files; users should prefer using the [groot](https://go-hep.org/x/hep/groot) package to open and create ROOT files
- [groot/root](https://go-hep.org/x/hep/groot/root): ROOT core interfaces (`Object`, `Named`, `ObjArray`, ...)
- [groot/rsrv](https://go-hep.org/x/hep/groot/rsrv): exposes HTTP end-point to manipulate ROOT files, plot and create histograms and graphs from files or trees
- [groot/rtree](https://go-hep.org/x/hep/groot/rtree): interface to decode, read, concatenate and iterate over ROOT Trees
- [groot/rtypes](https://go-hep.org/x/hep/groot/rtypes): `rtypes` contains the means to register types (ROOT ones and user defined ones) with the ROOT type factory system
- [groot/rvers](https://go-hep.org/x/hep/groot/rvers): `rvers` contains the ROOT version and the classes' versions `groot` is supporting and currently reading.

Interacting with ROOT files should be performed with the `groot` package:

```go
import "go-hep.org/x/hep/groot"

func F() {
	f1, err := groot.Open("some/file.root")
	f2, err := groot.Create("some/other.root")
}
```

`groot` has the needed bootstrap code that allows to automatically import the registration code for all the ROOT types `groot` knows how to handle.

Here is a quick and dirty Rosetta code for migrating to `groot`:

- `rootio.H1` → `groot/rhist.H1`
- `rootio.H2` → `groot/rhist.H2`
- `rootio.Graph` → `groot/rhist.Graph`
- `rootio.Tree` → `groot/rtree.Tree`
- `rootio.ChainOf` → `groot/rtree.ChainOf`
- `rootio.Scanner` → `groot/rtree.Scanner`
- `rootio.File` → `groot/riofs.File`
- `rootio.Directory` → `groot/riofs.Directory`
- `rootio.Open` → `groot.Open`
- `rootio.Create` → `groot.Create`

### groot/rsrv & root-srv

`root-srv` has been refactored to extract the pure plot/file interaction machinery from the GUI part.
The plot creation and the ROOT file interaction parts have been refactored into a new package [go-hep.org/x/hep/groot/rsrv](https://go-hep.org/x/hep/groot/rsrv) that contains a couple of HTTP end-points that can be reused in third-party packages or applications.

`rsrv` exposes a REST API that expects JSON requests (`OpenFileRequest`, `PlotH1Request`, `PlotTreeRequest`, ...) and returns JSON responses.
The HTTP end-points are attached to the [rsrv.Server](https://pkg.go.dev/go-hep.org/x/hep/groot/rsrv#Server) type.

`root-srv` can now open ROOT files served over `xrootd`.

## hbook

`hbook` now exposes a `Bin` method on `H{1,2}D` to retrieve a bin by its `(x,y)` coordinates.

## hplot

Following an issue raised on `gonum/plot`, `hplot` now creates histograms with transparent background by default.

## rio

Improved test coverage (by adding some more tests.)

## xrootd

The low-level bits for the following requests have been implemented:

- `kXR_query`
- `kXR_prepare`
- `kXR_endsess`
- `kXR_locate`
- `kXR_decrypt`
- `kXR_admin`

The first steps to support the `"host"` security provider have also been implemented.

We are still missing the implementation for the GSI authentification protocol: still waiting on [xrootd](http://xrootd.org) to provide specifications for this protocol (progress is tracked here: [issue-757](https://github.com/xrootd/xrootd/issues/757).)

Improved test coverage.

## AOB

We will try to have preliminary support for writing `TTrees` in the next release.
That should be fun.

Interoperability with [Apache Arrow Arrays](https://arrow.apache.org) is in the works.
It might even prove to be easier to support Apache Arrow first and then implement `TTrees` writing support on top of that.
We will see...

Stay tuned! (and, as always, any kind of help (reviews, patches, nice emails, constructive criticism) deeply appreciated.)
