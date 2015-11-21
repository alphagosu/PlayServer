// TODO: Which kind of test is better? How do you use them?
package test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
//
// 	h "github.com/alphagosu/PlayServer/handlers"
// )
//
// const (
// 	baseURL = "http://localhost:8080/"
// )

// func TestRootHandlerExternal(t *testing.T) {
// 	req, err := http.NewRequest("GET", baseURL, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	client := http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("Got status code: '%d', expected '%d'", resp.StatusCode, http.StatusOK)
// 	}
// }

// func TestRootHandlerInternal(t *testing.T) {
// 	req, err := http.NewRequest("GET", baseURL, nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	rr := httptest.NewRecorder()
// 	err = h.RootHandler(rr, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if rr.Code != http.StatusOK {
// 		t.Fatalf("Got status code: '%d', expected '%d'", rr.Code, http.StatusOK)
// 	}
// }
