package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	url1 "net/url"
)

// Persis
// STORE MAPIING IN LOCAL Key:Val and use during redirect
//         OR
// ADD SIMPLE MARIADB OR CLUSTERD
// ADD MEMCACHE OR REDIS

type URL struct {
	Url string `json:"url"`
}

type Short URL
type Long URL

var memstore map[string]string

func (url *URL) GetUrl() string {
	return url.Url
}

func redirect(w http.ResponseWriter, r *http.Request) {

	// Step 0.0 CHECK ALLOWED METHODS
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var url URL

	// GET BODY OD REQ
	// VALIDATE BODY -- TBD
	// Fill Body to URL

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		// DEBUG : fmt.Fprintf(w, err.Error(), url)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Step 1: Get Url

	fmt.Println("Memory Store: %+v", memstore)
	longUrl, ok := memstore[url.Url]

	fmt.Println("Short URL: %+v", url)
	fmt.Println("Long URL: %+v", longUrl)
	if !ok {
		// 404
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	// Step 2 : Get Parameter for url
	// 2.2 Validate for empty params.
	// Step 3 : Get the mapping from repository
	// Step 4 : redirect to the
	// s := r.Get
	// Redirect
	w.Header().Set("Location", longUrl)
	http.Redirect(w, r, longUrl, http.StatusSeeOther)
}

func shortUrl(w http.ResponseWriter, r *http.Request) {

	// Step 0.0 CHECK ALLOWED METHODS
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var url URL

	// GET BODY OD REQ
	// VALIDATE BODY -- TBD
	// Fill Body to URL

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		// DEBUG : fmt.Fprintf(w, err.Error(), url)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// VALIDATE VALID URI
	_, err = url1.ParseRequestURI(url.Url)
	if err != nil {
		fmt.Println("Invalid URL:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// DEBUG : SHAN
	// fmt.Fprintf(w, "URL: %+v", url)

	// longurl := url.GetLong()

	// Hash the long to short url
	hash := md5.Sum([]byte(url.Url))

	shortUrl := URL{hex.EncodeToString(hash[:])[:8]}

	// Check Colllision and regenerate if required - TBD

	// Persist in the DB or IN-MEM STORE
	memstore[shortUrl.Url] = url.Url

	fmt.Println("Memory Store: %+v", memstore)
	// SEND URL
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortUrl)
}

func main() {

	// INIT

	memstore = make(map[string]string)
	// ADD ROUTES
	// TBD :  MUX or GIN if need more flex.
	http.HandleFunc("/shorturl", shortUrl)
	http.HandleFunc("/redirect", redirect)

	go http.ListenAndServe(":80", nil)

	fmt.Println("Press <Enter> to stop")
	fmt.Scanln()
}
