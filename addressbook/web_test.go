package addressbook

import (
	"bytes"
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
	router = DefaultRouter()

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

	var response Addressbook
	parseBody(w, &response)

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

func TestAddContacts(t *testing.T) {
	post_data := []byte(`{"name": "Javier","email": "jgarcia@email.com"}`)
	req, err := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(post_data))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response Contact
	parseBody(w, &response)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusCreated, w.Code)
	}

	if name := response.Name; name != "Javier" {
		t.Errorf("Expected: Javier, Received: %s", name)
	}

	if email := response.Email; email != "jgarcia@email.com" {
		t.Errorf("Expected: jgarcia@email.com, Received: %s", email)
	}
}

func readBody(w *httptest.ResponseRecorder) string {
	body, _ := ioutil.ReadAll(w.Body)
	return string(body)
}

func parseBody(w *httptest.ResponseRecorder, response interface{}) {
	body := readBody(w)
	json.Unmarshal([]byte(body), &response)
}
