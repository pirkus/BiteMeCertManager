package main

import (
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "fmt"
    "crypto/rsa"
)

func loadPEMCert(certPath string) *x509.Certificate {
  intermediateRoot, err := ioutil.ReadFile(certPath)
  if err != nil {
    panic(err)
  }

  block, _ := pem.Decode([]byte(intermediateRoot))
  if block == nil {
    panic("Failed to parse certificate PEM.")
  }
  cert, err := x509.ParseCertificate(block.Bytes)
  if err != nil {
    panic("failed to parse certificate: " + err.Error())
  }

  if cert != nil {
    fmt.Print("There is something")
  }

  return cert
}

func loadAndDecryptPEMPrivateKey(keyPath string, password string) *rsa.PrivateKey {
  // Read the key from the filesystem
  intermediateKey, err := ioutil.ReadFile(keyPath)

  // Decode PEM blocks into bytes
  block, _ := pem.Decode([]byte(intermediateKey))
  if block == nil {
    panic("Failed to parse certificate PEM.")
  }

  // Unlock the key with a password
  decryptedPEMBlock, err := x509.DecryptPEMBlock(block, []byte(password))
  if err != nil {
    panic("Cannot decrypt private key. Probably using wrong password.")
  }

  // Parse the key into into RSA structure
  key,err := x509.ParsePKCS1PrivateKey(decryptedPEMBlock)
  if err != nil {
    panic("failed to parse certificate: " + err.Error())
  }

  return key
}
