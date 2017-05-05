package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/olahol/melody"
	"github.com/tarm/serial"
	"io/ioutil"
	"strings"
	"time"
)

var Melody = melody.New()

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	Melody.HandleMessage(func(s *melody.Session, msg []byte) {

		Melody.BroadcastFilter(msg, func(q *melody.Session) bool {
			log.Println(q.Request.URL.Path)
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	Melody.Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
}

func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for what is mostly likely the Arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "cu.usbmodem") {
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find a USB device that 'looks'
	// like an Arduino.
	return ""
}

func main() {
	// Start the listener

	c := &serial.Config{Name: findArduino(), Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	go loop(s)

	RunServer(LoadHTTP(), LoadHTTPS())
}

func loop(s *serial.Port) {
	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])

		Melody.BroadcastFilter(buf[:n], func(q *melody.Session) bool {
			return true
		})

		time.Sleep(50 * time.Millisecond)
	}

}
