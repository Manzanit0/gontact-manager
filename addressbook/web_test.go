package addressbook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	router = setupRouter()

	retCode := m.Run()
	os.Exit(retCode)
}

func TestListContacts(t *testing.T) {
	db.Add(New("Javier", "jgarcia@email.com"))

	req, err := http.NewRequest(http.MethodGet, "/list", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	response := parseBody(w)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	if len(response.Contacts) != 1 {
		t.Errorf("Expected 1 contact, got different amount.")
	}

	if name := response.Contacts[0].Name; name != "Javier" {
		t.Errorf("Expected: Javier, Received: %s", name)
	}

	if email := response.Contacts[0].Email; email != "jgarcia@email.com" {
		t.Errorf("Expected: jgarcia@email.com, Received: %s", email)
	}
}

func parseBody(w *httptest.ResponseRecorder) Addressbook {
	var response Addressbook
	body, _ := ioutil.ReadAll(w.Body)
	stringBody := string(body)
	json.Unmarshal([]byte(stringBody), &response)
	return response
}
