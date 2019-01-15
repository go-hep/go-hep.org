---
date: 2019-01-15T12:32:03+01:00
title: "Users of Go-HEP and Go in HENP"
weight: 30
---

This page tries to gather users of [Go-HEP](https://go-hep.org), papers and talks about Go-HEP and/or uses of [Go](https://golang.org) in science and HENP in particular.
Feel free to send pull requests [here](https://github.com/go-hep/go-hep.org) to complete this list.

## Go Research papers

The [Go](https://golang.org) project maintains a wiki page that gathers a few research papers written *about* Go or *with* Go:

- https://github.com/golang/go/wiki/ResearchPapers

## Papers and Talks

### 2019

### 2018

- ProIO: An Event-Based I/O Stream Format for Protobuf Messages (D. Blyth, J. Alcaraz, S. Binet, S.V. Chekanov)
  - [arXiv:1812.03967](https://arxiv.org/abs/1812.03967)
- Towards the ALICE Online-Offline (O2) control system, T. Mrnjavac for the ALICE collaboration
  - [indico](https://indico.cern.ch/event/587955/contributions/2935762/)
- Exploring polyglot software frameworks in ALICE with FairMQ and fer, S. Binet for the ALICE collaboration
  - [indico](https://indico.cern.ch/event/587955/contributions/2938059/)

### 2017

- Go-HEP: writing concurrent software with ease and Go (S. Binet)
  - [doi:10.1088 1742/6596/1085/5/052012](https://doi.org/10.1088/1742-6596/1085/5/052012)
	- [arXiV:1808.06529](https://arxiv.org/abs/1808.06529)
- Construction and First Tests of anin-beamPET Demonstrator Dedicated to the Ballistic Control of Hadrontherapy Treatments With 65 MeV Protons, E. Busato _et al_
  - [doi/10.1109/TRPMS.2017.2780447](https://doi.org/10.1109/TRPMS.2017.2780447)

### 2016

- The SoLid anti-neutrino detector's read-out system, N. Ryder
  - [doi/10.1109/NSSMIC.2016.8069730](https://doi.org/10.1109/NSSMIC.2016.8069730)

### 2012

- Can Go address the multicore issues of today and the manycore problems of tomorrow? (S. Binet)
  - [doi:10.1088 1742-6596/368/1/012017](https://doi.org/10.1088/1742-6596/368/1/012017)
	- [pdf](http://iopscience.iop.org/article/10.1088/1742-6596/368/1/012017/pdf)
- GoCxx: a tool to easily leverage C++ legacy code for multicore-friendly Go libraries and frameworks (S. Binet)
  - [doi:10.1088 1742-6596/396/5/052012](https://doi.org/10.1088/1742-6596/396/5/052012)
	- [pdf](http://iopscience.iop.org/article/10.1088/1742-6596/396/5/052012/pdf)

### 2011

- ng: What next-generation languages can teach us about HENP frameworks in the manycore era (S. Binet)
  - [doi:10.1088 1742-6596/331/4/042002](https://doi.org/10.1088/1742-6596/331/4/042002)
  - [pdf](http://iopscience.iop.org/article/10.1088/1742-6596/331/4/042002/pdf)

## Code

- ALICE, for the O2 Control system: https://github.com/AliceO2Group/Control
- ATLAS, TileCal HV analysis code: (private code)
- AVIRM group, DAQ+Monitoring: https://gitlab.in2p3.fr/avirm/analysis-go
- CMS, Data aggregation system: https://github.com/dmwm/das2go
- LSST, slow control: https://github.com/go-lsst/fcs-lpc-motor-ctl
- LSST, IT tools: https://github.com/airnandez/cluefs
- Master-1,2 lectures/exercizes: https://github.com/master-pfa-info
- ProIO: https://github.com/proio-org/go-proio
- QUBIC, database monitoring:  https://github.com/ziotom78/qutedb
- SoLiD, trigger: https://bitbucket.org/solidexperiment/readout-software/src/master/
- SoLiD, Monitoring: https://github.com/sbinet-solid/solid-mon-rpi