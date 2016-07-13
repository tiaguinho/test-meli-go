package examples

import (
	"encoding/json"
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/tiaguinho/mercadolibre-go-sdk"
	"log"
	"net/http"
	"strconv"
)

var (
	client    *meli.Client
	productID string
)

func init() {
	client = &meli.Client{ClientID: client_id, ClientSecret: client_secret}
}

//Home func
func Home(r render.Render, req *http.Request) {
	if client.MLToken.AccessToken == "" && req.URL.Query().Get("code") == "" {
		redirectURL, _ := client.GetAuthUrl("http://localhost:8080/", meli.AuthUrls["MLB"])

		if redirectURL != "" {
			r.Redirect(redirectURL, http.StatusFound)
		}
	} else {
		client.Authorize(req.URL.Query().Get("code"), "http://localhost:8080/")
	}

	r.HTML(http.StatusOK, "home", nil)
}

//Get example
func Get(r render.Render) {
	params := map[string]string{"access_token": client.MLToken.AccessToken}
	body, err := client.Get("sites/MLB/categories", params)

	var categories []Category
	if err != nil {
		log.Panicln(err)
	}

	json.Unmarshal(body, &categories)

	r.HTML(http.StatusOK, "methods/get", categories)
}

//Post example
func Post(r render.Render) {
	params := map[string]string{"access_token": client.MLToken.AccessToken}

	product := Product{
		ListingTypeID:     "free",
		Title:             "Golang SDK Title",
		Description:       "Golang SDK description product test",
		CategoryID:        "MLB50655",
		BuyingMode:        "buy_it_now",
		CurrencyID:        "BRL",
		Condition:         "new",
		Price:             100.00,
		AvailableQuantity: 1,
		Pictures: []Image{
			{
				Source: "http://i.stack.imgur.com/DJBD5.png",
			},
		},
	}

	resp, err := client.Post("items", product, params)
	var products []Product
	if err != nil {
		log.Panicln(err)
	}

	json.Unmarshal(resp, &product)
	productID = product.ID

	products = getProducts()

	r.HTML(http.StatusOK, "methods/post", products)
}

//Put example
func Put(r render.Render) {
	params := map[string]string{"access_token": client.MLToken.AccessToken}

	status := Status{"closed"}
	client.Put(fmt.Sprintf("items/%s", productID), status, params)

	products := getProducts()

	r.HTML(http.StatusOK, "methods/put", products)
}

//Delete example
func Delete(r render.Render, req *http.Request) {
	params := map[string]string{"access_token": client.MLToken.AccessToken}
	if req.URL.Query().Get("id") != "" {
		client.Delete(fmt.Sprintf("questions/%s", req.URL.Query().Get("id")), params)

		r.Redirect("/methods/delete", http.StatusFound)
	}

	params["seller_id"] = strconv.Itoa(client.MLToken.UserId)
	params["status"] = "unanswered"

	body, _ := client.Get("questions/search", params)

	var questions Questions
	json.Unmarshal(body, &questions)

	r.HTML(http.StatusOK, "methods/delete", questions)
}
