/*
Package reqcontext contains the request context. Each request will have its own instance of RequestContext filled by the
middleware code in the api-context-wrapper.go (parent package).

Each value here should be assumed valid only per request only, with some exceptions like the logger.
*/
package reqcontext

import (
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// RequestContext is the context of the request, for request-dependent parameters
type RequestContext struct {
	// ReqUUID is the request unique ID
	ReqUUID uuid.UUID
	Database database.AppDatabase
	// Logger is a custom field logger for the request
	Logger logrus.FieldLogger
	User *database.User
}
