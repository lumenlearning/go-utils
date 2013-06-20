/******************************************************************************
go-utils Source Code
Copyright (C) 2013 Lumen LLC.

This file is part of the go-utils Source Code.

go-utils is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

go-utils is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with go-utils.  If not, see <http://www.gnu.org/licenses/>.
*****************************************************************************/

package http

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Mainly created to support the use of the simultaneous use of the 
//   github.com/lumenlearning/go-utils/http/UserAgentClient and
//   net/http/Client types.
type Clientish interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Head(url string) (resp *http.Response, err error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// No fuss, no muss, just go get the content from this URL
func GetPageContent(url string) (content string, err error) {
	return GetPageContentWithClient(url, nil)
}

// Also, but use the Clientish I supplied.
func GetPageContentWithClient(url string, client Clientish) (content string, err error) {
	// Create a client if we didn't receive one.
	if client == nil {
		client = &http.Client{}
	}

	// Make the call
	resp, err := client.Get(url)
	if err != nil {
		return "", errors.New(fmt.Sprintf("client.Get => %v", err.Error()))
	}

	// Read the response
	respBytes, err := ReadResponseBody(resp)
	if err != nil {
		return "", errors.New(fmt.Sprintf("ReadResponseBody => %v", err.Error()))
	}

	// Convert to string
	respStr := string(*respBytes)

	return respStr, nil
}

func ReadResponseBody(resp *http.Response) (*[]byte, error) {
	// Make sure we close this when we're done
	defer resp.Body.Close()

	// Empty buffer for content
	body := []byte(nil)

	// Read the body content.  Do we trust the ContentLength value?
	// Maybe we should just ignore ContentLength and opt to just keep reading
	//   until there is nothing left to read.
	contentLength := resp.ContentLength
	if contentLength > 0 {
		body = make([]byte, contentLength)
		resp.Body.Read(body)
	} else {
		bio := bufio.NewReader(resp.Body)
		buf := make([]byte, 4096)

		for {
			n, err := bio.Read(buf)
			if n > 0 {
				body = append(body, buf[0:n]...)
			}
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, err
				}
			}
		}
	}

	return &body, nil
}
