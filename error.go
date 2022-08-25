package gomail

type MailError struct {
	message string
}

func (me *MailError) Error() string {
	return me.message
}

func NewError(msg string) *MailError {
	return &MailError{msg}
}
