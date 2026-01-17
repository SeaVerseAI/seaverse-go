// Copyright 2026 SeaVerse AI
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

// DefaultEndpoint is the default base URL for the account API
const DefaultEndpoint = "https://account.seaverse.ai"

// NewDefaultClient creates a new client with the default endpoint.
// This is a convenience function that uses DefaultEndpoint.
//
// Example:
//
//	client, err := v1.NewDefaultClient()
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewDefaultClient(opts ...ClientOption) (*Client, error) {
	return NewClient(DefaultEndpoint, opts...)
}
