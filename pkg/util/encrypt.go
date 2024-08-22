package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var Encrypt *Encryption

// AES 对称加密
type Encryption struct {
	key []byte
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

func (k *Encryption) SetKey(key string) {
	k.key = []byte(key)
}

// PKCS7 填充
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

// 移除 PKCS7 填充
func unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// AES CBC 模式加密
func (k *Encryption) AseEncrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(k.key)
	if err != nil {
		return "", err
	}

	// 使用 CBC 模式需要一个 IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 对明文进行填充，使其长度成为块大小的整数倍
	plainTextBytes := pad([]byte(plainText), aes.BlockSize)

	// 创建 CBC 模式的加密器
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainTextBytes))

	// 加密数据
	blockMode.CryptBlocks(cipherText, plainTextBytes)

	// 返回 Base64 编码的加密结果，IV 和密文一起传输
	return base64.StdEncoding.EncodeToString(append(iv, cipherText...)), nil
}

// AES CBC 模式解密
func (k *Encryption) AseDecrypt(cipherText string) (string, error) {
	// 解析 Base64 编码的密文
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// 创建 AES 密码块
	block, err := aes.NewCipher(k.key)
	if err != nil {
		return "", err
	}

	// 分离 IV 和密文
	if len(cipherTextBytes) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	// 创建 CBC 模式的解密器
	blockMode := cipher.NewCBCDecrypter(block, iv)

	// 解密数据
	plainText := make([]byte, len(cipherTextBytes))
	blockMode.CryptBlocks(plainText, cipherTextBytes)

	// 移除填充并返回明文
	return string(unpad(plainText)), nil
}
