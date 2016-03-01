package packethandler

import (
	"time"
)

const sessionTimeout time.Duration = time.Second * 150

// Prune deletes sessions which haven't been accessed for more than 150 seconds. Prune is run every 10 seconds.
func Prune() {
	for k, v := range Sessions {
		if time.Since(v.LastRequest) > sessionTimeout {
			delete(Sessions, k)
		}
	}
	time.Sleep(time.Second * 10)
}
