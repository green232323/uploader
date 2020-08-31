package handlers

import (
	"github.com/dnahurnyi/uploader/clientAPI/app/contracts"
	"net/http"
	"strings"

	"github.com/rs/zerolog"
)

type PortsHandler struct {
	Log    zerolog.Logger
	Parser contracts.Parser
}

func (h *PortsHandler) Handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.getPort(w, req)
	case http.MethodPost:
		h.loadPorts(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *PortsHandler) loadPorts(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Not implemented"))
}

func (h *PortsHandler) getPort(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/ports/")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Not implemented, got id: " + id))
}
