package main

import (
	"app/db"
	"app/messages"
	"app/types"
	"encoding/json"
	"fmt"
	"net/http"
)

//Handler handle webhook from tg
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &types.WebhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}
	fmt.Println("message", body)
	messages.HandleMessage(body)
}

func main() {
	//db.CreateTables()
	db.GetMessagesToSend()
	http.ListenAndServe(":8000", http.HandlerFunc(Handler))
}
