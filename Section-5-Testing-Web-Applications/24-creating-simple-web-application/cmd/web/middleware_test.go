package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	type args struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "empty", args: args{headerName: "", headerValue: "", addr: "", emptyAddr: false}},
		{name: "empty", args: args{headerName: "", headerValue: "", addr: "", emptyAddr: true}},
		{name: "forwarded", args: args{headerName: "X-Forwarded-For", headerValue: "192.3.2.1", addr: "", emptyAddr: false}},
		{name: "empty", args: args{headerName: "", headerValue: "", addr: "hello:world", emptyAddr: false}},
	}

	// create dummy handler to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, " not present")
		}
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})

	for _, tt := range tests {
		// create handler to test
		handlerToTest := app.addIpToContext(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)

		if tt.args.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(tt.args.headerName) > 0 {
			req.Header.Add(tt.args.headerName, tt.args.headerValue)
		}
		if len(tt.args.addr) > 0 {
			req.RemoteAddr = tt.args.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, contextUserKey, "whatever")

	ip := app.ipFromContext(ctx)
	if !strings.EqualFold("whatever", ip) {
		t.Error("Wrong value returned from context")
	}
}
