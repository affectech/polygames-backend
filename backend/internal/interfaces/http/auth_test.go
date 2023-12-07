package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"polygames/internal/core"
	"testing"
)

func Test_server_SignIn(t *testing.T) {
	type fields struct {
		core       core.Core
		router     chi.Router
		httpServer *http.Server
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				core:       tt.fields.core,
				router:     tt.fields.router,
				httpServer: tt.fields.httpServer,
			}
			s.SignIn(tt.args.w, tt.args.r)
		})
	}
}
