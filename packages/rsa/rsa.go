package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func main() {
	// 密文
	ciphertextBase64 := "your_encrypted_base64_string"

	// 密文解码
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		fmt.Println("Error decoding ciphertext:", err)
		return
	}

	// 加载私钥
	privateKeyPEM := []byte(`-----BEGIN RSA PRIVATE KEY-----
    your_private_key_here
    -----END RSA PRIVATE KEY-----`)

	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	if privateKeyBlock == nil {
		fmt.Println("Error decoding private key")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}

	// 解密
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	// 输出解密后的明文
	fmt.Println("Decrypted text:", string(plaintext))
}
