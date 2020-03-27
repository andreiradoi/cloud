package email

import (
	"github.com/fieldkit/cloud/server/data"
)

type Emailer interface {
	SendValidationToken(person *data.User, validationToken *data.ValidationToken) error
	SendRecoveryToken(person *data.User, recoveryToken *data.RecoveryToken) error
}
