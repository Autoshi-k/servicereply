package servicereply

import "fmt"

type ServiceReplyI interface {
	Error() string
	UserMessage() string
}

type ServiceReply struct {
	Status       Status      `json:"status"`
	ServiceError *Error      `json:"error"`
	Data         interface{} `json:"data"`
}

type Status string

const (
	StatusOK    Status = "OK"
	StatusError        = "error"
)

type Error struct {
	ErrorID ErrorID `json:"errorID"`
	Message string  `json:"message"`
	Err     error   `json:"-"`
}

type ErrorID int

const (
	InternalServiceError ErrorID = 0
)

func (sr ServiceReply) Error() string {
	if sr.ServiceError != nil {
		return fmt.Sprintf("[ERR status %d] msg: %s, err: %s", sr.ServiceError.ErrorID, sr.ServiceError.Message, sr.ServiceError.Err)
	} else {
		return ""
	}
}

func (sr ServiceReply) UserMessage() string {
	if sr.ServiceError != nil {
		return sr.ServiceError.Message
	} else {
		return ""
	}
}

func NewOKServiceReply(data interface{}) ServiceReplyI {
	return ServiceReply{
		Status:       StatusOK,
		ServiceError: nil,
		Data:         data,
	}
}

func NewInternalServiceError(err error) ServiceReplyI {
	return ServiceReply{
		Status: StatusError,
		ServiceError: &Error{
			ErrorID: InternalServiceError,
			Message: err.Error(),
			Err:     err,
		},
		Data: nil,
	}
}
