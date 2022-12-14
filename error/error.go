package error

type Error interface {
	error
	StatusCode() int
	Message() string
}

func New(message string, http int, originalErr error) ErrorType {
	return ErrorType{
		Msg:         message,
		HttpCode:    http,
		OriginalErr: originalErr,
	}
}

type ErrorType struct {
	Msg         string `json:"message"`
	HttpCode    int    `json:"httpCode"`
	OriginalErr error  `json:"originalErr"`
}

func (err ErrorType) Error() string {
	if err.OriginalErr != nil {
		return err.Msg + " " + err.OriginalErr.Error()
	}

	return err.Msg
}

func (err ErrorType) StatusCode() int {
	return err.HttpCode
}

func (err ErrorType) Message() string {
	return err.Msg
}
