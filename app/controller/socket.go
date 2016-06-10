package controller

import (
	//"fmt"
	"log"
	//"reflect"
	"net/http"

	"github.com/evanfeenstra/circuitSocket/app/wrappers/melody"

)

var ClientHost string

func Configure(url string) {
	ClientHost = url
}

func init() {
	melody.Melody.Upgrader.CheckOrigin = func(r *http.Request) bool { 
		if r.Header.Get("Origin") == ClientHost {
			return true 
		} else {
			log.Println("Unauthorized Websocket Orgin: ",r.Header.Get("Origin"))
			return false
		}
	}
}

func ReadWS(w http.ResponseWriter, r *http.Request) {

	//client gets on WS page
	/*var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	p := params.ByName("id")

	log.Println(melody.HubList[p])*/

	melody.Melody.HandleRequest(w, r)

	//melody.Melody.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}
