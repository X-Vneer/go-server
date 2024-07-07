package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"github.com/x-vneer/go-server/api"
	"github.com/x-vneer/go-server/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		logrus.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	database, err := tools.NeedDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	tokenDetails := (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		logrus.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logrus.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
