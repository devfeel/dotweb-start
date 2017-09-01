package encoding

import "encoding/base64"

//Base64Encode returns the base64 encoding of src.
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//Base64Decode returns the bytes represented by the base64 string s.
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
