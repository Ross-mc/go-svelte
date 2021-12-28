package users

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ross-mc/go-svelte-proj/responses"
)

type UserRegistrationRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupSuccessResponse struct {
	Username string `json:"username"`
}

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
		signUpRequest := UserRegistrationRequest{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &signUpRequest)
		if err != nil {
			responses.BadRequest(w)
		} else {
			res := SignupSuccessResponse{
				Username: signUpRequest.Username,
			}
			w.Header().Set("Content-Type", "application/json")
			resJson, _ := json.Marshal(res)
			w.Write(resJson)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responses.MethodNotAllowed(w)
	} else {
		io.WriteString(w, "You logged in")
	}
}
