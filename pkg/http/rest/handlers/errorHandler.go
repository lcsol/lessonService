package handlers

import (
	"net/http"
)

func (h *handler) serverError(rw http.ResponseWriter, err error) {
	h.errLog.Println(err)

	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
