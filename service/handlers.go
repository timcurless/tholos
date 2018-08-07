package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

var isHealthy = true

var v *vault.Client

func init() {
	var err error
	v, err = initializeVault()
	if err != nil {
		logrus.Errorf("Error initializing Vault: %v", err.Error())
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if isHealthy {
		h, _ := v.Sys().Health()
		data, _ := json.Marshal(healthCheckResponse{
			Status:        "UP",
			VaultInstance: h.ClusterName,
		})
		writeJSONResponse(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "DOWN"})
		writeJSONResponse(w, http.StatusServiceUnavailable, data)
	}
}

func GetAWSSTSToken(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TTL string `json:"ttl"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Invalid request. Error: %v", err.Error())
		data, _ := json.Marshal(errorResponse{ErrorStr: err.Error()})
		writeJSONResponse(w, http.StatusBadRequest, data)
	}

	data := make(map[string]interface{}, 1)
	data["ttl"] = body.TTL

	token, err := createSecret(v, "aws/sts/deploy", data)

	if err != nil {
		logrus.Errorf("Could not create a new STS Token. Error: %v", err.Error())
		data, _ := json.Marshal(errorResponse{ErrorStr: err.Error()})
		writeJSONResponse(w, http.StatusInternalServerError, data)
	}

	res, _ := json.Marshal(token.Data)

	writeJSONResponse(w, http.StatusCreated, res)
}

func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status        string `json:"status"`
	VaultInstance string `json:"vault_name"`
}

type errorResponse struct {
	ErrorStr string `json:"error"`
}

type AWSCred struct {
	AWSKeyID  string `json:"aws_access_key_id"`
	AWSSecret string `json:"aws_secret_access_key"`
}
