package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	URL          string
	Host         HostResource
	Image        ImageResource
	ImageVersion ImageVersionResource
	Site         SiteResource
	Tenant       TenantResource
	Subnet       SubnetResource
	BootImage    BootImageResource
	Debug        bool
}

// New client
func New(url string) *Client {
	c := Client{
		URL: url,
	}
	c.Host.Client = &c
	c.Image.Client = &c
	c.ImageVersion.Client = &c
	c.Site.Client = &c
	c.Tenant.Client = &c
	c.Subnet.Client = &c
	c.BootImage.Client = &c
	return &c
}

func (c Client) SetDebug() {
	c.Debug = true
}

// Create resource
func (c Client) Create(endp string, s interface{}) ([]byte, error) {
	url := c.URL + endp
	log.Printf("header: application/json, method: POST, url: %s", url)

	b, _ := json.MarshalIndent(&s, "", "  ")
	fmt.Printf("Payload:\n%s\n", string(b))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body:\n%s\n", string(body))

	return body, nil
}

// Get resource
func (c Client) Get(endp string, name string) ([]byte, error) {
	url := c.URL + endp + "/" + name
	log.Printf("url: %s", url)

	resp, err := http.Get(url + "?envelope=false")
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf("Get %s: failed with status code %d", url, resp.StatusCode)
	}

	defer resp.Body.Close()
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return cont, nil
}

// Exist resource
func (c Client) Exist(endp string, name string) (bool, error) {
	url := c.URL + endp + "/" + name
	log.Printf("url: %s", url)

	resp, err := http.Get(url + "?envelope=false")
	if err != nil {
		return false, err
	}

	switch resp.StatusCode {
	case 404:
		return false, nil
	case 200:
		return true, nil
	}
	return false, fmt.Errorf("Get %s: failed with status code %d", url, resp.StatusCode)
}

// All resources
func (c Client) All(endp string) ([]byte, error) {
	url := c.URL + endp
	log.Printf("url: %s", url)

	resp, err := http.Get(url + "?envelope=false")
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return cont, nil
}
