+++
date = "2016-11-25T09:11:34+01:00"
title = "Contributing"
weight = 7
+++

The `go-hep` projects eagerly accepts contributions from the community.

## Introduction

The `go-hep` project provides libraries and tools in Go for the HEP community, and we would like you to join us in improving `go-hep`'s quality and scope.
This document is for anyone who is contributing or interested in contributing.
Questions about `go-hep` or the use of its libraries can be directed to the [go-hep](mailto:go-hep@googlegroups.com) mailing list.

## Contributing

### Working Together

When contributing or otherwise participating, please:

- Be friendly and welcoming
- Be patient
- Be thoughtful
- Be respectful
- Be charitable
- Avoid destructive behavior

Excerpted from the [Go conduct document](https://golang.org/conduct).

### Reporting Bugs

When you encounter a bug, please open an issue on the corresponding repository.
Start the issue title with the repository/sub-repository name, like `fwk/hbooksvc: issue name`.
Be specific about the environment you encountered the bug in (_e.g.:_ operating system, Go compiler version, ...).
If you are able to write a test that reproduces the bug, please include it in the issue.
As a rule, we keep all tests OK and try to increase code coverage.

### Your First Code Contribution

If you are a new contributor, *thank you!*
Before your first merge, you will need to be added to the [CONTRIBUTORS](https://codeberg.org/go-hep/license/src/tree/main/CONTRIBUTORS) and [AUTHORS](https://codeberg.org/go-hep/license/src/tree/main/AUTHORS) files.
Open a pull request adding yourself to these files.
All `go-hep` code follows the BSD license in the [license document](https://codeberg.org/go-hep/license/src/tree/main/LICENSE).
We prefer that code contributions do not come with additional licensing.
For exceptions, added code must also follow a BSD license.

### Code Contribution

If it is possible to split a large pull request into two or more smaller pull requests, please try to do so.
Pull requests should include tests for any new code before merging.
It is ok to start a pull request on partially implemented code to get feedback, and see if your approach to a problem is sound.
You don't need to have tests, or even have code that compiles to open a pull request, although both will be needed before merge.
When tests use magic numbers, please include a comment explaining the source of the number.
Benchmarks are optional for new features, but if you are submitting a pull request justified by performance improvement, you will need benchmarks to measure the impact of your change, and the pull request should include a report from [benchcmp](https://pkg.go.dev/golang.org/x/tools/cmd/benchcmp) or, preferably, [benchstat](https://github.com/rsc/benchstat).

Commit messages also follow some rules.
They are best explained at the official [Go](https://golang.org) "Contributing guidelines" document:

[golang.org/doc/contribute.html](https://golang.org/doc/contribute.html#commit_changes)

For example:

```
fwk/hbooksvc: add support for booking S2D values

This CL adds support for saving and reading `hbook.S2D` values.
The existing implementation of the `hbooksvc` had only support
for `hbook.H1D` and `hbook.H2D` values.

Fixes go-hep/hep#42.
```

If the `CL` modifies multiple packages at the same time, include them in the commit message:

```
hbook/rootcnv,rootio: implement conversion from rootio.TH1x to hbook.H1D

This CL adds the ability to convert rootio.TH1{D,F,I} 1-dim histograms
to hbook.H1D.
Note that ROOT::TH1x histograms do not provide enough informations to
correctly (and completely) construct a hbook.H1D (missing sumwx and
sumwx2 informations for the bins.)

Fixes go-hep/hep#40.
```

Please always format your code with [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports).
Best is to have it invoked as a hook when you save your `.go` files.

Files in the `go-hep` repository don't list author names, both to avoid clutter and to avoid having to keep the lists up to date.
Instead, your name will appear in the change log and in the [CONTRIBUTORS](https://codeberg.org/go-hep/license/src/tree/main/CONTRIBUTORS) and [AUTHORS](https://codeberg.org/go-hep/license/src/tree/main/AUTHORS) files.

New files that you contribute should use the standard copyright header:

```
// Copyright 2019 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
```

Files in the repository are copyright the year they are added.
Do not update the copyright year on files that you change.

### Code Review

If you are a contributor, please be welcoming to new contributors.
[Here](http://sarah.thesharps.us/2014/09/01/the-gentle-art-of-patch-review/) is a good guide.

There are several terms code reviewers may use that you should become familiar with.

  * ` LGTM ` — looks good to me
  * ` SGTM ` — sounds good to me
  * ` CL ` — change list; a single commit in the repository
  * ` s/foo/bar/ ` — please replace ` foo ` with ` bar `; this is [sed syntax](http://en.wikipedia.org/wiki/Sed#Usage)
  * ` s/foo/bar/g ` — please replace ` foo ` with ` bar ` throughout your entire change

We follow the convention of requiring at least 1 reviewer to say LGTM before a merge.
When code is tricky or controversial, submitters and reviewers can request additional review from others and more LGTMs before merge.
You can ask for more review by saying PTAL in a comment in a pull request.
You can follow a PTAL with one or more @someone to get the attention of particular people.
If you don't know who to ask, and aren't getting enough review after saying PTAL, then PTAL @go-hep/developers will get more attention.
Also note that you do not have to be the pull request submitter to request additional review.

### What Can I Do to Help?

If you are looking for some way to help the `go-hep` project, there are good places to start, depending on what you are comfortable with.

- You can [search](https://codeberg.org/go-hep/hep/issues) for open issues in need of resolution.
- You can improve documentation, or improve examples.
- You can add and improve tests.
- You can improve performance, either by improving accuracy, speed, or both.
- You can suggest and implement new features that you think belong in `go-hep`.

### Style

We use [Go style](https://github.com/golang/go/wiki/CodeReviewComments).

---

This _"Contributing"_ guide has been extracted from our sister project [Gonum](https://gonum.org).
Its guide is [here](https://github.com/gonum/license/blob/master/CONTRIBUTING.md).
