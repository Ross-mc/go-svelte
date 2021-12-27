package users

import (
	"io"
	"net/http"
	"strings"

	"github.com/ross-mc/go-svelte-proj/responses"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/signup") {
		signUpHandler(w, r)
	}
	if strings.HasSuffix(r.URL.Path, "/login") {
		loginHandler(w, r)
	}

}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responses.MethodNotAllowed(w)
	} else {
		io.WriteString(w, "You signed up")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responses.MethodNotAllowed(w)
	} else {
		io.WriteString(w, "You logged in")
	}
}
