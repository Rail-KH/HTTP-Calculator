package application_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Rail-KH/HTTP-Calculator/internal/application"
)

func TestApp(t *testing.T) {
	testAppSuccess := []struct {
		name       string
		expression application.Request
		status     int
		result     application.Answer
	}{
		{
			name:       "result1",
			expression: application.Request{Expression: "1+1"},
			status:     200,
			result:     application.Answer{Result: 2},
		},

		{
			name:       "result2",
			expression: application.Request{Expression: "2+2*2"},
			status:     200,
			result:     application.Answer{Result: 6},
		},
	}
	for _, tc := range testAppSuccess {
		json_b, _ := json.Marshal(tc.expression)
		json_new := bytes.NewBuffer(json_b)
		req := httptest.NewRequest("POST", "/api/v1/calculate", json_new)
		w := httptest.NewRecorder()
		application.CalcHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		answer := application.Answer{}
		json.Unmarshal(w.Body.Bytes(), &answer)
		if tc.result.Result != answer.Result {
			t.Errorf("wrong result")
		}
		if tc.status != w.Code {
			t.Errorf("wrong status code")
		}

	}

	testAppUnvalid := []struct {
		name         string
		expression   application.Request
		status       int
		server_error application.ServerError
	}{
		{
			name:         "error1",
			expression:   application.Request{Expression: "1+1*"},
			status:       422,
			server_error: application.ServerError{Error: "Expression is not valid"},
		},

		{
			name:         "error2",
			expression:   application.Request{Expression: "((2+2-*(2"},
			status:       422,
			server_error: application.ServerError{Error: "Expression is not valid"},
		},
		{
			name:         "error3",
			expression:   application.Request{Expression: "100/0"},
			status:       422,
			server_error: application.ServerError{Error: "Expression is not valid"},
		},
	}
	for _, tc := range testAppUnvalid {
		json_b, _ := json.Marshal(tc.expression)
		json_new := bytes.NewBuffer(json_b)
		req := httptest.NewRequest("POST", "/api/v1/calculate", json_new)
		w := httptest.NewRecorder()
		application.CalcHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		server_error := application.ServerError{}
		json.Unmarshal(w.Body.Bytes(), &server_error)
		if tc.server_error.Error != server_error.Error {
			t.Errorf("wrong error")
		}
		if tc.status != w.Code {
			t.Errorf("wrong status code")
		}

	}

}
