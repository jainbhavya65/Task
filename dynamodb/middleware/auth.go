package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type Userbody struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func Routes(router *mux.Router) {
	router.HandleFunc("/signin", Signin).Methods("POST")
}

func CheckToken(w http.ResponseWriter, req *http.Request) (bool, string) {
	c, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return false, "nil"
		}
		return false, "nil"
	}

	tknStr := c.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, "nil"
		}
		return false, "nil"
	}
	if !tkn.Valid {
		return false, "nil"
	}
	return true, claims.Username
}

func Signin(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var u1 Userbody
		json.Unmarshal(reqBody, &u1)
		var u = Userbody {
			Username: "admin",
			Password: "password",
		}
		if u.Username != "admin" {
			http.Error(w, "Username and/or password do not match", http.StatusUnauthorized)
			return
		}
		bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
		err = bcrypt.CompareHashAndPassword(bs, []byte(u1.Password))
		if err != nil {
			http.Error(w, "password do not match", http.StatusUnauthorized)
			return
		}
		token, err := CreateToken(u.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			MaxAge:   54000,
			HttpOnly: true,
			Secure:   false,
			Domain:   "localhost",
		})
		json.NewEncoder(w).Encode("token: " + token)
}


func CreateToken(userName string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Hour * 15).Unix()
	atClaims["name"] = userName
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
