// Copyright 2015 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// deploy the latest version of the go-hep.org web site.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	ssh       = "/usr/bin/ssh"
	server    = "root@clrwebgohep.in2p3.fr"
	service   = "go-hep.service"
	remoteDir = "/srv/go-hep.org"
)

func main() {

	run("go", "generate")
	run("go", "build", "-tags", "netgo", "-v", "-o", "go-hep-serve", ".")
	run("hugo")
	run("scp", "-r", "./public", server+":"+remoteDir+"/.")

	srv := remote{server}
	srv.run("systemctl", "stop", service)
	run("scp", "./go-hep-serve", server+":"+remoteDir+"/.")
	srv.run("cd " + remoteDir + " && git pull")
	srv.run("systemctl", "restart", service)
	os.Remove("go-hep-serve")
}

func run(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	fmt.Printf("+ %s\n", strings.Join(c.Args, " "))
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}

type remote struct {
	server string
}

func (r *remote) run(cmd string, args ...string) {
	rargs := []string{r.server, "--", cmd}
	rargs = append(rargs, args...)

	c := exec.Command(ssh, rargs...)
	fmt.Printf("+ %s\n", strings.Join(c.Args, " "))
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
