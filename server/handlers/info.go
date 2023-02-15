package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/render"
)

type InfoResponse struct {
	Pod  Pod  `json:"pod"`
	Node Node `json:"node"`
}

type Pod struct {
	Name               string `json:"name"`
	IP                 string `json:"ip"`
	Namespace          string `json:"namespace"`
	ServiceAccountName string `json:"serviceAccountName"`
}

type Node struct {
	Name string `json:"name"`
}

var notAvailable string = "n/a"

func Info(w http.ResponseWriter, r *http.Request) {
	var ok bool

	// initialize pod object and set values if available
	pod := Pod{}
	pod.Name, ok = os.LookupEnv("POD_NAME")
	if !ok {
		pod.Name = notAvailable
	}
	pod.IP, ok = os.LookupEnv("POD_IP")
	if !ok {
		pod.IP = notAvailable
	}
	pod.Namespace, ok = os.LookupEnv("POD_NAMESPACE")
	if !ok {
		pod.Namespace = notAvailable
	}
	pod.ServiceAccountName, ok = os.LookupEnv("POD_SERVICE_ACCOUNT_NAME")
	if !ok {
		pod.ServiceAccountName = notAvailable
	}

	// initialize node object and set values if available
	node := Node{}
	node.Name, ok = os.LookupEnv("NODE_NAME")
	if !ok {
		node.Name = notAvailable
	}

	res := &InfoResponse{
		Pod:  pod,
		Node: node,
	}

	render.JSON(w, r, res)
}
