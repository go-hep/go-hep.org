---
title: "Go HEP release 0.15.0"
date: 2018-09-06T14:41:23+02:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1410416.svg)](https://doi.org/10.5281/zenodo.1410416)

Release [`v0.15.0`](https://github.com/go-hep/hep/tree/v0.15.0) is fresh from the oven.

This new release dropped explicit support for Go-1.8.x but, in turn, gained support for Go-1.11.x and "Go Modules".
In a nutshell, Go modules allow to explicitly declare what are the dependencies a given module needs for the `go build` tool to successfully build it.
And, more importantly, Go modules allow to explicitly declare what are the needed versions of these dependencies, essentially making a build completely reproducible.

You can find more informations about Go modules over there:

- https://research.swtch.com/vgo
- https://github.com/golang/go/wiki/Modules

Currently, modules are only tested in Travis-CI, on the Go `master` branch.
But as Go-1.12.x will get closer and modules get more ubiquitous, we'll gradually switch to "Go modules" being the mainstream way to build Go-HEP.

Do not hesitate to report any issues you encounter when building with `GO111MODULE=on` enabled.

## rootio

Another big news for the `v0.15.0` release is the support for writing ROOT files:

- writing `TObjStrings`, `TH1x`, `TH2x`, `TGraph`, `TGraph{,Assymm}Errors` is in,
- support for writing compressed ROOT files as well (including `lz4`, `lzma` and `zlib`)
- 2 new ROOT-related commands:
  - `cmd/yoda2root`: a command to convert [YODA](https://yoda.hepforge.org) files into ROOT ones (so: histograms and scatters)
  - `rootio/cmd/root-cp`: a command to extract objects from a ROOT file into a new ROOT file

To support writing `TH1x`, `TH2x` and `TGraphs`, `hbook` types have been modified to export most of their fields -- so one can create a `rootio.H1D` from a `hbook.H1D`.
This enabled `hbook/rootcnv` to gain 3 new functions:

- [rootcnv.FromH1D](https://godoc.org/go-hep.org/x/hep/hbook/rootcnv#FromH1D): a function that converts an `hbook.H1D` into a `rootio.H1D`, loosing a bit of informations along the way (ROOT isn't as precise as `hbook` or `YODA` are)
- [rootcnv.FromH2D](https://godoc.org/go-hep.org/x/hep/hbook/rootcnv#FromH2D): a function that converts `hbook.H2D`s into `rootio.H2D`s,
- [rootcnv.FromS2D](https://godoc.org/go-hep.org/x/hep/hbook/rootcnv#FromS2D): a function that converts `hbook.S2D` into `rootio.TGraphAsymmErrors`.

## rootio & xrootd

Finally, we have received 2 patches from [Paul Seyfert (a.k.a pseyfert)](https://github.com/pseyfert), our first "CERNois" committer :).
Paul enhanced the UI of `root-ls` and `xrd-ls` to better deal with nested directories (and how they are displayed) in both of these commands.
Thanks Paul!

## Examples

Without further ado, here is how you would create a ROOT file, with `lz4` compression, containing a `TObjString`:

```go
func main() {
	w, err := rootio.Create("out.root", rootio.WithLZ4(flate.BestCompression))
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	var (
		k = "my-objstring"
		v = rootio.NewObjString("Hello World from Go-HEP!")
	)

	err = w.Put(k, v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("wkeys: %d\n", len(w.Keys()))

	err = w.Close()
	if err != nil {
		log.Fatalf("could not close file: %v", err)
	}
}
```

and here is how you would use `root-cp`:

```
$> root-cp -h
Usage: root-cp [options] file1.root [file2.root [...]] out.root

ex:
 $> root-cp f.root out.root
 $> root-cp f1.root f2.root f3.root out.root
 $> root-cp f1.root:hist* f2.root:h2 out.root

options:

$> root-cp ./testdata/graphs.root:g* out.root
$> root-cp root://xrootd.example.org/file.root:hist.* out.root
```

More ROOT files writing examples can be found here:

- [Histo1D](https://godoc.org/go-hep.org/x/hep/rootio#example-Create--Histo1D)
- [Histo2D](https://godoc.org/go-hep.org/x/hep/rootio#example-Create--Histo2D)
- [GraphAsymmErrors](https://godoc.org/go-hep.org/x/hep/rootio#example-Create--GraphAsymmErrors)
- [Create WithZlib](https://godoc.org/go-hep.org/x/hep/rootio#example-Create--WithZlib)

## AOB

We will try to have preliminary support for writing `TTrees` in the next release.
That should be fun.

Interoperability with [Apache Arrow Arrays](https://arrow.apache.org) is still on the table.
It might even prove to be easier to support Apache Arrow first and then implement `TTrees` writing support on top of that.
We will see...

Stay tuned! (and, as always, any kind of help (reviews, patches, nice emails, constructive criticism) deeply appreciated.)
