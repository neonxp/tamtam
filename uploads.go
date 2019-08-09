package tamtam

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

type uploads struct {
	client *client
}

func newUploads(client *client) *uploads {
	return &uploads{client: client}
}

//GetUploadURL returns url to upload files
func (a *uploads) GetUploadURL(uploadType UploadType) (*UploadEndpoint, error) {
	result := new(UploadEndpoint)
	values := url.Values{}
	values.Set("type", string(uploadType))
	body, err := a.client.request(http.MethodPost, "uploads", values, nil)
	if err != nil {
		return result, err
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return result, json.NewDecoder(body).Decode(result)
}

//UploadMedia uploads file to TamTam server
func (a *uploads) UploadMedia(endpoint *UploadEndpoint, filename string) (*UploadedInfo, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("data", filename)
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := fh.Close(); err != nil {
			log.Println(err)
		}
	}()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	resp, err := http.Post(endpoint.Url, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	result := new(UploadedInfo)
	return result, json.NewDecoder(resp.Body).Decode(result)
}

//UploadPhoto uploads photos to TamTam server
func (a *uploads) UploadPhoto(filename string) (*PhotoTokens, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("data", filename)
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := fh.Close(); err != nil {
			log.Println(err)
		}
	}()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}

	endpoint, err := a.GetUploadURL(PHOTO)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()

	resp, err := http.Post(endpoint.Url, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	result := new(PhotoTokens)
	return result, json.NewDecoder(resp.Body).Decode(result)
}
