// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

func (s Enumerations) Len() int {
	return len(s)
}

func (s Enumerations) Less(i, j int) bool {
	return s[i].CName() < s[j].CName()
}

func (s Enumerations) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Structs) Len() int {
	return len(s)
}

func (s Structs) Less(i, j int) bool {
	return s[i].CName() < s[j].CName()
}

func (s Structs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Typedefs) Len() int {
	return len(s)
}

func (s Typedefs) Less(i, j int) bool {
	return s[i].CName() < s[j].CName()
}

func (s Typedefs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Functions) Len() int {
	return len(s)
}

func (s Functions) Less(i, j int) bool {
	return s[i].CName() < s[j].CName()
}

func (s Functions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
