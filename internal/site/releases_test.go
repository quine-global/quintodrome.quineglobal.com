package site

import (
	"testing"

	"maragu.dev/is"
)

func TestParseTag(t *testing.T) {
	cases := []struct {
		tag       string
		wantVer   string
		wantBuild int
	}{
		{"4.0.0", "4.0.0", 0},
		{"4.0.0-q", "4.0.0", 0},
		{"4.0.0q", "4.0.0", 0},
		{"4.0.0-q-canary.12", "4.0.0", 12},
		{"4.0.0-q-canary.1", "4.0.0", 1},
		{"1.0.0-alpha", "1.0.0-alpha", 0},
	}

	for _, c := range cases {
		t.Run(c.tag, func(t *testing.T) {
			gotVer, gotBuild := parseTag(c.tag)
			is.Equal(t, c.wantVer, gotVer)
			is.Equal(t, c.wantBuild, gotBuild)
		})
	}
}
