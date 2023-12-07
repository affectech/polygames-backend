package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"polygames/internal/core"
	"testing"
)

func Test_server_CreateGame(t *testing.T) {
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
			s.CreateGame(tt.args.w, tt.args.r)
		})
	}
}

func Test_server_GetGame(t *testing.T) {
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
			s.GetGame(tt.args.w, tt.args.r)
		})
	}
}

func Test_server_UpdateGame(t *testing.T) {
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
			s.UpdateGame(tt.args.w, tt.args.r)
		})
	}
}

func Test_server_DeleteGame(t *testing.T) {
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
			s.DeleteGame(tt.args.w, tt.args.r)
		})
	}
}
