package handler

import (
	"net/http"

	"github.com/arissupardiman/digitalent-microservice/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
