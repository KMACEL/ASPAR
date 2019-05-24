package pion

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/websocket"
)

// ProjectPages represents the examples loaded from pages.json.
type ProjectPages []*ProjectPage

// ProjectPage from pages.json.
type ProjectPage struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Type        string `json:"type"`
	IsJS        bool
	IsWASM      bool
}

const (
	index     = "pion/index.html"
	linkAspar = "aspar"
	contact   = "contact"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Pion is
func Pion() {
	addr := flag.String("address", ":8080", "Address to host the HTTP server on.")
	flag.Parse()

	log.Println("Listening on", *addr)
	err := serve(*addr)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func serve(addr string) error {
	// Load the pages
	pages, err := getPages()
	if err != nil {
		return err
	}

	// Load the templates
	homeTemplate := template.Must(template.ParseFiles(index))

	// Serve the required pages
	// DIY 'mux' to avoid additional dependencies
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if url == "/wasm_exec.js" {
			http.FileServer(http.Dir("./vendor-wasm/golang.org/misc/wasm/")).ServeHTTP(w, r)
			return
		}

		// Split up the URL. Expected parts:
		// 1: Base url
		// 2: "example"
		// 3: Example type: js or wasm
		// 4: Example folder, e.g.: data-channels
		// 5: Static file as part of the example
		parts := strings.Split(url, "/")
		if len(parts) > 4 &&
			parts[1] == linkAspar {
			pageType := parts[2]
			pageLink := parts[3]

			for _, page := range *pages {
				if page.Link != pageLink {
					continue
				}
				fiddle := filepath.Join(pageLink, contact)
				fiddle = filepath.Join("pion", fiddle)

				if len(parts[4]) != 0 {
					http.StripPrefix("/"+linkAspar+"/"+pageType+"/"+pageLink+"/", http.FileServer(http.Dir(fiddle))).ServeHTTP(w, r)
					return
				}

				temp := template.Must(template.ParseFiles("pion/aspar.html"))
				_, err = temp.ParseFiles(filepath.Join(fiddle, "contact.html"))
				if err != nil {
					panic(err)
				}

				data := struct {
					*ProjectPage
					JS bool
				}{
					page,
					pageType == "js",
				}

				err = temp.Execute(w, data)
				if err != nil {
					panic(err)
				}
				return
			}
		}

		// Serve the main page
		err = homeTemplate.Execute(w, pages)
		if err != nil {
			panic(err)
		}
	})
	webSocketRead()
	webSocketWrite("MERHABA")
	// Start the server
	return http.ListenAndServe(addr, nil)
}

func webSocketRead() {
	// Web socket
	http.HandleFunc("/socketread", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		for {
			// Read message from browser
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
		}
	})
}

func webSocketWrite(message string) {
	http.HandleFunc("/socketwrite", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		if err := conn.WriteMessage(1, []byte(message)); err != nil {
			return
		}

	})
}

// getPages loads the examples from the pages.json file.
func getPages() (*ProjectPages, error) {
	file, err := os.Open("pion/pages.json")
	if err != nil {
		return nil, fmt.Errorf("failed to list pages (please run in the pages folder): %v", err)
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	var pages ProjectPages
	err = json.NewDecoder(file).Decode(&pages)
	if err != nil {
		return nil, fmt.Errorf("failed to parse examples: %v", err)
	}

	for _, page := range pages {
		fiddle := filepath.Join(page.Link, "contact")
		js := filepath.Join(fiddle, "contact.js")
		if _, err := os.Stat("pion/" + js); !os.IsNotExist(err) {
			page.IsJS = true
		}
		wasm := filepath.Join(fiddle, "contact.wasm")
		if _, err := os.Stat(wasm); !os.IsNotExist(err) {
			page.IsWASM = true
		}
	}

	return &pages, nil
}
