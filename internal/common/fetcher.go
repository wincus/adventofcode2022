package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// GetData retrieves data used to solve day n
func GetData(n int) ([]string, error) {

	var data []string

	h, ok := os.LookupEnv("SESSION")

	if !ok {
		return nil, fmt.Errorf("SESSION env not found")
	}

	// https://adventofcode.com/2022/day/%v/input
	u := &url.URL{
		Scheme: "https",
		Host:   "adventofcode.com",
		Path:   fmt.Sprintf("%v/day/%v/input", 2022, n),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("cookie", fmt.Sprintf("session=%v", h))

	c := &http.Client{}

	res, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("could not read response body: %v", err)
	}

	defer res.Body.Close()

	data = strings.Split(string(b), "\n")

	log.Printf("got %v lines of input data", len(data))

	return data, nil

}

func ShowData(d []string) {

	for _, s := range d {
		fmt.Printf("%v\n", string(s))
	}

}

// remove empty lines
func Trim(s []string) []string {

	i := 0

	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}

	return s[:i]
}
