package cookie_test

import (
	cookie "github.com/derrickwilliams/gock-play/older"
	g "gopkg.in/h2non/gock.v0"
	"testing"
)

func TestGimmeCookie(t *testing.T) {
	g.New("https://www.google.com").ReplyFunc(func(r *g.Response) {
		r.BodyString("hello from boogers")
	})

	cookie.GimmeCookie(func(incoming string) {
		t.Logf("%s: \n Get you some ", incoming)
	})
}
