package kafka

import (
	"github.com/IBM/sarama"
)

type MockSyncProducer struct {
	Messages []*sarama.ProducerMessage
	Err      error
}

func (m *MockSyncProducer) TxnStatus() sarama.ProducerTxnStatusFlag {
	//TODO implement me
	panic("implement me")
}

func (m *MockSyncProducer) IsTransactional() bool {
	//TODO implement me
	panic("implement me")
}

func (m *MockSyncProducer) BeginTxn() error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSyncProducer) CommitTxn() error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSyncProducer) AbortTxn() error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSyncProducer) AddOffsetsToTxn(offsets map[string][]*sarama.PartitionOffsetMetadata, groupId string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockSyncProducer) AddMessageToTxn(msg *sarama.ConsumerMessage, groupId string, metadata *string) error {
	//TODO implement me
	panic("implement me")
}

func NewMockSyncProducer() *MockSyncProducer {
	return &MockSyncProducer{
		Messages: make([]*sarama.ProducerMessage, 0),
	}
}

func (m *MockSyncProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	if m.Err != nil {
		return 0, 0, m.Err
	}
	m.Messages = append(m.Messages, msg)
	return 0, 0, nil
}

func (m *MockSyncProducer) SendMessages(msgs []*sarama.ProducerMessage) error {
	if m.Err != nil {
		return m.Err
	}
	m.Messages = append(m.Messages, msgs...)
	return nil
}

func (m *MockSyncProducer) Close() error {
	return nil
}
