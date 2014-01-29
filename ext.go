// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"os"
	"strings"
)

type xmlDoc__ struct {
	*XmlDoc
}

func (v *xmlDoc__) Doc() *XmlDoc {
	return v.XmlDoc
}

func (d *XmlDoc) Print() error {
	return New(d.file).Save(os.Stdout)
}

func (d *XmlDoc) FindStruct(id string) *Struct {
	for _, v := range d.Structs {
		if v.Id() == id {
			return v
		}
	}
	return nil
}

func (d *XmlDoc) FindUnion(id string) *Union {
	for _, v := range d.Unions {
		if v.Id() == id {
			return v
		}
	}
	return nil
}

func (d *XmlDoc) FindField(id string) *Field {
	for _, v := range d.Fields {
		if v.Id() == id {
			return v
		}
	}
	return nil
}

type Ellipsis struct {
}

type Ellipses []*Ellipsis

func (f *Function) ReturnType() Type {
	return f.TypeOf(f.Returns())
}

func (f *FunctionType) ReturnType() Type {
	return f.TypeOf(f.Returns())
}

func (s *Struct) Fields() []*Field {
	return s.getFields(s)
}

func (s *Union) Fields() []*Field {
	return s.getFields(s)
}

func (d *XmlDoc) getFields(s Composite) (fields []*Field) {
	for _, id := range strings.Split(s.Members(), " ") {
		f := d.FindField(id)
		if f != nil { // not member functions ...
			fields = append(fields, f)
		}
	}
	return
}

func demangledName(t Mangled) string {
	return strings.Replace(t.Demangled(), "::.", "_", -1)
}

func (a *ArrayType) Len() int {
	return a.Size() / a.ElementType().Size()
}

func (v *EnumValue) Id() string {
	return sprint(v.parent.Id(), "_", v.serial)
}

func (v *EnumValue) File() string {
	return v.parent.File()
}

func (v *EnumValue) Line() string {
	return v.parent.Line()
}

func (v *EnumValue) Location() string {
	return v.parent.Location()
}

func (v *Argument) Id() string {
	return sprint(v.parent.Id(), "_", v.serial)
}

func (v *Argument) CName() string {
	if v.Name_ == "" {
		return sprint("v_", v.serial)
	}
	return v.Name_
}

func (s *Struct) CName() string {
	if s.Name_ != "" {
		return s.Name_
	}
	return demangledName(s)
}

func (s *Union) CName() string {
	if s.Name_ != "" {
		return s.Name_
	}
	return demangledName(s)
}
