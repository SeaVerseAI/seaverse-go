# account - SeaVerse Go SDK

SeaVerse Account Hub APIs for authentication and user management

> Part of the [seaverse-go](https://github.com/SeaVerseAI/seaverse-go) SDK family

## Installation

```bash
go get github.com/SeaVerseAI/seaverse-go/account/v1@latest
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    merchant "github.com/SeaVerseAI/seaverse-go/account/v1"
)

func main() {
    // Use the default client (recommended)
    client, err := merchant.NewDefaultClient()
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Example: Verify user token
    // Add your API calls here
}
```

## Default Endpoint

The default endpoint is: **`https://account.seaverse.ai`**

You can override it by creating a custom client:

```go
client, err := merchant.NewClient("https://custom.endpoint.com")
```

## Features

- ✅ **Type-safe** - Full type safety with Go generics
- ✅ **Context-aware** - Proper context.Context support
- ✅ **Idiomatic** - Follows Go best practices
- ✅ **Lightweight** - Minimal dependencies
- ✅ **Independent versioning** - Semantic versioning per package

## Documentation

For full API documentation, see the [SeaVerse Account Hub API Docs](https://account.seaverse.ai/docs).

## Versioning

This package follows [Semantic Versioning](https://semver.org/) and is independently versioned from other seaverse-go packages.

## Auto-Generated

⚠️ This SDK is auto-generated from OpenAPI specifications. **Do not edit manually.**

## Contributing

Contributions welcome! See [CONTRIBUTING.md](../../CONTRIBUTING.md) for guidelines.

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](../../LICENSE) for details.

---
Generated from [account-hub](https://github.com/SeaVerseAI/account-hub) OpenAPI specification • Copyright © 2026 SeaVerse AI
