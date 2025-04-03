package auth

import (
	"reflect"
	"testing"
)

func TestApiKey(t *testing.T) {
	type test struct {
		name    string
		headers map[string][]string
		wantErr bool
		want    string
	}

	tests := []test{
		{name: "Empty headers",
			headers: make(map[string][]string, 0),
			wantErr: true,
			want:    "",
		},
		{name: "No Authorization headers",
			headers: map[string][]string{"Bla1": {"1"}, "Bla2": {"1"}},
			wantErr: true,
			want:    "",
		},
		{name: "Empty Authorization headers",
			headers: map[string][]string{"Authorization": {}, "Bla": {"2"}},
			wantErr: true,
			want:    "",
		},
		{name: "No bearer Authorization headers",
			headers: map[string][]string{"Authorization": {"12312"}},
			wantErr: true,
			want:    "",
		},
		{name: "Nice authorization headers",
			headers: map[string][]string{"Authorization": {"ApiKey 12312"}},
			wantErr: false,
			want:    "12312",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			header, err := GetAPIKey(tc.headers)

			// Check error condition
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// If we don't expect an error, check the returned value
			if !tc.wantErr {
				if !reflect.DeepEqual(tc.want, header) {
					t.Errorf("expected: %v, got: %v", tc.want, header)
				}
			}
		})
	}
}
