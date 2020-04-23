package transmission

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOption(t *testing.T) {
	var tests = []struct {
		name string
		opt  Option
		want config
	}{
		{
			name: "auth",
			opt:  WithAuth("admin", "querty"),
			want: config{
				Username: "admin",
				Password: "querty",
			},
		},
		{
			name: "http_client",
			opt:  WithHTTPClient(http.DefaultClient),
			want: config{
				HTTPClient: http.DefaultClient,
			},
		},
		{
			name: "user_agent",
			opt:  WithUserAgent("go-transmission"),
			want: config{
				UserAgent: "go-transmission",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var c config
			tc.opt.apply(&c)
			if !cmp.Equal(tc.want, c) {
				t.Errorf("unexpected client configuration, diff = \n%s", cmp.Diff(tc.want, c))
			}
		})
	}
}
