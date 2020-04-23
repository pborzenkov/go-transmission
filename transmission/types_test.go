package transmission

import (
	"encoding/json"
	"testing"
)

func TestWeekday(t *testing.T) {
	var tests = []struct {
		name string
		days Weekday
		want string
	}{
		{
			name: "single_day",
			days: Tuesday,
			want: "Tuesday",
		},
		{
			name: "several_days",
			days: Monday | Friday | Saturday,
			want: "Monday, Friday, Saturday",
		},
		{
			name: "weekdays",
			days: Monday | Tuesday | Wednesday | Thursday | Friday,
			want: "Weekdays",
		},
		{
			name: "weekends",
			days: Saturday | Sunday,
			want: "Weekends",
		},
		{
			name: "everyday",
			days: Sunday | Monday | Tuesday | Wednesday | Thursday | Friday | Saturday,
			want: "Every Day",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if want, got := tc.want, tc.days.String(); want != got {
				t.Errorf("unexpected output, want = %q, got = %q", want, got)
			}
		})
	}
}

func TestEncryption(t *testing.T) {
	var tests = []struct {
		enc  Encryption
		want string
	}{
		{enc: EncryptionPreferred, want: "preferred"},
		{enc: EncryptionRequired, want: "required"},
		{enc: EncryptionTolerated, want: "tolerated"},
		{enc: 10, want: "Encryption(10)"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.want, func(t *testing.T) {
			if want, got := tc.want, tc.enc.String(); want != got {
				t.Errorf("unexpected output, want = %q, got = %q", want, got)
			}
		})
	}

	type testStruct struct {
		Enc Encryption `json:"enc"`
	}
	t.Run("marshal", func(t *testing.T) {
		data, err := json.Marshal(testStruct{Enc: EncryptionRequired})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if want, got := `{"enc":"required"}`, string(data); want != got {
			t.Errorf("unexpected output, want = %q, got = %q", want, got)
		}
		if _, err = json.Marshal(testStruct{Enc: 10}); err == nil {
			t.Errorf("expected marshal to fail")
		}
	})
	t.Run("unmarshal", func(t *testing.T) {
		var resp testStruct
		if err := json.Unmarshal([]byte(`{"enc":"tolerated"}`), &resp); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if want, got := EncryptionTolerated, resp.Enc; want != got {
			t.Errorf("unexpected output, want = %q, got = %q", want, got)
		}
		if err := json.Unmarshal([]byte(`{"enc":"random"}`), &resp); err == nil {
			t.Errorf("expected unmarshal to fail")
		}
	})
}
