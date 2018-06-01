+++
title = "Go HEP release 0.6"
date = "2017-09-08T10:32:55+02:00"
weight = 10
+++

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.887895.svg)](https://doi.org/10.5281/zenodo.887895)

Release [`v0.6`](https://github.com/go-hep/hep/tree/v0.6) is a small iteration (code-wise) but contains:

- a software paper submitted to the Journal of OpenSource Software (JOSS): [![JOSS Paper](http://joss.theoj.org/papers/0b007c81073186f7c61f95ea26ad7971/status.svg)](http://joss.theoj.org/papers/0b007c81073186f7c61f95ea26ad7971)
- a new `WIP` command, [rootio/cmd/root-dump](https://go-hep.org/x/hep/rootio/cmd/root-dump), to dump the content of ROOT files, including the entries of TTrees,
- documentation fixes and updates in the `rootio` package,
- still more work in the `Delaunay` and `Voronoi` area (thanks [Bastian](https://github.com/Bastiantheone)!)

One API change that is worth noting, is the retrieval of `rootio.Key`s from a `rootio.Directory`.
The `d4823f0` changes:

```go
type Directory interface {
    Get(namecycle string) (Object, bool)
}
```

to:

```go
type Directory interface {
    Get(namecycle string) (Object, error)
}
```

this allows to more quickly bubble-up errors while reducing boilerplate code at the `Directory.Get(...)` call site:

```diff
diff --git a/cmd/root2npy/main.go b/cmd/root2npy/main.go
index c81e81d..e3827fa 100644
--- a/cmd/root2npy/main.go
+++ b/cmd/root2npy/main.go
@@ -159,9 +159,9 @@ func main() {
        }
        defer f.Close()
 
-       obj, ok := f.Get(*tname)
-       if !ok {
-               log.Fatalf("no object named %q in file %q", *tname, *fname)
+       obj, err := f.Get(*tname)
+       if err != nil {
+               log.Fatal(err)
        }
```
