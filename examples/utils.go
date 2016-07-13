package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func getProducts() (products []Product) {
	params := map[string]string{"access_token": client.MLToken.AccessToken}

	body, err := client.Get(fmt.Sprintf("users/%d/items/search", client.MLToken.UserId), params)
	if err != nil {
		log.Println(err)
	}

	var ps ProductSearch
	json.Unmarshal(body, &ps)

	params["ids"] = strings.Join(ps.Results, ",")
	body, err = client.Get("items", params)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(body, &products)

	return
}
