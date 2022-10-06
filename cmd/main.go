package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/LiliyaD/Test_task_GasPrice/internal/journal"
	"github.com/LiliyaD/Test_task_GasPrice/internal/pkg/calculation"
)

var srcJSONLink = flag.String("src", "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json", "source JSON file")

func main() {
	flag.Parse()

	journal.New(false)
	journal.LogInfo("RUN Client")

	resp, err := http.Get(*srcJSONLink)
	if err != nil {
		journal.LogFatal(err)
	}

	fmt.Println(srcJSONLink)
	fmt.Println(*srcJSONLink)

	calculation.Parse(resp.Body)

	responce := calculation.Process()

	calculation.SaveJSON(&responce)
}
