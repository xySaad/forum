package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	database "forum/database"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Hand_register_get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "templates/register.html")
}

func Hand_register_post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		redirect("/register", "metthod not allowed", 405, w)
		return
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		redirect("/register", "format not acceptable", 406, w)
		return
	}
	if err := Verify(user); err != nil {
		redirect("/register", err.Error(), 409, w)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		redirect("/register", "internal server err", 500, w)
		return
	}
	uuid, err := uuid.NewV7()
	if err != nil {
		redirect("/register", "internal server err", 500, w)
		// plese handle this err m not in mood
		return
	}
	err = database.CreateUser(user.Username, user.Email, string(hashedPassword), uuid.String())
	if err != nil {
		redirect("/register", "internal server err", 500, w)
		// yess handle this also i hve no idea what to put here
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   uuid.String(),
		Expires: time.Now().Add(5 * time.Hour),
	})
	redirect("/", "registered succesfully", 200, w)
	fmt.Println(IsAuth(r))
}
