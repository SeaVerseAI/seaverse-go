# SeaVerse SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/SeaVerseAI/seaverse-go.svg)](https://pkg.go.dev/github.com/SeaVerseAI/seaverse-go)

Go packages for SeaVerse AI Platform services.

## Installation

Install individual packages as needed:

```bash
# Install signature package
go get github.com/SeaVerseAI/seaverse-go/signature@latest
```

## Available Packages

### [Signature](./signature)

HMAC-SHA256 signature generation and verification.

```go
import signature "github.com/SeaVerseAI/seaverse-go/signature/v1"

signer := signature.NewSigner("your-secret-key")
sig, err := signer.Sign(params)
```

[Documentation](./signature/README.md)

## Go Versions Supported

This library is compatible with Go 1.21 and above.

## Project Structure

This repository is organized as a monorepo, similar to [google-cloud-go](https://github.com/googleapis/google-cloud-go). Each package is independently versioned and can be installed separately:

```
seaverse-sdk-go/
├── signature/           # Signature package
│   ├── v1/             # Version 1 implementation
│   ├── go.mod          # Independent module
│   └── README.md
├── go.mod              # Root module
└── README.md
```

Each package has its own `go.mod` file and can be versioned independently.

## Contributing

Contributions are welcome. Please ensure:

1. No security keys or sensitive URLs in code
2. Follow the existing package structure
3. Add tests for new features
4. Update documentation

## Security

⚠️ **Important**: This is a public repository. Never commit:
- API keys
- Secret keys
- Private URLs
- Credentials of any kind

Use environment variables or secure configuration management for sensitive data.

## License

Apache 2.0
