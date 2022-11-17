package handler

import (
	"github.com/FlowerBirds/go-server-admin/config"
	"github.com/FlowerBirds/go-server-admin/util"
	"net/http"
)

func MakeListServerHandler(config *config.UdpConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		servers := config.GetServers()
		util.ResponseWithJson(servers, w)
	}
}

func MakeListClientHandler(config *config.UdpConfig) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		clients := config.GetClients()
		util.ResponseWithJson(clients, w)
	}
}
