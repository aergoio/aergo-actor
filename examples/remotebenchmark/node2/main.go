package main

import (
	"log"
	"runtime"

	"github.com/AsynkronIT/goconsole"
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo-actor/examples/remotebenchmark/messages"
	"github.com/aergoio/aergo-actor/mailbox"
	"github.com/aergoio/aergo-actor/remote"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 1)
	runtime.GC()

	remote.Start("127.0.0.1:8080")
	var sender *actor.PID
	props := actor.
		FromFunc(
			func(context actor.Context) {
				switch msg := context.Message().(type) {
				case *messages.StartRemote:
					log.Println("Starting")
					sender = msg.Sender
					context.Respond(&messages.Start{})
				case *messages.Ping:
					sender.Tell(&messages.Pong{})
				}
			}).
		WithMailbox(mailbox.Bounded(1000000))

	actor.SpawnNamed(props, "remote")

	console.ReadLine()
}
