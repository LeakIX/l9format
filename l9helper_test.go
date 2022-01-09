package l9format

import (
	"strings"
	"testing"
)

var TestEventsByFingerprint = []struct{
	fingerprint string
	Event L9Event}{
	{Event: L9Event{EventSource: "test", Summary: strings.Repeat("test2\ndate: test\nEtag: 223sd\n", 300)}, fingerprint: "bc2c0be9cd8fd9c83b2fbaae92f0685832b49f9ea5fa0ee8760c14ceeed86528"},
	{Event: L9Event{EventSource: "test", Summary: strings.Repeat("test2\nlast-modified: 2012\n", 300)}, fingerprint: "bc2c0be9cd8fd9c83b2fbaae92f0685832b49f9ea5fa0ee8760c14ceeed86528"},
	{Event: L9Event{EventSource: "test", Summary: strings.Repeat("test2\n", 300)}, fingerprint: "bc2c0be9cd8fd9c83b2fbaae92f0685832b49f9ea5fa0ee8760c14ceeed86528"},
	{Event: L9Event{EventSource: "test", Summary: strings.Repeat("test3", 300)}, fingerprint: "bc2c0be9cd8fd9c8f6333b94f6333b94f6333b94f6333b94f6333b94f6333b94"},
	{Event: L9Event{EventSource: "test", Summary: strings.Repeat("test3", 321)}, fingerprint: "bc2c0be9cd8fd9c85d35bfef5d35bfef5d35bfef5d35bfef5d35bfef5d35bfef"},

}

func TestL9Event_UpdateFingerprint(t *testing.T) {
	for _, testEvent := range TestEventsByFingerprint {
		event :=  testEvent.Event
		err := event.UpdateFingerprint()
		if err != nil {
			t.Error(err)
		}
		if len(event.EventFingerprint) != fingerPrintLength*2 {
			t.Errorf("fingerprint length IS not %d: %s", fingerPrintLength*2, event.EventFingerprint)
		}
		if testEvent.fingerprint != event.EventFingerprint {
			t.Errorf("fingerprint %s doesn't match base %s", event.EventFingerprint, testEvent.fingerprint)
		}
	}

}
