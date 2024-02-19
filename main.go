package simple_acme

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/go-acme/lego/v4/registration"
)

func GenerateCertificate(email string, CADirURL string, domains []string) {
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
		log.Fatal(err)
	}

	provider, err := route53.NewDNSProvider()
	if err != nil {
		log.Fatal(err)
	}
	err = client.Challenge.SetDNS01Provider(provider)
	if err != nil {
		log.Fatal(err)
	}

	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Fatal(err)
	}
	user.registration = reg

	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}
	certificate, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Fatal(err)
	}

	err = writeCertificate(certificate)
	if err != nil {
		log.Fatalln(err)
	}
}
