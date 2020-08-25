package template

type ExportedMethod interface {
	FormatMsg() string
	CheckList() bool
}

type NotificationCenter struct {
	ExportedMethod
}

func NewNotificationCenter(method ExportedMethod) *NotificationCenter {
	return &NotificationCenter{
		ExportedMethod: method,
	}
}

// NOTE: you shall not implement FormatMst() and CheckList()

func (noc *NotificationCenter) Send() string {
	if noc.CheckList() {
		return "Notification:" + noc.FormatMsg()
	}

	return "Nil"
}

type SMS struct {
}

func (sms *SMS) CheckList() bool {
	return true
}

func (sms *SMS) FormatMsg() string {
	return "SMS"
}
