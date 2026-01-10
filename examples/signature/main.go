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

package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	signature "github.com/SeaVerseAI/seaverse-sdk-go/signature/v1"
)

func main() {
	// Example 1: Basic Usage
	fmt.Println("=== Example 1: Basic Signature Generation ===")
	basicExample()

	fmt.Println()

	// Example 2: Signature Verification
	fmt.Println("=== Example 2: Signature Verification ===")
	verificationExample()

	fmt.Println()

	// Example 3: Real-world API Request
	fmt.Println("=== Example 3: API Request with Timestamp ===")
	apiRequestExample()
}

func basicExample() {
	// Create a signer with your secret key
	// Note: In production, use environment variables for secrets
	signer := signature.NewSigner("my-secret-key")

	// Prepare request parameters
	params := map[string]any{
		"user_id": 12345,
		"action":  "create",
		"amount":  99.99,
	}

	// Generate signature
	sig, err := signer.Sign(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parameters: %+v\n", params)
	fmt.Printf("Signature:  %s\n", sig)
}

func verificationExample() {
	signer := signature.NewSigner("verification-secret")

	// Original request parameters
	params := map[string]any{
		"action": "purchase",
		"amount": 199.99,
		"item":   "product-123",
	}

	// Generate signature
	sig, err := signer.Sign(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Original Signature: %s\n", sig)

	// Verify the signature
	valid, err := signer.Verify(params, sig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Signature Valid: %v\n", valid)

	// Try verifying with modified params
	modifiedParams := map[string]any{
		"action": "purchase",
		"amount": 299.99, // Changed amount
		"item":   "product-123",
	}

	valid, err = signer.Verify(modifiedParams, sig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Modified Params Valid: %v (expected: false)\n", valid)
}

func apiRequestExample() {
	signer := signature.NewSigner("api-secret-key")

	// Create API request with timestamp (to prevent replay attacks)
	params := map[string]any{
		"api_key":   "public-api-key-123",
		"action":    "query",
		"user_id":   67890,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	// Generate signature
	sig, err := signer.Sign(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("API Request:\n")
	fmt.Printf("  User ID:    %v\n", params["user_id"])
	fmt.Printf("  Action:     %v\n", params["action"])
	fmt.Printf("  Timestamp:  %v\n", params["timestamp"])
	fmt.Printf("  Signature:  %s\n", sig)

	// Simulate server-side verification
	fmt.Println("\nServer-side verification:")

	// Check timestamp freshness (within 5 minutes)
	ts, _ := strconv.ParseInt(params["timestamp"].(string), 10, 64)
	requestTime := time.Unix(ts, 0)
	age := time.Since(requestTime)

	fmt.Printf("  Request age: %v\n", age)

	if age > 5*time.Minute {
		fmt.Println("  ❌ Request too old (replay attack?)")
		return
	}

	// Verify signature
	valid, err := signer.Verify(params, sig)
	if err != nil {
		log.Fatal(err)
	}

	if valid {
		fmt.Println("  ✅ Signature verified successfully")
		fmt.Println("  ✅ Request accepted")
	} else {
		fmt.Println("  ❌ Invalid signature")
		fmt.Println("  ❌ Request rejected")
	}
}
