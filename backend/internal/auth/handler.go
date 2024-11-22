package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func LoginHandler(secretKey []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var creds LoginRequest
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			log.Printf("ERROR: Ungültige Login-Anfrage: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		response := LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    int64(15 * time.Minute / time.Second),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("ERROR: Fehler beim Senden der Antwort für Benutzer %s: %v", creds.Username, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("INFO: Benutzer %s erfolgreich angemeldet", creds.Username)

	}
}

func RefreshHandler(secretKey []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RefreshRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("ERROR: Ungültige Refresh-Anfrage: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		claims, err := VerifyRefreshToken(req.RefreshToken, secretKey)
		if err != nil {
			log.Printf("ERROR: Ungültiger Refresh Token: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken, err := GenerateAccessToken(claims.UserID, secretKey)
		if err != nil {
			log.Printf("ERROR: Fehler beim Generieren des Access Tokens für user_id %s: %v", claims.UserID, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: req.RefreshToken,
			ExpiresIn:    int64(15 * time.Minute / time.Second),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("ERROR: Fehler beim Senden der Antwort für user_id %s: %v", claims.UserID, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Printf("INFO: Access Token erfolgreich erneuert für user_id %s", claims.UserID)
	}
}
