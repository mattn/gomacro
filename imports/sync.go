// this file was generated by gomacro command: import "sync"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "sync"
	. "reflect"
)

func init() {
	Binds["sync"] = map[string]Value{
		"NewCond":	ValueOf(pkg.NewCond),
	}
	Types["sync"] = map[string]Type{
		"Cond":	TypeOf((*pkg.Cond)(nil)).Elem(),
		"Locker":	TypeOf((*pkg.Locker)(nil)).Elem(),
		"Mutex":	TypeOf((*pkg.Mutex)(nil)).Elem(),
		"Once":	TypeOf((*pkg.Once)(nil)).Elem(),
		"Pool":	TypeOf((*pkg.Pool)(nil)).Elem(),
		"RWMutex":	TypeOf((*pkg.RWMutex)(nil)).Elem(),
		"WaitGroup":	TypeOf((*pkg.WaitGroup)(nil)).Elem(),
	}
}