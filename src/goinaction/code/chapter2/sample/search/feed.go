package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	// 关键字defer 来安排调用Close 方法，可以保证这个函数一定会被调用。哪怕函
	// 数意外崩溃终止，也能保证关键字defer安排调用的函数会被执行
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	// 将文件解码到一个切片里,这个切片的每一项是一个指向一个Feed 类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
