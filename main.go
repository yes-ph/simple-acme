package simple_acme

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

func GenerateCertificate(provider challenge.Provider, email string, CADirURL string, domains []string, certificateFilename string, privateKeyFilename string, formatString string) error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	user := user{
		email: email,
		key:   privateKey,
	}

	config := lego.NewConfig(&user)

	config.CADirURL = CADirURL
	config.Certificate.KeyType = certcrypto.RSA2048

	client, err := lego.NewClient(config)
	if err != nil {
		return err
	}

	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		return err
	}

	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return err
	}
	user.registration = reg

	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}
	certificate, err := client.Certificate.Obtain(request)
	if err != nil {
		return err
	}

	err = writeCertificate(certificate, certificateFilename, privateKeyFilename, formatString)
	if err != nil {
		return err
	}

	return nil
}
