package kafka

type mockDLQ struct {
	called bool
	Topic  string
}

func (m *mockDLQ) SendToDLQ(val []byte) {
	m.called = true
}
