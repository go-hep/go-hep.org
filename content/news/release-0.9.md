---
title: "Go HEP release 0.9"
date: 2018-01-24T17:26:30+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1158687.svg)](https://doi.org/10.5281/zenodo.1158687)

Release [`v0.9`](https://github.com/go-hep/hep/tree/v0.9) is fresh from the oven.

This is a quick and simple release fix with the following main features:

- fix a broken link to the PDF LaTeX example (back from the multi- to mono-repo migration)
- fix a compilation failure in `rootio/cmd/root-srv` originating from an API modification in a 3rd-party dependency (`github.com/satori/go.uuid`)

One interesting new feature is the introduction of a simple `ROOT::TTree` to `CSV` file conversion command: `root2csv`.

```
$> go get go-hep.org/x/hep/cmd/root2csv
$> root2csv --help
root2csv -h
Usage of root2csv:
  -f string
    	path to input ROOT file name
  -o string
    	path to output CSV file name (default "output.csv")
  -t string
    	name of the tree to convert (default "tree")

$> go doc go-hep.org/x/hep/cmd/root2csv 
root2csv converts the content of a ROOT TTree to a CSV file.

    Usage of root2csv:
      -f string
        	path to input ROOT file name
      -o string
        	path to output CSV file name (default "output.csv")
      -t string
        	name of the tree to convert (default "tree")

By default, root2csv will write out a CSV file with ';' as a column
delimiter. root2csv ignores the branches of the TTree that are not supported
by CSV:

    - slices/arrays
    - C++ objects

Example:

    $> root-ls -t ./testdata/small-flat-tree.root
    === [./testdata/small-flat-tree.root] ===
    version: 60806
    TTree          tree                 my tree title (entries=100)
      Int32        "Int32/I"            TBranch
      Int64        "Int64/L"            TBranch
      UInt32       "UInt32/i"           TBranch
      UInt64       "UInt64/l"           TBranch
      Float32      "Float32/F"          TBranch
      Float64      "Float64/D"          TBranch
      Str          "Str/C"              TBranch
      ArrayInt32   "ArrayInt32[10]/I"   TBranch
      ArrayInt64   "ArrayInt64[10]/L"   TBranch
      ArrayUInt32  "ArrayInt32[10]/i"   TBranch
      ArrayUInt64  "ArrayInt64[10]/l"   TBranch
      ArrayFloat32 "ArrayFloat32[10]/F" TBranch
      ArrayFloat64 "ArrayFloat64[10]/D" TBranch
      N            "N/I"                TBranch
      SliceInt32   "SliceInt32[N]/I"    TBranch
      SliceInt64   "SliceInt64[N]/L"    TBranch
      SliceUInt32  "SliceInt32[N]/i"    TBranch
      SliceUInt64  "SliceInt64[N]/l"    TBranch
      SliceFloat32 "SliceFloat32[N]/F"  TBranch
      SliceFloat64 "SliceFloat64[N]/D"  TBranch

    $> root2csv -o out.csv -t tree -f testdata/small-flat-tree.root
    root2csv: scanning leaves...
    root2csv: >>> "ArrayInt32" [10]int32 not supported
    root2csv: >>> "ArrayInt64" [10]int64 not supported
    root2csv: >>> "ArrayInt32" [10]int32 not supported
    root2csv: >>> "ArrayInt64" [10]int64 not supported
    root2csv: >>> "ArrayFloat32" [10]float32 not supported
    root2csv: >>> "ArrayFloat64" [10]float64 not supported
    root2csv: >>> "SliceInt32" []int32 not supported
    root2csv: >>> "SliceInt64" []int64 not supported
    root2csv: >>> "SliceInt32" []int32 not supported
    root2csv: >>> "SliceInt64" []int64 not supported
    root2csv: >>> "SliceFloat32" []float32 not supported
    root2csv: >>> "SliceFloat64" []float64 not supported
    root2csv: scanning leaves... [done]

    $> head out.csv
    ## Automatically generated from "testdata/small-flat-tree.root"
    Int32;Int64;UInt32;UInt64;Float32;Float64;Str;N
    0;0;0;0;0;0;evt-000;0
    1;1;1;1;1;1;evt-001;1
    2;2;2;2;2;2;evt-002;2
    3;3;3;3;3;3;evt-003;3
    4;4;4;4;4;4;evt-004;4
    5;5;5;5;5;5;evt-005;5
    6;6;6;6;6;6;evt-006;6
    7;7;7;7;7;7;evt-007;7
```

## AOB

The next release cycle will see improvements in the `rootio` department (read *and* write.)

Stay tuned.
