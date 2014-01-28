// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccxml

type Argument struct {
	Name__
	Type__
	Location__
	File__
	Line__
	xmlDoc__

	ptrKind PtrKind
	parent  *Function
	serial  int
}
type Arguments []*Argument

type ArrayType struct {
	Id__
	Min__
	Max__
	Type__
	Size__
	Align__
	xmlDoc__
}
type ArrayTypes []*ArrayType

type Constructor struct {
	Id__
	Name__
	Artificial__
	Throw__
	Context__
	Access__
	Mangled__
	Demangled__
	Location__
	File__
	Line__
	Endline__
	Inline__
	xmlDoc__
}
type Constructors []*Constructor

type CvQualifiedType struct {
	Id__
	Type__
	Restrict__
	Const__
	Volatile__
	xmlDoc__
}
type CvQualifiedTypes []*CvQualifiedType

type Destructor struct {
	Id__
	Name__
	Artificial__
	Throw__
	Context__
	Access__
	Mangled__
	Demangled__
	Location__
	File__
	Line__
	Endline__
	Inline__
	xmlDoc__
}
type Destructors []*Destructor

type EnumValue struct {
	Name__
	Init__
	xmlDoc__

	parent *Enumeration
	serial int
}
type EnumValues []*EnumValue

type Enumeration struct {
	Id__
	Name__
	Context__
	Location__
	File__
	Line__
	Size__
	Align__
	Artificial__
	xmlDoc__
	EnumValues EnumValues `xml:"EnumValue"`
}
type Enumerations []*Enumeration

type Field struct {
	Id__
	Name__
	Type__
	Offset__
	Context__
	Access__
	Location__
	File__
	Line__
	Bits__
	xmlDoc__
}
type Fields []*Field

type File struct {
	Id__
	Name__
	xmlDoc__
}
type Files []*File

type Function struct {
	Id__
	Name__
	Returns__
	Throw__
	Context__
	Mangled__
	Location__
	File__
	Line__
	Extern__
	Attributes__
	Demangled__
	Endline__
	Inline__
	Static__
	Arguments Arguments `xml:"Argument"`
	Ellipses  Ellipses  `xml:"Ellipsis"`
	xmlDoc__
}
type Functions []*Function

type FunctionType struct {
	Id__
	Arguments Arguments `xml:"Argument"`
	Returns__
	xmlDoc__
}
type FunctionTypes []*FunctionType

type FundamentalType struct {
	Id__
	Name__
	Size__
	Align__
	xmlDoc__
}
type FundamentalTypes []*FundamentalType

type XmlDoc struct {
	CvsRevision__
	Namespaces       Namespaces       `xml:"Namespace"`
	Functions        Functions        `xml:"Function"`
	Structs          Structs          `xml:"Struct"`
	Typedefs         Typedefs         `xml:"Typedef"`
	FundamentalTypes FundamentalTypes `xml:"FundamentalType"`
	Enumerations     Enumerations     `xml:"Enumeration"`
	Unions           Unions           `xml:"Union"`
	Variables        Variables        `xml:"Variable"`
	PointerTypes     PointerTypes     `xml:"PointerType"`
	FunctionTypes    FunctionTypes    `xml:"FunctionType"`
	ArrayTypes       ArrayTypes       `xml:"ArrayType"`
	Unimplementeds   Unimplementeds   `xml:"Unimplemented"`
	Fields           Fields           `xml:"Field"`
	Destructors      Destructors      `xml:"Destructor"`
	OperatorMethods  OperatorMethods  `xml:"OperatorMethod"`
	Constructors     Constructors     `xml:"Constructor"`
	CvQualifiedTypes CvQualifiedTypes `xml:"CvQualifiedType"`
	ReferenceTypes   ReferenceTypes   `xml:"ReferenceType"`
	Files            Files            `xml:"File"`

	file string
}

type Namespace struct {
	Id__
	Name__
	Members__
	Mangled__
	Demangled__
	Context__
	xmlDoc__
}
type Namespaces []*Namespace

type OperatorMethod struct {
	Id__
	Name__
	Returns__
	Artificial__
	Throw__
	Context__
	Access__
	Mangled__
	Demangled__
	Location__
	File__
	Line__
	Endline__
	Inline__
	xmlDoc__
}
type OperatorMethods []*OperatorMethod

type PointerType struct {
	Id__
	Type__
	Size__
	Align__
	xmlDoc__
}
type PointerTypes []*PointerType

type ReferenceType struct {
	Id__
	Type__
	Size__
	Align__
	xmlDoc__
}
type ReferenceTypes []*ReferenceType

type Struct struct {
	Id__
	Name__
	Context__
	Mangled__
	Demangled__
	Location__
	File__
	Line__
	Artificial__
	Size__
	Align__
	Members__
	Bases__
	Incomplete__
	Attributes__
	Access__
	xmlDoc__
}
type Structs []*Struct

type Typedef struct {
	Id__
	Name__
	Type__
	Context__
	Location__
	File__
	Line__
	xmlDoc__
}
type Typedefs []*Typedef

type Unimplemented struct {
	Id__
	TreeCode__
	TreeCodeName__
	Node__
	xmlDoc__
}
type Unimplementeds []*Unimplemented

type Union struct {
	Id__
	Name__
	Context__
	Mangled__
	Demangled__
	Location__
	File__
	Line__
	Size__
	Align__
	Members__
	Bases__
	Artificial__
	Access__
	xmlDoc__
}
type Unions []*Union

type Variable struct {
	Id__
	Name__
	Type__
	Context__
	Location__
	File__
	Line__
	Extern__
	xmlDoc__
}
type Variables []*Variable

type Access__ struct {
	Access_ string `xml:"access,attr"`
}

func (s Access__) Access() string {
	return s.Access_
}

type Align__ struct {
	Align_ int `xml:"align,attr"`
}

func (s Align__) Align() int {
	return s.Align_
}

type Artificial__ struct {
	Artificial_ string `xml:"artificial,attr"`
}

func (s Artificial__) Artificial() string {
	return s.Artificial_
}

type Attributes__ struct {
	Attributes_ string `xml:"attributes,attr"`
}

func (s Attributes__) Attributes() string {
	return s.Attributes_
}

type Bases__ struct {
	Bases_ string `xml:"bases,attr"`
}

func (s Bases__) Bases() string {
	return s.Bases_
}

type Bits__ struct {
	Bits_ string `xml:"bits,attr"`
}

func (s Bits__) Bits() string {
	return s.Bits_
}

type Const__ struct {
	Const_ string `xml:"const,attr"`
}

func (s Const__) Const() string {
	return s.Const_
}

type Context__ struct {
	Context_ string `xml:"context,attr"`
}

func (s Context__) Context() string {
	return s.Context_
}

type CvsRevision__ struct {
	CvsRevision_ string `xml:"cvs_revision,attr"`
}

func (s CvsRevision__) CvsRevision() string {
	return s.CvsRevision_
}

type Demangled__ struct {
	Demangled_ string `xml:"demangled,attr"`
}

func (s Demangled__) Demangled() string {
	return s.Demangled_
}

type Endline__ struct {
	Endline_ string `xml:"endline,attr"`
}

func (s Endline__) Endline() string {
	return s.Endline_
}

type Extern__ struct {
	Extern_ string `xml:"extern,attr"`
}

func (s Extern__) Extern() string {
	return s.Extern_
}

type File__ struct {
	File_ string `xml:"file,attr"`
}

func (s File__) File() string {
	return s.File_
}

type Id__ struct {
	Id_ string `xml:"id,attr"`
}

func (s Id__) Id() string {
	return s.Id_
}

type Incomplete__ struct {
	Incomplete_ string `xml:"incomplete,attr"`
}

func (s Incomplete__) Incomplete() string {
	return s.Incomplete_
}

type Init__ struct {
	Init_ int `xml:"init,attr"`
}

func (s Init__) Init() int {
	return s.Init_
}

type Inline__ struct {
	Inline_ string `xml:"inline,attr"`
}

func (s Inline__) Inline() string {
	return s.Inline_
}

type Line__ struct {
	Line_ string `xml:"line,attr"`
}

func (s Line__) Line() string {
	return s.Line_
}

type Location__ struct {
	Location_ string `xml:"location,attr"`
}

func (s Location__) Location() string {
	return s.Location_
}

type Mangled__ struct {
	Mangled_ string `xml:"mangled,attr"`
}

func (s Mangled__) Mangled() string {
	return s.Mangled_
}

type Max__ struct {
	Max_ string `xml:"max,attr"`
}

func (s Max__) Max() string {
	return s.Max_
}

type Members__ struct {
	Members_ string `xml:"members,attr"`
}

func (s Members__) Members() string {
	return s.Members_
}

type Min__ struct {
	Min_ string `xml:"min,attr"`
}

func (s Min__) Min() string {
	return s.Min_
}

type Name__ struct {
	Name_ string `xml:"name,attr"`
}

func (s Name__) CName() string {
	return s.Name_
}

type Node__ struct {
	Node_ string `xml:"node,attr"`
}

func (s Node__) Node() string {
	return s.Node_
}

type Offset__ struct {
	Offset_ string `xml:"offset,attr"`
}

func (s Offset__) Offset() string {
	return s.Offset_
}

type Restrict__ struct {
	Restrict_ string `xml:"restrict,attr"`
}

func (s Restrict__) Restrict() string {
	return s.Restrict_
}

type Returns__ struct {
	Returns_ string `xml:"returns,attr"`
}

func (s Returns__) Returns() string {
	return s.Returns_
}

type Size__ struct {
	Size_ int `xml:"size,attr"`
}

func (s Size__) Bits() int {
	return s.Size_
}

func (s Size__) Size() int {
	return s.Size_ / 8
}

type Static__ struct {
	Static_ string `xml:"static,attr"`
}

func (s Static__) Static() string {
	return s.Static_
}

type Throw__ struct {
	Throw_ string `xml:"throw,attr"`
}

func (s Throw__) Throw() string {
	return s.Throw_
}

type TreeCode__ struct {
	TreeCode_ string `xml:"tree_code,attr"`
}

func (s TreeCode__) TreeCode() string {
	return s.TreeCode_
}

type TreeCodeName__ struct {
	TreeCodeName_ string `xml:"tree_code_name,attr"`
}

func (s TreeCodeName__) TreeCodeName() string {
	return s.TreeCodeName_
}

type Type__ struct {
	Type_ string `xml:"type,attr"`
}

func (s Type__) Type() string {
	return s.Type_
}

type Volatile__ struct {
	Volatile_ string `xml:"volatile,attr"`
}

func (s Volatile__) Volatile() string {
	return s.Volatile_
}
