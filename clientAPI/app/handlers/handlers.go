package handlers

import (
	"fmt"
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
	ctx := req.Context()

	parseFunc, prepErr := h.Parser.Parse(req.Body)
	if prepErr != nil {
		msg := "operation failed, can't start object parsing"
		h.Log.Error().Err(prepErr).Msg(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}
	counter := 0

	for {
		select {
		case <-ctx.Done():
			h.Log.Warn().Msgf("operation canceled. %d records processed", counter)
			return
		default:
			data, err := parseFunc()
			if err != nil {
				msg := "parse failed, %d records successfully processed"
				h.Log.Error().Err(err).Msgf(msg, counter)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(msg+", err: %v", counter, err)))
				return
			}
			if data == nil {
				// All data parsed
				msg := "success, %d records successfully processed"
				h.Log.Info().Msgf(msg, counter)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf(msg, counter)))
				return
			}
			// TODO: Replace by send to portDomain service
			fmt.Println("Port parsed: ", string(data.Present()))
			counter++
		}
	}
}

func (h *PortsHandler) getPort(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/ports/")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Not implemented, got id: " + id))
}
