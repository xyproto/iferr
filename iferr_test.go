package iferr

import (
	"strings"
	"testing"
)

func iferrStr(in string, pos int) (string, error) {
	return IfErr([]byte(in), pos)
}

func iferrOK(t *testing.T, fn string, off int, exp string) {
	const (
		fnPre   = "package main\nfunc foo() "
		fnPost  = " {}"
		actPre  = "if err != nil {\n\treturn "
		actPost = "\n}\n"
	)

	act, err := iferrStr(fnPre+fn, len(fnPre)+1+off)
	if err != nil {
		t.Errorf("iferr() is failed: %s for %q", err, fn)
		return
	}
	if !strings.HasPrefix(act, actPre) || !strings.HasSuffix(act, actPost) {
		t.Errorf("iferr() returns with unexpected prefix or suffix: %q", act)
		return
	}
	act = act[len(actPre) : len(act)-len(actPost)]
	if act != exp {
		t.Errorf("iferr() returns unexpected: actual=%q expect=%q", act, exp)
		return
	}
}

func TestIferr(t *testing.T) {
	iferrOK(t, `(interface{}, error)`, 0, `nil, err`)
	iferrOK(t, `(map[string]struct{}, error)`, 0, `nil, err`)
	iferrOK(t, `(chan bool, error)`, 0, `nil, err`)
	iferrOK(t, `(bool, error)`, 0, `false, err`)
	iferrOK(t, `(foo, error)`, 0, `foo{}, err`)
	iferrOK(t, `(*foo, error)`, 0, `nil, err`)
}
