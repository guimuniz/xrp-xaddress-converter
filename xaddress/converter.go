package xaddress

import (
	"context"
	"fmt"
	"strconv"
)

// EncodeAddressToXAddress encodes a classic r-address and destination tag into an X-Address.
// The tag parameter must be a base-10 string for a 32-bit unsigned integer.
// Pass an empty string for tag if no destination tag is needed.
// isTestnet indicates whether to use testnet (true) or mainnet (false) encoding.
func EncodeAddressToXAddress(_ context.Context, address, tag string, isTestnet bool) (*EncodedAddressResponse, error) {
	var tagPtr *uint32
	if tag != "" {
		t64, err := strconv.ParseUint(tag, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid tag: %w", err)
		}
		t := uint32(t64)
		tagPtr = &t
	}

	x, err := EncodeXAddress(address, tagPtr, isTestnet)
	if err != nil {
		return nil, err
	}
	return &EncodedAddressResponse{Address: x}, nil
}

// DecodeXAddressToAddress decodes an X-Address into its classic address and destination tag components.
func DecodeXAddressToAddress(_ context.Context, xAddress string) (*DecodedAddressResponse, error) {
	classic, tagPtr, isTest, err := DecodeXAddress(xAddress)
	if err != nil {
		return nil, err
	}
	var tagStr string
	if tagPtr != nil {
		tagStr = strconv.FormatUint(uint64(*tagPtr), 10)
	}
	return &DecodedAddressResponse{
		Account: classic,
		Tag:     tagStr,
		Test:    isTest,
	}, nil
}
