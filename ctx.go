package bctx

import "context"

type (
	Context    = context.Context
	CancelFunc = context.CancelFunc
)

var (
	Canceled         = context.Canceled
	DeadlineExceeded = context.DeadlineExceeded

	WithCancel   = context.WithCancel
	WithDeadline = context.WithDeadline
	WithTimeout  = context.WithTimeout

	Background = context.Background
	TODO       = context.TODO
)
