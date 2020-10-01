package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const authorization = "Authorization"

type Config struct {
	client  http.Client
	baseUrl string
}

func main() {
	cfg := Config{
		client:  http.Client{},
		baseUrl: "",
	}
	ln := len(os.Args)
	if ln < 3 {
		log.Fatal("missing required arguments (countercli method url [auth]")
	}
	path := os.Args[2]
	switch os.Args[1] {
	case "get":
		num, err := cfg.getCounter(path)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(num)
	case "create":
		st, err := cfg.postCounter(path)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("value: 0, token: %v", st)
	case "increment":
		if ln < 4 {
			log.Fatal("missing required arguments (countercli method url [auth]")
			return
		}
		token := "Bearer " + os.Args[3]
		foo, err := cfg.putCounter(path, token)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(foo)
	case "delete":
		if ln < 4 {
			log.Fatal("missing required arguments (countercli method url [auth]")
			return
		}
		token := "Bearer " + os.Args[3]
		err := cfg.deleteCounter(path, token)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("deleted successfully")
	default:
		log.Fatal("bruh moment")
	}
}

func (c *Config) deleteCounter(path string, authHeader string) error {
	req, err := http.NewRequest(http.MethodDelete, c.baseUrl+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set(authorization, authHeader)
	_, err = c.client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) putCounter(path string, authHeader string) (int, error) {
	req, err := http.NewRequest(http.MethodPut, c.baseUrl+path, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set(authorization, authHeader)
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	m := make(map[string]int)
	if err := json.Unmarshal(bod, &m); err != nil {
		return 0, err
	}
	if c.baseUrl == "" {
		path = "/" + strings.SplitAfterN(path, "/", 4)[3]
	}
	return m[path], nil

}

func (c *Config) postCounter(path string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, c.baseUrl+path, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(resp.Header.Get(authorization), "Bearer "), nil
}

func (c *Config) getCounter(path string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+path, nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode > 399 {
		return 0, errors.New(string(bod))
	}
	m := make(map[string]int)
	if err := json.Unmarshal(bod, &m); err != nil {
		return 0, err
	}
	if c.baseUrl == "" {
		path = "/" + strings.SplitAfterN(path, "/", 4)[3]
	}
	return m[path], nil
}
