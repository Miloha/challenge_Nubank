package main

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/Miloha/challenge_Nubank/managecard"
)

func main() {

	fmt.Println("Hello Nubank :9000")

	mit := managecard.NewCard()

	// register `mit` object with `rpc.DefaultServer`
	rpc.Register(mit)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	rpc.HandleHTTP()

	//	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	//		io.WriteString(res, "RPC SERVER LIVE!")
	//	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}
