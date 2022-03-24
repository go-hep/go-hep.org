---
title: "Go Hackathon - CERN"
date: 2022-03-24T17:00:00+01:00
weight: 10
---

[GoHACK Devpost](https://gohack.devpost.com/) is *LIVE!*


[Google](https://google.com) and [CERN](https://home.cern) are teaming up to host a [Go](https://golang.org) hackathon and improve scientific applications or libraries written in [Go](https://golang.org).

Either working on general scientific apps/libs, or working on 3 [CERN](https://home.cern)-related applications or libraries:

- [Reva](https://reva.link/): a gRPC server written in Go to provide interoperability between different storage and applications,
- [Golbd](https://github.com/cernops/golbd): a DNS load balancer written in Go to efficiently manage the thousands of computer nodes in the CERN Computer Center,
- [Go-HEP](https://go-hep.org)

More informations about the hackathon are available here:
- [https://gohack.devpost.com/](https://gohack.devpost.com/)

Don't hesitate to spread the word.

In preparation for this hackathon, we have setup 3 projects (they are
but suggestions):

## `groot-rntuple`: adding support for ROOT-7 `RNTuple` file format

- https://github.com/go-hep/hep/projects/5

`groot` should be able to read and write the new [RNTuple](https://github.com/root-project/root/blob/master/tree/ntuple/v7/doc/specifications.md) file format from the upcoming ROOT v7.

as the format is still experimental, code for reading/writing such files should live under:
- `go-hep.org/x/hep/groot/exp/rntup`

initially, support for the following items should be added:
- r/w support for "raw" file mode (ie: _sans_ r/w into a ROOT file, just a "bare" UNIX file)
- r/w support for [basic types](https://github.com/root-project/root/blob/master/tree/ntuple/v7/doc/specifications.md#basic-types)
- support for [envelopes](https://github.com/root-project/root/blob/master/tree/ntuple/v7/doc/specifications.md#envelopes)
- support for [std::string](https://github.com/root-project/root/blob/master/tree/ntuple/v7/doc/specifications.md#mapping-of-c-types-to-fields-and-columns)
- support for [std::vector<T> w/ T=builtin](https://github.com/root-project/root/blob/master/tree/ntuple/v7/doc/specifications.md#mapping-of-c-types-to-fields-and-columns)

support for [std::vector<T>](https://github.com/root-project/root/blob/master/tree/ntuple/v7/doc/specifications.md#mapping-of-c-types-to-fields-and-columns) where `T` is a non-builtin may be added later.

## `groot-member-wise-streaming`: adding support for reading/writing member-wise objects

- https://github.com/go-hep/hep/projects/6

ROOT/C++ has support for reading/writing data object-wise or member-wise.
ie:
```go
type Data struct {
    A float32
    B int32
    C string
    D []Pair
}

type Pair struct {
    P0, P1 float64
}
```

a collection of user `[]Data` of 3 `Data` items may be written as:
```
[A B C [[P0 P1] [P0 P1] [P0 P1]]] [A B C [[P0 P1] [P0 P1]]] [A B C []]]
```

or, in member-wise mode:
```
[A A A B B B C C C [[P0 P1] [P0 P1] [P0 P1]] [[P0 P1] [P0 P1]] []]]
```

or, in member-wise mode (but "full split"):
```
[A A A B B B C C C [[P0 P0 P0] [P1 P1 P1]] [[P0 P0] [P1 P1]] []]]
```

as of `Go-HEP@v0.30`, `groot` has reasonable support for object-wise r/w of user types, but member-wise support is sub-par.

ROOT uses member-wise mode for:

- `std::map<K,V>` (which is written as a triplet `[len, []K, []V]`)
- `TClonesArray`
- `std::vector<T>`

## `hplot-zoom-pan`: adding support for plot interactivity in `hplot`

- https://github.com/go-hep/hep/projects/7

`hplot` is built on top of [gonum.org/v1/plot](https://pkg.go.dev/gonum.org/v1/plot) to easily create publication quality plots.
`hplot` comes with a couple of HEP-oriented plotters.

at the moment (`Go-HEP@v0.31`), `hplot` only correctly handles "paper" plots: there is no facility to pan/zoom within the plot (even when the plot is displayed w/ the `vggio` or `vgsvg` backends).
there is no interactivity with an `hplot` plot.

it would be great to be able to:

- pan a plot with the mouse (and/or with some `vi`-like keyboard bindings?)
- zoom in/out a plot with the mouse wheel
- zoom in/out on the x- or y-axis with the mouse
- zoom in/out through a box with the mouse

a first iteration could be to implement functions to do this on top of the `gonum/plot/vg/vggio` backend, inside the `pawgo` program.


Let the fun begin!
