# Import

    go get github.com/yes-ph/simple-acme

# Usage

	import "github.com/go-acme/lego/v4/providers/dns/digitalocean"

	func main() {
		environment := os.Getenv("ENVIRONMENT")

		CADirURL := lego.LEDirectoryStaging

		if environment == "production" {
			CADirURL = lego.LEDirectoryProduction
		}

		provider, err := digitalocean.NewDNSProvider()
		if err != nil {
			log.Fatal(err)
		}

		err = simple_acme.GenerateCertificate(
			provider,
			"email@example.com",
			CADirURL,
			[]string{
				"example.com",
				"*.example.com",
			},
			"tls.crt",
			"tls.key",
			"",
		)

		if err != nil {
			log.Fatal(err)
		}
	}