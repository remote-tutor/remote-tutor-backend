package aws

import (
	awsDiagnostics "backend/diagnostics/aws"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"io/ioutil"
	"os"
	"time"
)

func GenerateSignedURL(fullURL, sourceIP string) (string, error) {
	privateKey, err := getPrivateKey(os.Getenv("SIGNED_URL_PRIVATE_KEY"))
	if err != nil {
		awsDiagnostics.WriteAWSSignedURLErr(err, "Generate Signed URL (reading private key)")
		return "", err
	}
	signer := sign.NewURLSigner(os.Getenv("CLOUDFRONT_USER_KEY"), privateKey)

	if os.Getenv("APP_ENV") == "development" {
		signedURL, err := signer.Sign(fullURL, time.Now().Add(5 * time.Second))
		if err != nil {
			awsDiagnostics.WriteAWSSignedURLErr(err, "Generate Signed URL (generating signed URL)")
			return "", err
		}
		return signedURL, nil
	}

	// Sign URL to be valid for from now, expires 30 minutes from now, and
	// restricted to the given IP address (IF THE MODE IS NOT DEVELOPMENT).
	policy := &sign.Policy{
		Statements: []sign.Statement{
			{
				Resource:  fullURL,
				Condition: sign.Condition{
					// Optional IP source address range
					IPAddress: &sign.IPAddress{SourceIP: sourceIP},
					// Required date the URL will expire after
					DateLessThan: &sign.AWSEpochTime{Time: time.Now().Add(45 * time.Second)},
				},
			},
		},
	}

	signedURL, err := signer.SignWithPolicy(fullURL, policy)
	if err != nil {
		awsDiagnostics.WriteAWSSignedURLErr(err, "Generate Signed URL (generating signed URL)")
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
