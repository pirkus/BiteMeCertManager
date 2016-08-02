package main

import (
	"BiteMeCertManager/testhelpers"
	"os"
	"testing"
)

const invalidCert = "freaking bollocks everywhere"

const cert = `-----BEGIN CERTIFICATE-----
MIIF4zCCA8ugAwIBAgICEAEwDQYJKoZIhvcNAQELBQAwgYkxCzAJBgNVBAYTAkdC
MRAwDgYDVQQIDAdFbmdsYW5kMRAwDgYDVQQKDAdTa3kgUGxjMQswCQYDVQQLDAJJ
UzEjMCEGA1UEAwwaU2t5IE1vYmlsZSBJbnRlcm1lZGlhdGUgQ0ExJDAiBgkqhkiG
9w0BCQEWFWRsLXNuc2NlcnRzQGJza3liLmNvbTAeFw0xNjA1MTcxNDEzMjRaFw0x
NzA1MjcxNDEzMjRaMIGBMQswCQYDVQQGEwJHQjEQMA4GA1UECAwHRW5nbGFuZDEP
MA0GA1UEBwwGTG9uZG9uMRAwDgYDVQQKDAdTa3kgUGxjMQswCQYDVQQLDAJJUzEP
MA0GA1UEAwwGc2VydmVyMR8wHQYJKoZIhvcNAQkBFhBwaXJrdXNAZ21haWwuY29t
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyiikOI8Bv5QVlSUviC9C
7/bShDFNdkV9/H5bSc4cUp5xH/ewfj0Orc3DnCVyzFxO0csvAt2jgcQqkGoJt9v/
v8qgdJDAutc6EfelrnkOpJkIpEbc6a+lAjH4ekKopQ4158Pfy8EPQ4Ixz4DzKyYG
kKioRPT+Nj5M+JNTUlp6T4e4Qqj666tBbLhAPgPuS/cjF/YoOOU6hA8tdybdzyTt
3o+Wg4lbRcqHozoC0CXKm14PDJ1YF3LagNTA3Gde1fF5E2GhTKRqvZ5SQLz6nXCE
mAgimFtZvBmxbCA1BC3iVJm1arEKMRb+LIl023EBAV3aoinUs2i7UG+dwxRXZCP+
PwIDAQABo4IBWTCCAVUwCQYDVR0TBAIwADARBglghkgBhvhCAQEEBAMCBkAwMwYJ
YIZIAYb4QgENBCYWJE9wZW5TU0wgR2VuZXJhdGVkIFNlcnZlciBDZXJ0aWZpY2F0
ZTAdBgNVHQ4EFgQUnz++fslbMppJvxmxv82bN0ssNa0wgbsGA1UdIwSBszCBsIAU
DXbTd8yqOv9muDFXpKT2ev/y4HChgZOkgZAwgY0xCzAJBgNVBAYTAkdCMRAwDgYD
VQQIDAdFbmdsYW5kMQ8wDQYDVQQHDAZMb25kb24xEDAOBgNVBAoMB1NreSBQbGMx
CzAJBgNVBAsMAklTMRYwFAYDVQQDDA1Ta3kgTW9iaWxlIENBMSQwIgYJKoZIhvcN
AQkBFhVkbC1zbnNjZXJ0c0Bic2t5Yi5jb22CAhAAMA4GA1UdDwEB/wQEAwIFoDAT
BgNVHSUEDDAKBggrBgEFBQcDATANBgkqhkiG9w0BAQsFAAOCAgEAfYJ0/72fGijj
HazXwghpqNxxM437ke2UdhZdAvXFr+dhXcmgw6cpmts1DJUDbY6pFq/B8Lrpf/PQ
dESqoRBL02LMBKZZmzcEpdAtGi/oyaPM2LNg9OlMbMmTFBaljni0KoYPVvegfs7F
IVJuTVFgfqNF84bTfSlIRjGoAJ9XlRs/ftbf427w44KQySIs/yxpeM8YWfqKC94e
1N8zuRm/gYyuIfZDQQ5bMVAFaNVcAdFzp6l5AM84/OyrbG9+LyZ0V8AY1bryJqfQ
j0FZjkTkP8zENMkJx6uRsPPrVDp3R2DvqcJV766kaLe5v9VeW76oLEScouGZSq0o
fDtB4CWdZRbAFHHmOmSs2Dpvdcd58Na8N8SEftBn8BvmoHsMT/PNK13L6hq2o7oo
aKO/jru/dCv6N01vu4DarRuG1F180kvS4JeYF/3G4oEL4iv1OrRQoYiru+DfQjMq
J8fl1nw0Ua6C/oqGp46g7StlU/ywDzKvj+PMn7v9CoX14wgdmDuey5igrpEcMtqf
oKTc42EDfZppjzMsKgazUt+BmqIkIn7GkvoC1hy8xMJYKkuEqv4IJnbZo9v+kCM/
zXBiAz0fN0sbTO9JcCEqgFSL/SgMAq09WkS1DW1dqYFlY08B422pjPt2QTMtb2Wk
Ke6vGPfqT9NV6XgCls6Kra/toU4qwOI=
-----END CERTIFICATE-----`

func TestMain(m *testing.M) {
	testhelpers.WriteFile("./test.cert.pem", cert)
	testhelpers.WriteFile("./invalid.cert.pem", invalidCert)

	exitCode := m.Run()

	testhelpers.RemoveFile("./test.cert.pem")
	testhelpers.RemoveFile("./invalid.cert.pem")

	os.Exit(exitCode)
}

func TestCertCanBeLoaded(t *testing.T) {
	loadedCert := loadPEMCert("./test.cert.pem")

	if loadedCert == nil {
		t.Errorf("Certificate could not be loaded.")
	}
}

func TestPanicsIfInvalidCertIsPassedIn(t *testing.T) {
	defer testhelpers.AssertPanic(t, "Kok")
	loadPEMCert("./invalid.cert.pem")
}
