package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/didil/k8s-pod-info-api/server"
	"github.com/didil/k8s-pod-info-api/server/handlers"
	"github.com/stretchr/testify/assert"
)

func TestInfoHandler(t *testing.T) {
	r := server.NewRouter()
	s := httptest.NewServer(r)
	defer s.Close()

	req, err := http.NewRequest(http.MethodGet, s.URL+"/api/v1/info", nil)
	assert.NoError(t, err)

	os.Setenv("POD_NAME", "my-pod")
	os.Setenv("POD_IP", "192.168.144.12")
	os.Setenv("POD_NAMESPACE", "my-namespace")
	os.Setenv("POD_SERVICE_ACCOUNT_NAME", "my-sa")
	os.Setenv("NODE_NAME", "my-node-name")

	cl := &http.Client{}
	resp, err := cl.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	info := &handlers.InfoResponse{}

	err = json.NewDecoder(resp.Body).Decode(info)
	assert.NoError(t, err)

	expectedInfo := &handlers.InfoResponse{
		Pod: handlers.Pod{
			Name:               "my-pod",
			IP:                 "192.168.144.12",
			Namespace:          "my-namespace",
			ServiceAccountName: "my-sa",
		},
		Node: handlers.Node{
			Name: "my-node-name",
		},
	}

	assert.Equal(t, expectedInfo, info)
}
