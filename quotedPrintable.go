// quotedPrintable implements an encoder for the quoted-printable
// wire format defined in RFC 2045.
package quotedPrintable

import (
  "strings"
  "encoding/hex"
)

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
