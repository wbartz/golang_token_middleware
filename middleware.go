package main

import (
	"net/http"

	"github.com/kpango/glg"

	jwt "github.com/dgrijalva/jwt-go"
)

func validateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenCookie, err := r.Cookie(tokenName)

	switch {
	case err == http.ErrNoCookie:
		glg.Printf("No token %v", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	case err != nil:
		glg.Printf("Cookie parse error %v", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if tokenCookie.Value == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			glg.Printf("Invalid token!")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		glg.Printf("Valid token!")
		next(w, r)
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			glg.Printf("Token Expired, get a new one.")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return

		default:
			glg.Printf("ValidationError error: %+v", vErr.Errors)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

	default:
		glg.Printf("Token parse error: %v", err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

}
