package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Download struct {
	Url           string
	TargetPath    string
	TotalSections int
}

func main() {
	fmt.Printf("Download Manager\n")
	startTime := time.Now()

	d := Download{
		Url:           "https://www.dropbox.com/s/l67iz22wvafq4w0/video.mp4?dl=1",
		TargetPath:    "video.mp4",
		TotalSections: 10,
	}

	err := d.Do()
	if err != nil {
		log.Fatalf("An Error occurred while downloading the file %s\n", err)
	}

	fmt.Printf("Download completed in %v seconds\n", time.Since(startTime).Seconds())
}

func (d Download) Do() error {
	fmt.Printf("Making a connection\n")

	// We're only going to make this request to check the size of the file
	// so we can partition appropriately
	r, err := d.getNewRequest("HEAD")
	if err != nil {
		return err
	}

	// call the actual function
	resp, err := http.DefaultClient.Do(r)

	// error handling
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("can't process, response is %v", resp.StatusCode)
	}

	// Print out the result
	fmt.Printf("Got %v\n", resp.StatusCode)

	// Now we see the size of the video file
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	// error handling
	if err != nil {
		return err
	}

	// make an array of size Total Sections,
	// each value will have an int array of size
	// two demarcating the start and end chunks of each section
	var sections = make([][2]int, d.TotalSections)

	// determine each section's size
	eachSize := size / d.TotalSections

	// example: if file size is 100 bytes, our section should like:
	// [[0 10] [11 21] [22 32] [33 43] [44 54] [55 65] [66 76] [77 87] [88 98] [99 99]]
	for i := range sections {
		if i == 0 {
			sections[i][0] = 0
		} else {
			sections[i][0] = sections[i-1][1] + 1
		}

		if i == d.TotalSections-1 {
			sections[i][1] = size - 1
		} else {
			sections[i][1] = sections[i][0] + eachSize
		}
	}
	fmt.Println(sections)

	// do download for each section
	var wg sync.WaitGroup
	for i, s := range sections {
		// capture the current value of i and s as they will continually change
		wg.Add(1)
		i := i
		s := s
		go func() {
			defer wg.Done()

			err = d.downloadSection(i, s)
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()

	fmt.Printf("size is %v bytes\n", size)

	err = d.mergeFiles(sections)
	if err != nil {
		return err
	}

	return nil
}

func (d Download) getNewRequest(method string) (*http.Request, error) {
	r, error := http.NewRequest(
		method,
		d.Url,
		nil,
	)

	if error != nil {
		return nil, error
	}
	r.Header.Set("User-Agent", "Download Manager v1")
	return r, nil
}

func (d Download) downloadSection(i int, s [2]int) error {
	r, err := d.getNewRequest("GET")
	if err != nil {
		return err
	}
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", s[0], s[1]))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	fmt.Printf("Downloaded %v bytes for section %v: %v\n", resp.Header.Get("Content-Length"), i, s)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("section-%v.tmp", i), b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (d Download) mergeFiles(section [][2]int) error {
	f, err := os.OpenFile(d.TargetPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := range section {
		b, err := ioutil.ReadFile(fmt.Sprintf("section-%v.tmp", i))
		if err != nil {
			return nil
		}
		n, err := f.Write(b)
		if err != nil {
			return nil
		}
		fmt.Printf("%v bytes merged\n", n)
	}
	return nil
}
