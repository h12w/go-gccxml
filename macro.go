// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strings"
)

var (
	reMacro = regexp.MustCompile(`#define (\S+) (.*)`)
)

type Macro struct {
	Name string
	Body string
}

type Macros []Macro

func DecodeMacros(r io.Reader) (ms Macros, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m := reMacro.FindStringSubmatch(scanner.Text())
		if len(m) == 3 {
			ms = append(ms, Macro{m[1], strings.TrimSpace(m[2])})
		}
	}
	err = scanner.Err()
	sort.Sort(ms)
	return
}

func (m Macro) HasPrefix(prefix string) bool {
	return strings.HasPrefix(m.Name, prefix)
}

func (m Macro) NoArg() bool {
	return !strings.Contains(m.Name, "(")
}

func (m Macro) IsString() bool {
	return len(m.Body) >= 2 && m.Body[0] == '"' && m.Body[len(m.Body)-1] == '"'
}

func (m Macro) ContainsNum() bool {
	for _, c := range m.Body {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func (m Macro) ContainsOthers(ms Macros) bool {
	for _, o := range ms {
		if strings.Contains(m.Body, o.Name) {
			return true
		}
	}
	return false
}

func (ms Macros) Constants(prefix string) (fms Macros) {
	for _, m := range ms {
		if m.HasPrefix(prefix) &&
			m.NoArg() &&
			(m.IsString() || m.ContainsNum() || m.ContainsOthers(ms)) {
			fms = append(fms, m)
		}
	}
	return
}

func (s Macros) Len() int {
	return len(s)
}

func (s Macros) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s Macros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
