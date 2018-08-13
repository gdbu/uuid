package uuid

import (
	"encoding/json"
	"testing"
)

func TestUUID(t *testing.T) {
	var (
		ts  testStruct
		nts testStruct
		bs  []byte
		err error
	)

	ts.UUID = newUUID(6)

	// Attempt to marshal the test struct
	if bs, err = json.Marshal(ts); err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}

	// Attempt to unmarshal the marshaled bytes
	if err = json.Unmarshal(bs, &nts); err != nil {
		t.Fatalf("Error unmarshaling: %v", err)
	}

	// Ensure test struct matches new test struct
	if ts != nts {
		t.Fatalf("Invalid value, expected %v and received %v", ts, nts)
	}
}

// testStruct is used to simulate a struct including a UUID
type testStruct struct {
	UUID UUID `json:"uuid"`
}
