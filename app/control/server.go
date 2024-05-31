package control

import (
	"fmt"
	"net/http"
	"strings"
)

type CrimeServer struct {
	CrimeService CrimeStore
}

func (c *CrimeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	crime := strings.TrimPrefix(r.URL.Path, "/crimes/")
	switch r.Method {
	case http.MethodPost:
		c.processStat(w, crime)
	case http.MethodGet:
		c.showStat(w, crime)
	}

}

func (c *CrimeServer) processStat(w http.ResponseWriter, crime string) {
	c.CrimeService.RecordStat(crime)
	w.WriteHeader(http.StatusAccepted)
}

func (c *CrimeServer) showStat(w http.ResponseWriter, crime string) {
	stat := c.CrimeService.GetCrimeStat(crime)

	if stat == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, stat)
}
