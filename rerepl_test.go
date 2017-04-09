package main

import (
	"strings"
	"testing"
)

func TestEvalLinet(t *testing.T) {
	// success, but no captures
	res, err := EvalLine("^foo$ foo")
	if err != nil {
		t.Errorf("error found: %s", err)
	}
	if res == nil {
		t.Errorf("result expected, but nil")
	}
	if !res.matched || len(res.captures) != 0 {
		t.Errorf("invalid result: %s", res)
	}

	// success and captures
	res, err = EvalLine("^([a-z]{3})\\sand\\s([a-z]{3})$ foo and bar")
	if err != nil {
		t.Errorf("error found: %s", err)
	}
	if res == nil {
		t.Errorf("result expected, but nil")
	}
	if !res.matched || len(res.captures) != 2 || res.captures[0] != "foo" || res.captures[1] != "bar" {
		t.Errorf("invalid result: %s", res)
	}

	// invalid input
	res, err = EvalLine("foo")
	if err == nil || !strings.Contains(err.Error(), "invalid input") {
		t.Errorf("error not found: %s", err)
	}

	// invalid regexp
	res, err = EvalLine("(foo foo")
	if err == nil || !strings.Contains(err.Error(), "invalid regexp") {
		t.Errorf("error not found: %s", err)
	}
}
