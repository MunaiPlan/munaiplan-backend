package client

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
)

type twoSeries struct {
	Depth []float64 `json:"Глубина"`
	A     []float64 `json:"A"`
	B     []float64 `json:"B"`
}

func TestPredictValidatesResponses(t *testing.T) {
	cases := map[string]struct {
		status int
		body   string
		want   error
	}{
		"ok":          {200, `{"A":[1,2],"B":[3,4]}`, nil},
		"missing":     {200, `{"A":[1,2]}`, ErrInvalidResponse},
		"short":       {200, `{"A":[1],"B":[3,4]}`, ErrInvalidResponse},
		"nan":         {200, `{"A":[NaN,2],"B":[3,4]}`, ErrInvalidResponse},
		"not ready":   {503, `{"detail":"Model artifacts are not ready"}`, ErrUnavailable},
		"bad request": {422, `{"detail":"x"}`, ErrInvalidResponse},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/predict/moment/" {
					t.Errorf("path %s", r.URL.Path)
				}
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.body))
			}))
			defer srv.Close()
			var out twoSeries
			err := NewTorqueAndDragClient(srv.URL+"/predict").Predict(context.Background(), Torque, requests.TorqueAndDragFromMLModelRequest{MD: []float64{0, 1}}, &out)
			if !errors.Is(err, tc.want) || (tc.want == nil && err != nil) {
				t.Fatalf("got %v, want %v", err, tc.want)
			}
			if tc.want == nil && out.B[1] != 4 {
				t.Fatalf("decoded %+v", out)
			}
		})
	}
	var out twoSeries
	err := NewTorqueAndDragClient("http://127.0.0.1:1").Predict(context.Background(), Torque, requests.TorqueAndDragFromMLModelRequest{}, &out)
	if !errors.Is(err, ErrUnavailable) {
		t.Fatalf("unreachable service must be ErrUnavailable, got %v", err)
	}
}
