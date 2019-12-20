---
title: "Go HEP release 0.22.0"
date: 2019-12-20T12:00:00+01:00
weight: 10
---
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3588414.svg)](https://doi.org/10.5281/zenodo.3588414)

Release [`v0.22.0`](https://github.com/go-hep/hep/tree/v0.22.0) is fresh from the oven.

This release contains a couple of `groot` related features and bugfixes.

## groot

Building on the work on tree-writing from `v0.21.00`, `groot` gained 3 new commands:

- [cmd/arrow2root](https://go-hep.org/x/hep/cmd/arrow2root) converts an [ARROW](https://arrow.apache.org) data file into a corresponding ROOT tree,

```
$> arrow2root testdata/primitives.file.data
$> root-ls -t output.root
=== [./output.root] ===
version: 61804
  TTree      tree         tree    (entries=10)
    bools    "bools/O"    TBranch
    int8s    "int8s/B"    TBranch
    int16s   "int16s/S"   TBranch
    int32s   "int32s/I"   TBranch
    int64s   "int64s/L"   TBranch
    uint8s   "uint8s/b"   TBranch
    uint16s  "uint16s/s"  TBranch
    uint32s  "uint32s/i"  TBranch
    uint64s  "uint64s/l"  TBranch
    float32s "float32s/F" TBranch
    float64s "float64s/D" TBranch
```

- [cmd/npy2root](https://go-hep.org/x/hep/cmd/npy2root) converts a [NumPy data file](https://numpy.org/neps/nep-0001-npy-format.html) into a correspondig ROOT tree,

```
$> npy2root -h
npy2root converts the content of a NumPy data file to a ROOT file and tree.

Usage: npy2root [OPTIONS] input.npy

The NumPy data file format is described here:

 https://docs.scipy.org/doc/numpy/neps/npy-format.html

Example:

 $> npyio-ls input.npy
 ================================================================================
 file: input.npy
 npy-header: Header{Major:1, Minor:0, Descr:{Type:<f8, Fortran:false, Shape:[2 3]}}
 data = [0 1 2 3 4 5]

 $> npy2root -o output.root -t mytree ./input.npy
 $> root-ls -t ./output.root
 === [./output.root] ===
 version: 61804
   TTree   mytree       mytree  (entries=2)
     numpy "numpy[3]/D" TBranch

 $> root-dump ./output.root
 >>> file[./output.root]
 key[000]: mytree;1 "mytree" (TTree)
 [000][numpy]: [0 1 2]
 [001][numpy]: [3 4 5]

Options:
  -o string
    	path to output ROOT file (default "output.root")
  -t string
    	name of the output ROOT tree (default "tree")
```

- [groot/cmd/root-merge](https://go-hep.org/x/hep/groot/cmd/root-merge) is the beginning of an `hadd`-like command, merging (for now) trees and graphs.

```
$> root-merge -h
Usage: root-merge [options] file1.root [file2.root [file3.root [...]]]

ex:
 $> root-merge -o out.root ./testdata/chain.flat.1.root ./testdata/chain.flat.2.root

options:
  -o string
    	path to merged output ROOT file (default "out.root")
  -v	enable verbose mode

$> root-merge ./testdata/chain.flat.*
$> root-ls -t ./out.root 
=== [./out.root] ===
version: 61804
  TTree    tree           my tree title (entries=10)
    B      "B/O"          TBranch
    Str    "Str/C"        TBranch
    I8     "I8/B"         TBranch
    I16    "I16/S"        TBranch
    I32    "I32/I"        TBranch
    I64    "I64/L"        TBranch
    U8     "U8/b"         TBranch
    U16    "U16/s"        TBranch
    U32    "U32/i"        TBranch
    U64    "U64/l"        TBranch
    F32    "F32/F"        TBranch
    F64    "F64/D"        TBranch
    ArrBs  "ArrBs[10]/O"  TBranch
    ArrI8  "ArrI8[10]/B"  TBranch
    ArrI16 "ArrI16[10]/S" TBranch
    ArrI32 "ArrI32[10]/I" TBranch
    ArrI64 "ArrI64[10]/L" TBranch
    ArrU8  "ArrU8[10]/b"  TBranch
    ArrU16 "ArrU16[10]/s" TBranch
    ArrU32 "ArrU32[10]/i" TBranch
    ArrU64 "ArrU64[10]/l" TBranch
    ArrF32 "ArrF32[10]/F" TBranch
    ArrF64 "ArrF64[10]/D" TBranch
    N      "N/I"          TBranch
    SliBs  "SliBs[N]/O"   TBranch
    SliI8  "SliI8[N]/B"   TBranch
    SliI16 "SliI16[N]/S"  TBranch
    SliI32 "SliI32[N]/I"  TBranch
    SliI64 "SliI64[N]/L"  TBranch
    SliU8  "SliU8[N]/b"   TBranch
    SliU16 "SliU16[N]/s"  TBranch
    SliU32 "SliU32[N]/i"  TBranch
    SliU64 "SliU64[N]/l"  TBranch
    SliF32 "SliF32[N]/F"  TBranch
    SliF64 "SliF64[N]/D"  TBranch
```

`groot/rtree` also saw some activity:

- fix for correctly writing `[]int8` and `[]uint8` branch data
- fix for correctly writing empty strings (subtle bug when the empty string was the first or the last to be written)
- fix for correctly writing trees under (deeply) nested directories
- [groot/rtree.Copy](https://godoc.org/go-hep.org/x/hep/groot/rtree#Copy) to easily copy trees from one place to another (possibly across files)
- support for writing leaves with n-dim arrays (`n>=2`)

The `root-print`, `root-cp` and `root-merge` commands have been improved to correctly handle (arbitrarily) nested directories.

`groot` gained support for compressing and decompressing data via the [Zstandard](https://facebook.github.io/zstd/) library.
`ROOT-6.20/00` should be released with this feature as well.
While working on the `rcompress` package (that handles all compression/decompression aspects of `groot`), a bit of optimization work has been carried out, resulting in ~10% optimization in compression speed and a ~40% improvement in memory usage during compression.

## hbook

`hbook` gained 2 new APIs:

- [hbook.H1D#FillN](https://godoc.org/go-hep.org/x/hep/hbook#H1D.FillN) that fills a 1D-histogram with a slice of data and its associated weights
- [hbook.H12D#FillN](https://godoc.org/go-hep.org/x/hep/hbook#H2D.FillN) that fills a 2D-histogram with a slice of data and its associated weights

## FUSE

Support for mounting ROOT files and remote XRootD location thru FUSE has been removed from `go-hep.org/x/hep`.
FUSE support on macOS is not open-source anymore and the provided feature thru Go-HEP was a bit flaky (at least, the tests were flaky).
The `xrdfuse` package, `xrd-fuse` and `root-fuse` commands have been migrated to [go-hep.org/x/exp](https://go-hep.org/x/exp).

We might consider using the 9p protocol instead of FUSE.

## AOB

See you in 2020!

## ChangeLog

* [8a851ee](/commit/8a851ee) all: update URL of NumPy array data file specification
* [934bdd7](/commit/934bdd7) groot/internal/rcompress: optimize compression
* [5b1a2dd](/commit/5b1a2dd) groot/internal/rcompress: fix roundtrip compression
* [bd257a1](/commit/bd257a1) groot/internal/rcompress: add compress benchmark
* [5801a56](/commit/5801a56) hep: update go.mod
* [08b1f75](/commit/08b1f75) groot/riofs: handle riofs.Dir in riofs.FileOf
* [c5200f1](/commit/c5200f1) groot/cmd/root-merge: first import
* [590a5c3](/commit/590a5c3) groot/rtree: make Writer implement root.Merger
* [dfecc37](/commit/dfecc37) groot/rhist: implement root.Merger interface for Graph{,Errors}
* [28fe971](/commit/28fe971) groot/root: introduce Merger interface
* [1437ad2](/commit/1437ad2) groot,xrootd: move FUSE related code to go-hep.org/x/exp
* [d45349f](/commit/d45349f) groot/internal/rcompress: streamline error strings
* [1005829](/commit/1005829) groot/internal/rcompress: use rtests.RunCxxROOT to run ROOT macros
* [54eedba](/commit/54eedba) groot/internal/rcompress: add support for flate.Best{Compression,Speed} to Zstd
* [7c7bdf3](/commit/7c7bdf3) groot/{riofs,internal/rcompress}: add support for zstd (de)compressor
* [1d81a65](/commit/1d81a65) groot/cmd/root-print: add test for '^dir'
* [4f47f41](/commit/4f47f41) groot/cmd/root-print: support nested-dirs
* [9db9056](/commit/9db9056) groot/cmd/root-cp: add support for nested-directories
* [5ee79ea](/commit/5ee79ea) groot/cmd/root-cp: add support for copying trees
* [11df8e6](/commit/11df8e6) cmd/npy2root: add support for n-dim arrays
* [8b51e1e](/commit/8b51e1e) groot/rtree: add support for leaves with n-dim arrays (n>=2)
* [049273a](/commit/049273a) cmd/npy2root: first import
* [45e282c](/commit/45e282c) hbook: add H2D.FillN(xs, ys, ws)
* [4811e2c](/commit/4811e2c) hbook: add H1D.FillN(xs, ws)
* [241c5af](/commit/241c5af) cmd/arrow2root,groot/rarrow: introduce rarrow.NewFlatTreeWriter
* [64dd534](/commit/64dd534) cmd/arrow2root: first import
* [6ecf64a](/commit/6ecf64a) groot/rtree: introduce Copy and CopyN for copying trees
* [c8dbaf0](/commit/c8dbaf0) groot: introduce groot/internal/rcmd
* [5c5a8e9](/commit/5c5a8e9) groot/rtree: test writing []int8 and [N]int8
* [2f7935b](/commit/2f7935b) groot/rtree: support writing trees w/ empty strings
* [3a56e78](/commit/3a56e78) groot/rbytes: make WriteString correctly handle empty strings as first value
* [ebe3617](/commit/ebe3617) groot/{riofs,rtree}: add support for writing n-tuples under nested directories

