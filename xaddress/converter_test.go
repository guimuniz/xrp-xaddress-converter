package xaddress

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeAndDecodeXAddress(t *testing.T) {
	address := "rng5ZxeWue9pggAPuGHZKXkYQQBmspKTdZ"
	tag := "2282235929"
	xAddress := "X75FD6PAsxLqhNCnHEUpYxWmfBE84hGXM4CriEbwFxxhrww"

	// Test encoding mainnet address with tag
	encodedAddress, err := EncodeAddressToXAddress(context.Background(), address, tag, false)
	assert.NoError(t, err)
	assert.Equal(t, xAddress, encodedAddress.Address)

	// Test decoding X-address back to classic address and tag
	decodedXAddress, err := DecodeXAddressToAddress(context.Background(), xAddress)
	assert.NoError(t, err)
	assert.Equal(t, address, decodedXAddress.Account)
	assert.Equal(t, tag, decodedXAddress.Tag)
	assert.Equal(t, false, decodedXAddress.Test)
}

func TestEncodeAddressWithoutTag(t *testing.T) {
	address := "rU6K7V3Po4snVhBBaU29sesqs2qTQJWDw1"

	// Test encoding without tag
	encodedAddress, err := EncodeAddressToXAddress(context.Background(), address, "", false)
	assert.NoError(t, err)
	assert.NotEmpty(t, encodedAddress.Address)

	// Decode and verify no tag
	decodedXAddress, err := DecodeXAddressToAddress(context.Background(), encodedAddress.Address)
	assert.NoError(t, err)
	assert.Equal(t, address, decodedXAddress.Account)
	assert.Equal(t, "", decodedXAddress.Tag)
}

func TestEncodeTestnetAddress(t *testing.T) {
	address := "rng5ZxeWue9pggAPuGHZKXkYQQBmspKTdZ"
	tag := "12345"

	// Test encoding testnet address
	encodedAddress, err := EncodeAddressToXAddress(context.Background(), address, tag, true)
	assert.NoError(t, err)
	assert.NotEmpty(t, encodedAddress.Address)

	// Decode and verify it's marked as testnet
	decodedXAddress, err := DecodeXAddressToAddress(context.Background(), encodedAddress.Address)
	assert.NoError(t, err)
	assert.Equal(t, address, decodedXAddress.Account)
	assert.Equal(t, tag, decodedXAddress.Tag)
	assert.Equal(t, true, decodedXAddress.Test)
}

func TestInvalidTag(t *testing.T) {
	address := "rU6K7V3Po4snVhBBaU29sesqs2qTQJWDw1"

	// Test with invalid tag (not a number)
	_, err := EncodeAddressToXAddress(context.Background(), address, "invalid", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid tag")
}

func TestInvalidXAddress(t *testing.T) {
	// Test decoding an invalid X-address
	_, err := DecodeXAddressToAddress(context.Background(), "InvalidXAddress123")
	assert.Error(t, err)
}

// ExampleEncodeAddressToXAddress demonstrates encoding a classic address with tag to X-Address.
func ExampleEncodeAddressToXAddress() {
	result, _ := EncodeAddressToXAddress(
		context.Background(),
		"rng5ZxeWue9pggAPuGHZKXkYQQBmspKTdZ",
		"2282235929",
		false,
	)
	fmt.Println(result.Address)
	// Output: X75FD6PAsxLqhNCnHEUpYxWmfBE84hGXM4CriEbwFxxhrww
}

// ExampleDecodeXAddressToAddress demonstrates decoding an X-Address to classic format.
func ExampleDecodeXAddressToAddress() {
	result, _ := DecodeXAddressToAddress(
		context.Background(),
		"X75FD6PAsxLqhNCnHEUpYxWmfBE84hGXM4CriEbwFxxhrww",
	)
	fmt.Printf("%s:%s", result.Account, result.Tag)
	// Output: rng5ZxeWue9pggAPuGHZKXkYQQBmspKTdZ:2282235929
}
