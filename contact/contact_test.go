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
}

//fakeMailRepo is a mock Mail Service Interface
type fakeMailRepo struct{}

func (faker *fakeMailRepo) SendEmail(emailInput *ses.SendEmailInput, jobsInput *ses.SendRawEmailInput) int {
	return 0
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
			ExpectedData:      expectedContactData(),
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

	var fakeEmail fakeMailRepo

	repo = New(templateFS, &fakeEmail)

	return repo, err
}

// configure api you want to test
func setUpRouter(engine *gin.Engine) {
	engine.POST(SEND_CONTACT_MAIL, repo.SendContactMail)
}

func expectedContactData() map[string]interface{} {
	contact := make(map[string]interface{})
	contact["name"] = "shruti"
	contact["email"] = "shruti@gmail.com"
	contact["project_info"] = "Describe small inaformation about my project"
	contact["reference"] = "Canopas Employee"
	contact["contact_type"] = "Chat or Email"
	return contact
}
