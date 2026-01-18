package mq

import "testing"

func TestTimerTaskControlQueueName(t *testing.T) {
	if got := TimerTaskControlQueueName("", "node1"); got != "node1.timer" {
		t.Fatalf("unexpected queue name: %s", got)
	}
	if got := TimerTaskControlQueueName("q_", "node1"); got != "q_node1.timer" {
		t.Fatalf("unexpected queue name with prefix: %s", got)
	}
	if got := TimerTaskControlQueueName("", ""); got != "" {
		t.Fatalf("expected empty, got: %s", got)
	}
}
