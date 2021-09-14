package logs

import (
	"context"
	"io"
	"time"
)

type LogOptions struct {
	W io.Writer

	MaxBackoff time.Duration
	AppName    string
	VMID       string
	RegionCode string
	Org        string
}

type LogStream interface {
	Stream(ctx context.Context, opts *LogOptions) <-chan LogEntry
}
