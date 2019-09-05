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

	"github.com/neonxp/tamtam/schemes"
)

type uploads struct {
	client *client
}

func newUploads(client *client) *uploads {
	return &uploads{client: client}
}

//UploadMedia uploads file to TamTam server
func (a *uploads) UploadMediaFromFile(uploadType schemes.UploadType, filename string) (*schemes.UploadedInfo, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	return a.UploadMediaFromReader(uploadType, fh)
}

//UploadMediaFromUrl uploads file from remote server to TamTam server
func (a *uploads) UploadMediaFromUrl(uploadType schemes.UploadType, u url.URL) (*schemes.UploadedInfo, error) {
	respFile, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer respFile.Body.Close()
	return a.UploadMediaFromReader(uploadType, respFile.Body)
}

func (a *uploads) UploadMediaFromReader(uploadType schemes.UploadType, reader io.Reader) (*schemes.UploadedInfo, error) {
	result := new(schemes.UploadedInfo)
	return result, a.uploadMediaFromReader(uploadType, reader, result)
}

//UploadPhotoFromFile uploads photos to TamTam server
func (a *uploads) UploadPhotoFromFile(filename string) (*schemes.PhotoTokens, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	result := new(schemes.PhotoTokens)
	return result, a.uploadMediaFromReader(schemes.PHOTO, fh, result)
}

//UploadPhotoFromUrl uploads photo from remote server to TamTam server
func (a *uploads) UploadPhotoFromUrl(u url.URL) (*schemes.PhotoTokens, error) {
	respFile, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer respFile.Body.Close()
	result := new(schemes.PhotoTokens)
	return result, a.uploadMediaFromReader(schemes.PHOTO, respFile.Body, result)
}

//UploadPhotoFromReader uploads photo from reader
func (a *uploads) UploadPhotoFromReader(reader io.Reader) (*schemes.PhotoTokens, error) {
	result := new(schemes.PhotoTokens)
	return result, a.uploadMediaFromReader(schemes.PHOTO, reader, result)
}

func (a *uploads) getUploadURL(uploadType schemes.UploadType) (*schemes.UploadEndpoint, error) {
	result := new(schemes.UploadEndpoint)
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

func (a *uploads) uploadMediaFromReader(uploadType schemes.UploadType, reader io.Reader, result interface{}) error {
	endpoint, err := a.getUploadURL(uploadType)
	if err != nil {
		return err
	}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("data", "file")
	if err != nil {
		return err
	}
	_, err = io.Copy(fileWriter, reader)
	if err != nil {
		return err
	}

	if err := bodyWriter.Close(); err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	if err := bodyWriter.Close(); err != nil {
		return err
	}
	resp, err := http.Post(endpoint.Url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	return json.NewDecoder(resp.Body).Decode(result)
}
