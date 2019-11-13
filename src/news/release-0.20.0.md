---
title: "Go HEP release 0.20.0"
date: 2019-11-13T17:00:00+01:00
weight: 10
---
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3540979.svg)](https://doi.org/10.5281/zenodo.3540979)

Release [`v0.20.0`](https://github.com/go-hep/hep/tree/v0.20.0) is fresh from the oven.

This release contains a major new `groot` feature: the ability to *write* (simple, flat) trees:

- [groot/rtree#NewWriter](https://godoc.org/go-hep.org/x/hep/groot/rtree#NewWriter)

and a couple of new features for [hplot](https://godoc.org/go-hep.org/x/hep/hplot).

## groot

- `groot` support has been bumped to `ROOT-6.18/04`
- `groot` can now create nested directories that are compatible with ROOT.
Previously, `groot` would create nested directories which could only be read back by `groot` (and `uproot`).

```go
w, err := groot.Create("subdirs.root")
if err != nil {
	log.Fatal(err)
}
defer w.Close()

dir1, err := w.Mkdir("dir1")
if err != nil {
	log.Fatal(err)
}

dir11, err := dir1.Mkdir("dir11")
if err != nil {
	log.Fatal(err)
}

err = dir11.Put("obj1", rbase.NewObjString("data-obj1"))
if err != nil {
	log.Fatal(err)
}

dir2, err := w.Mkdir("dir2")
if err != nil {
	log.Fatal(err)
}

err = dir2.Put("obj2", rbase.NewObjString("data-obj2"))
if err != nil {
	log.Fatal(err)
}

err = w.Close()
if err != nil {
	log.Fatal(err)
}

// Output:
// >> subdirs.root
// >> subdirs.root/dir1
// >> subdirs.root/dir1/dir11
// >> subdirs.root/dir1/dir11/obj1
// >> subdirs.root/dir2
// >> subdirs.root/dir2/obj2
```

- `cmd/root2csv` can now handle `TLeafElement` leaves of builtins

- `groot/riofs.File.SegmentMap` has been added to `riofs.File` to inspect the list of ROOT records in a file

- support for writing simple flat trees of builtins and arrays of builtins (no slices, yet) has been added:

[godoc:example](https://godoc.org/go-hep.org/x/hep/groot/rtree#example-package--CreateFlatNtuple)
```go
type Data struct {
	I32    int32
	F64    float64
	Str    string
	ArrF64 [5]float64
}
const (
	fname = "groot-flat-ntuple.root"
	nevts = 5
)

f, err := groot.Create(fname)
if err != nil {
	log.Fatalf("%+v", err)
}
defer f.Close()

var evt Data

wvars := []rtree.WriteVar{
	{Name: "I32", Value: &evt.I32},
	{Name: "F64", Value: &evt.F64},
	{Name: "Str", Value: &evt.Str},
	{Name: "ArrF64", Value: &evt.ArrF64},
}
tree, err := rtree.NewWriter(f, "mytree", wvars)
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
// -- filled tree with 5 entries
// -- read back ROOT file
// entry[0]: {I32:0 F64:0 Str:evt-0 ArrF64:[0 1 2 3 4]}
// entry[1]: {I32:1 F64:1 Str:evt-1 ArrF64:[1 2 3 4 5]}
// entry[2]: {I32:2 F64:2 Str:evt-2 ArrF64:[2 3 4 5 6]}
// entry[3]: {I32:3 F64:3 Str:evt-3 ArrF64:[3 4 5 6 7]}
// entry[4]: {I32:4 F64:4 Str:evt-4 ArrF64:[4 5 6 7 8]}
```

## hplot

- [hplot.VLine](https://godoc.org/go-hep.org/x/hep/hplot#VLine) has been added to easily create a vertical line on a plot
- [hplot.HLine](https://godoc.org/go-hep.org/x/hep/hplot#HLine) has been added to easily create a horitzontal line on a plot
- [hplot.NewBand](https://godoc.org/go-hep.org/x/hep/hplot#NewBand) has been added to easily create a band between two sets of data points on a plot

![vline-example](https://github.com/go-hep/hep/raw/master/hplot/testdata/vline_golden.png)
![hline-example](https://github.com/go-hep/hep/raw/master/hplot/testdata/hline_golden.png)
![band-example](https://github.com/go-hep/hep/raw/master/hplot/testdata/band_golden.png)

## AOB

We'll try to release `v0.21.0` shortly with the ability to write slices to a tree and perhaps structs.

Stay tuned! (and, as always, any kind of help (reviews, patches, nice emails, constructive criticism) deeply appreciated.)

