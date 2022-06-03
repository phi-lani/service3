package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/phi-lani/service3/business/sys/validate"
	"github.com/phi-lani/service3/foundation/web"
	"go.uber.org/zap"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		// return errors.New("untrusted error")
		return validate.NewRequestError(errors.New("untrusted error"), http.StatusBadRequest)
		// return web.NewShutdownError("restart service")
		// panic("testing panic")
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)

}
