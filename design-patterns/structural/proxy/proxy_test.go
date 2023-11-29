package proxy

import "testing"

func TestAgentTask_RentHouse(t *testing.T) {
	agent := NewAgentTask()
	agent.RentHouse("上海", 8000)
}
