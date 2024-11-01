package api

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	LOUNAAT_INFO_URL = "https://www.lounaat.info"
)

// Restaurant represents the details of a restaurant.
type Restaurant struct {
	Name     string
	Dishes   []string
	Distance string
}

func FetchLounaat(lat, lng float64) ([]Restaurant, error) {
	fullURL := constructURL(lat, lng)

	req, err := createRequest(fullURL)
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := parseResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return parseRestaurants(doc), nil
}

func constructURL(lat, lng float64) string {
	baseURL := fmt.Sprintf("%s/ajax/filter", LOUNAAT_INFO_URL)
	params := url.Values{}
	params.Add("view", "lahistolla")
	params.Add("day", "2")
	params.Add("page", "0")
	params.Add("coords[lat]", fmt.Sprintf("%f", lat))
	params.Add("coords[lng]", fmt.Sprintf("%f", lng))

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func createRequest(fullURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", LOUNAAT_INFO_URL)
	req.Header.Set("Accept-Encoding", "gzip")
	return req, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func parseResponseBody(resp *http.Response) (*goquery.Document, error) {

	reader, err := gzip.NewReader(resp.Body)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	return goquery.NewDocumentFromReader(reader)
}

func traverseNodes(selection *goquery.Selection, depth int) {
	selection.Each(func(i int, s *goquery.Selection) {
		// Print the current node details
		indent := strings.Repeat("  ", depth) // Create an indentation based on the depth
		fmt.Printf("%s<%s", indent, goquery.NodeName(s))

		// Print attributes
		for _, attr := range s.Get(0).Attr {
			fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
		}
		fmt.Printf(">\n")

		// Print text content if the node has any
		if text := strings.TrimSpace(s.Text()); text != "" {
			fmt.Printf("%s  Text: %s\n", indent, text)
		}

		// Recursively print child nodes
		traverseNodes(s.Children(), depth+1)

		fmt.Printf("%s</%s>\n", indent, goquery.NodeName(s)) // Close the current node
	})
}

func parseRestaurants(doc *goquery.Document) []Restaurant {
	var restaurants []Restaurant

	doc.Find(".menu.item").Each(func(i int, s *goquery.Selection) {
		restaurants = append(restaurants, parseRestaurant(s))
	})

	return restaurants
}

func parseRestaurant(s *goquery.Selection) Restaurant {
	restaurant := Restaurant{
		Name:     s.Find(".item-header h3 a").Text(),
		Distance: s.Find(".item-footer .dist").Text(),
		Dishes:   parseDishes(s),
	}
	return restaurant
}

func parseDishes(s *goquery.Selection) []string {
	var dishes []string
	s.Find(".item-body .dish").Each(func(j int, dishSelection *goquery.Selection) {
		dishes = append(dishes, dishSelection.Text())
	})
	return dishes
}
