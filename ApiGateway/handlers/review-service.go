package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/my-flights/ApiGateway/model"
	"github.com/my-flights/ApiGateway/utils"
)

const ReviewsServiceApi string = "/api/reviews"

func GetAllReviewsByAirline(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	airlineId, _ := strconv.ParseUint(params["id"], 10, 32)
	response, err := http.Get(utils.BaseReviewService.Next().Host + ReviewsServiceApi + "/" + strconv.FormatUint(uint64(airlineId), 10))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var request model.CreateReviewRequest
	data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewReader(data)).Decode(&request)

	req, _ := http.NewRequest(http.MethodPost, utils.BaseReviewService.Next().Host+ReviewsServiceApi+"/", bytes.NewReader(data))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
