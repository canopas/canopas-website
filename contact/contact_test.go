package contact

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
	"github.com/stretchr/testify/assert"
)

//go:embed templates/contact-email-template.html
var templateFS embed.FS

var repo *Template
var err error
var testRequests []utils.TestRequest

const (
	SEND_CONTACT_MAIL = "/api/send-contact-mail"
)

var contactDetailInput = ContactDetails{
	Name:        "shruti",
	Email:       "shruti@gmail.com",
	ProjectInfo: "Describe small inaformation about my project",
	Reference:   "Canopas Employee",
	ContactType: "Chat or Email",
	Token:       "xyz123",
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

func Test_init(t *testing.T) {
	repo, err = initializeRepo()
	assert.Nil(t, err)
	testRequests = []utils.TestRequest{
		{
			Url:               SEND_CONTACT_MAIL,
			Method:            "POST",
			Headers:           nil,
			Body:              contactDetailInput,
			ResponseCode:      http.StatusOK,
			ResponseTypeArray: false,
			ExpectedData:      "Contact mail sent successfully",
		},
	}
}

func TestAllAPIs(t *testing.T) {

	asserts := assert.New(t)
	engine := gin.New()

	setUpRouter(engine)

	for _, testData := range testRequests {

		w := httptest.NewRecorder()
		var req *http.Request
		var got interface{}

		if testData.Body != nil {
			requestByte, _ := json.Marshal(testData.Body)
			reqBodyData := bytes.NewReader(requestByte)
			req, err = http.NewRequest(testData.Method, testData.Url, reqBodyData)
		} else {
			req, err = http.NewRequest(testData.Method, testData.Url, nil)
		}

		asserts.NoError(err)

		engine.ServeHTTP(w, req)
		assert.Equal(t, testData.ResponseCode, w.Code)
		if testData.ResponseTypeArray {
			got = utils.GotArrayData(w, t)
		} else {
			got = utils.GotData(w, t)
		}

		assert.Equal(t, testData.ExpectedData, got)

	}
}

func initializeRepo() (*Template, error) {

	var utilsRepo stubUtilsRepo

	repo = New(templateFS, &utilsRepo)

	return repo, err
}

// configure api you want to test
func setUpRouter(engine *gin.Engine) {
	engine.POST(SEND_CONTACT_MAIL, repo.SendContactMail)
}
