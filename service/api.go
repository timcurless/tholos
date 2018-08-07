package service

import (
	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func initializeVault() (*vault.Client, error) {
	// Default assumes vault is at http://127.0.0.1:8200
	config := vault.DefaultConfig()

	v, err := vault.NewClient(config)

	if err != nil {
		logrus.Errorln("Error initializing vault: " + err.Error())
		return nil, err
	}

	status, _ := v.Sys().Health()
	logrus.Printf("Connected to Vault Cluster: %v, ID: %v", status.ClusterName, status.ClusterID)

	return v, nil
}

func createSecret(v *vault.Client, path string, data map[string]interface{}) (*vault.Secret, error) {
	logical := v.Logical()

	secret, err := logical.Write(path, data)
	if err != nil {
		logrus.Errorf("Error creating new secret: %v", err.Error())
		return nil, err
	}

	return secret, nil
}
