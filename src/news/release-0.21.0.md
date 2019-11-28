---
title: "Go HEP release 0.21.0"
date: 2019-11-28T16:00:00+01:00
weight: 10
---
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3556711.svg)](https://doi.org/10.5281/zenodo.3556711)

Release [`v0.21.0`](https://github.com/go-hep/hep/tree/v0.21.0) is fresh from the oven.

This release contains a major new `groot` feature: the ability to *write* (simple, flat) trees, including variable-length arrays (_a.k.a_ slices):

- [groot/rtree--CreateFlatNtuple](https://godoc.org/go-hep.org/x/hep/groot/rtree#example-package--CreateFlatNtuple)
- [groot/rtree--CreateFlatFromStruct](https://godoc.org/go-hep.org/x/hep/groot/rtree#example-package--CreateFlatNtupleFromStruct)
- [groot/rtree--CreateFlatNtupleWithLZMA](https://godoc.org/go-hep.org/x/hep/groot/rtree#example-package--CreateFlatNtupleWithLZMA)

## groot

- `groot` supports more ROOT-4 files (as created by Geant4: `TH{1,2}x` and `TTree`)
- fixed compilation on 32b systems
- add [groot/rtree.WriterVarsFromStruct](https://godoc.org/go-hep.org/x/hep/groot/rtree#WriteVarsFromStruct) to generate a slice of `rtree.WriterVars` from a user-provided `struct` that can be then used to fill a tree
- add `With{LZ4,LZMA,Zlib}` and `WithoutCompression` function to configure whether a tree should use compression (and what kind of compression, if any)
- add `WithBasketSize` to configure the basket size of trees/branches
- add auto-flushing of branches' baskets

[godoc:example](https://godoc.org/go-hep.org/x/hep/groot/rtree#example-package--CreateFlatNtupleFromStruct)
```go
type Data struct {
    I32    int32
    F64    float64
    Str    string
    ArrF64 [5]float64
    N      int32
    SliF64 []float64 `groot:"SliF64[N]"` // tell ROOT/C++ the leaf name and the leaf holding the count
}

const (
    fname = "struct.root"
    nevts = 5
)

f, err := groot.Create(fname)
if err != nil {
    log.Fatalf("%+v", err)
}
defer f.Close()

var evt Data

tree, err := rtree.NewWriter(f, "mytree", rtree.WriteVarsFromStruct(&evt))
if err != nil {
    log.Fatalf("could not create tree writer: %+v", err)
}

fmt.Printf("-- created tree %q:\n", tree.Name())
for i, b := range tree.Branches() {
    fmt.Printf("branch[%d]: name=%q, title=%q\n", i, b.Name(), b.Title())
}

for i := 0; i < nevts; i++ {
    evt.I32 = int32(i)
    evt.F64 = float64(i)
    evt.Str = fmt.Sprintf("evt-%0d", i)
    evt.ArrF64 = [5]float64{float64(i), float64(i + 1), float64(i + 2), float64(i + 3), float64(i + 4)}
    evt.N = int32(i)
    evt.SliF64 = []float64{float64(i), float64(i + 1), float64(i + 2), float64(i + 3), float64(i + 4)}[:i]
    _, err = tree.Write()
    if err != nil {
        log.Fatalf("could not write event %d: %+v", i, err)
    }
}
fmt.Printf("-- filled tree with %d entries\n", tree.Entries())

err = tree.Close()
if err != nil {
    log.Fatalf("could not write tree: %+v", err)
}

err = f.Close()
if err != nil {
    log.Fatalf("could not close tree: %+v", err)
}

// Output:
// -- created tree "mytree":
// branch[0]: name="I32", title="I32/I"
// branch[1]: name="F64", title="F64/D"
// branch[2]: name="Str", title="Str/C"
// branch[3]: name="ArrF64", title="ArrF64[5]/D"
// branch[4]: name="N", title="N/I"
// branch[5]: name="SliF64", title="SliF64[N]/D"
// -- filled tree with 5 entries
// -- read back ROOT file
// entry[0]: {I32:0 F64:0 Str:evt-0 ArrF64:[0 1 2 3 4] N:0 SliF64:[]}
// entry[1]: {I32:1 F64:1 Str:evt-1 ArrF64:[1 2 3 4 5] N:1 SliF64:[1]}
// entry[2]: {I32:2 F64:2 Str:evt-2 ArrF64:[2 3 4 5 6] N:2 SliF64:[2 3]}
// entry[3]: {I32:3 F64:3 Str:evt-3 ArrF64:[3 4 5 6 7] N:3 SliF64:[3 4 5]}
// entry[4]: {I32:4 F64:4 Str:evt-4 ArrF64:[4 5 6 7 8] N:4 SliF64:[4 5 6 7]}
```

## hplot

- [hplot.Function](https://godoc.org/go-hep.org/x/hep/hplot#Function) has been copied from [Gonum/plot](https://godoc.org/gonum.org/v1/plot/plotter#Function) with extra support for log-y axes (automatically discarding intervals where the returned value is invalid on log-y axes)

[godoc:example](https://godoc.org/go-hep.org/x/hep/hplot#example-Function--LogY)
```go
quad := hplot.NewFunction(func(x float64) float64 { return x * x })
quad.Color = color.RGBA{B: 255, A: 255}

fun := hplot.NewFunction(func(x float64) float64 {
    switch {
    case x < 6:
        return 20
    case 6 <= x && x < 7:
        return 0
    case 7 <= x && x < 7.5:
        return 30
    case 7.5 <= x && x < 9:
        return 0
    case 9 <= x:
        return 40
    }
    return 50
})
fun.LogY = true
fun.Color = color.RGBA{R: 255, A: 255}

p := hplot.New()
p.Title.Text = "Functions - Log-Y scale"
p.X.Label.Text = "X"
p.Y.Label.Text = "Y"

p.Y.Scale = plot.LogScale{}
p.Y.Tick.Marker = plot.LogTicks{}

p.Add(fun)
p.Add(quad)
p.Add(hplot.NewGrid())
p.Legend.Add("x^2", quad)
p.Legend.Add("fct", fun)
p.Legend.ThumbnailWidth = 0.5 * vg.Inch

p.X.Min = 5
p.X.Max = 10
p.Y.Min = 10
p.Y.Max = 100

err := p.Save(200, 200, "testdata/functions_logy.png")
if err != nil {
    log.Panic(err)
}
```
![functions-logy-example](https://github.com/go-hep/hep/raw/master/hplot/testdata/functions_logy_golden.png)

## clean-up

The whole GoHEP tree has been cleaned up to remove the use of [github.com/pkg/errors](https://godoc.org/github.com/pkg/errors) and use instead [x/xerrors](https//godoc.org/golang.org/x/xerrors).
`x/xerrors` will be phased out (in favor of `fmt` from the stdlib) when Go-1.13 will be the oldest supported release.

This is to gain nice error reports as documented in the [Go blog](https://blog.golang.org):

- https://blog.golang.org/go1.13-errors

## AOB

The GoHEP website has also gained a new [/dist](/dist) section, where binaries for Darwin, Freebsd, Linux and Windows are uploaded (for every new tagged release).

We'll try to release `v0.22.0` with the ability to write structs and nested trees.

Stay tuned! (and, as always, any kind of help (reviews, patches, nice emails, constructive criticism) deeply appreciated.)

## ChangeLog

* [a3bdc48](/commit/a3bdc48) groot/rhist: add test for ROOT-4 H1D
* [6c8c87a](/commit/6c8c87a) groot: add support for ROOT-4 histograms
* [c5fa1b7](/commit/c5fa1b7) groot/rtree: support G4-like tree
* [508bea9](/commit/508bea9) groot/rtree: TBranch-v6
* [2afe0e7](/commit/2afe0e7) groot/rtree: TBranchElement-v1
* [96ca4b3](/commit/96ca4b3) groot/rtree: TTree-v5
* [d1049cd](/commit/d1049cd) groot/rtree: TBranch-v10
* [f3ad4a5](/commit/f3ad4a5) hplot: add support for log-y based functions
* [0b630d8](/commit/0b630d8) ci: add x-compile release script
* [0fa7206](/commit/0fa7206) ci: display C++ ROOT version
* [c79f47a](/commit/c79f47a) groot/{riofs,rtree}: fix 32b build
* [b2bb9f0](/commit/b2bb9f0) ci: test 32b builds
* [5f73852](/commit/5f73852) groot/internal/rtests: add ROOTError wrap test
* [ac4885b](/commit/ac4885b) ci: add binary ROOT installation scripts
* [91c9502](/commit/91c9502) groot/rtree: add WriterVarsFromStruct helper function
* [7379a85](/commit/7379a85) groot/rtree: make NewScanVars support Slices and Arrays
* [9be2300](/commit/9be2300) groot/rtree: test WithXXX API and ability to correctly read back trees from C++ ROOT
* [a1a9024](/commit/a1a9024) groot/internal/rtests: introduce RunCxxROOT to (more easily) run ROOT macros
* [ac1dc83](/commit/ac1dc83) groot/{riofs,internal/rcompress,rtree}: introduce rcompress.Settings
* [f35da3b](/commit/f35da3b) groot/{riofs,rtree}: add WithXYZ API to configure how trees should be created
* [813c909](/commit/813c909) ci: add mk-release script to push versions to index.golang
* [d12963d](/commit/d12963d) fwk: apply golangci suggestions
* [295baa5](/commit/295baa5) fwk: simplify fowk.MsgStream interface
* [86f5de3](/commit/86f5de3) all: use x/xerrors
* [f9cbd76](/commit/f9cbd76) groot/rtree: improve diagnostic+doc for rtree.New(Tree)Scanner
* [c4a04e9](/commit/c4a04e9) groot/rarrow: fix Arrow-struct creation data-offset
* [b8a39b7](/commit/b8a39b7) groot/rtree: handle auto-flush
* [b545af5](/commit/b545af5) groot/rtree: add support var-len arrays to trees

