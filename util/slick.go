package util

import (
	"context"
	"net/http"
)

type Context struct {
	w   http.ResponseWriter
	r   *http.Request
	ctx context.Context
}
