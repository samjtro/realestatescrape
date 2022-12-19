package data

import (
	"fmt"
	"sort"

	"github.com/samjtro/realestatescrape/scrape"
)

var (
	priceList  []int
	totalPrice int
)

func init() {
	rdcListings := scrape.Scrape()

	for _, listings := range rdcListings {
		for _, listing := range listings {
			priceList = append(priceList, listing.Price)
			fmt.Println(listing.Sqft)
		}
	}

	for _, price := range priceList {
		totalPrice += price
	}

	sort.Ints(priceList)
}

func Print() {
	fmt.Printf("Mean: %d\n", MeanPrice())
	min, max := MinMaxPrice()
	fmt.Printf("Min: %d\nMax: %d\n", min, max)
}

func MeanPrice(fu) int {
	return totalPrice / len(priceList)
}

func MedianPrice() int {
	return priceList[len(priceList)/2]
}

func MinMaxPrice() (int, int) {
	var (
		min int
		max int
	)

	for i, price := range priceList {
		if i+1 < len(priceList) {
			if price < priceList[i+1] {
				min = price
			} else if price > priceList[i+1] {
				max = price
			}
		}
	}

	return min, max
}
