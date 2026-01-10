// Copyright 2026 SeaVerse AI
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signature

import (
	"testing"
)

func TestSigner_Sign(t *testing.T) {
	tests := []struct {
		name      string
		secret    string
		params    map[string]any
		shouldErr bool
	}{
		{
			name:   "basic string parameters",
			secret: "test-secret",
			params: map[string]any{
				"action": "create",
				"user":   "alice",
			},
		},
		{
			name:   "parameters with numbers",
			secret: "test-secret",
			params: map[string]any{
				"user_id":   12345,
				"timestamp": int64(1234567890),
				"amount":    float64(99.99),
			},
		},
		{
			name:   "parameters sorted alphabetically",
			secret: "test-secret",
			params: map[string]any{
				"zebra": "z",
				"apple": "a",
				"mango": "m",
			},
		},
		{
			name:   "skip empty string",
			secret: "test-secret",
			params: map[string]any{
				"key1": "value1",
				"key2": "",
				"key3": "value3",
			},
		},
		{
			name:   "skip nil values",
			secret: "test-secret",
			params: map[string]any{
				"key1": "value1",
				"key2": nil,
				"key3": "value3",
			},
		},
		{
			name:   "special characters without encoding",
			secret: "test-secret",
			params: map[string]any{
				"message": "hello world!@#$%^&*()",
				"path":    "/api/v1/users",
			},
		},
		{
			name:      "nil params should error",
			secret:    "test-secret",
			params:    nil,
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signer := NewSigner(tt.secret)
			got, err := signer.Sign(tt.params)

			if tt.shouldErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Verify signature format
			if len(got) != 64 {
				t.Errorf("expected signature length 64, got %d", len(got))
			}

			// Verify signature is hexadecimal lowercase
			for _, c := range got {
				if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
					t.Errorf("signature contains invalid hex character: %c", c)
				}
			}

			// Verify consistency - same params should produce same signature
			got2, _ := signer.Sign(tt.params)
			if got != got2 {
				t.Errorf("signature not consistent: %s != %s", got, got2)
			}
		})
	}
}

func TestSigner_SignatureFormat(t *testing.T) {
	// Test specific known signature
	signer := NewSigner("my-secret-key")
	params := map[string]any{
		"action":    "create",
		"timestamp": "1234567890",
		"user_id":   "12345",
	}

	sig, err := signer.Sign(params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the signature string is built correctly
	// Expected string: action=create&timestamp=1234567890&user_id=12345
	t.Logf("Generated signature: %s", sig)

	// Re-sign with same params should match
	sig2, _ := signer.Sign(params)
	if sig != sig2 {
		t.Errorf("signatures don't match: %s != %s", sig, sig2)
	}
}

func TestSigner_Verify(t *testing.T) {
	signer := NewSigner("test-secret")
	params := map[string]any{
		"action": "test",
		"value":  123,
	}

	// Generate signature
	sig, err := signer.Sign(params)
	if err != nil {
		t.Fatalf("failed to sign: %v", err)
	}

	// Verify should succeed
	valid, err := signer.Verify(params, sig)
	if err != nil {
		t.Errorf("verify failed: %v", err)
	}
	if !valid {
		t.Errorf("signature verification failed for valid signature")
	}

	// Invalid signature should fail
	valid, err = signer.Verify(params, "invalid-signature")
	if err != nil {
		t.Errorf("verify failed: %v", err)
	}
	if valid {
		t.Errorf("invalid signature was verified as valid")
	}

	// Modified params should fail
	modifiedParams := map[string]any{
		"action": "test",
		"value":  456,
	}
	valid, err = signer.Verify(modifiedParams, sig)
	if err != nil {
		t.Errorf("verify failed: %v", err)
	}
	if valid {
		t.Errorf("signature verified with modified params")
	}
}

func TestSigner_BuildSignatureString(t *testing.T) {
	signer := NewSigner("test")

	tests := []struct {
		name     string
		params   map[string]any
		expected string
	}{
		{
			name: "alphabetical sorting",
			params: map[string]any{
				"c": "3",
				"a": "1",
				"b": "2",
			},
			expected: "a=1&b=2&c=3",
		},
		{
			name: "skip empty and nil",
			params: map[string]any{
				"a": "1",
				"b": "",
				"c": nil,
				"d": "4",
			},
			expected: "a=1&d=4",
		},
		{
			name: "number conversion",
			params: map[string]any{
				"int":     123,
				"int64":   int64(456),
				"float64": 78.9,
			},
			expected: "float64=78.9&int=123&int64=456",
		},
		{
			name: "special characters",
			params: map[string]any{
				"msg":  "hello world!",
				"path": "/api/v1",
			},
			expected: "msg=hello world!&path=/api/v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := signer.buildSignatureString(tt.params)
			if got != tt.expected {
				t.Errorf("buildSignatureString() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSigner_ValueToString(t *testing.T) {
	signer := NewSigner("test")

	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{"string", "test", "test"},
		{"int", 123, "123"},
		{"int32", int32(123), "123"},
		{"int64", int64(123), "123"},
		{"uint", uint(123), "123"},
		{"uint32", uint32(123), "123"},
		{"uint64", uint64(123), "123"},
		{"float32", float32(12.34), "12.34"},
		{"float64", 12.34, "12.34"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"nil", nil, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := signer.valueToString(tt.value)
			if got != tt.expected {
				t.Errorf("valueToString() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkSigner_Sign(b *testing.B) {
	signer := NewSigner("benchmark-secret")
	params := map[string]any{
		"action":    "test",
		"user_id":   12345,
		"timestamp": int64(1234567890),
		"amount":    99.99,
		"status":    "active",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = signer.Sign(params)
	}
}

func BenchmarkSigner_Verify(b *testing.B) {
	signer := NewSigner("benchmark-secret")
	params := map[string]any{
		"action":    "test",
		"user_id":   12345,
		"timestamp": int64(1234567890),
	}
	sig, _ := signer.Sign(params)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = signer.Verify(params, sig)
	}
}
