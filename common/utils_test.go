package common

import (
	"fmt"
	"testing"
)

type MockSolutionParser struct {
	ExpectedContent string
}

func (m *MockSolutionParser) CreateSolutionInput(content string) (SolutionInput, error) {
	if m.ExpectedContent != content {
		return nil, fmt.Errorf("unexpected content: got %v, want %v", content, m.ExpectedContent)
	}
	return &MockSolutionInput{}, nil
}

type MockSolutionInput struct{}

func (m MockSolutionInput) Validate() error {
	return nil
}

func TestReadInput(t *testing.T) {
	mockParser := &MockSolutionParser{ExpectedContent: "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"}

	_, err := ReadInput("testdata", Day1, Input1, mockParser)
	if err != nil {
		t.Fatalf("ReadInput failed: %v", err)
	}
}
