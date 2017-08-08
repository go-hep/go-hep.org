+++
title = "Go HEP release 0.4"
date = "2017-08-08T17:32:55+02:00"
weight = 10
+++

Release `v0.4` is a small iteration (because of holidays.)

But still, it does bring interesting things to the table:

- [root2npy](https://go-hep.org/x/hep/cmd/root2npy), a new command to convert the content of simple `TTree`s into a [NumPy data file](https://docs.scipy.org/doc/numpy/neps/npy-format.html),
- migration from the `github.com/gonum/xyz` packages to their new location -- `gonum.org/v1/gonum/xyz`,

## `fastjet`

- preliminary work for implementing better/faster jet clustering, thanks [Bastian Wieck](https://github.com/Bastiantheone)

## `hplot`

- make the `hplot.H1D` type implement `gonum/plot.Thumbnailer` (for better integration with legends and labels)

## `lcio`

- fix a bug in the marshaling of `lcio.TrackerRawData`, thanks [David Blyth](https://github.com/decibelCooper)


