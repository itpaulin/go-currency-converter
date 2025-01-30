package terminal

import (
	"errors"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

type CoinValue struct {
	CurrencyFrom string
	Value float64
	CurrencyTo string
}
func GetData() CoinValue {

	var coin string
	var value string
	var finalCoin string

	var optionsCoins = []string{"BRL", "USD", "EUR", "JPY"}



	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a coin type ").
				Options(huh.NewOptions("BRL", "USD", "EUR", "JPY")...).
				Value(&coin),

			huh.NewInput().
				Title("Enter the value").
				Validate(func (str string) error {
					if str == "" {
						return errors.New("please enter a value")
					}
					if _, err := strconv.ParseFloat(str, 64); err != nil {
						return errors.New("please enter a valid number")
					}
					return nil
				}).
				Value(&value),
		), 

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose the final coin type").
				OptionsFunc(
					func() []huh.Option[string] {  
						var filteredCoins []string
						for _, option := range optionsCoins {
							if option != coin {
								filteredCoins = append(filteredCoins, option)
							}
						}
						return huh.NewOptions(filteredCoins...)
					}, &coin,
				).
				Value(&finalCoin),
		),
	)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	floatValue, _ := strconv.ParseFloat(value, 64)
	
	return CoinValue{CurrencyFrom: coin, Value: floatValue, CurrencyTo: finalCoin}
}