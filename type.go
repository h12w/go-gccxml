// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

import (
	"io"
)

type PtrKind int

/*
Type conversion rules:

	             (Num)
				 (Bool)
	             (Struct)
	|--- C.f(T p) ------------------ f(p T)
	|
	|                                           char*
	|                                           (String)
	|                                         ----------------- f(p string)
	|                                         |
	|                 In/InOut                | [1]T
	|                 (Slice)                 | (PtrWrapper)
	|               --------------- f(p []T) ------------------ f(p *T)
	|               |
	|               | Out                        [1]T
	|               | (SliceReturn)              (CPtrReturn)
	| --- C.f(T *p) ---------------- f() (p []T) --------------- f() (p T)
	|               |
	|               | void*
	|               | (PtrWrapper)
	|               --------------- f(p uintptr)
	|
	| --- C.f() *p
	| --- C.f(T **p)


	struct (ref/value type)

*/
const (
	NotSet       PtrKind = iota
	PtrArray             // including input array, and output buffer to write.
	PtrString            // like PtrArray, but type is char*
	PtrReference         // ptr to struct/union
	PtrTypedef           // Typedef a pointer
	PtrReturn            // return a value through the pointer
	PtrFunction
	PtrArrayArray  //
	PtrStringArray //
	PtrGeneral
)

type Named interface {
	Id() string
	CName() string
	File() string
	Line() string
	Location() string
}

type Type interface {
	Id() string
	Size() int
	Align() int
}

type Mangled interface {
	Mangled() string
	Demangled() string
}

// Struct or Union
type Composite interface {
	Named
	Members() string
	Bases() string
}

type Aliased interface {
	Base() Type
	Root() Type
}

func IsVoid(t Type) bool {
	if v, ok := ToFundamental(t); ok {
		return v.CName() == "void"
	}
	return false
}

func IsVoidPtr(t Type) bool {
	if pt, ok := ToPointer(t); ok {
		return IsVoid(pt.PointedType())
	}
	return false
}

func IsCString(t Type) bool {
	if pt, ok := ToPointer(t); ok {
		if ft, ok := ToFundamental(pt.PointedType()); ok {
			return ft.CName() == "char"
		}
	}
	return false
}

func IsStruct(t Type) bool {
	if at, ok := t.(Aliased); ok {
		return IsStruct(at.Root())
	}
	_, is := t.(*Struct)
	return is
}

func IsUnion(t Type) bool {
	if at, ok := t.(Aliased); ok {
		return IsUnion(at.Root())
	}
	_, is := t.(*Union)
	return is
}

func IsPointer(t Type) bool {
	_, is := ToPointer(t)
	return is
}

func ToTypedef(t Type) (*Typedef, bool) {
	td, ok := t.(*Typedef)
	return td, ok
}

func IsTypedef(t Type) bool {
	_, is := t.(*Typedef)
	return is
}

func IsFuncPtr(t Type) bool {
	if p, ok := ToPointer(t); ok {
		if _, ok := ToFuncType(p.PointedType()); ok {
			return true
		}
	}
	return false
}

func ToPointer(t Type) (*PointerType, bool) {
	switch ct := t.(type) {
	case Aliased:
		return ToPointer(ct.Root())
	case *PointerType:
		return ct, true
	}
	return nil, false
}

func ToFundamental(t Type) (*FundamentalType, bool) {
	switch ct := t.(type) {
	case Aliased:
		return ToFundamental(ct.Root())
	case *FundamentalType:
		return ct, true
	}
	return nil, false
}

func ToEnum(t Type) (*Enumeration, bool) {
	switch ct := t.(type) {
	case Aliased:
		return ToEnum(ct.Root())
	case *Enumeration:
		return ct, true
	}
	return nil, false
}

func ToComposite(t Type) (Composite, bool) {
	switch ct := t.(type) {
	case Aliased:
		return ToComposite(ct.Root())
	case Composite:
		return ct, true
	}
	return nil, false
}

func ToFuncType(t Type) (*FunctionType, bool) {
	switch ct := t.(type) {
	case Aliased:
		return ToFuncType(ct.Root())
	case *FunctionType:
		return ct, true
	}
	return nil, false
}

func IsFuncType(t Type) bool {
	_, is := ToFuncType(t)
	return is
}

func (a *Argument) CType() Type {
	return a.TypeOf(a.Type__.Type())
}

func (a *Argument) SetKind(k PtrKind) {
	a.ptrKind = k
}

func (a *Argument) IsCString() bool {
	return IsCString(a.CType())
}

func (a *Argument) PtrKind() PtrKind {
	if a.ptrKind != NotSet {
		return a.ptrKind
	}

	if IsTypedef(a.CType()) && IsPointer(a.CType()) {
		return PtrTypedef
	}
	if pt, ok := ToPointer(a.CType()); ok {
		return ptrKind(pt.PointedType())
	}
	return NotSet
}

func ptrKind(pointedType Type) PtrKind {
	switch v := pointedType.(type) {
	case Composite:
		return PtrReference
	case *FundamentalType:
		switch v.CName() {
		case "void":
			return PtrGeneral
		case "char":
			return PtrString
		default:
			return PtrReturn
		}
	case *FunctionType:
		return PtrFunction
	case *Typedef:
		return ptrKind(v.Base())
	case *CvQualifiedType:
		if v.Const() == "1" {
			k := ptrKind(v.Base())
			if k != PtrReturn {
				return k
			}
			// array kinds
			if IsCString(v.Base()) {
				return PtrStringArray
			} else if IsPointer(v.Base()) {
				return PtrArrayArray
			}
			return PtrArray
		} else {
			return PtrReturn
		}
	case *PointerType:
		return PtrReturn
	}
	return PtrGeneral
}

func (f *Field) CType() Type {
	return f.TypeOf(f.Type__.Type())
}

func rootType(v Type) Type {
	if t, ok := v.(Aliased); ok {
		return rootType(t.Base())
	}
	return v
}

func (d *Typedef) Size() int {
	return d.Root().Size()
}

func (d *Typedef) Align() int {
	return d.Root().Align()
}

func (t *Typedef) Base() Type {
	return t.TypeOf(t.Type())
}

func (t *Typedef) Root() Type {
	return rootType(t)
}

func (t *Typedef) IsFundamental() bool {
	_, is := ToFundamental(t.Root())
	return is
}

func (t *Typedef) IsEnum() bool {
	_, is := ToEnum(t.Root())
	return is
}

func (t *Typedef) IsFuncType() bool {
	_, is := ToFuncType(t.Root())
	return is
}

func (t *Typedef) IsComposite() bool {
	_, is := ToComposite(t.Root())
	return is
}

func (t *Typedef) IsPointer() bool {
	return IsPointer(t.Root())
}

func (t *Typedef) IsStructPtr() bool {
	pt, ok := t.Root().(*PointerType)
	if ok {
		return IsStruct(pt.PointedType())
	}
	return false
}

func (t *Typedef) IsUnionPtr() bool {
	pt, ok := t.Root().(*PointerType)
	if ok {
		return IsUnion(pt.PointedType())
	}
	return false
}

func (t *PointerType) PointedType() Type {
	return t.TypeOf(t.Type())
}

func (t *ArrayType) ElementType() Type {
	return t.TypeOf(t.Type())
}

func (d *CvQualifiedType) Size() int {
	return d.Root().Size()
}

func (d *CvQualifiedType) Align() int {
	return d.Root().Align()
}

func (t *CvQualifiedType) Base() Type {
	return t.TypeOf(t.Type())
}

func (t *CvQualifiedType) Root() Type {
	return rootType(t)
}

func (t *ReferenceType) Base() Type {
	return t.TypeOf(t.Type())
}

func (t *ReferenceType) Root() Type {
	return rootType(t)
}

func (v *Variable) CType() Type {
	return v.TypeOf(v.Type__.Type())
}

type CallbackInfo struct {
	CName     string // Type name of the function pointer
	CType     *FunctionType
	DataIndex int // index of void* argument in callback func type
	ArgIndex  int // index of argument of callback ptr
}

func (f *Function) HasCallback() (*CallbackInfo, bool) {
	args := f.Arguments
	var info CallbackInfo
	for i, a := range args {
		// is function pointer
		pt, isPt := ToPointer(a.CType())
		if !isPt {
			continue
		}
		ft, isFt := ToFuncType(pt.PointedType())
		if !isFt {
			continue
		}
		if i >= len(args)-1 || !IsVoidPtr(args[i+1].CType()) {
			continue
		}
		// contains an argument of void*
		dataIndex, hasDataArg := ft.dataIndex()
		if !hasDataArg {
			continue
		}
		// name of the function pointer type
		if td, ok := ToTypedef(a.CType()); ok {
			info.CName = td.CName()
		} else {
			info.CName = f.CName() + "_" + a.CName()
		}
		info.CType = ft
		info.DataIndex = dataIndex
		info.ArgIndex = i
		break
	}
	if info.CType == nil {
		return nil, false
	}
	return &info, true
}

func (f *FunctionType) Size() int {
	return 0
}

func (f *FunctionType) Align() int {
	return 0
}

func (f *FunctionType) dataIndex() (int, bool) {
	args := f.Arguments
	if len(args) > 0 {
		if IsVoidPtr(args[0].CType()) {
			return 0, true
		} else if IsVoidPtr(args[len(args)-1].CType()) {
			return len(args) - 1, true
		}
	}
	return -1, false
}

func (f *FunctionType) WriteCDecl(w io.Writer, funcName string) {
	ss := make([]string, len(f.Arguments))
	for i, a := range f.Arguments {
		ss[i] = decl(a.CType(), a.CName())
	}
	fpn(w, decl(f.ReturnType(), funcName+"("+join(ss, ", ")+")"))
	//	fpn(decl(f.ReturnType(), ""

}

func (f *FunctionType) WriteCallbackStub(w io.Writer, funcName, stubName string) {
	f.WriteCDecl(w, funcName)
	fp(w, " {")
	if IsVoid(f.ReturnType()) {
		fpn(w, "    ")
	} else {
		fpn(w, "    return ")
	}
	ss := make([]string, len(f.Arguments))
	for i, a := range f.Arguments {
		ss[i] = a.CName()
	}
	fp(w, stubName, "(", join(ss, ", "), ");")
	fp(w, "}")
}

func (d *Unimplemented) Size() int {
	return 0
}

func (d *Unimplemented) Align() int {
	return 0
}

// TODO: function type
func decl(ty Type, v string) string {
	switch t := ty.(type) {
	case *Struct:
		return sprint("struct ", t.Name_, " ", v)
	case *Union:
		return sprint("union ", t.Name_, " ", v)
	case *FundamentalType:
		return sprint(t.Name_, " ", v)
	case *Typedef:
		return sprint(t.Name_, " ", v)
	case *Enumeration:
		return sprint(t.Name_, " ", v)
	case *PointerType:
		ps := decl(t.PointedType(), "")
		if contains(ps, "[") {
			return decl(t.PointedType(), sprint("(*", v, ")"))
		}
		return sprint(ps, "*", v)
	case *ArrayType:
		return sprint(decl(t.ElementType(), ""), v, "[", t.Size(), "]")
	case *CvQualifiedType:
		return decl(t.Base(), v)
	case *ReferenceType:
		return sprint(decl(t.Base(), "&"+v))
	}
	return v
}

func (d *XmlDoc) TypeOf(id string) Type {
	for _, v := range d.Structs {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.FundamentalTypes {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.Typedefs {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.Unions {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.PointerTypes {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.Enumerations {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.FunctionTypes {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.ArrayTypes {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.CvQualifiedTypes {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.ReferenceTypes {
		if v.Id() == id {
			return v
		}
	}
	for _, v := range d.Unimplementeds {
		if v.Id() == id {
			return v
		}
	}
	return &Unimplemented{Id__: Id__{id}}
}
