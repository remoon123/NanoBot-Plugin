// Package main generates banner.go
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const banner = `package kanban

// Version ...
var Version = "%s"

// Banner ...
var Banner = "* QQ + NanoBot + Golang\n" +
	"* Version " + Version + " - %s\n" +
	"* Copyright © 2023 - %d FloatTech. All Rights Reserved.\n" +
	"* Project: https://github.com/FloatTech/NanoBot-Plugin"
`

const timeformat = `2006-01-02 15:04:05 +0800 CST`

func main() {
	f, err := os.Create("banner.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	vartag := bytes.NewBuffer(nil)
	vartagcmd := exec.Command("git", "tag", "--sort=committerdate")
	vartagcmd.Stdout = vartag
	err = vartagcmd.Run()
	if err != nil {
		panic(err)
	}
	s := strings.Split(vartag.String(), "\n")
	now := time.Now()
	_, err = fmt.Fprintf(f, banner, s[len(s)-2], now.Format(timeformat), now.Year())
	if err != nil {
		panic(err)
	}
}
