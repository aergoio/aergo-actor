package main

import (
	"runtime"

	"github.com/AsynkronIT/goconsole"
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo-actor/remote"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//empty actor just to have something to remote spawn
	props := actor.FromFunc(func(ctx actor.Context) {})
	remote.Register("remote", props)

	remote.Start("127.0.0.1:8080")

	console.ReadLine()
}
