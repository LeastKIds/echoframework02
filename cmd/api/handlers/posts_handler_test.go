package handlers

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostHandler(t *testing.T) {
	c := http.Client{}
	r, err := c.Get("http://127.0.0.1:8080/health-check")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
	assert.Equal(t, http.StatusOK, r.StatusCode, "X")
}

// func (s *EndToEndSuite) TestPostNoResult() {
// 	c := http.Client{}
// 	r, _ := c.Get("http://127.0.0.1:8080/post/5334")
// 	s.Equal(http.StatusOK, r.StatusCode)
// 	b, _ := io.ReadAll(r.Body)
// 	// json 값을 비교
// 	s.JSONEq(`{"status":"ok", "data":[]}`, string(b))
// }
