package handlers

import (
	"fmt"
	"github.com/dnahurnyi/uploader/clientAPI/app/contracts"
	"net/http"
	"strings"

	"github.com/rs/zerolog"
)

type PortsHandler struct {
	Log          zerolog.Logger
	Parser       contracts.Parser
	DomainClient contracts.DomainClient
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

	receiverFunc, prepErr := h.DomainClient.Send(ctx)
	if prepErr != nil {
		msg := "operation failed, can't connect to PortDomainService"
		h.Log.Error().Err(prepErr).Msg(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}

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
				if err := receiverFunc(nil, true); err != nil {
					msg := "parse failed and receiver close error. %d records successfully processed"
					h.Log.Error().Err(err).Msgf(msg, counter)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(fmt.Sprintf(msg+", err: %v", counter, err)))
					return
				}
				msg := "parse failed, %d records successfully processed"
				h.Log.Error().Err(err).Msgf(msg, counter)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(msg+", err: %v", counter, err)))
				return
			}
			if data == nil {
				// All data parsed, close the client
				if err := receiverFunc(nil, true); err != nil {
					msg := "receiver close error. %d records successfully processed"
					h.Log.Error().Err(err).Msgf(msg, counter)
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fmt.Sprintf(msg+", err: %v", counter, err)))
					return
				}
				msg := "success, %d records successfully processed"
				h.Log.Info().Msgf(msg, counter)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(fmt.Sprintf(msg, counter)))
				return
			}
			if err := receiverFunc(data, false); err != nil {
				if err := receiverFunc(nil, true); err != nil {
					msg := "receiver failed and receiver close error. %d records successfully processed"
					h.Log.Error().Err(err).Msgf(msg, counter)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(fmt.Sprintf(msg+", err: %v", counter, err)))
					return
				}
				msg := "send failed, %d records successfully processed"
				h.Log.Error().Err(err).Msgf(msg, counter)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(msg+", err: %v", counter, err)))
				return
			}
			counter++
		}
	}
}

func (h *PortsHandler) getPort(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/ports/")
	storable, err := h.DomainClient.Get(req.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("can't get port record. Err: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(storable.Present())
}
