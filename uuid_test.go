package uuid

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/hatchify/simply"
)

var (
	uuidSink   UUID
	stringSink string
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

func TestUUIDString(context *testing.T) {
	var (
		id                  UUID
		randomPrefix        string
		expectedSuffix      string
		expectedMSPredicate string
	)

	expectedSuffix = "-0600-00000000"
	expectedMSPredicate = "-0600-000000000000"

	id = newUUID(6)
	randomPrefix = strings.Split(id.String(), expectedSuffix)[0]

	test := simply.Target(id.String(), context, "Test standard UUID format")
	result := test.Equals(randomPrefix + expectedSuffix)
	test.Validate(result)

	test = simply.Target(id.ExtendedString(), context, "Test microsoft UUID format")
	result = test.Equals(randomPrefix + expectedMSPredicate)
	test.Validate(result)
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuidSink = New()
	}

	b.ReportAllocs()
}

func BenchmarkUUIDParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			uuidSink = New()
		}
	})

	b.ReportAllocs()
}

func BenchmarkUUIDString(b *testing.B) {
	u := New()
	for i := 0; i < b.N; i++ {
		stringSink = u.String()
	}

	b.ReportAllocs()
}

// testStruct is used to simulate a struct including a UUID
type testStruct struct {
	UUID UUID `json:"uuid"`
}
