package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func Test_GetAPIKey(t *testing.T) {
	type test struct {
		input   func(req *http.Request)
		expects string // error output
	}
	cases := []test{
		{
			input:   func(req *http.Request) { req.Header.Set("Authorization", "ApiKey HEREISAKEY123") },
			expects: "",
		},
		{
			input:   func(req *http.Request) { req.Header.Set("Authorizashun", "ApiKey HEREISAKEY123") },
			expects: "no authorization header included",
		},
		{
			input:   func(req *http.Request) { req.Header.Set("Authorization", "ApiKey") },
			expects: "malformed authorization header",
		},
		{
			input:   func(req *http.Request) { req.Header.Set("Authorization", "") },
			expects: "no authorization header included",
		},
		{
			input:   func(req *http.Request) { req.Header.Set("FakeHeader", "NoValue") },
			expects: "no authorization header included",
		},
	}
	for i, c := range cases {
		req, _ := http.NewRequest("GET", "/", nil)
		if c.input != nil {
			c.input(req)
		} else {
			t.Errorf("Invalid input function")
		}

		_, err := auth.GetAPIKey(req.Header)
		if err != nil && c.expects != err.Error() {
			t.Errorf("CASE %d: expected %s, got %s", i, c.expects, err.Error())
		}
	}
}
