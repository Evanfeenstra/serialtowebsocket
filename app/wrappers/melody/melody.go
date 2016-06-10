package melody

import (
	"strings"
	/*"log"
	"reflect"
	"encoding/json"*/

	"github.com/olahol/melody"

	"sync"
)


/* 
CHANGED MAXMESSAGESIZE IN CONFIG.GO IN OLAHOL MELODY
4096 max size gives about 500 max line length
*/

//  websocket hub
var Melody = melody.New()

// list of active pages (gives a unique ID to each client, in controller.index)
var HubList = map[string]int{} // trailing empty {} initializes zero value

// count of number of people on a page (to zero out HubList if page becomes empty)
var HubCount = map[string]int{}

// mutex for above 3 vars
var mutex = &sync.Mutex{}

func init() {

	Melody.HandleMessage(func(s *melody.Session, msg []byte) {

		Melody.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	Melody.HandleConnect(func(s *melody.Session) {

		p := s.Request.URL.Path
		p = strings.Split(p,"/")[2]

		mutex.Lock()
		HubList[p] = HubList[p] + 1
		HubCount[p] = HubCount[p] + 1
		mutex.Unlock()

	})

	Melody.HandleDisconnect(func(s *melody.Session) {

		p := s.Request.URL.Path
		p = strings.Split(p,"/")[2]

		mutex.Lock()
		HubCount[p] = HubCount[p] - 1 
		mutex.Unlock()

		//page became empty (no visitors)
		if HubCount[p] == 0 {
			mutex.Lock()
			delete(HubList, p)
			delete(HubCount, p)
			mutex.Unlock()
		}
	})
}
