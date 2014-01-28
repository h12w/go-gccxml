// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"fmt"
	"io"
	"strings"
)

func p(v ...interface{}) {
	fmt.Println(v...)
}

func sprint(v ...interface{}) string {
	return fmt.Sprint(v...)
}

type SSet struct {
	m map[string]struct{}
}

func NewSSet() SSet {
	return SSet{make(map[string]struct{})}
}

func (m *SSet) IsNil() bool {
	return m.m == nil
}

func (m *SSet) Add(ss ...string) {
	for _, s := range ss {
		m.m[s] = struct{}{}
	}
}

func (m *SSet) Del(s string) {
	delete(m.m, s)
}

func (m *SSet) Has(s string) bool {
	if m.m == nil {
		return false
	}
	_, has := m.m[s]
	return has
}

func (m *SSet) Slice() []string {
	ss := make([]string, 0, len(m.m))
	for s := range m.m {
		ss = append(ss, s)
	}
	return ss
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func fp(w io.Writer, v ...interface{}) {
	fmt.Fprint(w, v...)
	fmt.Fprintln(w)
}

func fpn(w io.Writer, v ...interface{}) {
	fmt.Fprint(w, v...)
}

func join(ss []string, sep string) string {
	return strings.Join(ss, sep)
}

func joins(ss ...string) string {
	return strings.Join(ss, " ")
}

func concat(ss ...string) string {
	return strings.Join(ss, "")
}
