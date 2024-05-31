package control

import (
	"crime-stats/app/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETCrimes(t *testing.T) {
	stub := mocks.NewMockCrimeService(map[string]int{
		"IdentityTheft": 20,
		"StolenVehicle": 10,
	})

	crimeServer := &CrimeServer{stub}
	t.Run("returns Identity Theft", func(t *testing.T) {
		request := newGetStatRequest("IdentityTheft")
		response := httptest.NewRecorder()

		crimeServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Stolen Vehicle", func(t *testing.T) {
		request := newGetStatRequest("StolenVehicle")
		response := httptest.NewRecorder()

		crimeServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetStatRequest("Assault")
		response := httptest.NewRecorder()

		crimeServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestCrimeStats(t *testing.T) {
	stats := mocks.NewMockCrimeService(map[string]int{})

	server := &CrimeServer{stats}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		crime := "StolenVehicle"
		request := newPostStatRequest(crime)
		// request, _ := http.NewRequest(http.MethodPost, "/crimes/IdentityTheft", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(stats.StatCalls) != 1 {
			t.Errorf("got %d calls to IdentityTheft want %d", len(stats.StatCalls), 1)
		}

		if stats.StatCalls[0] != crime {
			t.Errorf("did not store correct stat got %s want %s", stats.StatCalls[0], crime)
		}
	})
}

func newGetStatRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/crimes/%s", name), nil)
	return req
}

func newPostStatRequest(crime string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/crimes/%s", crime), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}
