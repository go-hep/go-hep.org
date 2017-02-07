// Copyright 2015 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

var (
	dirFlag = flag.String("dir", "./public", "directory of files to serve")

	goHepPkgs = []string{
		"brio",
		"csvutil",
		"fads",
		"fastjet",
		"fit",
		"fmom",
		"fwk",
		"hbook",
		"hepevt",
		"hepmc",
		"heppdt",
		"hplot",
		"lhef",
		"lhef2hepmc",
		"pawgo",
		"rio",
		"rootio",
		"sio",
		"slha",
	}
	cgoPkgs = []string{
		"croot",
	}
)

func main() {
	// redirect every http request to https
	go func() {
		for {
			err := http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))
			if err != nil {
				log.Printf("http error: %v\n", err)
			}
		}
	}()

	flag.Parse()

	dir, err := filepath.Abs(*dirFlag)
	if err != nil {
		log.Fatal(err)
	}

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Email:      "binet@cern.ch",
		HostPolicy: autocert.HostWhitelist("go-hep.org"),
		Cache:      autocert.DirCache("cert-cache"),
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.HandleFunc("/x/", goGetHandle)

	srv := http.Server{
		Addr:      ":https",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}

	err = srv.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}

func redirectToHttps(w http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/x/") {
		goGetHandle(w, req)
		return
	}

	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently,
	)
}

func goGetHandle(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	url = url[len("/x/"):]
	repo := url
	if strings.Contains(url, "/") {
		repo = filepath.Dir(url)
	}
	switch {
	case goGetHepPkg(repo):
		fmt.Fprintf(w, goGetHepTemplate, repo)
	case goGetCgoPkg(repo):
		fmt.Fprintf(w, goGetCgoTemplate, repo)
	default:
		http.NotFound(w, r)
		return
	}
}

func goGetHepPkg(pkg string) bool {
	for _, v := range goHepPkgs {
		if pkg == v {
			return true
		}
	}
	return false
}

func goGetCgoPkg(pkg string) bool {
	for _, v := range cgoPkgs {
		if pkg == v {
			return true
		}
	}
	return false
}

const goGetHepTemplate = `<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <meta name="go-import" content="go-hep.org/x/hep/%[1]s git https://github.com/go-hep/%[1]s">
  <meta name="go-source" content="go-hep.org/x/hep/%[1]s https://github.com/go-hep https://github.com/go-hep/%[1]s/tree/master{/dir} https://github.com/go-hep/%[1]s/blob/master{/dir}/{file}#L{line}">
  <meta http-equiv="refresh" content="0; url=https://godoc.org/go-hep.org/x/hep/%[1]s">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/go-hep.org/x/hep/%[1]s">move along</a>.
</body>
</html>
`

const goGetCgoTemplate = `<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <meta name="go-import" content="go-hep.org/x/cgo/%[1]s git https://github.com/go-hep/%[1]s">
  <meta name="go-source" content="go-hep.org/x/cgo/%[1]s https://github.com/go-hep https://github.com/go-hep/%[1]s/tree/master{/dir} https://github.com/go-hep/%[1]s/blob/master{/dir}/{file}#L{line}">
  <meta http-equiv="refresh" content="0; url=https://godoc.org/go-hep.org/x/cgo/%[1]s">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/go-hep.org/x/cgo/%[1]s">move along</a>.
</body>
</html>
`
