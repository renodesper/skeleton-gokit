package ctx

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetRequestID(t *testing.T) {
	requestID := uuid.New().String()
	ctx := context.Background()

	ctx = SetRequestID(ctx, requestID)
	c := GetRequestID(ctx)
	assert.Equal(t, requestID, c)
}

func TestSetRequestID(t *testing.T) {
	ctx := context.Background()

	// NOTE: Defined requestID
	requestID := "qwertyuiop"
	ctx = SetRequestID(ctx, requestID)
	s := GetRequestID(ctx)
	assert.Equal(t, requestID, s)

	// NOTE: Empty string as requestID
	ctx = SetRequestID(ctx, "")
	s = GetRequestID(ctx)
	assert.NotEmpty(t, s)
}

func TestExtractRequestID(t *testing.T) {
	ctx := context.Background()
	requestID := "qwertyuiop"
	header := map[string][]string{
		"X-Request-Id": {requestID},
	}

	ctx = ExtractRequestID(ctx, &http.Request{
		Header: header,
	})
	s := GetRequestID(ctx)
	assert.Equal(t, requestID, s)
}
