package server

import (
	"Blood-Transmutation/server/models"
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	BloodGroup1 string `json:"bloodGroup1"`
	BloodGroup2 string `json:"bloodGroup2"`
}

func api(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
  log.Println("Blood Group 1 : "+req.BloodGroup1)
  log.Println("Blood Group 2 : "+req.BloodGroup2)
	res := models.Level_4(req.BloodGroup1, req.BloodGroup2)
  log.Println("Response received from model")
	w.Write([]byte(res))
  log.Println("Response sent to client")
}
