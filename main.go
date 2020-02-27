package main

import (
	"fmt"
	"lion/cmd"
	"net/http"
)

const (
	path = "/webhooks"
)

var (
	WebhookEvents = map[string]string{
		"OrderAll":         "store/order/*",
		"ProductAll":       "store/product/*",
		"CategoryAll":      "store/category/*",
		"SKUAll":           "store/sku/*",
		"Customer":         "store/customer/*",
		"CartAll":          "store/cart/*",
		"CartLineItemAll":  "store/cart/lineItem/*",
		"ShipmentAll":      "store/shipment/*",
		"SubscriberAll":    "store/subscriber/*",
		"StoreUninstalled": "store/app/uninstalled",
		"StoreUpdated":     "store/information/updated",
	}
)


func Respond(w http.ResponseWriter, req *http.Request) {
	fmt.Println("here!")
	w.WriteHeader(405)
}

func main() {
	//http.HandleFunc(path, Respond)
	//fmt.Println("listening...")
	//err := http.ListenAndServe(":3000", nil)
	//if err != nil {
	//	panic(err)
	//}
	//for _, scope := range WebhookEvents {
	//	fmt.Println(scope)
	//}

	cmd.Execute()
}
