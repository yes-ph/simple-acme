package simple_acme

import (
	"fmt"
	"os"
	"time"

	"github.com/go-acme/lego/v4/certificate"
)

func writeCertificate(certificate *certificate.Resource, certificateFilename string, privateKeyFilename string, formatString string) error {
	now := time.Now()

	if formatString == "" {
		formatString = "2006-01-02_15-04-05"
	}

	folderName := now.Format(formatString)

	err := os.MkdirAll(folderName, 0755)
	if err != nil {
		return err
	}

	certFilePath := fmt.Sprintf("%s/%s", folderName, certificateFilename)

	if _, err := os.Lstat(certificateFilename); err == nil {
		err = os.Remove(certificateFilename)
		if err != nil {
			return err
		}
	}

	err = os.Symlink(certFilePath, certificateFilename)
	if err != nil {
		return err
	}

	err = os.WriteFile(certFilePath, certificate.Certificate, 0644)
	if err != nil {
		return err
	}

	keyFilePath := fmt.Sprintf("%s/%s", folderName, privateKeyFilename)

	err = os.WriteFile(keyFilePath, certificate.PrivateKey, 0644)
	if err != nil {
		return err
	}

	if _, err := os.Lstat(privateKeyFilename); err == nil {
		err = os.Remove(privateKeyFilename)
		if err != nil {
			return err
		}
	}

	err = os.Symlink(keyFilePath, privateKeyFilename)
	if err != nil {
		return err
	}

	return nil
}
