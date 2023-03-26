package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func websockethandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, []byte("hi i am golang websocket!"))
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}

}

func main() {
	http.HandleFunc("/socket", websockethandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:9092", nil))
}
