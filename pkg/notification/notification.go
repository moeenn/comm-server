package notification

type Notification struct {
	UserIds []string
	Payload []byte
}

// TODO: create a pending notification for the user, taking in []byte for payload
