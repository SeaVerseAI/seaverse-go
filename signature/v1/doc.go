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

// Package signature provides cryptographic signature utilities for SeaVerse SDK.
//
// # Installation
//
// To install this package:
//
//	go get github.com/SeaVerseAI/seaverse-go/signature@latest
//
// # Usage
//
// Basic usage for HMAC-SHA256 signing:
//
//	import signature "github.com/SeaVerseAI/seaverse-go/signature/v1"
//
//	// Create a signer with your secret key
//	signer := signature.NewSigner("your-secret-key")
//
//	// Prepare parameters
//	params := map[string]any{
//	    "user_id":   12345,
//	    "action":    "create",
//	    "timestamp": "1234567890",
//	}
//
//	// Generate signature
//	sig, err := signer.Sign(params)
//	if err != nil {
//	    // handle error
//	}
//
//	// Verify signature
//	valid, err := signer.Verify(params, sig)
//	if err != nil {
//	    // handle error
//	}
//
// # Signature Algorithm Specification
//
// The signature implementation follows these rules:
//
//   - Algorithm: HMAC-SHA256
//   - Encoding: UTF-8
//   - Sorting: Parameters sorted alphabetically by key (a-z)
//   - Special characters: Use original values without URL encoding
//   - Empty values: Empty strings, nil values are excluded from signature
//   - Numbers: All numeric types are converted to strings
//   - Format: key1=value1&key2=value2
//   - Output: Hexadecimal lowercase (64 characters)
//
// # Example
//
// Complete example showing all features:
//
//	package main
//
//	import (
//	    "fmt"
//	    "log"
//
//	    signature "github.com/SeaVerseAI/seaverse-go/signature/v1"
//	)
//
//	func main() {
//	    // Initialize signer
//	    signer := signature.NewSigner("my-secret-key")
//
//	    // Parameters with various types
//	    params := map[string]any{
//	        "user_id":   12345,                // int
//	        "amount":    99.99,                // float64
//	        "action":    "purchase",           // string
//	        "timestamp": "1234567890",         // string
//	        "empty":     "",                   // excluded
//	        "null":      nil,                  // excluded
//	    }
//
//	    // Generate signature
//	    signature, err := signer.Sign(params)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Signature: %s\n", signature)
//
//	    // Verify signature
//	    valid, err := signer.Verify(params, signature)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Valid: %v\n", valid)
//	}
package signature
