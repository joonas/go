// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package terminalstderr represents the imported interface "wasi:cli/terminal-stderr@0.2.0".
package terminalstderr

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	terminaloutput "github.com/wasmCloud/component-sdk-go/_examples/http-server/gen/wasi/cli/terminal-output"
)

// GetTerminalStderr represents the imported function "get-terminal-stderr".
//
//	get-terminal-stderr: func() -> option<terminal-output>
//
//go:nosplit
func GetTerminalStderr() (result cm.Option[terminaloutput.TerminalOutput]) {
	wasmimport_GetTerminalStderr(&result)
	return
}

//go:wasmimport wasi:cli/terminal-stderr@0.2.0 get-terminal-stderr
//go:noescape
func wasmimport_GetTerminalStderr(result *cm.Option[terminaloutput.TerminalOutput])
