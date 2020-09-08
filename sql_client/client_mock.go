package sql_client

import "fmt"

type clientMock struct {
	mocks map[string]Mock
}

type Mock struct {
	Query   string
	Args    []interface{}
	Error   error
	Columns []string
	Row     []interface{}
}

func AddMock(mock Mock) {

	if dbClient == nil {
		return
	}

	client, ok := dbClient.(*clientMock)
	if !ok {
		return
	}
	if client.mocks == nil {
		client.mocks = make(map[string]Mock, 0)
	}
	client.mocks[mock.Query] = mock
}

func (c *clientMock) QueryRow(query string, args ...interface{}) row {
	mock, ok := c.mocks[query]
	if !ok {
		fmt.Println("Error finding mock")
		return nil
	}

	if mock.Error != nil {
		fmt.Println("Mock error")
		return nil
	}

	result := rowMock{
		Columns: mock.Columns,
		Row:     mock.Row,
	}

	return &result
}
