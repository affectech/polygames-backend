package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"polygames/internal/core"
	"testing"
)

func Test_server_CreateGenre(t *testing.T) {
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
			s.CreateGenre(tt.args.w, tt.args.r)
		})
	}
}

func Test_server_GetGenre(t *testing.T) {
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
			s.GetGenre(tt.args.w, tt.args.r)
		})
	}
}

func Test_server_UpdateGenre(t *testing.T) {
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
			s.UpdateGenre(tt.args.w, tt.args.r)
		})
	}
}

func Test_server_DeleteGenre(t *testing.T) {
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
			s.DeleteGenre(tt.args.w, tt.args.r)
		})
	}
}
