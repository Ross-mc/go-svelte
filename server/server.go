package server

import (
	"net/http"
	"strings"

	"github.com/ross-mc/go-svelte-proj/users"
)

func Router(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/users") {
		users.UserRouter(w, r)
	}

}
