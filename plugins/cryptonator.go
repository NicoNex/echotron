/*
* Echotron-GO
* Copyright (C) 2018  Nicolò Santamaria
*/

package plugins

import (
	"gitlab.com/NicoNex/echotron"
	"fmt"
	"encoding/json"
	"strings"
	"time"
	)

type market struct {
	Market string `json:"market"`
	Price string `json:"price"`
	Volume float64 `json:"volume"`
}

type ticker struct {
	Base string `json:"base"`
	Target string `json:"target"`
	Price string `json:"price"`
	Volume string `json:"volume"`
	Change string `json:"change"`
	Markets []market `json:"markets,omitempty"`
}

type CryptonatorResponse struct {
	Data ticker `json:"ticker,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
	Success bool `json:"success"`
	Error string `json:"error"`
}

func GetCryptocurrencyData (cryptocurrency string, target string) CryptonatorResponse {
	var baseURL string = "https://api.cryptonator.com/api/ticker/"
	var response CryptonatorResponse

	url := fmt.Sprintf("%s%s-%s", baseURL, cryptocurrency, target)
	content := echotron.SendGetRequest(url)

	json.Unmarshal(content, &response)

	return response
}

func GetCryptocurrencyFullData (cryptocurrency string, target string) CryptonatorResponse {
	var baseURL string = "https://api.cryptonator.com/api/full/"
	var response CryptonatorResponse

	url := fmt.Sprintf("%s%s-%s", baseURL, cryptocurrency, target)
	content := echotron.SendGetRequest(url)

	json.Unmarshal(content, &response)
	return response
}

func CryptonatorGetData(cryptocurrency string, target string) string {
	var message strings.Builder
	var dataMessage string = "*Currency code*: %s\n*Target*: %s\n\n*Price*: %s *%s*\n\n*Total trade volume for the last 24 hours*: %s *%s*\n\n*Last hour price change*: %s *%s*\n\n*Last updated*: %s"

	data := GetCryptocurrencyData(cryptocurrency, target)

	if data.Success {
		message.WriteString(fmt.Sprintf(dataMessage, data.Data.Base, data.Data.Target, data.Data.Price, data.Data.Target, data.Data.Volume, data.Data.Target, data.Data.Change, data.Data.Target, time.Unix(data.Timestamp, 0)))
	} else {
		message.WriteString(data.Error)
	}
	return message.String()
}

func CryptonatorGetMarket(cryptocurrency string, target string) string {
	var message strings.Builder
	var marketMessage string = "*Market*\nThese are some useful data. You can check the prices of you favourite cryptocurrency from different sellers\n"

	data := GetCryptocurrencyFullData(cryptocurrency, target)

	if data.Success && len(data.Data.Markets) > 0 {
		message.WriteString(fmt.Sprintf(marketMessage))
		for i := 0; i < len(data.Data.Markets); i++ {
			message.WriteString(fmt.Sprintf("\n*Name of the exchange*: \"%s\"\n*Price*: %s *%s*\n*Volume*: %f *%s*\n", data.Data.Markets[i].Market, data.Data.Markets[i].Price, data.Data.Base, data.Data.Markets[i].Volume, data.Data.Target))
		}
	} else if !data.Success {
		message.WriteString(data.Error)
	} else {
		message.WriteString("No market data available")
	}
	return message.String()
}

