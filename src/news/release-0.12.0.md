---
title: "Go HEP release 0.12.0"
date: 2018-06-01T13:49:07+02:00
weight: 10
---

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.1257252.svg)](https://doi.org/10.5281/zenodo.1257252)

Release [`v0.12.0`](https://github.com/go-hep/hep/tree/v0.12.0) is fresh from the oven.

This release is first one to introduce preliminary support for [vgo](https://research.swtch.com/vgo), the official [Go](https://golang.org) way to handle versioning.
`vgo` is still in flux: the first Go version with experimental opt-in support should be `Go 1.11` (to be released in August 2018.)
Obviously, on the Go-HEP side, adjustments will probably still be required as the user story solidifies and experience is accumulated.

Nonetheless, it is still an interesting new development!

## geo/gdml

This release adds preliminary support for parsing [Geometry Description Markup Language (GDML)](http://cern.ch/gdml) files, a _de facto_ standard for describing (detector) geometries.
The documentation for this new package is here: [geo/gdml](https://godoc.org/go-hep.org/x/hep/geo/gdml).

Help wanted and (gladly) accepted to get this package in a shape where it could be used for detailed detector studies!
This is tracked here:

- https://github.com/go-hep/hep/projects/4

## hplot

`hplot` was slightly updated to cope with an interesting development percolating from upstream [gonum/plot](https://godoc.org/gonum.org/v1/plot), namely: the migration to a new PDF backend that allows to embed fonts inside the output PDF file.
No more PDFs that display weirdly on foreign computer. Yay!

This obviously means the resulting PDF files may be quite larger than with previous versions.
(You can't have your cake and eat it.)
You can use [vgpdf.Canvas.EmbedFonts](https://godoc.org/gonum.org/v1/plot/vg/vgpdf#Canvas.EmbedFonts) to get the old behaviour.

## rootio

This release adds preliminary support for chaining multiple `rootio.Tree`s into a logical view: the famed [rootio.Chain](https://godoc.org/go-hep.org/x/hep/rootio#Chain).

[Mohamed Amine El Gnaoui (a.k.a @maloft)](https://github.com/maloft), our new summer student [@LPC-Clermont](http://clrwww.in2p3.fr), provided the initial implementation: thanks!
More tests and benchmarks improvements yet to come :)

Another noteworthy change: `rootio/cmd/root-srv` dropped its dependency against `github.com/satori/go.uuid` in favor of `github.com/pborman/uuid`.
The latter exposes a more stable API.

## xrootd

This release adds yet another new package: `xrootd`.
This package will provide (eventually) a pure-Go implementation of an XRootD client as well as a server.

[Mikhail Ivchenko (a.k.a @EgorMatirov)](https://github.com/EgorMatirov), our [Google Summer of Code 2018](https://summerofcode.withgoogle.com/) student has been already hard at work, providing support for:

- the initial `xrootd` client, 
- the handshake with an XRootD-compliant server (`C++` or otherwise), and
- the `protocol` and `login` requests/responses.

The `dirlist` request/response is already in the pipe.
