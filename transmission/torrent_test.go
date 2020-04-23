package transmission

import (
	"encoding/json"
	"testing"
)

func TestIdentifier(t *testing.T) {
	var tests = []struct {
		name string
		ids  Identifier
		want string
	}{
		{
			name: "single_id",
			ids:  ID(1),
			want: `{"ids":1}`,
		},
		{
			name: "single_hash",
			ids:  Hash("abcde"),
			want: `{"ids":"abcde"}`,
		},
		{
			name: "mixed",
			ids:  IDs(ID(1), Hash("abcde")),
			want: `{"ids":[1,"abcde"]}`,
		},
		{
			name: "recenty_active",
			ids:  RecentlyActive(),
			want: `{"ids":"recently-active"}`,
		},
		{
			name: "all",
			ids:  All(),
			want: `{}`,
		},
	}

	type testStruct struct {
		IDs Identifier `json:"ids,omitempty"`
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(&testStruct{
				IDs: tc.ids,
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if want, got := tc.want, string(data); want != got {
				t.Errorf("unexpected output, want = %q, got = %q", want, got)
			}
		})
	}
}
