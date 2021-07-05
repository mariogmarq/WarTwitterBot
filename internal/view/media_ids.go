package view

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// Post an image returning his media_id
func (v *View) postImage(image string) (int64, error) {
	const endpoint = "https://upload.twitter.com/1.1/media/upload.json"
	file, _, err := prepareFile(image)
	if err != nil {
		return 0, err
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	f, err := w.CreateFormField("media")
	if err != nil {
		return 0, err
	}
	io.Copy(f, file)
	w.Close()

	req, _ := http.NewRequest("POST", endpoint, &buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := v.httpclient.Do(req)
	if err != nil {
		return 0, err
	}

	var r response
	data, _ := io.ReadAll(resp.Body)
	log.Println(string(data))
	json.Unmarshal(data, &r)

	return r.MediaId, nil
}

// Process the file for posting it, returning it and his data
func prepareFile(file string) (*os.File, os.FileInfo, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, nil, err
	}

	if !stat.Mode().IsRegular() {
		return nil, nil, errors.New("file is not regular")
	}

	return f, stat, nil
}

// Response made by the endpoint
type response struct {
	MediaIdString string `json:"media_id_string"`
	MediaId       int64  `json:"media_id"`
}
