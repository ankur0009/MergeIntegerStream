package main 

import (
	"fmt"
    "net/http"
    "io/ioutil"
    "os"
    "encoding/json"
)

// URL for the peloton API for reading streams
const pelotonReadStreamURL = "https://api.pelotoncycle.com/quiz/next/"

// Encapsulates the JSON payload for the stream. 
type Message struct {
	Current int
    Last int
    Stream string
}

func main() {
    http.HandleFunc("/", handleHTTPRequest)
    http.ListenAndServe(":8080", nil)
}

func handleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	current, last, stream := readStream("test1") 
    fmt.Fprintf(w, "Current:%d, Last:%d, Stream:%s", current, last, stream)
}

// Reads the stream with the given steam name, and returns <current integer, last integer, stream name> 
func readStream (stream string) (current int, last int, responseStream string) {
	// Builds the URL endpoint to fetch the stream.
	pelotonReadStreamURLWithEndpoint := ""
	pelotonReadStreamURLWithEndpoint += pelotonReadStreamURL
	pelotonReadStreamURLWithEndpoint += stream
	
	// Gets the JSON response from the stream.
	response, err := http.Get(pelotonReadStreamURLWithEndpoint)
	
	// Storos the unmarshalled JSON response.
	var m Message
	
	// Unmarshals the JSON response
    if err != nil {
        fmt.Printf("LOG_ERROR:%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("LOG_ERROR:%s", err)
            os.Exit(1)
        } else {
        		fmt.Printf("LOG:%s\n", string(contents))
        		
        		err := json.Unmarshal(contents, &m)
       		 	if err != nil {
            		fmt.Printf("LOG_ERROR:s", err)
            		os.Exit(1)
        	 	}
       } 	
    }
    
    // Returns the unmarshalled response elements. 
    return m.Current, m.Last, m.Stream 
}
