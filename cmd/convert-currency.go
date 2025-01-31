package cmd

func ConvertCurrency(input CoinValue, rates map[string]interface{}) float64 {

	if input.CurrencyFrom == "USD" {
		return input.Value * rates[input.CurrencyTo].(float64)
	}
	if input.CurrencyTo == "USD" {
		return input.Value / rates[input.CurrencyFrom].(float64)
	}

	valueInUSD := input.Value / rates[input.CurrencyFrom].(float64)

	return valueInUSD * rates[input.CurrencyTo].(float64)
	
}