package rpc2

import (
	"log"
	"reflect"
	"unicode"
	"unicode/utf8"
)

// Precompute the reflect type for error.  Can't use error directly
// because Typeof takes an empty interface value.  This is annoying.
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type handler struct {
	fn        reflect.Value
	argType   reflect.Type
	replyType reflect.Type
}

func addHandler(handlers map[string]*handler, mname string, handlerFunc interface{}) {
	if _, ok := handlers[mname]; ok {
		panic("rpc2: multiple registrations for " + mname)
	}

	method := reflect.ValueOf(handlerFunc)
	mtype := method.Type()
	// Method needs three ins: *client, *args, *reply.
	if mtype.NumIn() != 2 {
		log.Panicln("method", mname, "has wrong number of ins:", mtype.NumIn())
	}
	// Second arg need not be a pointer.
	argType := mtype.In(0)
	if !isExportedOrBuiltinType(argType) {
		log.Panicln(mname, "argument type not exported:", argType)
	}
	// Third arg must be a pointer.
	replyType := mtype.In(1)
	if replyType.Kind() != reflect.Ptr {
		log.Panicln("method", mname, "reply type not a pointer:", replyType)
	}
	// Reply type must be exported.
	if !isExportedOrBuiltinType(replyType) {
		log.Panicln("method", mname, "reply type not exported:", replyType)
	}
	// Method needs one out.
	if mtype.NumOut() != 1 {
		log.Panicln("method", mname, "has wrong number of outs:", mtype.NumOut())
	}
	// The return type of the method must be error.
	if returnType := mtype.Out(0); returnType != typeOfError {
		log.Panicln("method", mname, "returns", returnType.String(), "not error")
	}
	handlers[mname] = &handler{
		fn:        method,
		argType:   argType,
		replyType: replyType,
	}
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

// Is this an exported - upper case - name?
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}
