package identifier

// MockIdentifier mockNanoGenerator Mock of IIdentifier
type MockIdentifier struct{}

func NewMockIdentifier() MockIdentifier {
	return MockIdentifier{}
}

func (m MockIdentifier) NewNanoID() string {
	return "newnanoid"
}

func (m MockIdentifier) NewUUIDv5(name string) string {
	return "newuuidv5"
}
