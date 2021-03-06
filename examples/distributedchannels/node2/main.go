package main

import (
	"log"
	"runtime"

	"github.com/AsynkronIT/goconsole"
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo-actor/examples/distributedchannels/messages"
	"github.com/aergoio/aergo-actor/remote"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	remote.Start("127.0.0.1:8080")
	//create the channel
	channel := make(chan *messages.MyMessage)

	//create an actor receiving messages and pushing them onto the channel
	props := actor.FromFunc(func(context actor.Context) {
		if msg, ok := context.Message().(*messages.MyMessage); ok {
			channel <- msg
		}
	})
	actor.SpawnNamed(props, "MyMessage")

	//consume the channel just like you use to
	go func() {
		for msg := range channel {
			log.Println(msg)
		}
	}()

	console.ReadLine()
}
