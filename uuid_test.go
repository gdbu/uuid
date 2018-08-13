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

	if bs, err = json.Marshal(ts); err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}

	if err = json.Unmarshal(bs, &nts); err != nil {
		t.Fatalf("Error unmarshaling: %v", err)
	}

	if ts != nts {
		t.Fatalf("Invalid value, expected %v and received %v", ts, nts)
	}
}

type testStruct struct {
	UUID UUID `json:"uuid"`
}
