package function

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Handle a serverless request
func Handle(req []byte) string {

	// read in url
	reqString := strings.Trim(string(req), " \n")
	if _, err := url.ParseRequestURI(reqString); err != nil {
		response := []struct {
			Error   string
			Message string
		}{{
			"error parsing URL",
			err.Error(),
		}}
		output, _ := json.Marshal(response)
		return string(output)
	}

	// call crawler synchronously
	resp, err := http.Post("http://gateway.openfaas:8080/function/openfaas-imagecrawler", "application/json", bytes.NewBuffer(req))
	if err != nil {
		response := []struct {
			Error   string
			Message string
		}{{
			"error crawling URL",
			err.Error(),
		}}
		output, _ := json.Marshal(response)
		return string(output)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	client := &http.Client{}

	// send result to exif feed asynchronously
	exifReq, _ := http.NewRequest("POST", "http://gateway.openfaas:8080/async-function/openfaas-exiffeed", bytes.NewBuffer(body))
	exifReq.Header.Set("X-Callback-Url", "http://gateway:8080/function/openfaas-elastic")
	exipResp, err := client.Do(exifReq)
	if err != nil {
		response := []struct {
			Error   string
			Message string
		}{{
			"error sending to EXIF feed",
			err.Error(),
		}}
		output, _ := json.Marshal(response)
		return string(output)
	}
	defer exipResp.Body.Close()

	// send result to nsfw feed asynchronously
	nsfwReq, _ := http.NewRequest("POST", "http://gateway.openfaas:8080/async-function/openfaas-opennsfwfeed", bytes.NewBuffer(body))
	nsfwReq.Header.Set("X-Callback-Url", "http://gateway:8080/function/openfaas-elastic")
	nsfwResp, err := client.Do(nsfwReq)
	if err != nil {
		response := []struct {
			Error   string
			Message string
		}{{
			"error sending to NSFW feed",
			err.Error(),
		}}
		output, _ := json.Marshal(response)
		return string(output)
	}
	defer nsfwResp.Body.Close()

	// send result to inception feed asynchronously
	inceptionReq, _ := http.NewRequest("POST", "http://gateway.openfaas:8080/async-function/openfaas-inceptionfeed", bytes.NewBuffer(body))
	inceptionReq.Header.Set("X-Callback-Url", "http://gateway:8080/function/openfaas-elastic")
	inceptionResp, err := client.Do(inceptionReq)
	if err != nil {
		response := []struct {
			Error   string
			Message string
		}{{
			"error sending to NSFW feed",
			err.Error(),
		}}
		output, _ := json.Marshal(response)
		return string(output)
	}
	defer inceptionResp.Body.Close()

	return string(body)

}
