package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradpreston/bookings/internal/models"
)

type postData struct {
	key string
	value string
}

var theTests = []struct {
	name string
	url string
	method string
	expectedStatusCode int
} {
	{ "home", "/", "GET", 200, },
	{ "about", "/about", "GET", 200, },
	{ "gq", "/generals-quarters", "GET", 200, },
	{ "ms", "/majors-suite", "GET", 200, },
	{ "sa", "/search-availability", "GET", 200, },
	{ "contact", "/contact", "GET", 200, },
	// { "mr", "/make-reservation", "GET", []postData{}, 200, },
	// { "post-sa", "/search-availability", "POST", []postData{
	// 	{ key: "start", value: "01-01-2020" },
	// 	{ key: "end", value: "01-02-2020" },
	// 	}, http.StatusOK, 
	// },
	// { "post-sa-json", "/search-availability-json", "POST", []postData{
	// 	{ key: "start", value: "01-01-2020" },
	// 	{ key: "end", value: "01-02-2020" },
	// 	}, http.StatusOK, 
	// },
	// { "post-mr", "/make-reservation", "POST", []postData{
	// 	{ key: "first_name", value: "John" },
	// 	{ key: "last_name", value: "Smith" },
	// 	{ key: "email", value: "me@here.com" },
	// 	{ key: "phone", value: "123-456-7890" },
	// 	}, http.StatusOK, 
	// },
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal()
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID: 1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Reservation handler returned wrong response code: got %d wanted %d", rr.Code, 200)
	}

	// test case where reservation is not in session (reset eerything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d wanted %d", rr.Code, 200)
	}

	// test with non-existant room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d wanted %d", rr.Code, 200)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}