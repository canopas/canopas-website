package leave

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

//go:embed templates/new-leave-email-template.html
var templateFS embed.FS

var repo *LeaveRepository
var err error

var leaveRequest = LeaveData{
	Name:     "riya",
	Date:     "3 jan 2023",
	Status:   1,
	Reason:   "Casual Leave",
	Receiver: "riya@canopas.com",
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

func TestSendLeaveRequest(t *testing.T) {
	engine := gin.New()
	engine.POST("/api/leave/new", repo.SendLeaveRequest)
	requestByte, _ := json.Marshal(leaveRequest)
	reqBodyData := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/api/leave/new", reqBodyData)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotData(w, t)

	expected := "Leave request has been sent successfully"

	assert.Equal(t, expected, got)
}
func TestSendUpdateLeaveMail(t *testing.T) {
	engine := gin.New()
	engine.POST("/api/leave/update", repo.SendUpdateLeaveMail)
	requestByte, _ := json.Marshal(leaveRequest)
	reqBodyData := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/api/leave/update", reqBodyData)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	assert.EqualValues(t, http.StatusOK, w.Code)
	got := utils.GotData(w, t)

	expected := "Update leave request has been sent successfully"

	assert.Equal(t, expected, got)
}
