package main

import (
  "time"
  "crypto/x509"
  "io/ioutil"
  "encoding/pem"
  "fmt"
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509/pkix"
  "math/big"
)

func main() {
  template := &x509.Certificate {
    IsCA: false,
    BasicConstraintsValid : true,
    SubjectKeyId : []byte{1,2,3},
    SerialNumber : big.NewInt(1234),
    Subject : pkix.Name{
      Country : []string{"GB"},
      Organization: []string{"Sky Plc"},
    },
    NotBefore : time.Now(),
    NotAfter : time.Now().AddDate(5,5,5),
    // see http://golang.org/pkg/crypto/x509/#KeyUsage
    ExtKeyUsage : []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
    KeyUsage : x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign,
  }

  newCertPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
    panic(err)
  }

  intermediateCaPrivKey := loadAndDecryptPEMPrivateKey("/Users/dev/filips_playground/certificates/root/ca/intermediate/private/intermediate.key.pem", "R0MDaQzQ33G9WKrIPjGKBdnynHf/yDsvL2gZasAvsd0=")
  intermediateCaCert := loadPEMCert("/Users/dev/filips_playground/certificates/root/ca/intermediate/certs/intermediate.cert.pem")

  if intermediateCaPrivKey != nil {
    fmt.Println("We have got the key!")
  }

  cert, err := x509.CreateCertificate(rand.Reader, template, intermediateCaCert, &newCertPrivateKey.PublicKey, intermediateCaPrivKey)
  if err != nil {
        fmt.Println(err)
  }

  pkey := x509.MarshalPKCS1PrivateKey(newCertPrivateKey)
  ioutil.WriteFile("private.key", encryptToPEM(pkey, "RSA PRIVATE KEY", "password"), 0777)
  fmt.Println("private key saved to private.key")

  pubkey, _ := x509.MarshalPKIXPublicKey(&newCertPrivateKey.PublicKey)
  ioutil.WriteFile("public.key.pem", pubkey, 0777)
  fmt.Println("public key saved to public.key.pem")

  ioutil.WriteFile("cert.pem", encryptToPEM(cert, "CERTIFICATE", ""), 0777)
  fmt.Println("certificate saved to cert.pem")
}

func encryptToPEM(data []byte, blockType string, password string) []byte {
  if password != "" {
    encryptedBlock, err := x509.EncryptPEMBlock(rand.Reader, blockType, data, []byte(password), x509.PEMCipherAES256)
    if err != nil {
      panic(err)
    }

    pemBlock := pem.EncodeToMemory(encryptedBlock)
    if pemBlock == nil {
      panic("Failed to encode the certificate to PEM format.")
    }
    return pemBlock
  }

  pemBlock := pem.EncodeToMemory(&pem.Block{Type: blockType, Bytes: data})
  if pemBlock == nil {
    panic("Failed to encode the certificate to PEM format.")
  }
  fmt.Println("Using password: ", password)

  return pemBlock
}
