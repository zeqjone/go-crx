package crx

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io/ioutil"

	"github/zeqjone/crx/pb"

	"github.com/golang/protobuf/proto"
)

const symbols = "abcdefghijklmnopqrstuvwxyz"

// ID returns the extension id.
func ID(filename string) (id string, err error) {
	if !isCRX3(filename) {
		if !isCRX2(filename) {
			return id, ErrUnsupportedFileFormat
		} else {
			return idCrx2(filename)
		}
	}

	crx, err := ioutil.ReadFile(filename)
	if err != nil {
		return id, err
	}

	var (
		headerSize = binary.LittleEndian.Uint32(crx[8:12])
		metaSize   = uint32(12)
		v          = crx[metaSize : headerSize+metaSize]
		header     pb.CrxFileHeader
		signedData pb.SignedData
	)

	if err := proto.Unmarshal(v, &header); err != nil {
		return id, err
	}
	if err := proto.Unmarshal(header.SignedHeaderData, &signedData); err != nil {
		return id, err
	}

	idx := strIDx()
	sid := fmt.Sprintf("%x", signedData.CrxId[:16])
	buf := bytes.NewBuffer(nil)
	for _, char := range sid {
		index := idx[char]
		buf.WriteString(string(symbols[index]))
	}
	return buf.String(), nil
}

func strIDx() map[rune]int {
	index := make(map[rune]int)
	src := "0123456789abcdef"
	for i, char := range src {
		index[char] = i
	}
	return index
}

func idCrx2(filename string) (id string, err error) {
	crx, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var (
		pkLenth    = binary.LittleEndian.Uint32(crx[8:12])
		sigLenth   = binary.LittleEndian.Uint32(crx[12:16])
		metaSize   = uint32(16)
		headLength = pkLenth + sigLenth + metaSize
		v          = crx[metaSize:headLength]
	)
	hasded := sha256.Sum256(v[:pkLenth])
	strhashed := fmt.Sprintf("%x", hasded)

	buf := bytes.NewBuffer(nil)
	idx := strIDx()
	for _, char := range strhashed[:32] {
		index := idx[char]
		buf.WriteString(string(symbols[index]))
	}
	return buf.String(), nil
}
