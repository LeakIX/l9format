package l9format

import (
	"strings"
	"testing"
)

var TestEventsByFingerprint = map[string]L9Event{
	"bc2c0be9cd8fd9c886876c74fe77280819bd01b4a72ac148e2c94bf475b0fd34": {EventSource: "test",Summary: strings.Repeat("test\n",300)},
	"bc2c0be9cd8fd9c886876c74fe77280819bd01b4a72ac148e2c94bf431f15032": {EventSource: "test",Summary: strings.Repeat("test\n",299)},
	"5a2c67df7b758a1eb5b5f87eed5cadde5abff03ee094719ed1bbd7fe1282b070": {EventSource: "testing",Summary: strings.Repeat("test\n",299)},
	"5a2c67df7b758a1eb5b5f87eed5caddeed5caddeed5caddeed5caddefe9efc5e": {EventSource: "testing",Summary: strings.Repeat("test\n",2)},
	"5a2c67df7b758a1eb5b5f87eed5cadde5abff03e5abff03e5abff03e988de430": {EventSource: "testing",Summary: strings.Repeat("test\n",3)},
}

func TestL9Event_UpdateFingerprint(t *testing.T) {
	for fingerprint, event := range TestEventsByFingerprint {
		err := event.UpdateFingerprint()
		if err != nil {
			t.Error(err)
		}
		if len(event.EventFingerprint) != fingerPrintLength*2 {
			t.Errorf("fingerprint length IS not %d: %s", fingerPrintLength*2, event.EventFingerprint)
		}
		if fingerprint != event.EventFingerprint {
			t.Errorf("fingerprint %s doesn't match base %s", event.EventFingerprint, fingerprint)
		}
	}

}

