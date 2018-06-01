+++
title = "Go HEP release 0.2"
date = "2017-04-25T16:39:10+02:00"
weight = 10
+++

Two weeks after release `v0.1`, here is `v0.2`:

- Changelog: https://github.com/go-hep/hep/releases/v0.2
- Sources (zip): https://github.com/go-hep/hep/archive/v0.2.zip
- Sources (tar): https://github.com/go-hep/hep/archive/v0.2.tar.gz
- DOI: https://doi.org/10.5281/zenodo.556591

## What's new

New `fastjet` algorithms have been contributed by [Bastian Wieck](https://github.com/Bastiantheone) (thanks!):

- `EeGenKtAlgorithm`
- `EeKtAlgorithm`

`fastjet` is now also able to compute exclusive jets (thanks again [Bastian Wieck](https://github.com/Bastiantheone))

A new [lcio](https://godoc.org/go-hep.org/x/hep/lcio) package, with initial read/write support for (most) of the [LCIO event data model](https://github.com/iLCSoft/LCIO), has been released:

- `MCParticle`
- `SimTrackerHit`
- `SimCalorimeterHit`
- `LCFloatVec`
- `LCIntVec`
- `LCStrVec`
- `RawCalorimeterHit`
- `CalorimeterHit`
- `TrackerData`
- `TrackerHit`
- `TrackerHitPlane`
- `TrackerHitZCylinder`
- `TrackerPulse`
- `TrackerRawData`
- `Track`
- `Cluster`
- `Vertex`
- `ReconstructedParticle`
- `LCGenericObject`
- `LCRelation`

See the [`lcio-ls`](https://go-hep.org/x/hep/lcio/cmd/lcio-ls) command and the [examples](https://godoc.org/go-hep.org/x/hep/lcio#pkg-examples) for more informations.

## Next steps

Release `v0.3` (trying to maintain our 2-weeks release schedule) should see real work on providing support for reading `TTrees` with user classes.

In the meantime, please give `v0.2` a spin, file [issues](https://github.com/go-hep/hep/issues), send patches (via [pull requests](https://github.com/go-hep/hep/pulls)) and/or discuss anything `Go-HEP` related on the [go-hep](mailto:go-hep@googlegroups.com) mailing list.

