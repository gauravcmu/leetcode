package main

import (
	"fmt"
	"io/ioutil"

	"testing"
)
var fileinput  string = `DEPEND A B C
DEPEND B E
DEPEND E F
DEPEND C B
INSTALL A
LIST
INSTALL B
LIST
END`

func TestDependencyBuild(t *testing.T) {
	tempfile, err := ioutil.TempFile(".", "tdd")
	if err != nil {
		t.Fatalf("failed to create tempfile:%v", err)
	}
	fmt.Printf("tempfile: %v", tempfile)
	tempfile.WriteString(fileinput)

	dep, installed, err := buildDependency(tempfile.Name())
	t.Logf("%v %v %v", dep, installed, err)
	err  = tempfile.Close()
	if err != nil {
		t.Logf("failed to delete tempfile:%v", err)
	}
}
