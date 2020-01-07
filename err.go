package bctx

import (
	"context"
	"sync"
)

type CancelErrFunc func(error)

// WithCancelError returns a context that can return a custom error for cancelation.
// the error returned from `Err()` would return true for `xerrors.Is(err, context.Canceled)`.
func WithCancelError(ctx Context) (Context, CancelErrFunc) {
	ctx, fn := WithCancel(ctx)
	ectx := &errCtx{Context: ctx, fn: fn}
	return ectx, ectx.cancel
}

type errCtx struct {
	context.Context
	fn  CancelFunc
	mux sync.Mutex
	err error
}

func (ec *errCtx) Err() error {
	ec.mux.Lock()
	err := ec.err
	ec.mux.Unlock()
	if err == nil {
		err = ec.Context.Err()
	}
	return err
}

func (ec *errCtx) cancel(err error) {
	ec.mux.Lock()
	if ec.err == nil {
		ec.fn()
		ec.err = cancelError{err}
	}
	ec.mux.Unlock()
}

type cancelError struct {
	e error
}

func (ce cancelError) Error() string {
	return ce.e.Error()
}

func (ce cancelError) Is(oe error) bool {
	return oe == ce.e || oe == Canceled
}

func (ce cancelError) Unwrap() error {
	return ce.e
}
