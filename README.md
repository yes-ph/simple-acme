# Import

    go get github.com/yes-ph/simple-acme

# Usage

    environment := os.Getenv("ENVIRONMENT")

	CADirURL := lego.LEDirectoryStaging

	if environment == "production" {
		CADirURL = lego.LEDirectoryProduction
	}

	simple_acme.GenerateCertificate(
		"email@example.com",
		CADirURL,
		[]string{
			"example.com",
			"*.example.com",
		},
	)

	if err != nil {
		log.Fatal(err)
	}