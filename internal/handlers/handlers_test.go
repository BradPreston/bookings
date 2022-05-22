package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key string
	value string
}

var theTests = []struct {
	name string
	url string
	method string
	params []postData
	expectedStatusCode int
} {
	{ "home", "/", "GET", []postData{}, 200, },
	{ "about", "/about", "GET", []postData{}, 200, },
	{ "gq", "/generals-quarters", "GET", []postData{}, 200, },
	{ "ms", "/majors-suite", "GET", []postData{}, 200, },
	{ "sa", "/search-availability", "GET", []postData{}, 200, },
	{ "contact", "/contact", "GET", []postData{}, 200, },
	{ "mr", "/make-reservation", "GET", []postData{}, 200, },
	{ "post-sa", "/search-availability", "POST", []postData{
		{ key: "start", value: "01-01-2020" },
		{ key: "end", value: "01-02-2020" },
		}, http.StatusOK, 
	},
	{ "post-sa-json", "/search-availability-json", "POST", []postData{
		{ key: "start", value: "01-01-2020" },
		{ key: "end", value: "01-02-2020" },
		}, http.StatusOK, 
	},
	{ "post-mr", "/make-reservation", "POST", []postData{
		{ key: "first_name", value: "John" },
		{ key: "last_name", value: "Smith" },
		{ key: "email", value: "me@here.com" },
		{ key: "phone", value: "123-456-7890" },
		}, http.StatusOK, 
	},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal()
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal()
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}