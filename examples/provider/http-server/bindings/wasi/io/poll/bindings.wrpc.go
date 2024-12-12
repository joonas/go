// Generated by `wit-bindgen-wrpc-go` 0.11.0. DO NOT EDIT!
package poll

import (
	bytes "bytes"
	context "context"
	binary "encoding/binary"
	errors "errors"
	fmt "fmt"
	io "io"
	slog "log/slog"
	math "math"
	sync "sync"
	atomic "sync/atomic"
	wrpc "wrpc.io/go"
)

// `pollable` represents a single I/O event which may be ready, or not.
type Pollable interface{}

// Return the readiness of a pollable. This function never blocks.
//
// Returns `true` when the pollable is ready, and `false` otherwise.
func Pollable_Ready(ctx__ context.Context, wrpc__ wrpc.Invoker, self wrpc.Borrow[Pollable]) (r0__ bool, err__ error) {
	var buf__ bytes.Buffer
	write0__, err__ := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
		n := len(v)
		if n > math.MaxUint32 {
			return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
		}
		if err = func(v int, w io.Writer) error {
			b := make([]byte, binary.MaxVarintLen32)
			i := binary.PutUvarint(b, uint64(v))
			slog.Debug("writing string byte length", "len", n)
			_, err = w.Write(b[:i])
			return err
		}(n, w); err != nil {
			return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
		}
		slog.Debug("writing string bytes")
		_, err = w.Write([]byte(v))
		if err != nil {
			return fmt.Errorf("failed to write string bytes: %w", err)
		}
		return nil
	}(string(self), &buf__)
	if err__ != nil {
		err__ = fmt.Errorf("failed to write `self` parameter: %w", err__)
		return
	}
	if write0__ != nil {
		err__ = errors.New("unexpected deferred write for synchronous `self` parameter")
		return
	}
	var w__ wrpc.IndexWriteCloser
	var r__ wrpc.IndexReadCloser
	w__, r__, err__ = wrpc__.Invoke(ctx__, "wasi:io/poll@0.2.0", "pollable.ready", buf__.Bytes())
	if err__ != nil {
		err__ = fmt.Errorf("failed to invoke `[method]pollable.ready`: %w", err__)
		return
	}
	defer func() {
		if err := r__.Close(); err != nil {
			slog.ErrorContext(ctx__, "failed to close reader", "instance", "wasi:io/poll@0.2.0", "name", "[method]pollable.ready", "err", err)
		}
	}()
	if cErr__ := w__.Close(); cErr__ != nil {
		slog.DebugContext(ctx__, "failed to close outgoing stream", "instance", "wasi:io/poll@0.2.0", "name", "[method]pollable.ready", "err", cErr__)
	}
	r0__, err__ = func(r io.ByteReader) (bool, error) {
		slog.Debug("reading bool byte")
		v, err := r.ReadByte()
		if err != nil {
			slog.Debug("reading bool", "value", false)
			return false, fmt.Errorf("failed to read bool byte: %w", err)
		}
		switch v {
		case 0:
			return false, nil
		case 1:
			return true, nil
		default:
			return false, fmt.Errorf("invalid bool value %d", v)
		}
	}(r__)
	if err__ != nil {
		err__ = fmt.Errorf("failed to read result 0: %w", err__)
		return
	}
	return
}

// `block` returns immediately if the pollable is ready, and otherwise
// blocks until ready.
//
// This function is equivalent to calling `poll.poll` on a list
// containing only this pollable.
func Pollable_Block(ctx__ context.Context, wrpc__ wrpc.Invoker, self wrpc.Borrow[Pollable]) (err__ error) {
	var buf__ bytes.Buffer
	write0__, err__ := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
		n := len(v)
		if n > math.MaxUint32 {
			return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
		}
		if err = func(v int, w io.Writer) error {
			b := make([]byte, binary.MaxVarintLen32)
			i := binary.PutUvarint(b, uint64(v))
			slog.Debug("writing string byte length", "len", n)
			_, err = w.Write(b[:i])
			return err
		}(n, w); err != nil {
			return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
		}
		slog.Debug("writing string bytes")
		_, err = w.Write([]byte(v))
		if err != nil {
			return fmt.Errorf("failed to write string bytes: %w", err)
		}
		return nil
	}(string(self), &buf__)
	if err__ != nil {
		err__ = fmt.Errorf("failed to write `self` parameter: %w", err__)
		return
	}
	if write0__ != nil {
		err__ = errors.New("unexpected deferred write for synchronous `self` parameter")
		return
	}
	var w__ wrpc.IndexWriteCloser
	var r__ wrpc.IndexReadCloser
	w__, r__, err__ = wrpc__.Invoke(ctx__, "wasi:io/poll@0.2.0", "pollable.block", buf__.Bytes())
	if err__ != nil {
		err__ = fmt.Errorf("failed to invoke `[method]pollable.block`: %w", err__)
		return
	}
	defer func() {
		if err := r__.Close(); err != nil {
			slog.ErrorContext(ctx__, "failed to close reader", "instance", "wasi:io/poll@0.2.0", "name", "[method]pollable.block", "err", err)
		}
	}()
	if cErr__ := w__.Close(); cErr__ != nil {
		slog.DebugContext(ctx__, "failed to close outgoing stream", "instance", "wasi:io/poll@0.2.0", "name", "[method]pollable.block", "err", cErr__)
	}
	return
}

// Poll for completion on a set of pollables.
//
// This function takes a list of pollables, which identify I/O sources of
// interest, and waits until one or more of the events is ready for I/O.
//
// The result `list<u32>` contains one or more indices of handles in the
// argument list that is ready for I/O.
//
// If the list contains more elements than can be indexed with a `u32`
// value, this function traps.
//
// A timeout can be implemented by adding a pollable from the
// wasi-clocks API to the list.
//
// This function does not return a `result`; polling in itself does not
// do any I/O so it doesn't fail. If any of the I/O sources identified by
// the pollables has an error, it is indicated by marking the source as
// being reaedy for I/O.
func Poll(ctx__ context.Context, wrpc__ wrpc.Invoker, in []wrpc.Borrow[Pollable]) (r0__ []uint32, err__ error) {
	var buf__ bytes.Buffer
	write0__, err__ := func(v []wrpc.Borrow[Pollable], w interface {
		io.ByteWriter
		io.Writer
	}) (write func(wrpc.IndexWriter) error, err error) {
		n := len(v)
		if n > math.MaxUint32 {
			return nil, fmt.Errorf("list length of %d overflows a 32-bit integer", n)
		}
		if err = func(v int, w io.Writer) error {
			b := make([]byte, binary.MaxVarintLen32)
			i := binary.PutUvarint(b, uint64(v))
			slog.Debug("writing list length", "len", n)
			_, err = w.Write(b[:i])
			return err
		}(n, w); err != nil {
			return nil, fmt.Errorf("failed to write list length of %d: %w", n, err)
		}
		slog.Debug("writing list elements")
		writes := make(map[uint32]func(wrpc.IndexWriter) error, n)
		for i, e := range v {
			write, err := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
				n := len(v)
				if n > math.MaxUint32 {
					return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
				}
				if err = func(v int, w io.Writer) error {
					b := make([]byte, binary.MaxVarintLen32)
					i := binary.PutUvarint(b, uint64(v))
					slog.Debug("writing string byte length", "len", n)
					_, err = w.Write(b[:i])
					return err
				}(n, w); err != nil {
					return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
				}
				slog.Debug("writing string bytes")
				_, err = w.Write([]byte(v))
				if err != nil {
					return fmt.Errorf("failed to write string bytes: %w", err)
				}
				return nil
			}(string(e), w)
			if err != nil {
				return nil, fmt.Errorf("failed to write list element %d: %w", i, err)
			}
			if write != nil {
				writes[uint32(i)] = write
			}
		}
		if len(writes) > 0 {
			return func(w wrpc.IndexWriter) error {
				var wg sync.WaitGroup
				var wgErr atomic.Value
				for index, write := range writes {
					wg.Add(1)
					w, err := w.Index(index)
					if err != nil {
						return fmt.Errorf("failed to index nested list writer: %w", err)
					}
					write := write
					go func() {
						defer wg.Done()
						if err := write(w); err != nil {
							wgErr.Store(err)
						}
					}()
				}
				wg.Wait()
				err := wgErr.Load()
				if err == nil {
					return nil
				}
				return err.(error)
			}, nil
		}
		return nil, nil
	}(in, &buf__)
	if err__ != nil {
		err__ = fmt.Errorf("failed to write `in` parameter: %w", err__)
		return
	}
	if write0__ != nil {
		err__ = errors.New("unexpected deferred write for synchronous `in` parameter")
		return
	}
	var w__ wrpc.IndexWriteCloser
	var r__ wrpc.IndexReadCloser
	w__, r__, err__ = wrpc__.Invoke(ctx__, "wasi:io/poll@0.2.0", "poll", buf__.Bytes())
	if err__ != nil {
		err__ = fmt.Errorf("failed to invoke `poll`: %w", err__)
		return
	}
	defer func() {
		if err := r__.Close(); err != nil {
			slog.ErrorContext(ctx__, "failed to close reader", "instance", "wasi:io/poll@0.2.0", "name", "poll", "err", err)
		}
	}()
	if cErr__ := w__.Close(); cErr__ != nil {
		slog.DebugContext(ctx__, "failed to close outgoing stream", "instance", "wasi:io/poll@0.2.0", "name", "poll", "err", cErr__)
	}
	r0__, err__ = func(r wrpc.IndexReadCloser, path ...uint32) ([]uint32, error) {
		var x uint32
		var s uint
		for i := 0; i < 5; i++ {
			slog.Debug("reading list length byte", "i", i)
			b, err := r.ReadByte()
			if err != nil {
				if i > 0 && err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return nil, fmt.Errorf("failed to read list length byte: %w", err)
			}
			if s == 28 && b > 0x0f {
				return nil, errors.New("list length overflows a 32-bit integer")
			}
			if b < 0x80 {
				x = x | uint32(b)<<s
				if x == 0 {
					return nil, nil
				}
				vs := make([]uint32, x)
				for i := range vs {
					slog.Debug("reading list element", "i", i)
					vs[i], err = func(r io.ByteReader) (uint32, error) {
						var x uint32
						var s uint8
						for i := 0; i < 5; i++ {
							slog.Debug("reading u32 byte", "i", i)
							b, err := r.ReadByte()
							if err != nil {
								if i > 0 && err == io.EOF {
									err = io.ErrUnexpectedEOF
								}
								return x, fmt.Errorf("failed to read u32 byte: %w", err)
							}
							if s == 28 && b > 0x0f {
								return x, errors.New("varint overflows a 32-bit integer")
							}
							if b < 0x80 {
								return x | uint32(b)<<s, nil
							}
							x |= uint32(b&0x7f) << s
							s += 7
						}
						return x, errors.New("varint overflows a 32-bit integer")
					}(r)
					if err != nil {
						return nil, fmt.Errorf("failed to read list element %d: %w", i, err)
					}
				}
				return vs, nil
			}
			x |= uint32(b&0x7f) << s
			s += 7
		}
		return nil, errors.New("list length overflows a 32-bit integer")
	}(r__, []uint32{0}...)
	if err__ != nil {
		err__ = fmt.Errorf("failed to read result 0: %w", err__)
		return
	}
	return
}