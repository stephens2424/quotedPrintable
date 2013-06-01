// quotedPrintable implements an encoder for the quoted-printable
// wire format defined in RFC 2045.
package quotedPrintable

import (
	"encoding/hex"
	"strings"
)

func NewEncoder() QuotedPrintableEncoder {
	return QuotedPrintableEncoder{}
}

type QuotedPrintableEncoder struct{}

func (encoder QuotedPrintableEncoder) TransferEncodingType() string {
	return "quoted-printable"
}

func (encoder QuotedPrintableEncoder) Encode(src []byte) []byte {
	return Encode(src)
}

// Encodes bytes to quoted printable bytes. The resulting slice may
// be longer than the input slice.
func Encode(src []byte) []byte {
	var dstBuilder []byte
	if len(src) == 0 {
		return dstBuilder
	}

	lineCounter := 0
	lineMax := 76
	for _, b := range src {
		encoded := quotedPrintableEncodeByte(b)
		lineCounter += len(encoded)
		if lineCounter > lineMax {
			dstBuilder = append(dstBuilder, []byte("=\r\n")...)
			lineCounter = len(encoded)
		}
		dstBuilder = append(dstBuilder, encoded...)
	}

	return dstBuilder
}

// Encodes bytes to a quoted printable string.
func EncodeToString(src []byte) string {
	return string(Encode(src))
}

func quotedPrintableEncodeByte(b byte) []byte {
	if b == 9 || b == 32 || (b >= 33 && b <= 60) || (b >= 62 && b <= 126) {
		return []byte{b}
	}
	hex := "=" + strings.ToUpper(hex.EncodeToString([]byte{b}))
	return []byte(hex)
}
