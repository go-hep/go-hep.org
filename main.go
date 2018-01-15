// Copyright 2015 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/tls"
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

var (
	dirFlag = flag.String("dir", "./public", "directory of files to serve")

	goHepPkgs = []string{
		"hep",
		"hep/brio",
		"hep/csvutil",
		"hep/fads",
		"hep/fastjet",
		"hep/fit",
		"hep/fmom",
		"hep/fwk",
		"hep/hbook",
		"hep/hepevt",
		"hep/hepmc",
		"hep/heppdt",
		"hep/hplot",
		"hep/lhef",
		"hep/lhef2hepmc",
		"hep/pawgo",
		"hep/rio",
		"hep/rootio",
		"hep/sio",
		"hep/slha",
	}
	cgoPkgs = []string{
		"cgo",
		"cgo/croot",
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

	go http.ListenAndServe(":http", m.HTTPHandler(nil))
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
	var data = struct {
		Pkg  string
		Repo string
	}{
		Pkg:  repo,
		Repo: repo,
	}
	switch {
	case goGetHepPkg(repo):
		data.Repo = "hep"
		goGetHepTemplate.Execute(w, data)
		return
	case goGetCgoPkg(repo):
		repo = repo[len("cgo/"):]
		data.Pkg = repo
		data.Repo = repo
		goGetCgoTemplate.Execute(w, data)
		return
	default:
		http.NotFound(w, r)
		return
	}
}

func goGetHepPkg(pkg string) bool {
	for _, v := range goHepPkgs {
		if strings.HasPrefix(pkg, v) {
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

var goGetHepTemplate = template.Must(template.New("x/hep").Parse(`<!DOCTYPE html>
<html>
<head>
 <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
 <meta name="go-import" content="go-hep.org/x/{{.Repo}} git https://github.com/go-hep/{{.Repo}}"/>
 <meta name="go-source" content="go-hep.org/x/{{.Repo}} https://github.com/go-hep/{{.Repo}}/ https://github.com/go-hep/{{.Repo}}/tree/master{/dir} https://github.com/go-hep/{{.Repo}}/blob/master{/dir}/{file}#L{line}"/>
 <meta http-equiv="refresh" content="0; url=https://godoc.org/go-hep.org/x/{{.Pkg}}"/>
</head>
<body>
Nothing to see here; <a href="https://godoc.org/go-hep.org/x/{{.Pkg}}">move along</a>.
</body>
</html>
`))

var goGetCgoTemplate = template.Must(template.New("x/cgo").Parse(`<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
  <meta name="go-import" content="go-hep.org/x/cgo/{{.Repo}} git https://github.com/go-hep/{{.Repo}}"/>
  <meta name="go-source" content="go-hep.org/x/cgo/{{.Repo}} https://github.com/go-hep/{{.Repo}}/ https://github.com/go-hep/{{.Repo}}/tree/master{/dir} https://github.com/go-hep/{{.Repo}}/blob/master{/dir}/{file}#L{line}"/>
  <meta http-equiv="refresh" content="0; url=https://godoc.org/go-hep.org/x/cgo/{{.Pkg}}"/>
</head>
<body>
Nothing to see here; <a href="https://godoc.org/go-hep.org/x/cgo/{{.Pkg}}">move along</a>.
</body>
</html>
`))
