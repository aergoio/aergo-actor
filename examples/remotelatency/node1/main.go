package main

import (
	"time"

	console "github.com/AsynkronIT/goconsole"
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo-actor/examples/remotelatency/messages"
	"github.com/aergoio/aergo-actor/remote"

	"runtime"
)

// import "runtime/pprof"

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	messageCount := 1000000

	remote.Start("127.0.0.1:8081", remote.WithEndpointWriterBatchSize(10000))

	remote := actor.NewPID("127.0.0.1:8080", "remote")
	remote.
		RequestFuture(&messages.Start{}, 5*time.Second).
		Wait()

	for i := 0; i < messageCount; i++ {
		message := &messages.Ping{
			Time: makeTimestamp(),
		}
		remote.Tell(message)
		if i%1000 == 0 {
			time.Sleep(500)
		}
	}
	console.ReadLine()
}
