package main

import (
	"log"

	console "github.com/AsynkronIT/goconsole"
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo-actor/cluster"
	"github.com/aergoio/aergo-actor/cluster/consul"
	"github.com/aergoio/aergo-actor/examples/cluster/shared"
	"github.com/aergoio/aergo-actor/remote"
)

func Logger(next actor.ActorFunc) actor.ActorFunc {
	fn := func(context actor.Context) {
		switch context.Message().(type) {
		case *actor.Started:
			log.Printf("actor started " + context.Self().String())
		case *actor.Stopped:
			log.Printf("actor stopped " + context.Self().String())
		}
		next(context)
	}

	return fn
}

func main() {

	//this node knows about Hello kind
	remote.Register("Hello", actor.FromProducer(func() actor.Actor {
		return &shared.HelloActor{}
	}).WithMiddleware(Logger))

	cp, err := consul.New()
	if err != nil {
		log.Fatal(err)
	}
	cluster.Start("mycluster", "127.0.0.1:8080", cp)

	hello := shared.GetHelloGrain("MyGrain")

	res, err := hello.SayHello(&shared.HelloRequest{Name: "Roger"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Message from grain: %v", res.Message)
	console.ReadLine()
}
