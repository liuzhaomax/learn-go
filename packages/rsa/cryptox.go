package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
)

func MD5(byt []byte) string {
	hash := md5.New()
	_, _ = hash.Write(byt)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func MD5Str(str string) string {
	return MD5([]byte(str))
}

func SHA1(byt []byte) string {
	hash := sha1.New()
	_, _ = hash.Write(byt)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func SHA1Str(str string) string {
	return SHA1([]byte(str))
}

func SHA1MD5Str(str string) string {
	return SHA1Str(MD5Str(str))
}

func BASE64Encode(byt []byte) string {
	encoded := base64.StdEncoding.EncodeToString(byt)
	return encoded
}

func BASE64EncodeStr(str string) string {
	encoded := BASE64Encode([]byte(str))
	return encoded
}

func BASE64Decode(str string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	return decoded, err
}

func BASE64DecodeStr(str string) (string, error) {
	decoded, err := BASE64Decode(str)
	if err == nil {
		decodedStr := string(decoded)
		return decodedStr, nil
	}
	return "", err
}

// RSA加密

func GenRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, err
}

func PublicKeyToString(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyPem)
	return publicKeyStr, err
}

func PrivateKeyToString(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	privateKeyStr := base64.StdEncoding.EncodeToString(privateKeyPem)
	return privateKeyStr, nil
}

func PublicKeyB64StrToStruct(publicKeyStr string) (*rsa.PublicKey, error) {
	publicKeyPem, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PEM block in base64 format containing the private key")
	}
	block, _ := pem.Decode(publicKeyPem)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parsed key is not an RSA public key")
	}
	return rsaPublicKey, nil
}

func PrivateKeyB64StrToStruct(privateKeyStr string) (*rsa.PrivateKey, error) {
	privateKeyPem, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode PEM block in base64 format containing the private key")
	}
	block, _ := pem.Decode(privateKeyPem)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// x509.ParsePKCS1PrivateKey 返回的是 *rsa.PrivateKey 类型，
	// 而不是一个通用的 crypto.PrivateKey，
	// 所以在类型断言时不需要使用 .(*rsa.PrivateKey)
	return privateKey, nil
}

func RSADecrypt(privateKey *rsa.PrivateKey, encryptedStr string) (string, error) {
	cipherTextB64, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		return "", err
	}
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherTextB64)
	if err != nil {
		return "", err
	}
	return string(decryptedBytes), nil
}

func RSAEncrypt(publicKey *rsa.PublicKey, str string) (string, error) {
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(str))
	if err != nil {
		return "", err
	}
	encryptedStr := base64.StdEncoding.EncodeToString(encryptedBytes)
	return encryptedStr, nil
}

// password加密

func GetPwdEncodingOpts() *password.Options {
	return &password.Options{
		SaltLen:      16,
		Iterations:   64,
		KeyLen:       16,
		HashFunction: md5.New,
	}
}

func GetEncodedPwd(pwd string) (string, string) {
	salt, encodedPwd := password.Encode(pwd, GetPwdEncodingOpts())
	return salt, encodedPwd
}

func VerifyEncodedPwd(pwdHeldRaw string, salt string, pwdTarget string) bool {
	return password.Verify(pwdHeldRaw, salt, pwdTarget, GetPwdEncodingOpts())
}
