package crx

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strings"

	"crx/pb"

	"github.com/golang/protobuf/proto"
)

// Unpack unpacks chrome extension into some directory.
func Unpack(filename string) error {
	if !isCRX3(filename) {
		if !isCRX2(filename) {
			return ErrUnsupportedFileFormat
		}
		return UnpackCrx2(filename)
	}
	crx, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	var (
		headerSize = binary.LittleEndian.Uint32(crx[8:12])
		metaSize   = uint32(12)
		v          = crx[metaSize : headerSize+metaSize]
		header     pb.CrxFileHeader
		signedData pb.SignedData
	)

	if err := proto.Unmarshal(v, &header); err != nil {
		return err
	}
	if err := proto.Unmarshal(header.SignedHeaderData, &signedData); err != nil {
		return err
	}

	if len(signedData.CrxId) != 16 {
		return ErrUnsupportedFileFormat
	}

	data := crx[len(v)+int(metaSize):]
	reader := bytes.NewReader(data)
	size := int64(len(data))

	unpacked := strings.TrimRight(filename, crxExt)
	return Unzip(reader, size, unpacked)
}

type crx2Header struct {
	Flag          string
	FormatVersion uint32
	PublicKey     string
	Signature     string
}

// unpackCrx2
// crx 打包格式文档：http://www.dre.vanderbilt.edu/~schmidt/android/android-4.0/external/chromium/chrome/common/extensions/docs/crx.html
func UnpackCrx2(filename string) error {
	crx, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	var (
		pkLenth    = binary.LittleEndian.Uint32(crx[8:12])
		sigLenth   = binary.LittleEndian.Uint32(crx[12:16])
		metaSize   = uint32(16)
		headLength = pkLenth + sigLenth + metaSize
		v          = crx[metaSize:headLength]
	)
	header := &crx2Header{
		Flag:          string(crx[0:4]),
		FormatVersion: binary.LittleEndian.Uint32(crx[4:8]),
		PublicKey:     fmt.Sprintf("%x", v[:pkLenth]),
		Signature:     fmt.Sprintf("%x", v[pkLenth:]),
	}
	fmt.Printf("header: %#v", header)

	data := crx[headLength:]
	reader := bytes.NewReader(data)
	size := int64(len(data))
	unpacked := strings.TrimRight(filename, crxExt)

	// hasded := sha256.Sum256([]byte("hello world\n"))

	return Unzip(reader, size, unpacked)
}
