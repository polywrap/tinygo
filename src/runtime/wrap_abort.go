package runtime

import (
	"unsafe"
)

//go:wasm-module wrap
//export __wrap_abort
func __wrap_abort(
	msgPtr,
	msgLen,
	filePtr,
	fileLen,
	line,
	column uint32,
)

func wrapAbort(msg string) {
	msgPtr := uint32(uintptr(unsafe.Pointer(&msg)))
	msgLen := uint32(len(msg))

	file := "wrap_abort.go"
	filePtr := uint32(uintptr(unsafe.Pointer(&file)))
	fileLen := uint32(len(file))

	line := uint32(12)
	column := uint32(0)

	__wrap_abort(
		msgPtr,
		msgLen,
		filePtr,
		fileLen,
		line,
		column,
	)

	abort()
}
