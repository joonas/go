// Generated by `wit-bindgen-wrpc-go` 0.8.0. DO NOT EDIT!
// server package contains wRPC bindings for `server` world
package server

import (
	exports__wrpc__keyvalue__store "github.com/wasmCloud/provider-sdk-go/examples/keyvalue-inmemory/bindings/exports/wrpc/keyvalue/store"
	wrpc "wrpc.io/go"
)

func Serve(s wrpc.Server, h0 exports__wrpc__keyvalue__store.Handler) (stop func() error, err error) {
	stops := make([]func() error, 0, 1)
	stop = func() error {
		for _, stop := range stops {
			if err := stop(); err != nil {
				return err
			}
		}
		return nil
	}
	stop0, err := exports__wrpc__keyvalue__store.ServeInterface(s, h0)
	if err != nil {
		return
	}
	stops = append(stops, stop0)
	stop = func() error {
		if err := stop0(); err != nil {
			return err
		}
		return nil
	}
	return
}
