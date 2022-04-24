package context

import (
	"context"

	"github.com/apex/log"
)

type key int

const logFields key = 0

func getFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

// FromContext : context 데이터 필드에 있는 값을 리턴.
func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// WithField : context 에 데이터 필드를 추가 후 리턴. 
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, log.Fields{key: value})
}

// WithFields : 많은 필드를 추가.
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	return context.WithValue(ctx, logFields, f)
}