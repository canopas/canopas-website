package contact

import (
	"bytes"
	"embed"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//go:embed templates/contact-email-template.html
var templateFS embed.FS

var repo *Template

func TestSendContactMail(t *testing.T) {
	engine := gin.New()
	repo := New(templateFS, &stubUtilsRepo{})

	engine.POST("/api/send-contact-mail", repo.SendContactMail)

	contactDetailInput := ContactDetails{
		Name:        "shruti",
		Email:       "shruti@gmail.com",
		ProjectInfo: "Describe small inaformation about my project",
		Reference:   "Canopas Employee",
		ContactType: "Chat or Email",
		Token:       "xyz123",
	}
	requestByte, _ := json.Marshal(contactDetailInput)
	reqBodyData := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/api/send-contact-mail", reqBodyData)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}
	w := httptest.NewRecorder()

	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	var response string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error decoding response JSON: %v", err)
	}
	expected := "Contact mail sent successfully"

	assert.Equal(t, expected, response)
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
