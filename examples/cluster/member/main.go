package main

import (
	"fmt"
	"log"
	"time"

	console "github.com/AsynkronIT/goconsole"
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo-actor/cluster"
	"github.com/aergoio/aergo-actor/cluster/consul"
	"github.com/aergoio/aergo-actor/examples/cluster/shared"
	"github.com/aergoio/aergo-actor/remote"
)

const (
	timeout = 1 * time.Second
)

func main() {
	//this node knows about Hello kind
	remote.Register("Hello", actor.FromProducer(func() actor.Actor {
		return &shared.HelloActor{}
	}))

	cp, err := consul.New()
	if err != nil {
		log.Fatal(err)
	}
	cluster.Start("mycluster", "127.0.0.1:8081", cp)

	sync()
	async()

	console.ReadLine()

	cluster.Shutdown(true)
}

func sync() {
	hello := shared.GetHelloGrain("abc")
	options := cluster.NewGrainCallOptions().WithTimeout(5 * time.Second).WithRetry(5)

	res, err := hello.SayHelloWithOpts(&shared.HelloRequest{Name: "GAM"}, options)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Message from SayHello: %v", res.Message)
	for i := 0; i < 10000; i++ {
		x := shared.GetHelloGrain(fmt.Sprintf("hello%v", i))
		x.SayHello(&shared.HelloRequest{Name: "GAM"})
	}
	log.Println("Done")
}

func async() {
	hello := shared.GetHelloGrain("abc")
	c, e := hello.AddChan(&shared.AddRequest{A: 123, B: 456})

	for {
		select {
		case <-time.After(100 * time.Millisecond):
			log.Println("Tick..") //this might not happen if res returns fast enough
		case err := <-e:
			log.Fatal(err)
		case res := <-c:
			log.Printf("Result is %v", res.Result)
			return
		}
	}
}
