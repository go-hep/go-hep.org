+++
date = "2016-11-25T09:11:34+01:00"
title = "Contributing"
weight = 30
+++

The `go-hep` project eagerly accept contributions from the community.

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
Be specific about the environment you encountered the bug in.
If you are able to write a test that reproduces the bug, please include it in the issue.
As a rule we keep all tests OK.

### Suggesting Enhancements

If the scope of the enhancement is small, open an issue.
If it is large, such as suggesting a new repository, sub-repository, or interface refactoring, then please start a discussion on [the go-hep list](https://groups.google.com/forum/#!forum/go-hep).

### Your First Code Contribution

If you are a new contributor, thank you!
Before your first merge, you will need to be added to the [CONTRIBUTORS](https://github.com/go-hep/license/blob/master/CONTRIBUTORS) and [AUTHORS](https://github.com/go-hep/license/blob/master/AUTHORS) file.
Open a pull request adding yourself to them.
All `go-hep` code follows the BSD license in the [license document](https://github.com/go-hep/license/blob/master/LICENSE).
We prefer that code contributions do not come with additional licensing.
For exceptions, added code must also follow a BSD license.

### Code Contribution

If it is possible to split a large pull request into two or more smaller pull requests, please try to do so.
Pull requests should include tests for any new code before merging.
It is ok to start a pull request on partially implemented code to get feedback, and see if your approach to a problem is sound.
You don't need to have tests, or even have code that compiles to open a pull request, although both will be needed before merge.
When tests use magic numbers, please include a comment explaining the source of the number.
Benchmarks are optional for new features, but if you are submitting a pull request justified by performance improvement, you will need benchmarks to measure the impact of your change, and the pull request should include a report from [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp) or, preferably, [benchstat](https://github.com/rsc/benchstat).

### Code Review

If you are a contributor, please be welcoming to new contributors.
[Here](http://sarah.thesharps.us/2014/09/01/the-gentle-art-of-patch-review/) is a good guide.

There are several terms code reviewers may use that you should become familiar with.

  * ` LGTM ` — looks good to me
  * ` SGTM ` — sounds good to me
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

- You can [search](https://github.com/issues?utf8=%E2%9C%93&q=is%3Aopen+is%3Aissue+user%3Ago-hep) for open issues in need of resolution.
- You can improve documentation, or improve examples.
- You can add and improve tests.
- You can improve performance, either by improving accuracy, speed, or both.
- You can suggest and implement new features that you think belong in `go-hep`.

### Style

We use [Go style](https://github.com/golang/go/wiki/CodeReviewComments).

---

This _"Contributing"_ guide has been extracted from our sister project [Gonum](https://gonum.org).
Its guide is [here](https://github.com/gonum/license/blob/master/CONTRIBUTING.md).