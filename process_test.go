package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/juntaki/pp"
	"golang.org/x/tools/imports"
)

func Test(t *testing.T) {
	filename := "./test.txt"

	f, err := os.Open(filename)
	defer f.Close()
	in := f

	src, err := ioutil.ReadAll(in)

	target := filename
	res, err := imports.Process(target, src, nil)
	pp.Println(err)
	println(string(res))
}
