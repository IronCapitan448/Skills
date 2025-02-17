/* This program reads data from marvel gateway API
and returns titles start with avenger
Things I learned from this are
1. how to use public key and private key
2. query params in net/htpp package of go lang
3. converting a string using crypto md5.
4. Encoding and decoding of json
5. using json.decode instead of marshall and unmarshall which
reduces conversion of bytes
6. filtering json output within child items and when
json has arrays in them
7. Handling url and get method in net/http
8. have to close body to stop memory leaks in code
*/

package main

import (
	"crypto/md5"   // Importing md5 package for hashing
	"encoding/hex" // Importing hex package to convert hash to string
	"encoding/json"
	"fmt" // Importing fmt package for printing logs

	// Importing ioutil for reading API response
	"log"
	"net/http" // Importing http package to make GET requests
	"net/url"  // Importing url package to construct URL with parameters
	"time"     // Importing time package to generate a timestamp
)

const baseurl = "http://gateway.marvel.com/v1/public/comics"

// baseURL    = "http://gateway.marvel.com/v1/public/comics"
const publickey = "0f7aafd577a99485555dc5a643f47926"
const privatekey = "b4e249d6edd8ccaec400d3186f4a74ae9cd51fc3"

type Titles struct {
	Title string `json:"title"`
}
type Results struct {
	Result []Titles `json:"results"`
}
type Apiresponse struct {
	Data Results `json:"data"`
}

func generateMD5Hash(timestamp, privateKey, publicKey string) string {
	unhash := timestamp + privateKey + publicKey
	hashvalue := md5.Sum([]byte(unhash))
	return hex.EncodeToString(hashvalue[:])
}
func generatefullurl(string) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	hash := generateMD5Hash(timestamp, privatekey, publickey)
	params := url.Values{}
	params.Add("apikey", publickey)
	params.Add("hash", hash)
	params.Add("ts", timestamp)
	//params.Add("limit", "5")
	params.Add("format", "comic")
	params.Add("formatType", "comic")
	params.Add("title", "avengers")
	fmt.Println(params)
	fullurl := fmt.Sprintf("%s?%s", baseurl, params.Encode())
	fmt.Println(fullurl)
	//fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullurl)
	if err != nil {
		log.Fatal(err)

	}
	defer resp.Body.Close()
	//looping thru created structs to get to final titles in .data.results[].title
	if resp.StatusCode == http.StatusOK {
		Apiresponse := Apiresponse{}
		json.NewDecoder(resp.Body).Decode(&Apiresponse)
		//fmt.Printf("%v", Apiresponse)
		/* results := results{}
		Marveltitles := Titles{}
		fmt.Printf("%v", Marveltitles) */
		fmt.Println("Titles:")
		for _, item := range Apiresponse.Data.Result {
			fmt.Println(item.Title)
		}
	}

}
func main() {
	generatefullurl(baseurl)
}
