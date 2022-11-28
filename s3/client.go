package s3

import "fmt"

// Client implements methods that interacts with s3
type Client struct {
}

// New creates new client
func New() *Client {
	return &Client{}
}

// UploadFile uploads file to s3
func (c *Client) UploadFile(string) error {
	fmt.Println("File uploaded to s3")
	return nil
}
