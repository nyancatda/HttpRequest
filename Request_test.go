package HttpRequest

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	Body, HttpResponse, err := GetRequest("https://www.google.com/", []string{"Accept-Language:en-US,en;q=0.5"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(Body))
	fmt.Println(HttpResponse.Status)
}
