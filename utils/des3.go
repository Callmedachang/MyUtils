//加密工具类，用了3des和base64
package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

const key = `h*.d;cy7x_12dkx?#j39fdl!` //api数据加密、解密key
//des3 + base64 encrypt
func DesBase64Encrypt(origData []byte) []byte {
	result, err := TripleDesEncrypt(origData, []byte(key))
	if err != nil {
		panic(err)
	}
	return []byte(base64.StdEncoding.EncodeToString(result))
}

func DesBase64Decrypt(crypted []byte) []byte {
	result, err := base64.StdEncoding.DecodeString(string(crypted))
	if err != nil {
		panic(err)
	}
	origData, err := TripleDesDecrypt(result, []byte(key))
	if err != nil {
		panic(err)
	}
	return origData
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
