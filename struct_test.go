package telegram_test

import (
	"testing"

	"github.com/copartner6412/telegram"
)

type TestStruct struct {
	Name      string  `json:"name"`
	Age       int     `json:"age"`
	IsStudent bool    `json:"is_student"`
	Address   *string `json:"address"`
	Nested    Nested  `json:"nested"`
}

type Nested struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func TestStructToMapStringByJSONTag(t *testing.T) {
	address := "123 Main St"
	input := TestStruct{
		Name:      "John Doe",
		Age:       25,
		IsStudent: true,
		Address:   &address,
		Nested: Nested{
			City:  "San Francisco",
			State: "CA",
		},
	}

	expected := map[string]string{
		"name":       "John Doe",
		"age":        "25",
		"is_student": "true",
		"address":    "123 Main St",
		"nested":     `{"city":"San Francisco","state":"CA"}`,
	}

	output := make(map[string]string)
	err := telegram.StructToMapStringByJSONTag(output, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(output) != len(expected) {
		t.Fatalf("output map length mismatch: expected %d, got %d", len(expected), len(output))
	}

	for key, expectedValue := range expected {
		actualValue, ok := output[key]
		if !ok {
			t.Errorf("missing key in output: %s", key)
		} else if actualValue != expectedValue {
			t.Errorf("value mismatch for key %s: expected %q, got %q", key, expectedValue, actualValue)
		}
	}
}
