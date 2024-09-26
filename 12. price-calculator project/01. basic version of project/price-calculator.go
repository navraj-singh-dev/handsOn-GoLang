package main

import "fmt"

func main() {

	// create prices slice
	pricesWithoutTax := []float64{10, 20, 30}

	// tax percentages that will be applied on these prices
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// create a map to store these prices with applied taxes
	appliedTaxes := make(map[float64][]float64)

	// loop over both arrays and calculate the tax applied price
	for _, taxRates := range taxRates {
		// for every iteration create a new slice which will become
		// value for a single key in map
		calculatedTaxes := make([]float64, len(pricesWithoutTax))

		for priceIndex, priceValue := range pricesWithoutTax {
			// multiply the tax with the price using the formula and get the tax applied price
			calculatedTaxes[priceIndex] = priceValue * (1 + taxRates)
		}

		// as calculatedTaxes slice is filled we can set it as value to a map's key
		appliedTaxes[taxRates] = calculatedTaxes
	}

	// print the prices which are inclusive of tax rates
	fmt.Println(appliedTaxes)
}
