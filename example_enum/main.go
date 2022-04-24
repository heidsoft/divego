package main

import "fmt"

type ActionParameter string
const (
	ActionParameterTargetName        = ActionParameter("targetName")
	ActionParameterAlarmName         = ActionParameter("alarmName")
	ActionParameterOldStatus         = ActionParameter("oldStatus")
	ActionParameterNewStatus         = ActionParameter("newStatus")
	ActionParameterTriggeringSummary = ActionParameter("triggeringSummary")
	ActionParameterDeclaringSummary  = ActionParameter("declaringSummary")
	ActionParameterEventDescription  = ActionParameter("eventDescription")
	ActionParameterTarget            = ActionParameter("target")
	ActionParameterAlarm             = ActionParameter("alarm")
)

func main() {
	fmt.Println(string(ActionParameterAlarm))
}
