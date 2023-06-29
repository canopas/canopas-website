package notification

import (
	"bytes"
	"embed"
	"net/http"
	"net/http/httptest"
	"testing"
	"utils"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/tj/assert"
)

//go:embed templates/invitation-email-template.html
var templateFS embed.FS

var repo *NotificationRepository
var err error
var testRequests []utils.TestRequest

// stubUtilsRepo is a mock Utils Service Interface
type stubUtilsRepo struct{}

func (faker *stubUtilsRepo) SendEmail(emailInput *ses.SendEmailInput, jobsInput *ses.SendRawEmailInput) int {
	return 0
}

func (faker *stubUtilsRepo) VerifyRecaptcha(token string) (bool, error) {
	return true, nil
}

func (faker *stubUtilsRepo) SaveJobsToSpreadSheet(input []string) {
	/** this is stub method for adding jobs details in google spreadsheet */
}

func (faker *stubUtilsRepo) getInvitationEmailTemplate(input []string) string {
	return ""
}

func Test_SendNotificationMail(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/invitation", bytes.NewBuffer([]byte(`{"Receiver":"test@canopas.com","CompanyName":"Canopas","SpaceLink":"https://test.com"}`)))
	assert.NoError(t, err)

	engine := gin.New()
	setUpRouter(engine)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// configure api you want to test
func setUpRouter(engine *gin.Engine) {
	engine.POST("/api/invitation", repo.SendInvitationMail)
}
