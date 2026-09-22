package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	} else {
		return d.priorityLevel
	}
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch not := n.(type) {
	case directMessage:
		return not.senderUsername, not.importance()
	case groupMessage:
		return not.groupName, not.importance()
	case systemAlert:
		return not.alertCode, not.importance()
	default:
		return "", 0
	}
}

func main() {
	d := directMessage{
		senderUsername: "Kekkonen",
		messageContent: "Huutikset penalle",
		priorityLevel:  100,
		isUrgent:       true,
	}
	msg, importance := processNotification(d)
	fmt.Println(msg)
	fmt.Println(importance)
}
