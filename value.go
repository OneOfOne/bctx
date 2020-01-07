package bctx

import (
	"sync"
)

// WithValue returns a context that can hold multiple values and safe for concurrent access.
func WithValue(parent Context, key, val interface{}) Context {
	var vctx *valCtx
	if vctx, _ := parent.(*valCtx); vctx == nil {
		vctx = &valCtx{Context: parent}
	}
	vctx.set(key, val)
	return vctx
}

type ValueSetter func(key, val interface{})

func WithMultiValue(ctx Context) (Context, ValueSetter) {
	vctx := &valCtx{Context: ctx}
	return vctx, vctx.set
}

type valCtx struct {
	Context
	mux sync.RWMutex
	m   map[interface{}]interface{}
}

func (v *valCtx) Value(key interface{}) interface{} {
	v.mux.RLock()
	val, ok := v.m[key]
	v.mux.RUnlock()
	if ok {
		return val
	}

	return v.Context.Value(key)
}

func (v *valCtx) set(key, val interface{}) {
	v.mux.Lock()
	if v.m == nil {
		v.m = make(map[interface{}]interface{})
	}
	v.m[key] = val
	v.mux.Unlock()
}
