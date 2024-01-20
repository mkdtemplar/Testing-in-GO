package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_Home(t *testing.T) {
	type args struct {
		url                string
		expectedStatusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "home", args: args{url: "/", expectedStatusCode: http.StatusOK}},
		{name: "404", args: args{url: "/finsh", expectedStatusCode: http.StatusNotFound}},
	}

	var app application
	routes := app.routes()

	// create test web server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	for _, tt := range tests {
		resp, err := ts.Client().Get(ts.URL + tt.args.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != tt.args.expectedStatusCode {
			t.Errorf("for %s: expected %d, but got %d", tt.name, tt.args.expectedStatusCode, resp.StatusCode)
		}
	}
}

func Test_application_render(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		r    *http.Request
		t    string
		data *TemplateData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &application{}
			if err := a.render(tt.args.w, tt.args.r, tt.args.t, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
