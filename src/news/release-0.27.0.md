---
title: "Go HEP release 0.27.0"
date: 2020-05-26T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3836718.svg)](https://doi.org/10.5281/zenodo.3836718)

Release [`v0.27.0`](https://github.com/go-hep/hep/tree/v0.27.0) is out of the oven.

This release contains a major `groot` performance improvement.
As can be seen from [groot-bench](https://github.com/go-hep/groot-bench), `groot` can now read ROOT files (compressed or not) faster than (naive?) C++/ROOT `6.20/04` code (w/ `TTree::Branch` or `TTreeReader`):

```
name                               time/op
ReadCMS/GoHEP/Zlib-8               19.2s ± 1%
ReadCMS/ROOT-TreeBranch/Zlib-8     37.5s ± 1%
ReadCMS/ROOT-TreeReader/Zlib-8     26.1s ± 3%
ReadCMS/ROOT-TreeReaderMT/Zlib-8   25.6s ± 5%  (ROOT::EnableImplicitMT())
```

This was achieved by:
- re-engineering the reading code into dedicated `rleafXyz`, `rbasket`, ... types
- introducing a `rbasket` (concurrent) pre-fetcher

There are still a few low-hanging fruits to reap performances wise (reducing memory usage, reusing `rbasket` buffers, reusing `rbasket` buffers across decompression goroutines, etc...)

## cmd

- GoHEP gained a new command, `cmd/podio-gen`, that can generate Go types according to [PODIO](https://github.com/AIDASoft/podio) descriptions:

```
$> podio-gen -h
podio-gen generates a complete EDM from a PODIO YAML file definition.

Usage: podio-gen [OPTIONS] edm.yaml

Example:

  $> podio-gen -p myedm -o out.go -r 'edm4hep::->edm_,ExNamespace::->exns_' edm.yaml

Options:
  -o string
    	path to the output file containing the generated code (default "out.go")
  -p string
    	package name for the PODIO generated types (default "podio")
  -r string
    	comma-separated list of rewrite rules (e.g., 'edm4hep::->edm_')
```

This command should be useful in the connection with [Key4HEP](https://github.com/key4hep).

## fit

`fit` gained the ability to fit multivariate functions, thanks to [@JCCPort (Josh Porter)](https://github.com/JCCPort):

- https://godoc.org/go-hep.org/x/hep/fit#example-CurveND--Plane
- https://godoc.org/go-hep.org/x/hep/fit#CurveND


## fmom

`fmom` gained a couple of functions (under the friendly pressure and w/ contributions from [@rmadar (Romain Madar)](https://github.com/rmadar)):
- `fmom.Dot(p1, p2 P4) P4`
- `fmom.Boost(p P4, vec r3.Vec) P4`
- `fmom.BoostOf(p P4) r3.Vec`

## groot

`groot/rtree` gained a new type [rtree.Reader](https://go-hep.org/x/hep/groot/rtree#Reader), that allows to more easily (and faster) read data from a tree

- https://godoc.org/go-hep.org/x/hep/groot/rtree#example-Reader

```go
func ExampleReader() {
	f, err := groot.Open("../testdata/simple.root")
	if err != nil {
		log.Fatalf("could not open ROOT file: %+v", err)
	}
	defer f.Close()

	o, err := f.Get("tree")
	if err != nil {
		log.Fatalf("could not retrieve ROOT tree: %+v", err)
	}
	t := o.(rtree.Tree)

	var (
		v1 int32
		v2 float32
		v3 string

		rvars = []rtree.ReadVar{
			{Name: "one", Value: &v1},
			{Name: "two", Value: &v2},
			{Name: "three", Value: &v3},
		}
	)

	r, err := rtree.NewReader(t, rvars)
	if err != nil {
		log.Fatalf("could not create tree reader: %+v", err)
	}
	defer r.Close()

	err = r.Read(func(ctx rtree.RCtx) error {
		fmt.Printf("evt[%d]: %v, %v, %v\n", ctx.Entry, v1, v2, v3)
		return nil
	})
	if err != nil {
		log.Fatalf("could not process tree: %+v", err)
	}

	// Output:
	// evt[0]: 1, 1.1, uno
	// evt[1]: 2, 2.2, dos
	// evt[2]: 3, 3.3, tres
	// evt[3]: 4, 4.4, quatro
}

```

`root-dump` has been updated to use that new way of reading data.


Experimental support for `ROOT::RNtuple` was added -in a still very WIP fashion- into the `groot/exp/rntup` package.
When `ROOT7` is out (or when `ROOT::RNtuple` is more stable), this package is expected to graduate to `groot/rntup`.

The `rtree.Reader` type has also gained the ability to declare and evaluate user provided functions, taking a list of leaf names and the function to evaluate on the tree data:

- https://godoc.org/go-hep.org/x/hep/groot/rtree#example-Reader--WithFormulaFunc
- https://godoc.org/go-hep.org/x/hep/groot/rtree#example-Reader--WithFormulaFromUser

```go
func ExampleReader_withFormulaFunc() {
	f, err := groot.Open("../testdata/simple.root")
	if err != nil {
		log.Fatalf("could not open ROOT file: %+v", err)
	}
	defer f.Close()

	o, err := f.Get("tree")
	if err != nil {
		log.Fatalf("could not retrieve ROOT tree: %+v", err)
	}
	t := o.(rtree.Tree)

	var (
		data struct {
			V1 int32   `groot:"one"`
			V2 float32 `groot:"two"`
			V3 string  `groot:"three"`
		}
		rvars = rtree.ReadVarsFromStruct(&data)
	)

	r, err := rtree.NewReader(t, rvars)
	if err != nil {
		log.Fatalf("could not create tree reader: %+v", err)
	}
	defer r.Close()

	f64, err := r.FormulaFunc(
		[]string{"one", "two", "three"},
		func(v1 int32, v2 float32, v3 string) float64 {
			return float64(v2*10) + float64(1000*v1) + float64(100*len(v3))
		},
	)
	if err != nil {
		log.Fatalf("could not create formula: %+v", err)
	}

	fstr, err := r.FormulaFunc(
		[]string{"one", "two", "three"},
		func(v1 int32, v2 float32, v3 string) string {
			return fmt.Sprintf(
				"%q: %v, %q: %v, %q: %v",
				"one", v1, "two", v2, "three", v3,
			)
		},
	)
	if err != nil {
		log.Fatalf("could not create formula: %+v", err)
	}

	f1 := f64.Func().(func() float64)
	f2 := fstr.Func().(func() string)

	err = r.Read(func(ctx rtree.RCtx) error {
		v64 := f1()
		str := f2()
		fmt.Printf("evt[%d]: %v, %v, %v -> %g | %s\n", ctx.Entry, data.V1, data.V2, data.V3, v64, str)
		return nil
	})
	if err != nil {
		log.Fatalf("could not process tree: %+v", err)
	}

	// Output:
	// evt[0]: 1, 1.1, uno -> 1311 | "one": 1, "two": 1.1, "three": uno
	// evt[1]: 2, 2.2, dos -> 2322 | "one": 2, "two": 2.2, "three": dos
	// evt[2]: 3, 3.3, tres -> 3433 | "one": 3, "two": 3.3, "three": tres
	// evt[3]: 4, 4.4, quatro -> 4644 | "one": 4, "two": 4.4, "three": quatro
}
```

`groot` also gained a new command: `root-split`.

- https://godoc.org/go-hep.org/x/hep/groot/cmd/root-split

```
$> root-split -h
Usage: root-split [options] file.root

ex:
 $> root-split -o out.root -n 10 ./testdata/chain.flat.1.root

options:
  -n int
    	number of events to split into (default 100)
  -o string
    	path to output ROOT files (default "out.root")
  -t string
    	input tree name to split (default "tree")
  -v	enable verbose mode
```

`root-split` allows to split a given tree from a file into `n` files.

## hbook

`hbook` gained a new "sub-package", `hbook/ntup/ntroot` that provides convenience functions to expose `ROOT` trees as `hbook` ntuples.

`H1D` can now be subtracted, added, added and scaled (thanks `@rmadar`)

- https://godoc.org/go-hep.org/x/hep/hbook#AddH1D
- https://godoc.org/go-hep.org/x/hep/hbook#AddScaledH1D
- https://godoc.org/go-hep.org/x/hep/hbook#SubH1D

`hbook/yodacnv` now supports [YODA](https://yoda.hepforge.org) format version 2 (and its YAML-based metadata description.)


## hplot

Many improvement to `hplot` have been applied thanks again to the interesting suggestions from `@rmadar` (and his contributions.)
Namely, a new plotter, `HStack` has been provided:

![hstack-example](https://github.com/go-hep/hep/raw/master/hplot/testdata/hstack_golden.png)
![hstack-band-example](https://github.com/go-hep/hep/raw/master/hplot/testdata/hstack_band_golden.png)

A nice convenience tool to automatically generate (with `pdflatex` by default) `LaTeX` PDF plots from the `gonum.org/v1/plot/vg/vgtex` backend, has been also added to the new `hplot` concept of a `Figure`:

- https://godoc.org/go-hep.org/x/hep/hplot#Fig
- https://godoc.org/go-hep.org/x/hep/hplot#Figure
- https://godoc.org/go-hep.org/x/hep/hplot/htex#Handler

Also, a new `RatioPlot` plotter has been added as well:

![ratio-plot](https://github.com/go-hep/hep/raw/master/hplot/testdata/diff_plot_golden.png)

## lhef

`lhef` was improved to be able to handle version 3 of the Les Houches File format.


## sliceop/f64s

`sliceop/f64s` is a new package to apply some "map-reduce"-esque concepts and operations to slices of `float64` values.

- https://godoc.org/go-hep.org/x/hep/sliceop/f64s

The `Filter`, `Find`, `Map` and `Take` functions are available.

## AOB

That's all for today.
Writing trees with structured data is still on the roadmark, together with (probably) some performance improvements on the writing side of things.
Also, already in, is a feature similar to `TTreeFriends`: `rtree.Join(trees []rtree.Tree)`.
