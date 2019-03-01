package errors

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

// ErrorHandler defines an instance for rendering errors
type ErrorHandler struct {
	Values               map[string]ErrorDto
	defaultError         string
	enableReportToSentry bool
	RavenClientInstance  *raven.Client
}

// ErrorHandlerBehavior wraps all errors for easier message access
type ErrorHandlerBehavior interface {
	Error(e error, extraTags map[string]string) ErrorDto
	ErrorNoSentry(e error) ErrorDto
}

var once sync.Once
var ErrorHandlerInstance *ErrorHandler

// GetInstance returns a global ErrorHandler instance
func GetInstance(enviroment, sentryDSN, version string, reportToSentry bool) *ErrorHandler {

	once.Do(func() {

		var err error
		var ravenClientInstance *raven.Client
		if ravenClientInstance, err = raven.New(sentryDSN); err != nil {
			logrus.Panic(err.Error())
			os.Exit(2)
		}

		ravenClientInstance.SetEnvironment(enviroment)
		ravenClientInstance.SetDSN(sentryDSN)
		ravenClientInstance.SetRelease(version)

		ErrorHandlerInstance = &ErrorHandler{
			Values:               values,
			defaultError:         "MM0000000",
			enableReportToSentry: reportToSentry,
			RavenClientInstance:  ravenClientInstance,
		}

	})

	return ErrorHandlerInstance
}

func (eh *ErrorHandler) ErrorNoSentry(err error) ErrorDto {
	rawErrorMsg := err.Error()
	errorDto := eh.Values[eh.defaultError]
	logrus.Info("errorHandler: " + rawErrorMsg)
	if len(rawErrorMsg) == 0 {
		errorDto.ComponentMsg = "Undefined error code. No error code found"
	} else {
		msg := strings.Split(rawErrorMsg, "|")
		if len(msg) == 0 {
			errorDto.ComponentMsg = "Undefined error code. No error code found"
		} else {
			var id, ok = msg[0], false
			if errorDto, ok = eh.Values[id]; ok {
				errorDto.ComponentMsg = msg[1]
				if len(msg) > 2 {
					errorDto.SentryCode = msg[2]
				}
			} else {
				errorDto = eh.Values[eh.defaultError]
				errorDto.ComponentMsg = "this error doesn't exist. Msg " + rawErrorMsg
			}
		}
	}

	return errorDto
}

func (eh *ErrorHandler) Error(err error, extraTags map[string]string) ErrorDto {

	errorDto := eh.ErrorNoSentry(err)

	tags := eh.getSentryErrorsTags(extraTags)
	tags["id"] = errorDto.ID
	tags["status"] = strconv.Itoa(errorDto.Status)
	tags["sentryCode"] = errorDto.SentryCode

	if eh.enableReportToSentry {
		go func() { eh.RavenClientInstance.CaptureError(err, tags) }()
	}
	return errorDto
}

func (eh *ErrorHandler) getSentryErrorsTags(extraTags map[string]string) map[string]string {
	dst := make(map[string]string, len(extraTags))
	for k, v := range extraTags {
		dst[k] = v
	}

	return dst
}
