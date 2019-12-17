package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	for _, c := range []struct {
		In  []string
		Out string
		Err string
	}{
		{In: []string{}, Out: "", Err: "usage: cidr CIDR"},
		{In: []string{""}, Out: "", Err: "usage: cidr CIDR"},
		{In: []string{"", "192.0.2.0"}, Out: "", Err: "invalid CIDR address: 192.0.2.0"},
		{In: []string{"", "192.0.2.0/24"}, Out: "192.0.2.0 - 192.0.2.255", Err: ""},
	} {
		t.Run(fmt.Sprint(c.In), func(t *testing.T) {
			var o, e bytes.Buffer

			run(c.In, &o, &e)

			out := strings.TrimSpace(o.String())
			err := strings.TrimSpace(e.String())
			if out != c.Out {
				t.Errorf("out: got %q, expected %q", out, c.Out)
			}
			if err != c.Err {
				t.Errorf("err: got %q, expected %q", err, c.Err)
			}
		})
	}

}

func TestCalcValid(t *testing.T) {
	for cidr, expected := range map[string]string{
		"192.0.2.0/24":   "192.0.2.0 - 192.0.2.255",
		"192.0.2.1/24":   "192.0.2.0 - 192.0.2.255",
		"192.0.2.137/24": "192.0.2.0 - 192.0.2.255",
		"192.0.2.255/24": "192.0.2.0 - 192.0.2.255",
		"192.0.2.0/25":   "192.0.2.0 - 192.0.2.127",
		"192.0.2.1/25":   "192.0.2.0 - 192.0.2.127",
		"192.0.2.127/25": "192.0.2.0 - 192.0.2.127",
		"192.0.2.128/25": "192.0.2.128 - 192.0.2.255",
	} {
		actual, err := calc(cidr)
		if err != nil {
			t.Fatal(err)
		}

		if actual != expected {
			t.Errorf("got %q, expected %q", actual, expected)
		}
	}
}

func TestCalcInvalid(t *testing.T) {
	for _, cidr := range []string{
		"192.0.2.0",
		"192.0.2.0/256",
		"192.0.2.0/-1",
	} {
		actual, err := calc(cidr)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if actual != "" {
			t.Errorf("got unexpected result: %q", actual)
		}
	}
}
