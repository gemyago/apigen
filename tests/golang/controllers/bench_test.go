package controllers

import (
	"fmt"
	"strconv"
	"testing"
)

type intPathSegment int

func (s intPathSegment) String() string {
	return strconv.Itoa(int(s))
}

type stringPathSegment string

func (s stringPathSegment) String() string {
	return string(s)
}

type bindingContext struct {
	parent   *bindingContext
	path     fmt.Stringer
	level    int
	maxLevel int
}

func (c *bindingContext) fork(path fmt.Stringer) *bindingContext {
	return &bindingContext{
		parent:   c,
		level:    c.level + 1,
		maxLevel: c.maxLevel,
		path:     path,
	}
}

func call(ctx *bindingContext) {
	if ctx.level < ctx.maxLevel {
		call(ctx.fork(intPathSegment(ctx.level)))
	}
}

func BenchmarkAllocs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call(&bindingContext{level: 0, maxLevel: 10, path: stringPathSegment("root")})
	}
}
