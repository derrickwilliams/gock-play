package cookie

import (
	"io/ioutil"
	"net/http"
)

func GimmeCookie(p func(string)) {
	r, err := http.Get("https://www.google.com?q=cookies")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	p(string(body))
}
