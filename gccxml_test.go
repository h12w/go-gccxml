// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"testing"
)

func TestConstants(t *testing.T) {
	g := New("testdata/test_constants.h")
	ms, _ := g.Macros()
	ms = ms.Constants("GOGCCXML_")
	table := map[string]string{
		"GOGCCXML_MAJOR_VERSION": "0",
		"GOGCCXML_MINOR_VERSION": "0",
		"GOGCCXML_PATCHLEVEL":   "1",
		"GOGCCXML_42":  "(42)",
	}
	if len(table) != len(ms) {
		t.Fatalf("expected [%d] items. got [%d]\n", len(table), len(ms))
	}

	for _, m := range ms {
		if table[m.Name] != m.Body {
			t.Fatalf("expected %q for item %q. got %q\n", table[m.Name], m.Name, m.Body)
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
