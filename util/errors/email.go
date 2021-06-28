package errors

import (
	"fmt"
	"net/http"

	"gitlab.com/renodesper/gokit-microservices/util/error"
)

var (
	FailedMissingSMTPServerConfig   = error.NewError(http.StatusBadRequest, "PG4001", fmt.Errorf("SMTP server config is empty"))
	FailedMissingSMTPPortConfig     = error.NewError(http.StatusBadRequest, "PG4002", fmt.Errorf("SMTP port config is empty"))
	FailedMissingSMTPUser           = error.NewError(http.StatusBadRequest, "PG4003", fmt.Errorf("SMTP user is empty"))
	FailedMissingSMTPSenderIdentity = error.NewError(http.StatusBadRequest, "PG4004", fmt.Errorf("SMTP sender identity is empty"))
	FailedMissingSMTPSenderEmail    = error.NewError(http.StatusBadRequest, "PG4005", fmt.Errorf("SMTP sender email is empty"))
	FailedMissingSMTPReceiverEmail  = error.NewError(http.StatusBadRequest, "PG4006", fmt.Errorf("no receiver email configured"))

	FailedGenerateEmailBody    = error.NewError(http.StatusBadRequest, "PG4007", fmt.Errorf("failed in generating html mail body"))
	FailedCreateEmailDirectory = error.NewError(http.StatusBadRequest, "PG4008", fmt.Errorf("failed in creating mail directory"))
	FailedWriteEmail           = error.NewError(http.StatusBadRequest, "PG4009", fmt.Errorf("failed in writing html mail"))

	FailedGeneratePlainText = error.NewError(http.StatusBadRequest, "PG4010", fmt.Errorf("failed in generating plain text mail body"))
	FailedWritePlainText    = error.NewError(http.StatusBadRequest, "PG4011", fmt.Errorf("failed in writing plain text mail"))

	FailedGenerateEmail = error.NewError(http.StatusBadRequest, "PG4012", fmt.Errorf("failed in generating html mail"))
	FailedSendEmail     = error.NewError(http.StatusBadRequest, "PG4013", fmt.Errorf("failed in sending mail"))
)
