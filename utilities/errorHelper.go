package utilities

type Errors interface {
	Errors() []string
}
type ErrorMsgs struct {
	Messsages []string
}

func (e *ErrorMsgs) Errors() []string { return e.Messsages }

func New(text []string) Errors { return &ErrorMsgs{text} }
