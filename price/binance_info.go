package price

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func BinancePrice(symbol string) (float64, error) {
	var tradeInfo []BinanceResponse
	myReq := BinanceRequest{
		Symbol: symbol,
		Limit:  "1",
	}

	url := buildRequest[BinanceRequest](myReq)
	req, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer req.Body.Close()

	data, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(data, &tradeInfo)
	price, _ := strconv.ParseFloat(tradeInfo[0].Price, 64)
	return price, nil
}

func buildRequest[T any](ct T) string {
	addr := BINANCE + TRADE
	ask := structIter[T](ct)

	base, _ := url.Parse(addr)
	params := url.Values{}
	for _, q := range ask {
		params.Add(q[0], q[1])
	}

	base.RawQuery = params.Encode()
	return base.String()
}

func structIter[T any](ct T) [][]string {
	v := reflect.ValueOf(ct)
	typeOfS := v.Type()

	params := [][]string{}
	for i := 0; i < v.NumField(); i++ {
		name := typeOfS.Field(i).Name
		d := v.Field(i).Interface().(string)
		params = append(params, []string{strings.ToLower(name), d})
	}
	return params
}
