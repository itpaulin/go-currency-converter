package main

import (
	"fmt"
	"sync"

	"github.com/itpaulin/go-currency-converter/api"
	"github.com/itpaulin/go-currency-converter/cmd"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var rates map[string]interface{}
	go func() {
		defer wg.Done()
		rates = api.GetRates()
	}()


	var userInput cmd.CoinValue
    go func() {
		defer wg.Done()
		userInput = cmd.GetData()
		
	}()

	wg.Wait()

	cmd.ConvertCurrency(userInput, rates)

	fmt.Printf("The value of %.2f %s in %s is %.2f\n", userInput.Value, userInput.CurrencyFrom, userInput.CurrencyTo, cmd.ConvertCurrency(userInput, rates))
}