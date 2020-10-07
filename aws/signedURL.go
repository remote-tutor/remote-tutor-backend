package aws

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"io/ioutil"
	"os"
	"time"
)

func GenerateSignedURL(fullURL string) (string, error) {
	privateKey, err := getPrivateKey("aws/pk-APKAJD7IMVQH4KL6QTAQ.pem")
	if err != nil {
		return "", err
	}
	signer := sign.NewURLSigner(os.Getenv("CLOUDFRONT_USER_KEY"), privateKey)
	signedURL, err := signer.Sign(fullURL, time.Now().Add(10 * time.Second))
	if err != nil {
		return "", err
	}
	return signedURL, nil
}

func getPrivateKey(path string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(b)
	der, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return der, nil
}
