---
title: "Go HEP release 0.10"
date: 2018-02-15T19:22:57+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1173718.svg)](https://doi.org/10.5281/zenodo.1173718)

Release [`v0.10`](https://github.com/go-hep/hep/tree/v0.10) is fresh from the oven.

This release brings quite a few improvements in the `rootio` area:

- support for `ROOT` files compressed with LZMA
- support for `ROOT` files compressed with LZ4
- support scanning `ROOT` trees with branches holding multiple leaves
- support for `ROOT` trees holding `bool`, `[X]bool` and `std::vector<bool>`
- support for `ROOT` files holding `TBranch` version 10
- support for `ROOT` files holding `TBranchElement` version 8
- support for `ROOT` files holding `TH1` version 5
- support for `ROOT` trees and branches produced with `ROOT-6.12`

The `rootio/cmd/root-srv` command dropped support for Go-1.6 and earlier.
Please upgrade to a newer Go version.

In other news, the `csvutil/csvdriver` and `hbook/ntup/ntcsv` gained the ability to read remote CSV files (over HTTP.)

## AOB

Better support for `TStreamerElements` and thus the ability to read nested `std::vector<std::vector<T>>` data or the ability to run `root-dump` and `root-diff` on all `ROOT` files, has started.

Google Summer of Code has started and Go-HEP is part of it, thanks to the CERN-HSF umbrella.
Check out the 2 proposals:

- [Go-HEP/xrootd](http://hepsoftwarefoundation.org/gsoc/2018/proposal_GoHEPxrootd.html)
- [Go-HEP/FPGA](http://hepsoftwarefoundation.org/gsoc/2018/proposal_GoHEPreconfigure.html)

Stay tuned.
