// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/jit"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package jit

import (
	r "reflect"
	"github.com/cosmos72/gomacro/imports"
	amd64 "github.com/cosmos72/gomacro/jit/amd64"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/jit"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/jit"] = imports.Package{
	Binds: map[string]r.Value{
		"ADD":	r.ValueOf(ADD),
		"ADD_ASSIGN":	r.ValueOf(ADD_ASSIGN),
		"AND":	r.ValueOf(AND),
		"AND_ASSIGN":	r.ValueOf(AND_ASSIGN),
		"ASM_SUPPORTED":	r.ValueOf(ASM_SUPPORTED),
		"ASSIGN":	r.ValueOf(ASSIGN),
		"Bool":	r.ValueOf(Bool),
		"DEC":	r.ValueOf(DEC),
		"Float32":	r.ValueOf(Float32),
		"Float64":	r.ValueOf(Float64),
		"INC":	r.ValueOf(INC),
		"Int":	r.ValueOf(Int),
		"Int16":	r.ValueOf(Int16),
		"Int32":	r.ValueOf(Int32),
		"Int64":	r.ValueOf(Int64),
		"Int8":	r.ValueOf(Int8),
		"Invalid":	r.ValueOf(Invalid),
		"IsLeaf":	r.ValueOf(IsLeaf),
		"KHi":	r.ValueOf(KHi),
		"KLo":	r.ValueOf(KLo),
		"MMAP_SUPPORTED":	r.ValueOf(MMAP_SUPPORTED),
		"MUL":	r.ValueOf(MUL),
		"MUL_ASSIGN":	r.ValueOf(MUL_ASSIGN),
		"MakeConst":	r.ValueOf(MakeConst),
		"MakeMem":	r.ValueOf(MakeMem),
		"MakeParam":	r.ValueOf(MakeParam),
		"MakeReg":	r.ValueOf(MakeReg),
		"MakeVar":	r.ValueOf(MakeVar),
		"NEG":	r.ValueOf(NEG),
		"NOT":	r.ValueOf(NOT),
		"Name":	r.ValueOf(Name),
		"NewComp":	r.ValueOf(NewComp),
		"NewExpr1":	r.ValueOf(NewExpr1),
		"NewExpr2":	r.ValueOf(NewExpr2),
		"NoRegId":	r.ValueOf(NoRegId),
		"OR":	r.ValueOf(OR),
		"OR_ASSIGN":	r.ValueOf(OR_ASSIGN),
		"Ptr":	r.ValueOf(Ptr),
		"QUO":	r.ValueOf(QUO),
		"QUO_ASSIGN":	r.ValueOf(QUO_ASSIGN),
		"REM":	r.ValueOf(REM),
		"REM_ASSIGN":	r.ValueOf(REM_ASSIGN),
		"RHi":	r.ValueOf(RHi),
		"RLo":	r.ValueOf(RLo),
		"RSP":	r.ValueOf(RSP),
		"RVAR":	r.ValueOf(RVAR),
		"SHL":	r.ValueOf(SHL),
		"SHL_ASSIGN":	r.ValueOf(SHL_ASSIGN),
		"SHR":	r.ValueOf(SHR),
		"SHR_ASSIGN":	r.ValueOf(SHR_ASSIGN),
		"SUB":	r.ValueOf(SUB),
		"SUB_ASSIGN":	r.ValueOf(SUB_ASSIGN),
		"SUPPORTED":	r.ValueOf(SUPPORTED),
		"SizeOf":	r.ValueOf(SizeOf),
		"Uint":	r.ValueOf(Uint),
		"Uint16":	r.ValueOf(Uint16),
		"Uint32":	r.ValueOf(Uint32),
		"Uint64":	r.ValueOf(Uint64),
		"Uint8":	r.ValueOf(Uint8),
		"Uintptr":	r.ValueOf(Uintptr),
		"XOR":	r.ValueOf(XOR),
		"XOR_ASSIGN":	r.ValueOf(XOR_ASSIGN),
		"ZERO":	r.ValueOf(ZERO),
	}, Types: map[string]r.Type{
		"Arg":	r.TypeOf((*Arg)(nil)).Elem(),
		"Code":	r.TypeOf((*Code)(nil)).Elem(),
		"Comp":	r.TypeOf((*Comp)(nil)).Elem(),
		"Const":	r.TypeOf((*Const)(nil)).Elem(),
		"Expr":	r.TypeOf((*Expr)(nil)).Elem(),
		"Expr1":	r.TypeOf((*Expr1)(nil)).Elem(),
		"Expr2":	r.TypeOf((*Expr2)(nil)).Elem(),
		"Inst1":	r.TypeOf((*Inst1)(nil)).Elem(),
		"Inst2":	r.TypeOf((*Inst2)(nil)).Elem(),
		"Kind":	r.TypeOf((*Kind)(nil)).Elem(),
		"Mem":	r.TypeOf((*Mem)(nil)).Elem(),
		"Op1":	r.TypeOf((*Op1)(nil)).Elem(),
		"Op2":	r.TypeOf((*Op2)(nil)).Elem(),
		"Reg":	r.TypeOf((*Reg)(nil)).Elem(),
		"RegId":	r.TypeOf((*RegId)(nil)).Elem(),
		"Save":	r.TypeOf((*Save)(nil)).Elem(),
		"Size":	r.TypeOf((*Size)(nil)).Elem(),
		"SoftReg":	r.TypeOf((*SoftReg)(nil)).Elem(),
		"SoftRegId":	r.TypeOf((*SoftRegId)(nil)).Elem(),
		"Stmt":	r.TypeOf((*Stmt)(nil)).Elem(),
		"Stmt1":	r.TypeOf((*Stmt1)(nil)).Elem(),
		"Stmt2":	r.TypeOf((*Stmt2)(nil)).Elem(),
	}, Proxies: map[string]r.Type{
		"Arg":	r.TypeOf((*P_github_com_cosmos72_gomacro_jit_Arg)(nil)).Elem(),
		"Expr":	r.TypeOf((*P_github_com_cosmos72_gomacro_jit_Expr)(nil)).Elem(),
	}, Untypeds: map[string]string{
		"ASM_SUPPORTED":	"bool:true",
		"MMAP_SUPPORTED":	"bool:true",
		"Name":	"string:amd64",
		"SUPPORTED":	"bool:true",
	}, 
	}
}

// --------------- proxy for github.com/cosmos72/gomacro/jit.Arg ---------------
type P_github_com_cosmos72_gomacro_jit_Arg struct {
	Object	interface{}
	Const_	func(interface{}) bool
	Kind_	func(interface{}) amd64.Kind
	RegId_	func(interface{}) amd64.RegId
}
func (P *P_github_com_cosmos72_gomacro_jit_Arg) Const() bool {
	return P.Const_(P.Object)
}
func (P *P_github_com_cosmos72_gomacro_jit_Arg) Kind() amd64.Kind {
	return P.Kind_(P.Object)
}
func (P *P_github_com_cosmos72_gomacro_jit_Arg) RegId() amd64.RegId {
	return P.RegId_(P.Object)
}

// --------------- proxy for github.com/cosmos72/gomacro/jit.Expr ---------------
type P_github_com_cosmos72_gomacro_jit_Expr struct {
	Object	interface{}
	Const_	func(interface{}) bool
	Kind_	func(interface{}) amd64.Kind
}
func (P *P_github_com_cosmos72_gomacro_jit_Expr) Const() bool {
	return P.Const_(P.Object)
}
func (P *P_github_com_cosmos72_gomacro_jit_Expr) Kind() amd64.Kind {
	return P.Kind_(P.Object)
}