package auth

import (
	"net/http"
	"testing"
)

func TestGetAIPKey(t *testing.T) {
	cases := []struct {
		name           string
		header         http.Header
		expectErr      bool
		expectedApiKey string
	}{
		{
			name:           "Valid Header",
			header:         http.Header{"Authorization": []string{"ApiKey secret"}},
			expectErr:      false,
			expectedApiKey: "secret",
		},
		{
			name:           "Missing Authorization",
			header:         http.Header{},
			expectErr:      true,
			expectedApiKey: "",
		},
		{
			name:           "Missing ApiKey key",
			header:         http.Header{"Authorization": []string{}},
			expectErr:      true,
			expectedApiKey: "",
		},
		{
			name:           "Missing ApiKey value",
			header:         http.Header{"Authorization": []string{"ApiKey"}},
			expectErr:      true,
			expectedApiKey: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(c.header)
			if (err == nil) == c.expectErr {
				t.Errorf("Expected err == %v, but got %v", c.expectErr, err)
			}

			if apiKey != c.expectedApiKey {
				t.Errorf("Expected apiKey = %v, but got %v", c.expectedApiKey, apiKey)
			}
		})
	}
}
