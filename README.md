# XRP X-Address Converter

[![Go Reference](https://pkg.go.dev/badge/github.com/guimuniz/xrp-xaddress-converter.svg)](https://pkg.go.dev/github.com/guimuniz/xrp-xaddress-converter)
[![Go Report Card](https://goreportcard.com/badge/github.com/guimuniz/xrp-xaddress-converter)](https://goreportcard.com/report/github.com/guimuniz/xrp-xaddress-converter)

A Go library for converting between XRP Ledger classic addresses (r-address format) and X-Addresses.

## What are X-Addresses?

X-Addresses are a newer address format for the XRP Ledger that combines the classic address and destination tag into a single value. This simplifies the payment process and reduces errors caused by missing or incorrect destination tags.

**Example:**
- Classic address: `rng5ZxeWue9pggAPuGHZKXkYQQBmspKTdZ` with tag `2282235929`
- Equivalent X-Address: `X75FD6PAsxLqhNCnHEUpYxWmfBE84hGXM4CriEbwFxxhrww`

## Installation

```bash
go get github.com/guimuniz/xrp-xaddress-converter
```

## Usage

### Import the Package

```go
import (
    "context"
    "fmt"
    
    "github.com/guimuniz/xrp-xaddress-converter/xaddress"
)
```

### Convert Classic Address to X-Address

```go
// Convert address with tag
address := "rng5ZxeWue9pggAPuGHZKXkYQQBmspKTdZ"
tag := "2282235929"
isTestnet := false // true for testnet, false for mainnet

result, err := xaddress.EncodeAddressToXAddress(context.Background(), address, tag, isTestnet)
if err != nil {
    panic(err)
}

fmt.Println(result.Address) 
// Output: X75FD6PAsxLqhNCnHEUpYxWmfBE84hGXM4CriEbwFxxhrww
```

### Convert X-Address to Classic Address

```go
xAddr := "X75FD6PAsxLqhNCnHEUpYxWmfBE84hGXM4CriEbwFxxhrww"

result, err := xaddress.DecodeXAddressToAddress(context.Background(), xAddr)
if err != nil {
    panic(err)
}

fmt.Printf("Address: %s\n", result.Account)
// Output: Address: rng5ZxeWue9pggAPuGHZKXkYQQBmspKTdZ

fmt.Printf("Tag: %s\n", result.Tag)
// Output: Tag: 2282235929

fmt.Printf("Is Testnet: %v\n", result.Test)
// Output: Is Testnet: false
```

### Convert Address Without Tag

```go
address := "rU6K7V3Po4snVhBBaU29sesqs2qTQJWDw1"

// Use empty string for tag when there's no destination tag
result, err := xaddress.EncodeAddressToXAddress(context.Background(), address, "", false)
if err != nil {
    panic(err)
}

fmt.Println(result.Address)
// Output: XVfC9CTCJh6GN2x8bnrw3LtdbqiVCUFyQVMzRrMGUZpokKH
```

For detailed API documentation, see the [package documentation](https://pkg.go.dev/github.com/guimuniz/xrp-xaddress-converter/xaddress).

## Development

### Run Tests

```bash
make test
```

### Format Code

```bash
make fmt
```

### Build

```bash
make build
```

## Dependencies

- [github.com/xyield/xrpl-go](https://github.com/xyield/xrpl-go) - For Base58Check encoding and address operations

## Contributing

Contributions are welcome! Feel free to open issues or pull requests.

## Releasing a New Version

To publish a new version of this library:

1. Create a release branch from `develop`:
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b release/vX.Y.Z
   ```
2. Update the `CHANGELOG.md` with the new version and changes
3. Commit your changes:
   ```bash
   git add CHANGELOG.md
   git commit -m "chore: prepare release vX.Y.Z"
   git push origin release/vX.Y.Z
   ```
4. Open a pull request from `release/vX.Y.Z` to `master`
5. After PR approval and merge, create and push the version tag:
   ```bash
   git checkout master
   git pull origin master
   git tag -a vX.Y.Z -m "Release vX.Y.Z: Brief description"
   git push origin vX.Y.Z
   ```
6. The module will be automatically indexed by pkg.go.dev

Users can then install the specific version:
```bash
go get github.com/guimuniz/xrp-xaddress-converter@vX.Y.Z
```

## License

MIT License - see the [LICENSE](LICENSE) file for details.

## References

- [XRP Ledger X-Address Format](https://xrpaddress.info/)
- [XRPL Address Encoding](https://xrpl.org/accounts.html#addresses)
