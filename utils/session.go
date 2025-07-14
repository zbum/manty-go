package utils

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
)

type SessionData struct {
	UserID   int
	Username string
}

var (
	sessions   = make(map[string]SessionData)
	sessionsMu sync.RWMutex
)

func generateSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func SetSession(w http.ResponseWriter, userID int, username string) {
	sessionID := generateSessionID()
	sessionsMu.Lock()
	sessions[sessionID] = SessionData{UserID: userID, Username: username}
	sessionsMu.Unlock()

	cookie := &http.Cookie{
		Name:     "MANTY_SESSION",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
}

func GetSession(r *http.Request) *SessionData {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil
	}
	sessionsMu.RLock()
	defer sessionsMu.RUnlock()
	if session, ok := sessions[cookie.Value]; ok {
		return &session
	}
	return nil
}
