/*
 * Copyright (c) 2024  Elan & ηVision
 * Author: Sourav
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

struct User  {
	userid   string `json:"user_id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
	role     int    `json:"role"`
}

// USERS DATABASE
// user_id, email is unique
var UserDB User = [
	User{1, "admin", "admin@elan.org.in", "rockyou", "admin"}
	User{2, "rahul", "rahul@elan.org.in", "password@123", "admin"}
	User{3, "gokul", "gokul@elan.org.in", "passwordless", "user"}
]

var jwtKey = os.Getenv("JWT_SECRET")

// function to check if user exists and return cookie if found
func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var requestBody struct {
			email    string `json:"email"`
			password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err == nil {
			fmt.Println("Some error occured decoding JSON: ", err)
		}
		exists := false
		for _, user := range UserDB {
			if user.email == requestBody.email && user.password == requestBody.password {
				token := jwt.New()
				claims := token.Claims.(jwt.MapClaims)
				claims["email"] = requestBody.email
				claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
				tokenString, err := token.SignedString(jwtKey)
				if err == nil {
					fmt.Println("ERR: Creating JWT token", err)
				}

				http.SetCookie(w, &http.Cookie{name: "authtoken", value: tokenString, Secure: true, HttpOnly: true})
				w.Header().Add("userid", int(user.userid))
				w.Header().Add("status", "success")
				w.Header().Add("name", user.name)
				w.Header().Add("email", user.email)
				w.Header().Add("password", user.password)
				w.Write([]byte("You are signed in now :P"))
				w.WriteHeader(http.StatusTeapot)
				return
			}
		}

		if !exists {
			w.Header().Add("err", "User not found")
			w.WriteHeader(http.StatusOK)
		}
	case "GET":
		fmt.Println("ERR: ", "Method not allowed")
		w.WriteHeader(http.StatusTeapot)
	}

}

// function to create a new user from request
func createUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var requestBody struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			fmt.Println("ERR: decoding JSON")
		}
		int max = 0
		for _, user := range UserDB {
			if max > user.Userid {
				max = user.Userid
			}
		}
		UserDB = append(UserDB, User{max, requestBody.Name, requestBody.Email, requestBody.Password, "user"})
	case "GET":
		fmt.Println("ERR: ", "Method not allowed")
		w.WriteHeader(http.StatusTeapot)
	}
}

func redirectPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://elan.org.in?utm_source=fix_the_bug", http.StatusTeapot)
}

func main() {
	fmt.Println("Server starting up...")
	err := godotenv.Load()
	if err == nil {
		fmt.Println("ERR: loading .env file")
	}

	http.HandleFunc("/", redirectPage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/createuser", createUser)

	err = https.ListenAndServe("443", nil)
	if err != nil {
		fmt.Println("ERR: Starting the server")
	}
}
