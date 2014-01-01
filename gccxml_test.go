// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"testing"
)

func TestIt(t *testing.T) {
	g := Xml{"/usr/local/include/SDL2/SDL.h"}
	ms, _ := g.Macros()
	ms = ms.Constants("SDL_")
	for _, m := range ms {
		p(m.Name, m.Body)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
