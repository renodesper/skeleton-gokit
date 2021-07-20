package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	FailedMissingSMTPServerConfig   = error.NewError(http.StatusBadRequest, "EM1001", fmt.Errorf("smtp server config is empty"))
	FailedMissingSMTPPortConfig     = error.NewError(http.StatusBadRequest, "EM1002", fmt.Errorf("smtp port config is empty"))
	FailedMissingSMTPUser           = error.NewError(http.StatusBadRequest, "EM1003", fmt.Errorf("smtp user is empty"))
	FailedMissingSMTPSenderIdentity = error.NewError(http.StatusBadRequest, "EM1004", fmt.Errorf("smtp sender identity is empty"))
	FailedMissingSMTPSenderEmail    = error.NewError(http.StatusBadRequest, "EM1005", fmt.Errorf("smtp sender email is empty"))
	FailedMissingSMTPReceiverEmail  = error.NewError(http.StatusBadRequest, "EM1006", fmt.Errorf("no receiver email configured"))

	FailedGenerateEmailBody    = error.NewError(http.StatusBadRequest, "EM1007", fmt.Errorf("failed in generating html mail body"))
	FailedCreateEmailDirectory = error.NewError(http.StatusBadRequest, "EM1008", fmt.Errorf("failed in creating mail directory"))
	FailedWriteEmail           = error.NewError(http.StatusBadRequest, "EM1009", fmt.Errorf("failed in writing html mail"))

	FailedGeneratePlainText = error.NewError(http.StatusBadRequest, "EM1010", fmt.Errorf("failed in generating plain text mail body"))
	FailedWritePlainText    = error.NewError(http.StatusBadRequest, "EM1011", fmt.Errorf("failed in writing plain text mail"))

	FailedGenerateEmail = error.NewError(http.StatusBadRequest, "EM1012", fmt.Errorf("failed in generating html mail"))
	FailedSendEmail     = error.NewError(http.StatusBadRequest, "EM1013", fmt.Errorf("failed in sending mail"))
)
