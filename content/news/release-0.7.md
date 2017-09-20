---
title: "Go HEP release 0.7"
date: 2017-09-20T14:12:20+02:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.897927.svg)](https://doi.org/10.5281/zenodo.897927)

Release [`v0.7`](https://github.com/go-hep/hep/tree/v0.7) is fresh from the oven with a few breaking API changes.

## cmd/yoda

The `cmd/rio2yoda`, `cmd/root2yoda` and `cmd/yoda2rio` commands now support gzipped [YODA](https://yoda.hepforge.org) files.

The `cmd/yoda2rio` is also more lenient when dealing with `YODA` objects that are not yet supported on the `hbook` side (`Counter`, `S1D`, ...) (thanks [Lukas Heinrich](https://github.com/lukasheinrich) for the report.)

## fastjet

[Bastian](https://github.com/Bastiantheone) added the ability to plot the Voronoi diagram from a set of 2D-points.

## hbook

`hbook/yodacnv` was modified to support the more lenient approach with regard to unsupported (yet!) `YODA` objects.

## hplot

`hplot` has seen the most user-facing work:

- a new default style that is more in-line with current aesthetic standards (_ie:_ `matplotlib`-like)
- `hplot.New` uses this new default style and thus is now able to ensure it won't fail loading the fonts
- `hplot.New` thus only returns `*hplot.Plot`, without an `error` value
- `hplot.NewH1D` is also able to ensure no fonts-loading error will araise and thus only return a `*hplot.H1D`

So, where you were doing:

```go
p, err := hplot.New()
if err != nil {
	log.Fatal(err)
}

h, err := hplot.NewH1D(h1d)
if err != nil {
	log.Fatal(err)
}
```

you can now just write:

```go
p := hplot.New()
h := hplot.NewH1D(h1d)
```

And here are examples of the new default style:

![correlations](/images/news/release-0.7/go-correlations.png)
![results](/images/news/release-0.7/go-results.png)

## pawgo

`pawgo` has been slightly updated to support plotting 2-dim histograms.
`pawgo` was also fixed to correctly handle `YODA` files converted to `RIO` (thanks [Lukas Heinrich](https://github.com/lukasheinrich) for the report.)


## rootio

`rootio` gained 2 new commands:

- `rootio/cmd/root-diff`: a command to print the differences between 2 ROOT files, including the content of their `TTrees`,
- `rootio/cmd/root-print`: a command to print histograms contained in ROOT files into PDF, PNG, ... files.

`rootio` now also defines and exports the `rootio.H1` and `rootio.H2` interfaces that are implemented by the `TH1x` and `TH2x` (respectively) concrete types.
