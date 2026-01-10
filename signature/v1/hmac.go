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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
)

// Signer provides HMAC-SHA256 signature generation and verification.
type Signer struct {
	secret string
}

// NewSigner creates a new Signer with the provided secret key.
//
// The secret key will be used as the HMAC key for all signing operations.
// Keep the secret key secure and never expose it in client-side code or public repositories.
func NewSigner(secret string) *Signer {
	return &Signer{
		secret: secret,
	}
}

// Sign generates a HMAC-SHA256 signature from the provided parameters.
//
// Signature generation follows these rules:
//   - Encoding: UTF-8
//   - Sorting: keys sorted alphabetically (a-z)
//   - Special characters: use original values without URL encoding
//   - Empty values: empty strings, nil values are excluded from signature
//   - Numbers: converted to strings
//   - Format: key1=value1&key2=value2
//   - Output: hexadecimal lowercase
//
// Example:
//
//	signer := signature.NewSigner("my-secret-key")
//	params := map[string]any{
//	    "user_id":   12345,
//	    "action":    "create",
//	    "timestamp": "1234567890",
//	}
//	sig, err := signer.Sign(params)
//	if err != nil {
//	    // handle error
//	}
//	// sig will be a 64-character hexadecimal string
func (s *Signer) Sign(params map[string]any) (string, error) {
	if params == nil {
		return "", fmt.Errorf("params cannot be nil")
	}

	// Build the signature string
	sigString := s.buildSignatureString(params)

	// Generate HMAC-SHA256
	h := hmac.New(sha256.New, []byte(s.secret))
	h.Write([]byte(sigString))

	// Return hex encoded signature
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Verify validates a signature against the provided parameters.
//
// Returns true if the signature is valid, false otherwise.
// Uses constant-time comparison to prevent timing attacks.
//
// Example:
//
//	valid, err := signer.Verify(params, receivedSignature)
//	if err != nil {
//	    // handle error
//	}
//	if !valid {
//	    // signature is invalid
//	}
func (s *Signer) Verify(params map[string]any, signature string) (bool, error) {
	expected, err := s.Sign(params)
	if err != nil {
		return false, err
	}
	return hmac.Equal([]byte(expected), []byte(signature)), nil
}

// buildSignatureString creates the canonical string for signing.
//
// Rules:
// 1. Sort keys alphabetically
// 2. Skip nil, empty string values
// 3. Convert all values to strings
// 4. Join with & separator in key=value format
func (s *Signer) buildSignatureString(params map[string]any) string {
	// Extract and sort keys
	keys := make([]string, 0, len(params))
	for k, v := range params {
		// Skip empty values (nil, empty string)
		if v == nil || v == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build signature string
	pairs := make([]string, 0, len(keys))
	for _, k := range keys {
		v := params[k]
		strValue := s.valueToString(v)

		// Skip if converted value is empty
		if strValue == "" {
			continue
		}

		pairs = append(pairs, fmt.Sprintf("%s=%s", k, strValue))
	}

	result := ""
	for i, pair := range pairs {
		if i > 0 {
			result += "&"
		}
		result += pair
	}

	return result
}

// valueToString converts various types to string for signature generation.
func (s *Signer) valueToString(v any) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}
