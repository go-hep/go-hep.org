// Copyright 2015 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

var (
	dirFlag  = flag.String("dir", "./public", "directory of files to serve")
	distFlag = flag.String("dist-dir", "./dist", "directory of distribution files to serve")

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
		"hep/xrootd",
	}
	expPkgs = []string{
		"exp",
		"exp/vgshiny",
	}
	cgoPkgs = []string{
		"cgo",
		"cgo/croot",
	}
)

func main() {
	flag.Parse()

	dir, err := filepath.Abs(*dirFlag)
	if err != nil {
		log.Fatal(err)
	}

	distDir, err := filepath.Abs(*distFlag)
	if err != nil {
		log.Fatal(err)
	}

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Email:      "binet@cern.ch",
		HostPolicy: autocert.HostWhitelist("go-hep.org", "www.go-hep.org"),
		Cache:      autocert.DirCache("cert-cache"),
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(dir)))
	mux.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir(distDir))))
	mux.HandleFunc("/x/", goGetHandle)
	mux.HandleFunc("/commit/", commitHandle)

	go http.ListenAndServe(":http", m.HTTPHandler(http.HandlerFunc(redirectToHttps)))
	srv := http.Server{
		Addr:           ":https",
		Handler:        m.HTTPHandler(mux),
		TLSConfig:      m.TLSConfig(),
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 30,
		IdleTimeout:    time.Minute * 5,
		MaxHeaderBytes: 1 << 20,
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

func commitHandle(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	sha := url[len("/commit/"):]
	data := struct {
		Commit string
		Repo   string
	}{sha, "hep"}
	commitTemplate.Execute(w, data)
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
	case goGetExpPkg(repo):
		data.Repo = "exp"
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

func goGetExpPkg(pkg string) bool {
	for _, v := range expPkgs {
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
 <meta name="go-import" content="go-hep.org/x/{{.Repo}} git https://codeberg.org/go-hep/{{.Repo}}"/>
 <meta name="go-source" content="go-hep.org/x/{{.Repo}} https://codeberg.org/go-hep/{{.Repo}}/ https://codeberg.org/go-hep/{{.Repo}}/src/branch/main{/dir} https://codeberg.org/go-hep/{{.Repo}}/src/branch/main{/dir}/{file}#L{line}"/>
 <meta http-equiv="refresh" content="0; url=https://pkg.go.dev/go-hep.org/x/{{.Pkg}}"/>
</head>
<body>
Nothing to see here; <a href="https://pkg.go.dev/go-hep.org/x/{{.Pkg}}">move along</a>.
</body>
</html>
`))

var goGetExpTemplate = template.Must(template.New("x/exp").Parse(`<!DOCTYPE html>
<html>
<head>
 <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
 <meta name="go-import" content="go-hep.org/x/{{.Repo}} git https://codeberg.org/go-hep/{{.Repo}}"/>
 <meta name="go-source" content="go-hep.org/x/{{.Repo}} https://codeberg.org/go-hep/{{.Repo}}/ https://codeberg.org/go-hep/{{.Repo}}/src/branch/main{/dir} https://codeberg.org/go-hep/{{.Repo}}/src/branch/main{/dir}/{file}#L{line}"/>
 <meta http-equiv="refresh" content="0; url=https://pkg.go.dev/go-hep.org/x/{{.Pkg}}"/>
</head>
<body>
Nothing to see here; <a href="https://pkg.go.dev/go-hep.org/x/{{.Pkg}}">move along</a>.
</body>
</html>
`))

var goGetCgoTemplate = template.Must(template.New("x/cgo").Parse(`<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
  <meta name="go-import" content="go-hep.org/x/cgo/{{.Repo}} git https://codeberg.org/go-hep/{{.Repo}}"/>
  <meta name="go-source" content="go-hep.org/x/cgo/{{.Repo}} https://codeberg.org/go-hep/{{.Repo}}/ https://codeberg.org/go-hep/{{.Repo}}/src/branch/main{/dir} https://codeberg.org/go-hep/{{.Repo}}/src/branch/main{/dir}/{file}#L{line}"/>
  <meta http-equiv="refresh" content="0; url=https://pkg.go.dev/go-hep.org/x/cgo/{{.Pkg}}"/>
</head>
<body>
Nothing to see here; <a href="https://pkg.go.dev/go-hep.org/x/cgo/{{.Pkg}}">move along</a>.
</body>
</html>
`))

var commitTemplate = template.Must(template.New("commit").Parse(`<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="content-type" content="text/html; charset=utf-8"/>
  <meta http-equiv="refresh" content="0; url=https://codeberg.org/go-hep/{{.Repo}}/commit/{{.Commit}}"/>
</head>
<body>
Nothing to see here; <a href="https://codeberg.org/go-hep/{{.Repo}}/commit/{{.Commit}}">move along</a>.
</body>
</html>
`))
