package actor

import (
	"fmt"

	"github.com/aergoio/aergo-actor/eventstream"
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
			plog.Debug().Interface("reason", supervisorEvent.Reason).Interface("receiver", supervisorEvent.Child).
				Interface("directive", fmt.Stringer(supervisorEvent.Directive).String()).Msg("Supervisor handles a failing child")
		}
	})
}
