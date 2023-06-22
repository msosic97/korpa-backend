// handler func za kreiranje user-a
// biznis logika samo za user
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/msosic97/korpa-backend/database/user"
	"github.com/msosic97/korpa-backend/models"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

    // Parse the request body to get user data
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
	defer r.Body.Close()

    // Validate the user data (optional)
    if !user.Validate() {
        http.Error(w, "Invalid user data", http.StatusBadRequest)
        return 
    }

    // Perform the user registration logic
    err = db.RegisterUser(user)
    if err != nil {
        http.Error(w, "User registration failed", http.StatusInternalServerError)
        return
    }

    // Return a success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User registered successfully"))
}