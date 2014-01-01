// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"sort"
)

func (d *XmlDoc) prepare() {
	d.setDoc()
	d.prepareTypes()
}

func (d *XmlDoc) prepareTypes() {
	for _, v := range d.Functions {
		for i, a := range v.Arguments {
			a.parent = v
			a.serial = i
		}
	}
	for _, v := range d.FunctionTypes {
		for i, a := range v.Arguments {
			a.serial = i
		}
	}
	for _, v := range d.Enumerations {
		for i, a := range v.EnumValues {
			a.parent = v
			a.serial = i
		}
	}
	d.findAlias()
	d.sort()
}

func (d *XmlDoc) findAlias() {
	m := NewSSet()
	for _, td := range d.Typedefs {
		if !m.Has(td.Id()) {
			var t Type = td
			for {
				m.Add(t.Id())
				if a, ok := t.(Aliased); ok {
					t = a.Base()
				} else {
					break
				}
			}
			switch s := t.(type) {
			case *Struct:
				s.aliasName = td.CName()
			case *Union:
				s.aliasName = td.CName()
			}
		}
	}
}

func (d *XmlDoc) sort() {
	sort.Sort(d.Enumerations)
	sort.Sort(d.Structs)
	sort.Sort(d.Typedefs)
	sort.Sort(d.Functions)
}

func (v *xmlDoc__) setXmlDoc(d *XmlDoc) {
	v.XmlDoc = d
}

func (d *XmlDoc) setDoc() {
	for _, v := range d.Namespaces {
		v.setXmlDoc(d)
	}
	for _, v := range d.Functions {
		v.setXmlDoc(d)
		for _, a := range v.Arguments {
			a.setXmlDoc(d)
		}
	}
	for _, v := range d.Structs {
		v.setXmlDoc(d)
	}
	for _, v := range d.FundamentalTypes {
		v.setXmlDoc(d)
	}
	for _, v := range d.Typedefs {
		v.setXmlDoc(d)
	}
	for _, v := range d.Unions {
		v.setXmlDoc(d)
	}
	for _, v := range d.Variables {
		v.setXmlDoc(d)
	}
	for _, v := range d.PointerTypes {
		v.setXmlDoc(d)
	}
	for _, v := range d.Enumerations {
		v.setXmlDoc(d)
		for _, a := range v.EnumValues {
			a.setXmlDoc(d)
		}
	}
	for _, v := range d.FunctionTypes {
		v.setXmlDoc(d)
		for _, a := range v.Arguments {
			a.setXmlDoc(d)
		}
	}
	for _, v := range d.ArrayTypes {
		v.setXmlDoc(d)
	}
	for _, v := range d.Fields {
		v.setXmlDoc(d)
	}
	for _, v := range d.Destructors {
		v.setXmlDoc(d)
	}
	for _, v := range d.OperatorMethods {
		v.setXmlDoc(d)
	}
	for _, v := range d.Constructors {
		v.setXmlDoc(d)
	}
	for _, v := range d.CvQualifiedTypes {
		v.setXmlDoc(d)
	}
	for _, v := range d.ReferenceTypes {
		v.setXmlDoc(d)
	}
	for _, v := range d.Files {
		v.setXmlDoc(d)
	}
}
