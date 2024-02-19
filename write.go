package simple_acme

import (
	"fmt"
	"os"
	"time"

	"github.com/go-acme/lego/v4/certificate"
)

func writeCertificate(certificate *certificate.Resource) error {
	now := time.Now()
	folderName := now.Format("2006-01-02_15-04-05")

	err := os.MkdirAll(folderName, 0755)
	if err != nil {
		return err
	}

	certFilePath := fmt.Sprintf("%s/%s", folderName, "tls.crt")

	err = os.WriteFile(certFilePath, certificate.Certificate, 0644)
	if err != nil {
		return err
	}

	keyFilePath := fmt.Sprintf("%s/%s", folderName, "tls.key")

	err = os.WriteFile(keyFilePath, certificate.PrivateKey, 0644)
	if err != nil {
		return err
	}

	if _, err := os.Lstat("tls.crt"); err == nil {
		err = os.Remove("tls.crt")
		if err != nil {
			return err
		}
	}

	err = os.Symlink(certFilePath, "tls.crt")
	if err != nil {
		return err
	}

	if _, err := os.Lstat("tls.key"); err == nil {
		err = os.Remove("tls.key")
		if err != nil {
			return err
		}
	}

	err = os.Symlink(keyFilePath, "tls.key")
	if err != nil {
		return err
	}

	return nil
}
