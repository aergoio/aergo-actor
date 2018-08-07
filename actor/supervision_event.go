package actor

import (
	"github.com/aergoio/aergo-actor/eventstream"
	"github.com/aergoio/aergo-actor/log"
)

//SupervisorEvent is sent on the EventStream when a supervisor have applied a directive to a failing child actor
type SupervisorEvent struct {
	Child     *PID
	Reason    interface{}
	Directive Directive
}

var (
	supervisionSubscriber *eventstream.Subscription
)

func init() {
	supervisionSubscriber = eventstream.Subscribe(func(evt interface{}) {
		if supervisorEvent, ok := evt.(*SupervisorEvent); ok {
			plog.Debug("[SUPERVISION]", log.Stringer("actor", supervisorEvent.Child), log.Stringer("directive", supervisorEvent.Directive), log.Object("reason", supervisorEvent.Reason))
		}
	})
}
