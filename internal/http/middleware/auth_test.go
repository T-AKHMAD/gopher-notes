package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBearerToken(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   string
		wantOK bool
	}{

		{
			name:   "no header",
			header: "",
			want:   "",
			wantOK: false,
		},
		{
			name:   "bearer ok",
			header: "Bearer abc",
			want:   "abc",
			wantOK: true,
		},
		{
			name:   "bearer trims spaces",
			header: "Bearer    abc   ",
			want:   "abc",
			wantOK: true,
		},
		{
			name:   "wrong scheme",
			header: "Basic abc",
			want:   "",
			wantOK: false,
		},
		{
			name:   "bearer empty token",
			header: "Bearer ",
			want:   "",
			wantOK: false,
		},
		{
			name:   "bearer no space after prefix",
			header: "Bearerabc",
			want:   "",
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			if tt.header != "" {
				r.Header.Set("Authorization", tt.header)
			}

			got, ok := BearerToken(r)
			if ok != tt.wantOK {
				t.Fatalf("ok=%v, want %v", ok, tt.wantOK)
			}
			if got != tt.want {
				t.Fatalf("token=%q, want %q", got, tt.want)
			}
		})
	}
}
