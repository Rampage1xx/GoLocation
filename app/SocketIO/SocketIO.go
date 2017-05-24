package SocketIO

import (
	socketIO "github.com/googollee/go-socket.io"
	"net/http"
	"strings"

	"fmt"
	"github.com/revel/revel"
)

var (
	sioServer   *socketIO.Server
	revelHandle http.Handler
)

func handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// route socketio requests to the socketio handler
	// and send everything else to the revel handler
	if strings.HasPrefix(path, "/socket.io/1/") {
		sioServer.ServeHTTP(w, r)
	} else {
		revelHandle.ServeHTTP(w, r)
	}
}

func PatchServer() {
	// create socketio server config


	// create socketio server
	sioServer, err := socketIO.NewServer(nil)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	// register global and namespace handlers
	registerHandlers(sioServer)

	// store a reference to revel's old http.Handler
	revelHandle = revel.Server.Handler

	// replace revel.Server.Handler with our new handler
	revel.Server.Handler = http.HandlerFunc(handle)
}