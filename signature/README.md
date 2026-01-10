# Signature Package

[![Go Reference](https://pkg.go.dev/badge/github.com/SeaVerseAI/seaverse-go/signature.svg)](https://pkg.go.dev/github.com/SeaVerseAI/seaverse-go/signature)

Go package for HMAC-SHA256 signature generation and verification.

## Installation

```bash
go get github.com/SeaVerseAI/seaverse-go/signature@latest
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    signature "github.com/SeaVerseAI/seaverse-go/signature/v1"
)

func main() {
    // Create a signer with your secret key
    signer := signature.NewSigner("your-secret-key")

    // Prepare parameters
    params := map[string]any{
        "user_id":   12345,
        "action":    "create",
        "timestamp": "1234567890",
    }

    // Generate signature
    sig, err := signer.Sign(params)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Signature: %s\n", sig)

    // Verify signature
    valid, err := signer.Verify(params, sig)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Valid: %v\n", valid)
}
```

## Signature Algorithm Specification

The signature implementation follows these rules:

- ✓ **Algorithm**: HMAC-SHA256
- ✓ **Encoding**: UTF-8
- ✓ **Sorting**: Parameters sorted alphabetically by key (a-z)
- ✓ **Special characters**: Use original values without URL encoding
- ✓ **Empty values**: Empty strings, nil values are excluded from signature
- ✓ **Numbers**: All numeric types are converted to strings
- ✓ **Format**: `key1=value1&key2=value2`
- ✓ **Output**: Hexadecimal lowercase (64 characters)

## Examples

### Basic Signing

```go
signer := signature.NewSigner("my-secret-key")

params := map[string]any{
    "action":    "purchase",
    "user_id":   12345,
    "amount":    99.99,
    "timestamp": "1234567890",
}

sig, err := signer.Sign(params)
// sig: "a1b2c3d4e5f6..." (64 hex characters)
```

### Signature Verification

```go
valid, err := signer.Verify(params, receivedSignature)
if !valid {
    // Handle invalid signature
}
```

### Parameter Types

The package supports various parameter types:

```go
params := map[string]any{
    "string_val":  "hello",
    "int_val":     123,
    "int64_val":   int64(456),
    "float_val":   99.99,
    "bool_val":    true,
    "empty_val":   "",     // excluded from signature
    "nil_val":     nil,    // excluded from signature
}
```

## Security Best Practices

1. **Never expose your secret key** in client-side code or public repositories
2. **Use environment variables** or secure configuration management for secret keys
3. **Use HTTPS** when transmitting signatures over the network
4. **Implement timestamp validation** to prevent replay attacks
5. **Use unique request IDs** (nonce) for additional security

## Example with Timestamp Validation

```go
import (
    "strconv"
    "time"
)

func createSignedRequest(signer *signature.Signer, action string) (map[string]any, string, error) {
    params := map[string]any{
        "action":    action,
        "timestamp": strconv.FormatInt(time.Now().Unix(), 10),
        "nonce":     generateNonce(), // implement your nonce generator
    }

    sig, err := signer.Sign(params)
    return params, sig, err
}

func validateTimestamp(timestamp string, maxAge time.Duration) bool {
    ts, err := strconv.ParseInt(timestamp, 10, 64)
    if err != nil {
        return false
    }

    requestTime := time.Unix(ts, 0)
    return time.Since(requestTime) <= maxAge
}
```

## Testing

Run tests:

```bash
cd signature
go test -v ./...
```

Run benchmarks:

```bash
go test -bench=. -benchmem ./...
```

## License

Apache 2.0
