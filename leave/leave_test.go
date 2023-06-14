package leave

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

//go:embed templates/new-leave-email-template.html
var templateFS embed.FS

var repo *LeaveRepository
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

func (faker *stubUtilsRepo) getNewLeaveEmailTemplate(input []string) string {
	return ""
}

func Test_SendLeaveRequest(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/leave/new", bytes.NewBuffer([]byte(`{"Name":"test","Date":"3 jan 2023","Status":"request","Receiver":"test@canopas.com"}`)))
	assert.NoError(t, err)

	engine := gin.New()
	setUpRouter(engine)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_SendUpdateLeaveMail(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/leave/update", bytes.NewBuffer([]byte(`{"Name":"test","Date":"3 jan 2023","Status":"approve","Receiver":"test@canopas.com"}`)))
	assert.NoError(t, err)

	engine := gin.New()
	setUpRouter(engine)
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// configure api you want to test
func setUpRouter(engine *gin.Engine) {
	engine.POST("/api/leave/new", repo.SendLeaveRequest)
	engine.POST("/api/leave/update", repo.SendUpdateLeaveMail)
}
