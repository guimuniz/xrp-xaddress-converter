package xaddress

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	addresscodec "github.com/xyield/xrpl-go/address-codec"
)

const (
	accountIDLen = addresscodec.AccountAddressLength
	prefixLen    = 2
	xPayloadLen  = 31

	flagIndex     = prefixLen + accountIDLen
	tagStartIndex = flagIndex + 1
)

func encodeAccountID(accountID []byte) (string, error) {
	if len(accountID) != accountIDLen {
		return "", fmt.Errorf("accountID must be %d bytes, got %d", accountIDLen, len(accountID))
	}
	out := addresscodec.Encode(accountID, []byte{addresscodec.AccountAddressPrefix}, accountIDLen)
	if out == "" {
		return "", errors.New("failed to encode classic address")
	}
	return out, nil
}

func decodeAccountID(classic string) ([]byte, error) {
	_, accountID, err := addresscodec.DecodeClassicAddressToAccountID(classic)
	if err != nil {
		return nil, err
	}
	if len(accountID) != accountIDLen {
		return nil, fmt.Errorf("unexpected_account_length: got %d, want %d", len(accountID), accountIDLen)
	}
	return accountID, nil
}

var (
	prefixMain = []byte{0x05, 0x44}
	prefixTest = []byte{0x04, 0x93}
)

// EncodeXAddress builds an X-address from a classic address, optional tag, and network flag.
// If tag is nil, no destination tag is encoded. If provided, tag is encoded as 32-bit little-endian.
func EncodeXAddress(classic string, tag *uint32, isTest bool) (string, error) {
	accountID, err := decodeAccountID(classic)
	if err != nil {
		return "", fmt.Errorf("invalid classic address: %w", err)
	}

	prefix := prefixMain
	if isTest {
		prefix = prefixTest
	}

	payload := make([]byte, 0, xPayloadLen)

	payload = append(payload, prefix...)
	payload = append(payload, accountID...)

	var flag byte
	var tagLE [4]byte
	if tag != nil {
		flag = 1
		binary.LittleEndian.PutUint32(tagLE[:], *tag)
	}

	payload = append(payload, flag)
	payload = append(payload, tagLE[:]...)
	payload = append(payload, make([]byte, 4)...)

	return addresscodec.Base58CheckEncode(payload), nil
}

// DecodeXAddress decodes an X-address into its classic address, optional destination tag, and network flag.
func DecodeXAddress(xaddr string) (classic string, tag *uint32, isTest bool, err error) {
	payload, err := addresscodec.Base58CheckDecode(xaddr)
	if err != nil {
		return "", nil, false, err
	}

	if len(payload) != xPayloadLen {
		return "", nil, false, fmt.Errorf("unexpected_payload_length: got %d, want %d", len(payload), xPayloadLen)
	}

	switch {
	case bytes.Equal(payload[:prefixLen], prefixMain):
		isTest = false
	case bytes.Equal(payload[:prefixLen], prefixTest):
		isTest = true
	default:
		return "", nil, false, errors.New("invalid X-address: bad prefix")
	}

	accountID := payload[prefixLen : prefixLen+accountIDLen]
	flag := payload[flagIndex]

	switch flag {
	case 0:
		for i := tagStartIndex; i < xPayloadLen; i++ {
			if payload[i] != 0 {
				return "", nil, false, errors.New("remaining bytes must be zero")
			}
		}
		tag = nil
	case 1:
		tmp := binary.LittleEndian.Uint32(payload[tagStartIndex : tagStartIndex+4])
		v := tmp
		tag = &v
	default:
		return "", nil, false, errors.New("unsupported X-address flag (64-bit tags not supported)")
	}

	classic, err = encodeAccountID(accountID)
	if err != nil {
		return "", nil, false, err
	}

	return classic, tag, isTest, nil
}
