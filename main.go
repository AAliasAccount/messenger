package main 


import (
	//	"io"
	"net/http"
	"fmt"
	"encoding/json"

	// Router and middleware
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	// Used for rendering
	"github.com/unrolled/render"
)


type RecivedMessage struct {
	From 	string `json:"From"`
	Message string `json:"Message"`

}

type IDMessage struct  { 
	ID 		int8
	Message RecivedMessage
}


var messages = []IDMessage{
	//{ID: 1, From: "me", Message: "Please work"},
	//{ID: 2, From: "you", Message: "it worked"},
}

var id int8 = 1

func assignID(msg RecivedMessage){
	var from = msg.From
	var message = msg.Message

	var newID int8 = id
	id ++
	
	var newMessage = IDMessage{newID, RecivedMessage{from, message}}
	messages = append(messages, newMessage)
	fmt.Println(messages)
}


func main() {


	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	re := render.New()

	// Set up routes 

	// Get full Convo
	router.Get("/messages", func(w http.ResponseWriter, r *http.Request ){ 
		re.JSON(w, http.StatusOK, messages)
	})

	// Add new message
	router.Post("/newMessage", func(w http.ResponseWriter, r *http.Request ){
		var msgStruct RecivedMessage
		if err := json.NewDecoder(r.Body).Decode(&msgStruct); err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			fmt.Println("Error decoding JSON:", err)
			return
		}
		assignID(msgStruct)
	})


	http.ListenAndServe(":8000", router)
}
