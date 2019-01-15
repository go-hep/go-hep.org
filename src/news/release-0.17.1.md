---
title: "Go HEP release 0.17.1"
date: 2019-01-14T12:00:00+01:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.2540596.svg)](https://doi.org/10.5281/zenodo.2540596)

Release [`v0.17.1`](https://github.com/go-hep/hep/tree/v0.17.1) is fresh-ish from the oven.

This new release continues the refactoring and consolidation work of the [groot](https://go-hep.org/x/hep/groot) package (meant to replace `rootio`.)

## brio

- `brio/cmd/brio-gen` gained some documentation

## groot

- `groot/rdict` has seen a bunch of work to create user-type `StreamerInfo`s,
- `groot/cmd/root-gen-streamer`: new command to automatically generate a `StreamerInfo` and its `StreamerElement`s given a ROOT or user type,
- `groot/rdict` now properly handles streamers with arrays of builtins, and properly visits `std::vector<T>` fields (where `T` is a builtin)
- `groot/cmd/root-dump`: now properly handle `root.List` values
- `groot` dropped the use of `gofrs/uuid` and replaced it with `hashicorp/go-uuid`. `gofrs/uuid` dropped support for Go modules and broke the build of Go-HEP.

Here is usage example of the new [root-gen-streamer](https://go-hep.org/x/hep/groot/cmd/root-gen-streamer) command:

```sh
$> root-gen-streamer -help
Usage: root-gen-streamer [options]

ex:
 $> root-gen-streamer -p image -t Point -o streamers_gen.go
 $> root-gen-streamer -p go-hep.org/x/hep/hbook -t Dist0D,Dist1D,Dist2D -o foo_streamer_gen.go

options:
  -o string
    	output file name
  -p string
    	package import path
  -t string
    	comma-separated list of type names
  -v	enable verbose mode

```

## xrootd

The `xrootd/client` and `xrootd/server` packages have been merged into a single package, `xrootd`.
This simplified the import structure as well as reduced the amount of boilerplate code that was duplicated between the two packages.

For users of `xrootd/client`:

```diff
diff --git a/xrootd/cmd/xrd-ls/main.go b/xrootd/cmd/xrd-ls/main.go
index 784a5c7..9961802 100644
--- a/xrootd/cmd/xrd-ls/main.go
+++ b/xrootd/cmd/xrd-ls/main.go
@@ -31,7 +31,7 @@ import (
        "text/tabwriter"
 
        "github.com/pkg/errors"
-       xrdclient "go-hep.org/x/hep/xrootd/client"
+       "go-hep.org/x/hep/xrootd"
        "go-hep.org/x/hep/xrootd/xrdfs"
        "go-hep.org/x/hep/xrootd/xrdio"
 )
@@ -93,7 +93,7 @@ func xrdls(name string, long, recursive bool) error {
 
        ctx := context.Background()
 
-       c, err := xrdclient.NewClient(ctx, url.Addr, url.User)
+       c, err := xrootd.NewClient(ctx, url.Addr, url.User)
        if err != nil {
                return errors.Errorf("could not create client: %v", err)
        }
```

For users of `xrootd/server`:

```diff
diff --git a/xrootd/cmd/xrd-srv/main.go b/xrootd/cmd/xrd-srv/main.go
index 2a09dbd..230bf31 100644
--- a/xrootd/cmd/xrd-srv/main.go
+++ b/xrootd/cmd/xrd-srv/main.go
@@ -14,7 +14,7 @@ import (
        "os"
        "os/signal"
 
-       "go-hep.org/x/hep/xrootd/server"
+       "go-hep.org/x/hep/xrootd"
 )
 
 func init() {
@@ -56,7 +56,7 @@ func main() {
                log.Fatalf("could not listen on %q: %v", *addr, err)
        }
 
-       srv := server.New(server.NewFSHandler(baseDir), func(err error) {
+       srv := xrootd.NewServer(xrootd.NewFSHandler(baseDir), func(err error) {
                log.Printf("an error occured: %v", err)
        })
```

## AOB

Support for writing `TTree`s didn't make it under the X-mas tree.
This has been converted to a NYE resolution, though.

Before tackling this big item, support for reading `TClonesArray` is on the way (and tracked in the `sbinet/hep#issue-419` branch.)

A proposal for data frames is in the works in the [gonum/exp](https://github.com/gonum/exp) repository.
Feel free to comment on the associated [pull request (#19)](https://github.com/gonum/exp/pull/19) or on the [gonum-dev](https://groups.google.com/d/msg/gonum-dev/Ktm7SPGN0BY/L2-cIgzvDQAJ) forum.

Finally, this is GSoC proposal season.
I (`@sbinet`) will probably send a `GDML` oriented proposal.
Feel free to send or discuss yours on the `go-hep` mailing list.

Stay tuned! (and, as always, any kind of help (reviews, patches, nice emails, constructive criticism) deeply appreciated.)
