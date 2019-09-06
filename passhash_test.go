package passhash

import "testing"

func printError(e string, r string, t *testing.T) {
	t.Errorf("\nExpected: %s\nReceived: %s", e, r)
}

func TestSayHi(t *testing.T) {
	expected := "Hello Brett!"

	response := sayHi("Brett")

	if response != expected {
		printError(expected, response, t)
	}
}
