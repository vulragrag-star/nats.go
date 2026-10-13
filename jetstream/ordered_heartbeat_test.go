package jetstream

import (
	"testing"
	"time"
)

// Ordered Consume/Messages parse opts with ordered=true (5s heartbeat) but the
// inner pull consumer re-parses with ordered=false (Expires/2 = 15s). The fix
// propagates PullHeartbeat(consumeOpts.Heartbeat); these tests lock the defaults.
func TestOrderedConsumeDefaultHeartbeat(t *testing.T) {
	t.Parallel()
	ordered, err := parseConsumeOpts(true)
	if err != nil {
		t.Fatalf("parseConsumeOpts(true): %v", err)
	}
	if ordered.Heartbeat != 5*time.Second {
		t.Fatalf("ordered Consume default Heartbeat = %v, want 5s", ordered.Heartbeat)
	}
	plain, err := parseConsumeOpts(false)
	if err != nil {
		t.Fatalf("parseConsumeOpts(false): %v", err)
	}
	if plain.Heartbeat != DefaultExpires/2 {
		t.Fatalf("plain Consume default Heartbeat = %v, want %v", plain.Heartbeat, DefaultExpires/2)
	}
}

func TestOrderedMessagesDefaultHeartbeat(t *testing.T) {
	t.Parallel()
	ordered, err := parseMessagesOpts(true)
	if err != nil {
		t.Fatalf("parseMessagesOpts(true): %v", err)
	}
	if ordered.Heartbeat != 5*time.Second {
		t.Fatalf("ordered Messages default Heartbeat = %v, want 5s", ordered.Heartbeat)
	}
	plain, err := parseMessagesOpts(false)
	if err != nil {
		t.Fatalf("parseMessagesOpts(false): %v", err)
	}
	if plain.Heartbeat != DefaultExpires/2 {
		t.Fatalf("plain Messages default Heartbeat = %v, want %v", plain.Heartbeat, DefaultExpires/2)
	}
}
