package main

import (
	"log"
	"net/http"
)

// home displays the home page for the site.
func home(w http.ResponseWriter, r *http.Request) {
	// Ensure only "/" is matched
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// showSnippet returns a specific snippet to display.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// createSnippet creates a new snippet.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		// * Note: The following methods must be called in this
		// * order for them to work. Otherwise defaults will be
		// * called which cannot be changed later.
		// Inform user of allowed methods
		w.Header().Set("Allow", http.MethodPost)
		// If response is JSON, must set that in the Header map or it
		// will be sent as plaintext.

		// Send response
		http.Error(w, "Method Not Allowed", 405)
		// Above error is a shortcut for the following two lines
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		return
	}

	// Success
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Register paths
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Start webserver at port 4000
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
