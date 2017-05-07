package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/rs/cors"
	"os"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("so.On chat message golang: " + msg)
			m := make(map[string]interface{})
			m["a"] = "привет"
			e := so.Emit("cn1111", m)
			//Это не проблема
			fmt.Println("\n\n")

			b := make(map[string]string)
			b["u-a"] = "китайское Содержание" //Это не может быть китайским
			m["b-c"] = b
			e = so.Emit("cn2222", m)
			log.Println(e)

			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		// Socket.io acknowledgement example
		// The return type may vary depending on whether you will return
		// For this example it is "string" type
		so.On("chat message with ack", func(msg string) string {
			return "from golang: " + msg
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	mux := http.NewServeMux()

	var frontendUrl string
	if os.Getenv("HOME") == "/root" {
		frontendUrl = "http://zhanat.site:8081"
	} else {
		frontendUrl = "http://localhost:3000"
	}
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{frontendUrl},
		AllowCredentials: true,
	}).Handler(mux)

	mux.Handle("/socket.io/", server)
	mux.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", handler))
}

