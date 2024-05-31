package control

import (
	"crime-stats/app/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingStatsAndRetreivingThem(t *testing.T) {
	stats := mocks.NewMockCrimeService(map[string]int{})
	server := CrimeServer{stats}
	crime := "IdentityTheft"

	server.ServeHTTP(httptest.NewRecorder(), newPostStatRequest(crime))
	server.ServeHTTP(httptest.NewRecorder(), newPostStatRequest(crime))
	server.ServeHTTP(httptest.NewRecorder(), newPostStatRequest(crime))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetStatRequest(crime))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
