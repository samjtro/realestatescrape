package data

import (
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
		}
	}

	for _, price := range priceList {
		totalPrice += price
	}
}

func MeanPrice() int {
	return totalPrice / len(priceList)
}

/*func MedianPrice() {

}

func PriceOutliers() {

}*/
