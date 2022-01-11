---
date: 2019-01-15T12:32:03+01:00
title: "Users of Go-HEP and Go in HENP"
weight: 8
---

This page tries to gather users of [Go-HEP](https://go-hep.org), papers and talks about Go-HEP and/or uses of [Go](https://golang.org) in science and HENP in particular.
Feel free to send pull requests [here](https://github.com/go-hep/go-hep.org) to complete this list.

## Go Research papers

The [Go](https://golang.org) project maintains a wiki page that gathers a few research papers written *about* Go or *with* Go:

- https://github.com/golang/go/wiki/ResearchPapers

## Papers and Talks

### 2021

- `groot`: reading `ROOT` data, with `Go`, faster than `ROOT` (plenay talk [@Journées Informatiques IN2P3/IRFU](https://indico.in2p3.fr/event/25008))
  - [talk](https://talks.sbinet.org/2021/2021-11-16-ji-groot/talk.slide)

### 2020

- `ROOT-I/O` and `groot` (status report @LPC-Dev)
  - [talk - part 1](https://talks.sbinet.org/2020/2020-06-15-lpc-dev-rootio/talk.slide)
  - [talk - part 2](https://talks.sbinet.org/2020/2020-06-26-lpc-dev-groot/talk.slide)

### 2019

- Go-HEP: what can Go do for science?
  - [talks:pdf](https://github.com/sbinet/talks/raw/main/2019/2019-10-22-golab-io-gohep/slides.pdf)
	- [talk:video](https://invidious.fdn.fr/watch?v=OOzMBZjLkH0&list=PLGN1AjiJJv0mHK60uQxVXjYg-G8d4YyqQ&index=3&t=0s)
	- [conference:GoLabIO-2019](https://golab.io/agenda/session/98759)

### 2018

- ProIO: An Event-Based I/O Stream Format for Protobuf Messages (D. Blyth, J. Alcaraz, S. Binet, S.V. Chekanov)
  - [arXiv:1812.03967](https://arxiv.org/abs/1812.03967)
  - [doi:10.1016 j.cpc.2019.03.018](https://doi.org/10.1016/j.cpc.2019.03.018)
- Towards the ALICE Online-Offline (O2) control system, T. Mrnjavac for the ALICE collaboration
  - [indico](https://indico.cern.ch/event/587955/contributions/2935762/)
- Exploring polyglot software frameworks in ALICE with FairMQ and fer, S. Binet for the ALICE collaboration
  - [arXiv:1901.04834](https://arxiv.org/abs/1901.04834)
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

### 2015

- docker & HEP: Containerization of applications for development, distribution and preservation, S. Binet and B. Couturier
  - [doi:10.1088 1742-6596/664/2/022007](https://doi.org/10.1088/1742-6596/664/2/022007)
	- [pdf](http://iopscience.iop.org/article/10.1088/1742-6596/664/2/022007/pdf)

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
- ALICE, for the O2 framework: https://github.com/sbinet-alice/fer
- ATLAS, TileCal HV analysis code: (private code)
- ATLAS, DAQ for the Front-end test board v2 (ATLAS-BNL): https://gitlab.cern.ch/BNL-ATLAS/larphase2/fetb2/gofetb2daq
- AVIRM group, DAQ+Monitoring: https://gitlab.in2p3.fr/avirm/analysis-go
- CMS, Data aggregation system: https://github.com/dmwm/das2go
- LSST, slow control: https://github.com/go-lsst/fcs-lpc-motor-ctl
- LSST, IT tools: https://github.com/airnandez/cluefs
- Master-1,2 lectures/exercizes: https://github.com/master-pfa-info
- ProIO: https://github.com/proio-org/go-proio
- QUBIC, database monitoring:  https://github.com/ziotom78/qutedb
- SoLiD, trigger: https://bitbucket.org/solidexperiment/readout-software/src/master/
- SoLiD, Monitoring: https://github.com/sbinet-solid/solid-mon-rpi
