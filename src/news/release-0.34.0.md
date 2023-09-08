---
title: "Go HEP release 0.34.0"
date: 2023-09-08T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.8329416.svg)](https://doi.org/10.5281/zenodo.8329416)

Release [`v0.34.0`](https://github.com/go-hep/hep/tree/v0.34.0) is out of the oven.

This release contains no major API breakage, but all around improvements and bug fixes.

## hplot

- [hplot/vgop](https://pkg.go.dev/go-hep.org/x/hep@v0.34.0/hplot/vgop): a new package to implement JSON serialization of `gonum/plot` canvases

```go
p := hplot.New()
// ...
err := hplot.Save(p, 10*vg.Centimeter, 20*vg.Centimeter, "plot.json")
```

- `hplot` is now using `gonum/plot@v0.14.0` which has the necessary infrastructure to get nicer timeseries axes. See the provided examples ([here](https://pkg.go.dev/go-hep.org/x/hep@v0.34.0/hplot#Ticks)) for the `hplot.Ticks` type.

- `hplot` has gained some basic capability to display the legend (a color palette with numbers) of a 2D plot. See the provided [example](https://pkg.go.dev/go-hep.org/x/hep@v0.34.0/hplot#example-H2D-WithLegend).
- added a convenience forward function for `gonum/plot/plotter.NewLinePoints` (as `hplot.NewLinePoints`)

## groot

- added support for ROOT `6.28/04`
- [groot/rjson](https://pkg.go.dev/go-hep.org/x/hep@v0.34.0/groot/rjson): a new package to implement JSROOT-compatible JSON serialization of `groot`'s histogram types (`rhist.H{1,2}x`)

```go
h, err := f.Get("h1d")
raw, err := rjson.Marshal(h)
```

See the complete example [here](https://pkg.go.dev/go-hep.org/x/hep@v0.34.0/groot/rjson#example-Marshal).

- `groot/rtree` is now able to read trees created by Geant4 in multithreaded mode (see [here](https://github.com/go-hep/hep/issues/989) for more details).
- added support for the "new" `TLeafG` branch.

## Changelog


* [6e9be99e](/commit/6e9be99e) all: bump gioui@v0.30, x/{crypto,exp,image,text,tools}@latest
* [970796df](/commit/970796df) hplot{,vgop}: add support for Save("foo.json")
* [e1cbaa61](/commit/e1cbaa61) hplot/vgop: first import
* [301b96c6](/commit/301b96c6) groot/rjson: first import
* [d280b3c9](/commit/d280b3c9) groot: implement RSlicer for TH{1,2}x
* [8795baec](/commit/8795baec) groot/rbytes: introduce RSlicer interface
* [8631271d](/commit/8631271d) cmd/root2arrow: use latest arrow-cat for tests
* [1d490da0](/commit/1d490da0) cmd/root2npy: use internal/diff
* [0e69116b](/commit/0e69116b) groot: add tests for TLeafG
* [85fd5fd9](/commit/85fd5fd9) groot: add support for TLeafG
* [7cb2fac8](/commit/7cb2fac8) groot: bump to ROOT 6.28/04
* [548d277c](/commit/548d277c) groot/rtree: handle parallel-merged TBaskets
* [404ed51c](/commit/404ed51c) ci: bump staticcheck@2023.1.5
* [d41b03d0](/commit/d41b03d0) hplot: add NewLinePoints forward
* [b6113fad](/commit/b6113fad) hplot: add timeseries examples with epok
* [948b0093](/commit/948b0093) hplot: add automatic legend creation to H2D
* [d26c986b](/commit/d26c986b) hplot: add support for legend to Figure
* [02946662](/commit/02946662) hplot: forward plot.Legend
* [d0eea6ab](/commit/d0eea6ab) fwk: bump to gonuts/commander@v0.4.1
* [fa91d5d0](/commit/fa91d5d0) all: bump Gonum@v0.14.0, Gioui@v0.2.0
* [233c6a47](/commit/233c6a47) all: bump Go-1.21, drop Go-1.19
* [d85e15f7](/commit/d85e15f7) groot/cmd/root-gen-streamer: use an in-memory buffer for generation
* [c00ff1a0](/commit/c00ff1a0) hplot: update examples
* [8c9cc408](/commit/8c9cc408) hplot: update embedmd invocation
* [673ab4c5](/commit/673ab4c5) ci: improve mk-release to setup symlinks


