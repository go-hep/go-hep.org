+++
title = "Go HEP release 0.5"
date = "2017-09-04T10:32:55+02:00"
weight = 10
+++

Release [`v0.5`](https://github.com/go-hep/hep/tree/v0.5) has seen some interesting work.

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.884416.svg)](https://doi.org/10.5281/zenodo.884416)

## General

- Go-HEP has migrated all its [Gonum](https://gonum.org) imports to use `"gonum.org/v1/gonum/{mat,stat,...}"`, instead of the old `"github.com/gonum/{mat64,stat,...}"`
- Go-HEP has also migrated all its use of `gonum/plot` from `"github.com/gonum/plot/..."` to the new `"gonum.org/v1/plot/..."`
- Go-1.6's support has been dropped from the Travis-CI matrix
- Go-1.9's support has been added to the Travis-CI matrix

## rootio

The `rootio` package has seen most of the work: performance improvements for reading ROOT files, a new `root2npy` command and the beginning of the work to provide write support.

The performances of `rootio` have been assessed for the [ACAT conference](https://indico.cern.ch/event/567550/).
The talk has all the details and can be read [here](https://talks.godoc.org/github.com/go-hep/talks/2017/2017-08-24-go-hep-acat/talk.slide).
In a nutshell, `rootio` is now "just" 2 times slower than C++ ROOT to read files, and uses less than 10 times the VMem of C++ ROOT.
More work to match the read speed of C++ ROOT is of course needed but we think we know how to recoup most of performance delta.

[`cmd/root2npy`](https://go-hep.org/x/hep/cmd/root2npy) is a simple command to quickly convert a ROOT file containing a TTree into a [NumPy array data file format](https://docs.scipy.org/doc/numpy/neps/npy-format.html).

[`cmd/root-srv`](https://go-hep.org/x/hep/rootio/cmd/root-srv) has been updated to support `https` (thanks [Michaël Ughetto](https://github.com/mughetto)).

[`cmd/root2yoda`](https://go-hep.org/x/hep/cmd/root2yoda) is a simple command to quickly convert a ROOT file containing histograms and profiles into a [YODA](https://yoda.hepforge.org/) file.

## hbook & csvutil

`csvutil/csvdriver` has been improved to correctly handle CSV headers.
You need to tell the driver that the CSV file contains a header describing the columns of the CSV file, but once that's done, you can use the correct columns' names in your SQL queries.

`hbook/ntup/ntcsv` is a new package that leverages the new feature of `csvdriver` to easily create `hbook/ntup.Ntuple` values from a CSV file, with the correct columns' names.

## fit

Initial support for automatically computing the Hessian matrix has been added.
If the user does not provide a Hess function, Gonum's [diff/fd.Hessian](https://pkg.go.dev/gonum.org/v1/gonum/diff/fd#Hessian) will be used

## fastjet

More of the GSoC work for the faster jet clustering has been integrated.
Thanks [Bastian Wieck](https://github.com/Bastiantheone).

## AOB

With release `v0.5`, Go-HEP gained its first external user: [eicio](https://github.com/decibelCooper/eicio), _"an exploratory project testing ideas about forward-looking IO solutions to suit the needs of the Electron Ion Collider (EIC) community."_
Welcome `eicio`!
