package rest

import (
	"GriffinBackend/price"
	"GriffinBackend/rdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func getBinanceTrade(c *gin.Context) {
	currNeed := []string{price.ETHEREUM, price.POLYGON, price.USDC}
	result := map[string]float64{}
	for _, m := range currNeed {
		p, err := price.BinancePrice(m)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message": PRICE_GET_ERROR,
			})
			return
		}
		result[m] = p
	}

	// send price information
	p := price.PriceInformation{
		Ethereum: result[price.ETHEREUM],
		Polygon:  result[price.POLYGON],
		USDC:     result[price.USDC],
	}
	c.JSON(http.StatusOK, p)
}

func postPayment(c *gin.Context, db *redis.Client) {
	eid, err := handleParamEmployerId(c)
	cnt, err := handleParamPostPay(c)
	if err != nil {
		return
	}
	err = rdb.JsonArrAppend(
		db,
		fmt.Sprintf(HISTORICAL_PAYMENT_KEY, eid),
		HISTORICAL_PAYMENT_PATH,
		&cnt,
	)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": DATABASE_APPEND_FAIL,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": DATABASE_APPEND_SUCCESS,
		})
	}
}

func getPayment(c *gin.Context, db *redis.Client) {
	eid, err := handleParamEmployerId(c)
	if err != nil {
		return
	}

	var pr [][]rdb.Payment
	err = rdb.JsonGet(
		db,
		fmt.Sprintf(HISTORICAL_PAYMENT_KEY, eid),
		HISTORICAL_PAYMENT_PATH,
		&pr,
	)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": DATABASE_GET_FAIL,
		})
	} else {
		record := pr[0][1:]
		c.JSON(http.StatusOK, record)
	}
}

func getPaymentMonthly(c *gin.Context, db *redis.Client) {
	eid, err := handleParamEmployerId(c)
	if err != nil {
		return
	}

	var pr [][]rdb.Payment
	err = rdb.JsonGet(
		db,
		fmt.Sprintf(HISTORICAL_PAYMENT_KEY, eid),
		HISTORICAL_PAYMENT_PATH,
		&pr,
	)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": DATABASE_GET_FAIL,
		})
	} else {
		record := pr[0][1:]
		p := monthlyPayment(record)
		pS := monthlyPaymentStruct(p)
		c.JSON(http.StatusOK, pS)
	}
}
