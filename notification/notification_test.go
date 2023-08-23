package notification

import (
	"bytes"
	"embed"
	"encoding/json"
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

var invitationRequest = NotificationData{
	Receiver:    "test@canopas.com",
	CompanyName: "Canopas",
	SpaceLink:   "https://test.com",
}

func TestInit(t *testing.T) {
	repo = New(templateFS, &stubUtilsRepo{})
}

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

func TestSendNotificationMail(t *testing.T) {
	engine := gin.New()
	engine.POST("/api/invitation", repo.SendInvitationMail)
	requestByte, _ := json.Marshal(invitationRequest)
	reqBodyData := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/api/invitation", reqBodyData)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotData(w, t)

	expected := "Invitation mail has been sent successfully"

	assert.Equal(t, expected, got)
}

func TestSendAcceptenceMail(t *testing.T) {
	engine := gin.New()
	engine.POST("/api/acceptence", repo.SendAcceptenceMail)
	requestByte, _ := json.Marshal(invitationRequest)
	reqBodyData := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/api/acceptence", reqBodyData)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotData(w, t)

	expected := "Acceptence mail has been sent successfully"

	assert.Equal(t, expected, got)
}
