package actor

func (ref *ActorProcess) MsgNum() int32 {
	return 0
}

func (ref *deadLetterProcess) MsgNum() int32 {
	return 0
}

func (e *EventStreamProcess) MsgNum() int32 {
	return 0
}

func (ref *futureProcess) MsgNum() int32 {
	return 0
}

func (g *guardianProcess) MsgNum() int32 {
	return 0
}

// MsgNum returns a number of messages those are queued in mailbox
func (pid *PID) MsgNum() int32 {
	return 0
}
