// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

/*
#include <wchar.h>
*/
import "C"

import (
	"reflect"
)

func init() {
	initTypeMap()
}

// According to http://golang.org/cmd/cgo.
var gccCgoNumMap = map[string]string{
	"char":                   "C.char",
	"signed char":            "C.schar",
	"unsigned char":          "C.uchar",
	"short int":              "C.short",
	"short unsigned int":     "C.ushort",
	"wchar_t":                "C.wchar_t",
	"int":                    "C.int",
	"unsigned int":           "C.uint",
	"long int":               "C.long",
	"long unsigned int":      "C.ulong",
	"long long int":          "C.longlong",
	"long long unsigned int": "C.ulonglong",
	"float":                  "C.float",
	"double":                 "C.double",
	"complex float":          "complex64",
	"complex double":         "complex128",
}

type NumKind int

const (
	SignedInt NumKind = iota
	UnsignedInt
	Float
	Complex
)

type NumInfo struct {
	Kind NumKind
	Bits int
}

var cgoNumMap = map[string]NumInfo{}

// Initialize GoTypes in NumTypeMap
func initTypeMap() {
	cgoNumMap["C.char"] = GetNumInfo(C.char(0))
	cgoNumMap["C.schar"] = GetNumInfo(C.schar(0))
	cgoNumMap["C.uchar"] = GetNumInfo(C.uchar(0))
	cgoNumMap["C.short"] = GetNumInfo(C.short(0))
	cgoNumMap["C.ushort"] = GetNumInfo(C.ushort(0))
	cgoNumMap["C.wchar_t"] = GetNumInfo(C.wchar_t(0))
	cgoNumMap["C.int"] = GetNumInfo(C.int(0))
	cgoNumMap["C.uint"] = GetNumInfo(C.uint(0))
	cgoNumMap["C.long"] = GetNumInfo(C.long(0))
	cgoNumMap["C.ulong"] = GetNumInfo(C.ulong(0))
	cgoNumMap["C.longlong"] = GetNumInfo(C.longlong(0))
	cgoNumMap["C.ulonglong"] = GetNumInfo(C.ulonglong(0))
	cgoNumMap["C.float"] = GetNumInfo(C.float(0))
	cgoNumMap["C.double"] = GetNumInfo(C.double(0))
	cgoNumMap["complex64"] = GetNumInfo(complex64(0))
	cgoNumMap["complex128"] = GetNumInfo(complex128(0))
}

func NumInfoFromGccName(gccName string) NumInfo {
	return cgoNumMap[gccCgoNumMap[gccName]]
}

func NumCgoNameFromGccName(gccName string) string {
	return gccCgoNumMap[gccName]
}

func GetNumInfo(v interface{}) (t NumInfo) {
	t.Bits = reflect.TypeOf(v).Bits()
	switch v.(type) {
	case float32, float64, C.float, C.double:
		t.Kind = Float
	case complex64, complex128:
		t.Kind = Complex
	default:
		if sprint(reflect.ValueOf(-1).Convert(reflect.TypeOf(v)).Interface())[0] == '-' {
			t.Kind = SignedInt
		} else {
			t.Kind = UnsignedInt
		}
	}
	return
}
