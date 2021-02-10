package handlers

import (
	"net/http"
)

func (lh *LessonHandler) serverError(rw http.ResponseWriter, err error) {
	lh.errLog.Println(err)

	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
