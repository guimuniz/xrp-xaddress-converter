package xaddress

// EncodedAddressResponse represents the result of encoding a classic address to X-Address format.
type EncodedAddressResponse struct {
	Address string `json:"address"`
}

// DecodedAddressResponse represents the result of decoding an X-Address to classic address format.
type DecodedAddressResponse struct {
	Account string `json:"account"` // Classic r-address
	Tag     string `json:"tag"`     // Destination tag (empty if none)
	Test    bool   `json:"test"`    // Whether this is a testnet address
}
