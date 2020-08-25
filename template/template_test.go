package template_test

import (
	"template"
	"testing"
)

func TestTemplate(t *testing.T) {
	noc1 := template.NewNotificationCenter(&template.SMS{})
	if noc1.Send() != "Notification:SMS" {
		t.Fatal(noc1.Send())
	}
}
