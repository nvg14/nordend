package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetFeedFromURL(url string) (res []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		u := strings.Split(url, "https")
		url := "http" + u[1]
		resp, err = http.Get(url)
		if err != nil {
			return res, fmt.Errorf("GET error: %v", err)
		}

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("sead body: %v", err)
	}

	return res, nil
}

func RemoveDuplicateString(entries []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range entries {
		if _, val := keys[entry]; !val {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func CheckIfExists(search string, items []string) bool {
	for _, item := range items {
		if search == item {
			return true
		}
	}
	return false
}

func MakeUrl(host, part string, queryStrings map[string]string) string {
	url := host + part
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("MakeUrl", err)
		os.Exit(500)
		return ""
	}

	q := req.URL.Query()

	for k, v := range queryStrings {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

func Paginate(listSize int, page int, offset int) (int, int) {
	if page > listSize {
		page = listSize
	}

	end := (page * offset) + offset
	if end > listSize {
		end = listSize
	}

	start := page * offset

	return start, end
}
