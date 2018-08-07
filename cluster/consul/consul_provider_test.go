package consul

import (
	"log"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/aergoio/aergo-actor/cluster"
	"github.com/aergoio/aergo-actor/eventstream"
)

func TestRegisterMember(t *testing.T) {
	if testing.Short() {
		return
	}

	p, _ := New()
	defer p.Shutdown()
	err := p.RegisterMember("mycluster", "127.0.0.1", 8000, []string{"a", "b"}, nil, &cluster.NilMemberStatusValueSerializer{})
	/*if err != nil {
		log.Fatal(err)
	}*/
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(*net.OpError); ok && err.Op == "dial" {
			log.Println("maybe there is a no running consul server. skip consul test")
			return
		}
		log.Fatal(err)
	}
}

func TestRefreshMemberTTL(t *testing.T) {
	if testing.Short() {
		return
	}

	p, _ := New()
	defer p.Shutdown()
	err := p.RegisterMember("mycluster", "127.0.0.1", 8000, []string{"a", "b"}, nil, &cluster.NilMemberStatusValueSerializer{})
	/*if err != nil {
		log.Fatal(err)
	}*/
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(*net.OpError); ok && err.Op == "dial" {
			log.Println("maybe there is a no running consul server. skip consul test")
			return
		}
		log.Fatal(err)
	}
	p.MonitorMemberStatusChanges()
	eventstream.Subscribe(func(m interface{}) {
		log.Printf("Event %+v", m)
	})
	time.Sleep(60 * time.Second)
}
