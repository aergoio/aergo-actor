package router

func (m *mockProcess) MsgNum() int32 {
	m.Called()
	return 0
}
