package examples

//Category struct
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ProductSearch struct
type ProductSearch struct {
	Results []string `json:"results"`
}

//Product struct
type Product struct {
	ID                string  `json:"id,omitempty"`
	ListingTypeID     string  `json:"listing_type_id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	CategoryID        string  `json:"category_id"`
	BuyingMode        string  `json:"buying_mode"`
	CurrencyID        string  `json:"currency_id"`
	Condition         string  `json:"condition"`
	Price             float32 `json:"price"`
	AvailableQuantity int32   `json:"available_quantity"`
	Pictures          []Image `json:"pictures"`
	Status            string  `json:"status"`
}

//Image struct
type Image struct {
	Source string `json:"source"`
}

//Status struct
type Status struct {
	Status string `json:"status"`
}

//Questions response
type Questions struct {
	Questions []Question
}

//Question struct
type Question struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
