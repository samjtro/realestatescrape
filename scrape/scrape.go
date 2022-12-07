package scrape

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

var (
	cityStateFlag = "Cheyenne_WY"
	priceMaxFlag  = 200000
	pagesFlag     = 5

	rdcResults []Listings
	c          = colly.NewCollector()

	url = "https://www.realtor.com/realestateandhomes-search/%s/type-single-family-home/price-na-%d/pg-%d" //CityState, Price, Page #
)

type Listing struct {
	Status  string
	Address string
	Price   int
	/*Bed   int
	Bath    int
	Sqft    int
	LotSize int*/
}

type Listings []Listing

func init() {
	c.OnError(func(_ *colly.Response, err error) {
		log.Fatal(err)
	})
}

func Scrape() {
	for i := 1; i <= pagesFlag; i++ {
		url = fmt.Sprintf(url, cityStateFlag, priceMaxFlag, i)
		listings := ScrapeRDCHelper(url)
		rdcResults = append(rdcResults, listings)
	}

	for _, listings := range rdcResults {
		for _, listing := range listings {
			fmt.Println(listing.Status, listing.Address, listing.Price)
		}
	}
}

/*
for states:     'Colorado' or 'California'
for cities:     'city1-city2_state'   i.e.    'Los-Angeles_CA' or 'Denver_CO'
*/
func ScrapeRDCHelper(url string) Listings {
	var listings Listings

	c.OnHTML("li.jsx-1881802087.component_property-card", func(e *colly.HTMLElement) {
		var listing Listing
		var err error

		listing.Price, err = strconv.Atoi(UnformatPrice(e.ChildText("span.Price__Component-rui__x3geed-0.gipzbd")))

		if err != nil {
			log.Fatal(err)
		}

		listing.Status = e.ChildText("span.jsx-3853574337.statusText")
		listing.Address = e.ChildText("div.jsx-11645185.address.ellipsis.srp-page-address.srp-address-redesign") + " " + e.ChildText("div.jsx-11645185.address-second.ellipsis")

		/*bbsl := e.ChildText("ul.jsx-946479843.property-meta.list-unstyled.property-meta-srpPage")
		unorderedList := RDCUnmarshallPropertyMeta(bbsl)

		if len(unorderedList) >= 1 {
			listing.Bed = unorderedList[0]
		}

		if len(unorderedList) >= 2 {
			listing.Bath = unorderedList[1]
		}

		if len(unorderedList) >= 3 {
			listing.Sqft = unorderedList[2]
		}

		if len(unorderedList) >= 4 {
			listing.LotSize = unorderedList[3]
		}

		sqft, err := strconv.Atoi((e.ChildText("span.jsx-946479843.meta-value")))

		if err != nil {
			log.Fatal(err)
		}*/

		listings = append(listings, listing)
		//fmt.Println(listing)
	})

	c.Visit(url)

	return listings
}

func UnformatPrice(price string) string {
	var newInt []rune

	for i, c := range price {
		if c == '$' {
			if i > 0 {
				break
			}
		}

		if c == '$' {
			continue
		} else if c == ',' {
			continue
		} else if c == '+' {
			continue
		} else {
			newInt = append(newInt, c)
		}
	}

	return string(newInt)
}
