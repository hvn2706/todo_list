package server

import (
	"net/http"
	"todo_list/logger"
)

func ready(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("ok"))
	if err != nil {
		logger.Errorf("===== Write failed: %+v", err.Error())
		return
	}
	logger.Infof("===== Ready ok")
}
