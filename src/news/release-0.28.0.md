---
title: "Go HEP release 0.28.0"
date: 2020-10-19T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.4106583.svg)](https://doi.org/10.5281/zenodo.4106583)

Release [`v0.28.0`](https://github.com/go-hep/hep/tree/v0.28.0) is out of the oven.

This release contains a major new `groot` related feature: the ability to write non-split user types, by way of `TBranchElement` and `TLeafElement`.

- [write-full-split type](https://pkg.go.dev/go-hep.org/x/hep@v0.28.0/groot/rtree#example-package-CreateEventNtupleFullSplit)
- [write-non-split type](https://pkg.go.dev/go-hep.org/x/hep@v0.28.0/groot/rtree#example-package-CreateEventNtupleNoSplit)

## csvutil

- add support for writing slices (thanks [Chinmaya Krishnan Mahesh `(chin123)`](https://github.com/chin123)!)

## groot

- add support for writing large files (_i.e.:_ with size > 2Gb)
- changed the way `groot` streamers are created and handled (read/write) from `ROOT` ones
- removed `rtree.Scanner` and `rtree.TreeScanner` types, in favor of `rtree.Reader`
- add support for reading `std::bitset<N>`
- add initial support for reading `std::map<K,V>`
- introduce `rtree/rfunc.Formula` to easily create and use "user-based" formulae (thanks [Romain Madar `(rmadar)`](https://github.com/rmadar) for the improvements!)

## hplot

- use `gonum/plot/vg`'s [Gio](https://gioui.org/) backend for `pawgo`
- various fixes and improvements to `hplot` plotters (thanks [Romain Madar `(rmadar)`](https://github.com/rmadar)!)

## AOB

That's all for today.
Next cycle will probably see some work on the structured tree writing (mostly consolidation and adding missing features) and perhaps some performance improvements on the writing side of things.

## ChangeLog

* [653c8f36](/commit/653c8f36) all: update Go-mod
* [fe7367f2](/commit/fe7367f2) groot/rtree: test any, arr-of(any), vec(any) and vec(vec(any))
* [0db30afb](/commit/0db30afb) groot/rtree: test c-var-len-arrays
* [de9193a3](/commit/de9193a3) groot/rtree: implement writing of structured-event ntuples
* [33c9bd26](/commit/33c9bd26) groot/rdict: streamline StreamerOf, add support for slices
* [2e69fce9](/commit/2e69fce9) groot/rtree: add support for no-split to WriteVarsFromStruct
* [43b8c835](/commit/43b8c835) groot/rmeta: add stdint to CxxBuiltins
* [60cf05d2](/commit/60cf05d2) groot/rdict: add StreamerLoop visitor
* [2bfd292b](/commit/2bfd292b) groot/{cmd/root-ls,rcmd}: refactor root-ls as a rcmd.List
* [c5e9a68b](/commit/c5e9a68b) groot/{,rarrow,rcmd,riofs,rtree}: properly handle slices of n-dim arrays
* [ee067e76](/commit/ee067e76) groot/{rarrow,rtree}: reduce Leaf API
* [cdbafa1f](/commit/cdbafa1f) groot/{,rarrow,rtree}: better leaf-kind/shape/type handling
* [ae4a840d](/commit/ae4a840d) groot/rtree: simplify Tree/Branch/Leaf interfaces
* [16ed6fb7](/commit/16ed6fb7) cmd,groot: remove rtree.{Tree,}Scanner
* [dba77171](/commit/dba77171) groot/rsql/rsqldrv: use rtree.Reader in lieu of Scanner
* [d6ee5e87](/commit/d6ee5e87) groot/rsrv: use rtree.Reader in lieu of Scanner
* [eed0f08c](/commit/eed0f08c) groot/rtree: use Reader in lieu of Scanner
* [a3988c10](/commit/a3988c10) groot/cmd/root-gen-datareader: use Reader in lieu of Scanner
* [1c0b336c](/commit/1c0b336c) groot/rcmd: use rtree.Reader in lieu of rtree.Scanner
* [cf246da1](/commit/cf246da1) groot/rtree: fix rchain entry offsets computation
* [4e85dd24](/commit/4e85dd24) groot/rarrow: use rtree.Reader instead of rtree.Scanner
* [6a063617](/commit/6a063617) cmd/root2npy: handle ndim arrays
* [0be35e71](/commit/0be35e71) groot/{rarrow,rcmd,rdict,riofs}: properly handle ndim-arrays
* [4d49a780](/commit/4d49a780) groot/rtree: use flattenArrayType in rleaf
* [30b9ba32](/commit/30b9ba32) groot/rtree: streamline use of rtree.ReadVarsFromStruct in scanners
* [c0d6fad2](/commit/c0d6fad2) groot: introduce new x-flat-tree file, testing more ROOT builtins
* [20003c6f](/commit/20003c6f) groot/rdict: std-bitset cosmetics
* [06d30c23](/commit/06d30c23) groot/{rcmd,rdict,riofs,rtree}: handle map<K,V>
* [bd316477](/commit/bd316477) groot/rdict: fix std::string streamer usage, add preliminary support for std::map-member-wise streamer
* [51ce0f59](/commit/51ce0f59) groot/{rdict,riofs}: add support for std::map<K,V> with ROOT-Cling dict
* [2381e6ae](/commit/2381e6ae) groot/{rbytes,rcmd,rdict,riofs,rtree}: add support for bitset
* [400e2b6f](/commit/400e2b6f) groot/internal/rtests: add support for rootcling dictionary generation
* [7536fc78](/commit/7536fc78) groot/{rcmd,rtree}: handle Geant4 recovery baskets
* [f4d329eb](/commit/f4d329eb) groot/rdict: add initial test for ObjectFrom
* [43282104](/commit/43282104) groot/rdict: impl r/w rmeta.AnyP, rmeta.Anyp
* [cabbd2cc](/commit/cabbd2cc) groot/rdict: test r/w-streamer elem round-trip
* [012823d7](/commit/012823d7) groot/rdict: streamline rmeta.ULong64+rmeta.Long64 handling
* [e675e8f2](/commit/e675e8f2) groot/rdict: support finding counters across multi-hop inheritance hierarchies
* [664ec75b](/commit/664ec75b) groot/rdict: streamline type generation, support n-dim arrays
* [631f18d4](/commit/631f18d4) groot/{rdict,cmd/root-gen-streamer}: handle array dims via StreamerElement.ArrayDims
* [edce8855](/commit/edce8855) groot/{rbytes,rdict}: add ArrayDims() []int to StreamerElement
* [d5b19154](/commit/d5b19154) groot/{rcmd,rdict,rtree}: introduce R/W-streamers
* [19855690](/commit/19855690) groot/rbytes: extend StreamerInfo interface, introduce Binder and Counter interfaces
* [f50205ac](/commit/f50205ac) groot/{rbytes,rdict}: add Encoder/Decoder interfaces for ObjectWise/MemberWise kinds
* [ee55be2b](/commit/ee55be2b) groot/{rbytes,rdict}: add BuildStreamers API
* [8bb70b75](/commit/8bb70b75) groot/rtree: split r/w test of []intXX (extract and skip []int8)
* [711290aa](/commit/711290aa) groot/rtree: workaround for std::vector<TString>
* [b6db3d51](/commit/b6db3d51) groot/rdict: add test for parsing var-len D32 array
* [3e8885f3](/commit/3e8885f3) groot/{rdict,rvers}: add streamer-info for THashTable
* [c23f8cc9](/commit/c23f8cc9) groot/rcont: implement TRefTable
* [a882ddeb](/commit/a882ddeb) groot/rbytes: export r/w std::vector<std::string>
* [4672fa25](/commit/4672fa25) groot/rmeta: add TypeName2Enum
* [a4fbb9e4](/commit/a4fbb9e4) groot/rmeta: add STLstdstring to ESTLType enum
* [1dd68ee6](/commit/1dd68ee6) groot: add C++ streamer for TString
* [78f2a495](/commit/78f2a495) groot/rbytes: remove test of (internal detail) cap-size in ResizeXYZ
* [b3f3f622](/commit/b3f3f622) hep: bump github.com/klauspost/compress@v1.10.11
* [b9646023](/commit/b9646023) all: migrate to Gonum-v0.8
* [a39f822a](/commit/a39f822a) groot: bump to ROOT-6.22/02
* [5c4a32ea](/commit/5c4a32ea) groot/riofs: add support for writing big files
* [79d17b1e](/commit/79d17b1e) ci: bump to Go-1.15.x
* [73346da2](/commit/73346da2) xroot: update tests for new ccxrootd in2p3 config
* [d55d1b51](/commit/d55d1b51) groot/{rcont,rhist,riofs}: use rtests macro to emulate C++ rootls
* [4b4fa012](/commit/4b4fa012) groot: bump to ROOT-6.22/00
* [1a9e6f77](/commit/1a9e6f77) hplot: correctly handle negative bin content in log scale
* [7d2c4375](/commit/7d2c4375) hep: bump to gokrb5@v8.4.1
* [3f458e3a](/commit/3f458e3a) groot/riofs: handle files with no embedded streamer for std::string
* [98270f3f](/commit/98270f3f) groot/rtree: introduce Reader.Reset
* [abe8a2be](/commit/abe8a2be) groot/{rcmd,rtree}: handle reading of vector<Char_t>
* [095a9bc6](/commit/095a9bc6) ci: add apt-get update
* [0e658332](/commit/0e658332) hep: add official pkg.go.dev badge
* [f4c3901f](/commit/f4c3901f) xrootd/xrdio: implement correct URI parser
* [63dcd821](/commit/63dcd821) ci: bump golangci-lint@v1.28
* [47aa4362](/commit/47aa4362) pawgo: test plot H2D
* [d898ce9f](/commit/d898ce9f) hplot/cmd/iplot,pawgo: handle multiple plot-windows, blocking app.Main
* [e39ab99f](/commit/e39ab99f) all: bump klauspost/compress@v1.10.10
* [0f5496cf](/commit/0f5496cf) csvutil: add support for writing slices
* [b0b8bd0b](/commit/b0b8bd0b) groot: bump to go-mmap/mmap@v0.4.0
* [234a493a](/commit/234a493a) xrootd: do not segfault sending data with nil clients
* [aa37c691](/commit/aa37c691) hplot/cmd/iplot,pawgo: use vg-gio + Gio
* [ebbd2a74](/commit/ebbd2a74) groot/riofs: use go-mmap@v0.2.0
* [04360d78](/commit/04360d78) hbook,hplot: introduce hbook.Count, use in hplot.BinnedErrBand
* [521c323f](/commit/521c323f) groot/rdict: handle more range-specs from streamers for D32/F16
* [dc20b078](/commit/dc20b078) groot/riofs: use go-mmap/mmap
* [0f13550f](/commit/0f13550f) xrootd{,/xrdio}: do not segfault closing nil clients or nil files
* [1e01356f](/commit/1e01356f) all: bump to gokrb5@v8.4.0
* [45f0797a](/commit/45f0797a) groot/rtree: add support for rmeta.Long64 rstreamer
* [122ca822](/commit/122ca822) ci: add a timeout=10m for golangci-lint
* [5616b92d](/commit/5616b92d) ci: add back golangci-lint
* [24781f7c](/commit/24781f7c) hplot: properly handle hstacks with bands and log-y
* [4cf183c8](/commit/4cf183c8) groot/internal/rcompress: streamline handling of uncompressible data blocks
* [364dcc3f](/commit/364dcc3f) hep: update gonum/{gonum,plot}
* [3050bc0c](/commit/3050bc0c) groot/cmd/root-ls: clip branch name and title at 60 and 50 chars (resp.) for better display
* [69f0c8d5](/commit/69f0c8d5) groot/{riofs,rtree}: fix and test concurrent writing of trees in different ROOT files
* [ec2e5b95](/commit/ec2e5b95) ci: (temporarily) disable golangci-lint action
* [fcea6526](/commit/fcea6526) {cmd/root2arrow,groot/rarrow}: handle leaves with embedded std::vector<T>
* [e2ffbc7c](/commit/e2ffbc7c) cmd/root2npy: use rtree.Reader
* [6c0e8ee8](/commit/6c0e8ee8) groot/{rtree,rcmd}: add support for recovered baskets
* [4e246266](/commit/4e246266) ci: make travis ci less verbose
* [3b702ec4](/commit/3b702ec4) groot: implement new rfunc-gen default type name convention
* [384b9ef3](/commit/384b9ef3) groot/rdict: add generation of streamer info checksum
* [4d1e5f08](/commit/4d1e5f08) groot/rtree: fix handling of rleaf-element with a "This" StreamerSTL
* [5e981213](/commit/5e981213) groot/rtree: add formula test for slices
* [351d7f17](/commit/351d7f17) groot: add support for extracting copyright year from code templates
* [c771d0c9](/commit/c771d0c9) groot/cmd/root-gen-rfunc: first import
* [e1d8871a](/commit/e1d8871a) hplot: fix HStack band display logic
* [a8592066](/commit/a8592066) ci: remove lint scheduled/cron
* [508d731b](/commit/508d731b) groot/rtree: lift WriteVar out of Writer
* [9448e168](/commit/9448e168) groot/{rcmd,rtree}: make Copy take a Reader instead of Tree
* [1b6e3cf3](/commit/1b6e3cf3) groot/rtree{,/rfunc}: introduce rfunc.Formula protocol
* [0334332a](/commit/0334332a) hep: add pkg.go.dev badge
* [488e761e](/commit/488e761e) groot/{riofs,rtree}: add Join of trees, associated reader and test files
* [3d790e70](/commit/3d790e70) groot/rtree: simplify Tree interface
* [32ae501f](/commit/32ae501f) groot/rtree: fix typo in Chain documentation, rename tchain into chain
* [f94677ef](/commit/f94677ef) ci: extract lint-check into its own workflow
* [50964db8](/commit/50964db8) ci: fix action/cache restore key syntax
* [56b78f79](/commit/56b78f79) csvutil/csvdriver,hbook/ntup/ntcsv: use ql@v1.1.0 with driver.ConnBeginTx implementation
* [56834a47](/commit/56834a47) all: fix lint
* [98694bc7](/commit/98694bc7) ci: weave in golangci-lint
* [96b87a43](/commit/96b87a43) csvutil/csvdriver: staticcheck fixes
* [54e0051a](/commit/54e0051a) csvutil/csvdriver: implement new database/sql/driver.XyzContext
* [47234700](/commit/47234700) groot/rsql/rsqldrv: apply staticcheck fixes
* [889375c3](/commit/889375c3) groot/rsql/rsqldrv: migrate to new database/sql/driver.XyzContext interfaces

