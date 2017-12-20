---
title: "Go HEP release 0.8"
date: 2017-12-20T15:11:09+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1120266.svg)](https://doi.org/10.5281/zenodo.1120266)

Release [`v0.8`](https://github.com/go-hep/hep/tree/v0.8) is fresh from the oven.

No API changes in `go-hep` itself, just a few modifications to cater with changes in `gonum/v1/gonum` and `gonum/v1/plot`:

- `gonum/v1/gonum/stat/distuv` distributions replaced the name of their `rand.Rand` source field from `Source` to `Src` (to align with the standard library)
- `gonum/v1/gonum/stat/...` packages now use `golang.org/x/exp/rand.Rand` instead of `math/rand.Rand`

The biggest news (and where the work mostly happened between `v0.7` and `v0.8`) is that there's now a [Jupyter](https://jupyter.org) instance with a Go kernel installed with Go-HEP and Gonum libraries pre-installed.
This has been packaged up at [go-hep/binder](https://github.com/go-hep/binder).
You can try it there:
[![Binder](https://mybinder.org/badge.svg)](https://mybinder.org/v2/gh/go-hep/binder/master)

Here are some examples:

- [01-display-data](https://mybinder.org/v2/gh/go-hep/binder/master?filepath=examples%2F01-display-data.ipynb)
- [02-gonum-stat](https://mybinder.org/v2/gh/go-hep/binder/master?filepath=examples%2F02-gonum-stat.ipynb)
- [03-go-hep-plot](https://mybinder.org/v2/gh/go-hep/binder/master?filepath=examples%2F03-go-hep-hplot.ipynb)

## fmom

`fmom` gained a new top-level function `fmom.InvMass(p1, p2 P4) float64` : 
```go
// InvMass computes the invariant mass of two incoming 4-vectors p1 and p2.
func InvMass(p1, p2 P4) float64 {
       p := Add(p1, p2)
       return p.M()
}
```

## hplot

`hplot` gained a new top-level function `Show`:

```go
// Show displays the plot according to format, returning the raw bytes and
// an error, if any.
//
// If format is the empty string, then "png" is selected.
// The list of accepted format strings is the same one than from
// the gonum.org/v1/plot/vg/draw.NewFormattedCanvas function.
func Show(p *Plot, w, h vg.Length, format string) ([]byte, error) { ... }
```

`Show` is especially useful for the new support for Go in Jupyter notebooks (via the [Neugram](https://neugram.io) interpreter)

See https://github.com/go-hep/binder for more details and examples on using Go and Go-HEP in Jupyter.

[David Blyth](https://github.com/decibelcooper) provided a new `hplot.HInfoStdDev` `HInfoStyle` value and made them bitset-like.
Thanks David!

## rootio

- `cmd/root-dump` can now recursively handle `rootio.Directory` and thus walk an entire ROOT file.
- `cmd/root-ls` got the same handling of `rootio.Directory`.

