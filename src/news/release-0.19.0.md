---
title: "Go HEP release 0.19.0"
date: 2019-05-31T0:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.3236446.svg)](https://doi.org/10.5281/zenodo.3236446)

Release [`v0.19.0`](https://github.com/go-hep/hep/tree/v0.19.0) is fresh from the oven.

This release contains 3 new `groot`-related packages:

- [groot/rsql/rsqldrv](https://pkg.go.dev/go-hep.org/x/hep/groot/rsql/rsqldrv): a package to present a ROOT File+Tree as a SQL database, by way of implementing the `database/sql/driver` interface for `groot`
- [groot/rsql](https://pkg.go.dev/go-hep.org/x/hep/groot/rsql): a convenience package to scan (a la `TTree::Scan` and `TTree::Draw`) `rtree.Tree` and create `hbook.H1D` or `hbook.H2D`
- [groot/rarrow](https://pkg.go.dev/go-hep.org/x/hep/groot/rarrow): a package to present `rtree.Tree` as [Apache Arrow's](https://arrow.apache.org) Tables and Records. This will be very useful for interoperability with DataScience and/or Machine Learning toolkits.

## groot

Providing `groot/rarrow` made a new command possible - [cmd/root2arrow](https://pkg.go.dev/go-hep.org/x/hep/cmd/root2arrow) - that can convert a ROOT TTree stored in a ROOT file, into a sequence of Arrow records, stored into an Arrow file (or Arrow stream.)

```
$> go doc go-hep.org/x/hep/cmd/root2arrow
root2arrow converts the content of a ROOT TTree to an ARROW file.

    Usage of root2arrow:
      -o string
        	path to output ARROW file name (default "output.data")
      -stream
        	enable ARROW stream (default is to create an ARROW file)
      -t string
        	name of the tree to convert (default "tree")

    $> root2arrow -o foo.data -t tree ../../groot/testdata/simple.root
    $> arrow-ls ./foo.data
    version: V4
    schema:
      fields: 3
        - one: type=int32
        - two: type=float32
        - three: type=utf8
    records: 1
    $> arrow-cat ./foo.data
    version: V4
    record 1/1...
      col[0] "one": [1 2 3 4]
      col[1] "two": [1.1 2.2 3.3 4.4]
      col[2] "three": ["uno" "dos" "tres" "quatro"]
```

Here is an example of `rsql.Scan`:

```go
func main() {
	f, err := groot.Open("../testdata/simple.root")
	if err != nil {
	    log.Fatal(err)
	}
	defer f.Close()
	
	o, err := f.Get("tree")
	if err != nil {
	    log.Fatal(err)
	}
	
	tree := o.(rtree.Tree)
	
	var (
	    v1s []int32
	    v2s []float64
	    v3s []string
	)
	
	err = rsql.Scan(tree, "SELECT (one, two, three) FROM tree", func(x int32, y float64, z string) error {
	    v1s = append(v1s, x)
	    v2s = append(v2s, y)
	    v3s = append(v3s, z)
	    return nil
	})
	
	if err != nil {
	    log.Fatal(err)
	}
	
	fmt.Printf("tree[%q]: %v\n", "one", v1s)   // tree["one"]: [1 2 3 4]
	fmt.Printf("tree[%q]: %v\n", "two", v2s)   // tree["two"]: [1.1 2.2 3.3 4.4]
	fmt.Printf("tree[%q]: %q\n", "three", v3s) // tree["three"]: ["uno" "dos" "tres" "quatro"]
}
```

And one example of `rsql.ScanH1D`:

```go
func main() {
	f, err := groot.Open("../testdata/simple.root")
	if err != nil {
	    log.Fatal(err)
	}
	defer f.Close()
	
	o, err := f.Get("tree")
	if err != nil {
	    log.Fatal(err)
	}
	
	tree := o.(rtree.Tree)
	
	h, err := rsql.ScanH1D(tree, "SELECT two FROM tree", nil)
	if err != nil {
	    log.Fatal(err)
	}
	
	fmt.Printf("entries: %v\n", h.Entries()) // entries: 4
	fmt.Printf("x-mean: %v\n", h.XMean())    // x-mean: 2.75
}
```

## AOB

Stay tuned! (and, as always, any kind of help (reviews, patches, nice emails, constructive criticism) deeply appreciated.)

