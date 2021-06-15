package ctx

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type CtxKey int
type JwtKey string
type ApiKey string

const (
	// CtxRequestID ...
	CtxRequestID CtxKey = iota
	CtxJwtToken  JwtKey = "jwtToken"
	CtxApiKey    ApiKey = "apiKey"
)

// GetRequestID ...
func GetRequestID(ctx context.Context) string {
	res, _ := ctx.Value(CtxRequestID).(string)
	return res
}

// SetRequestID ...
func SetRequestID(ctx context.Context, val string) context.Context {
	if val == "" {
		val = uuid.New().String()
	}
	return context.WithValue(ctx, CtxRequestID, val)
}

// ExtractRequestID ...
func ExtractRequestID(ctx context.Context, r *http.Request) context.Context {
	ctx = SetRequestID(ctx, r.Header.Get("X-Request-Id"))
	return ctx
}

// GetJwtToken ...
func GetJwtToken(ctx context.Context) string {
	res, _ := ctx.Value(CtxJwtToken).(string)
	return res
}

// SetJwtToken ...
func SetJwtToken(ctx context.Context, val string) context.Context {
	if val == "" {
		return ctx
	}

	authorization := strings.TrimSpace(val)
	if authorization == "Bearer" {
		return ctx
	}

	slices := strings.Split(authorization, " ")
	jwtToken := slices[1]

	return context.WithValue(ctx, CtxJwtToken, jwtToken)
}

func ExtractJwtToken(ctx context.Context, r *http.Request) context.Context {
	ctx = SetJwtToken(ctx, r.Header.Get("Authorization"))
	return ctx
}

// GetApiKey ...
func GetApiKey(ctx context.Context) string {
	res, _ := ctx.Value(CtxApiKey).(string)
	return res
}

// SetApiKey ...
func SetApiKey(ctx context.Context, val string) context.Context {
	if val == "" {
		return ctx
	}

	return context.WithValue(ctx, CtxApiKey, val)
}

func ExtractApiKey(ctx context.Context, r *http.Request) context.Context {
	ctx = SetApiKey(ctx, r.Header.Get("X-Api-Key"))
	return ctx
}
