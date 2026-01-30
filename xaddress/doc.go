// Package xaddress provides utilities for converting between XRP Ledger classic addresses
// (r-addresses with destination tags) and X-Addresses.
//
// X-Addresses are a newer address format that combines the classic address and destination tag
// into a single value, simplifying the payment process and reducing errors caused by missing
// or incorrect destination tags.
//
// # Example Usage
//
// Encoding a classic address to X-Address:
//
//	encoded, err := xaddress.EncodeAddressToXAddress(ctx, "rU6K7V3Po4snVhBBaU29sesqs2qTQJWDw1", "12345", false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(encoded.Address) // X-Address format
//
// Decoding an X-Address to classic format:
//
//	decoded, err := xaddress.DecodeXAddressToAddress(ctx, "XVfC9CTCJh6GN2x8bnrw3LtdbqiVCUFyQVMzRrMGUZpokKH")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Address: %s, Tag: %s\n", decoded.Account, decoded.Tag)
//
// # Network Support
//
// The package supports both mainnet and testnet addresses. Use the isTestnet parameter
// to specify the network when encoding addresses.
package xaddress
