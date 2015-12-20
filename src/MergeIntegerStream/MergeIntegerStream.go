package main 

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "encoding/json"
)

// URL for the Peloton API for reading integer streams.
const pelotonReadStreamURL = "https://api.pelotoncycle.com/quiz/next/"

// Array that stores all numbers recieved from the Peloton integer streams. 
// Numbers that are generated (output) are removed from the array.
var allInputNumbers = make([]int, 0)

// Index for empty spots in the "allNumbers" array. 
// An empty spot is created, when a number is generated "output". 
// Default value is -1, which means there is no empty spot.
var emptySpotIndex = -1

// Current and last lowest numbers.
var currentLowest, lastLowest = -1, 0

// Encapsulates the JSON payload for the stream. 
type Message struct {
	Current int
    Last int
    Stream string
}

func main() {
	// Listens to URL requests at port 8080.
    http.HandleFunc("/", handleHTTPRequest)
    http.ListenAndServe(":8080", nil)
}

func handleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// Gets the first and second stream names. 
	// Note: Query path is ignored and no checks provided to validate the query path.
	//       Additional checks can be added.
	Stream1 := r.URL.Query().Get("stream1")
	Stream2 := r.URL.Query().Get("stream2")
	
	// Process only if the stream names are not empty. 
	if (Stream1 != "" && Stream2 != "") {
		// Reads the "Message" from the API.  
		current1, last1, stream1 := readStream(Stream1) 
		fmt.Printf("LOG:Current:%d, last%d, Stream:%s\n", current1, last1, stream1)		
    	current2, last2, stream2 := readStream(Stream2) 
    	fmt.Printf("LOG:Current:%d, last%d, Stream:%s\n", current2, last2, stream2)
    	
    	// Appends the current number(s) from each stream to "allInputNumbers".
    	// Note: The last number is ignored, size it is zero for the first request
    	//       and subseqeuntly has already been appended by the previous request.
    	// IMP NOTE: If you use different stream names for subsequent requests, the data
    	//           would all be merged into a single list. In other words, we are not
    	//           maintaining data per stream pair.   	
   		appendNumber(current1)
    	appendNumber(current2)
    	
    	// Gets the current lowest nubmer.
    	currentLowest := lowestNumber()
    	// Outputs the last and current lowest numbers.
    	fmt.Fprintf(w, "{\"last\" : %d, \"current\" : %d}", lastLowest, currentLowest)
    	
    	// Updates the last and current lowset numbers, respectively.
    	lastLowest = currentLowest;
    	currentLowest = -1;
    }
}

func appendNumber(number int) {
	if(emptySpotIndex != -1) {
		allInputNumbers[emptySpotIndex] = number;
		emptySpotIndex = -1
	} else {
		allInputNumbers = append(allInputNumbers, number)
	}
	
	for _, v := range allInputNumbers {
		fmt.Printf("%d,",v)
	}
	
	fmt.Printf("\n")
}

func lowestNumber () (lowest int) {
	lowest = allInputNumbers[0]
	emptySpotIndex = 0;
	
	for i, v := range allInputNumbers {
		if(v < lowest) {
			lowest = v
			emptySpotIndex = i;
		}
	}
	return lowest
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
