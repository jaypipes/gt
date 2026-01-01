package context

import (
	"context"
)

type ContextKey string

// ContextModifier sets some value on the context
type ContextModifier func(context.Context) context.Context

// New returns a new context.Context modified with zero or more options.
func New(mods ...ContextModifier) context.Context {
	ctx := context.Background()
	for _, mod := range mods {
		ctx = mod(ctx)
	}
	return ctx
}

// FromEnv returns a new context.Context populated from the environs or default
// option values
func FromEnv() context.Context {
	ctx := context.Background()
	ll := EnvOrDefaultLogLevel()
	logLevelVar.Set(ll)
	ctx = context.WithValue(ctx, logLevelKey, ll)
	useLogfmt := EnvOrDefaultLogLogfmt()
	if useLogfmt {
		ctx = WithLogLogfmt()(ctx)
	}
	return ctx
}
