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

## API

### `EncodeAddressToXAddress`

Converts a classic XRP address and destination tag to X-Address format.

```go
func EncodeAddressToXAddress(ctx context.Context, address, tag string, isTestnet bool) (*EncodedAddressResponse, error)
```

**Parameters:**
- `ctx`: Context for cancellation control
- `address`: Classic XRP address (r-address format)
- `tag`: Destination tag as decimal string (use "" for no tag)
- `isTestnet`: `true` for testnet, `false` for mainnet

**Returns:**
- `*EncodedAddressResponse`: Contains the generated X-Address
- `error`: Error if conversion fails

### `DecodeXAddressToAddress`

Converts an X-Address back to classic address and destination tag.

```go
func DecodeXAddressToAddress(ctx context.Context, xAddress string) (*DecodedAddressResponse, error)
```

**Parameters:**
- `ctx`: Context for cancellation control
- `xAddress`: The X-Address to decode

**Returns:**
- `*DecodedAddressResponse`: Contains the classic address, tag, and testnet flag
- `error`: Error if decoding fails

## Types

```go
type EncodedAddressResponse struct {
    Address string `json:"address"`
}

type DecodedAddressResponse struct {
    Account string `json:"account"` // Classic address
    Tag     string `json:"tag"`     // Tag (empty if none)
    Test    bool   `json:"test"`    // true if testnet
}
```

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

## License

MIT License - see the [LICENSE](LICENSE) file for details.

## References

- [XRP Ledger X-Address Format](https://xrpaddress.info/)
- [XRPL Address Encoding](https://xrpl.org/accounts.html#addresses)
