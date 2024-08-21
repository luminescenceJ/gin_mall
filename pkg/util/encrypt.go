package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

var Encrypt *Encryption

// AES 对称加密
type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	unpadNum := int(dst[len(dst)-1])
	if len(dst) < unpadNum {
		return nil, errors.New("填充字节数超过数据长度")
	}
	str := dst[:len(dst)-unpadNum]
	return str, nil
}

func (k *Encryption) AseEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		return src
	}
	NewSrcByte := PadPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

func (k *Encryption) SetKey(key string) {
	k.key = key
}
