package remote

import "sync/atomic"

func (m *endpointWriterMailbox) Len() int32 {
	return atomic.LoadInt32(&m.hasMoreMessages)
}

func (ref *process) MsgNum() int32 {
	return ref.MsgNum()
}
