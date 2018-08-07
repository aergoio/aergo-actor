package main

import (
	"flag"
	"runtime"

	"log"

	console "github.com/AsynkronIT/goconsole"
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo-actor/examples/remoterouting/messages"
	"github.com/aergoio/aergo-actor/mailbox"
	"github.com/aergoio/aergo-actor/remote"
)

var (
	flagBind = flag.String("bind", "localhost:8100", "Bind to address")
	flagName = flag.String("name", "node1", "Name")
)

type remoteActor struct {
	name  string
	count int
}

func (a *remoteActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.Ping:
		context.Respond(&messages.Pong{})
	}
}

func newRemoteActor(name string) actor.Producer {
	return func() actor.Actor {
		return &remoteActor{
			name: name,
		}
	}
}

func newRemote(bind, name string) {
	remote.Start(bind)
	props := actor.
		FromProducer(newRemoteActor(name)).
		WithMailbox(mailbox.Bounded(10000))

	actor.SpawnNamed(props, "remote")

	log.Println(name, "Ready")
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()

	flag.Parse()

	newRemote(*flagBind, *flagName)

	console.ReadLine()
}
