package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"

	"github.com/Miloha/challenge_Nubank/managecard"
)

func main() {

	fmt.Println("hello 9000")

	mit := managecard.NewCart()

	// register `mit` object with `rpc.DefaultServer`
	rpc.Register(mit)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// sample test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}
