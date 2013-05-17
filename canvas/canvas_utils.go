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

package canvas

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func CallAPI(srv, api, auth string, cln *http.Client) (*http.Response, error) {
	// Put together the full URL for the request
	url := "https://"+srv+"/api/v1/"+api

	return AuthorizedCall(url, auth, cln)
}

func AuthorizedCall(url, auth string, cln *http.Client) (*http.Response, error) {
	// Get an http.Client if one was not provided
	if cln == nil {
		cln = new(http.Client)
	}

	// Get an http.Request object ready to go, and add
	//   the Authorization header to it.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer <"+auth+">")

	// Make the request
	resp, err := cln.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the http.Response object.
	return resp, nil
}

func ReadResponse(resp *http.Response) (*[]byte, error) {
	contentLength := resp.ContentLength

	body := []byte(nil)
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

func GetObjFromJSON(data *[]byte) (*interface{}, error) {
	var obj interface{}
	err := json.Unmarshal(*data, &obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func GetNextLink(resp *http.Response) (string, error) {
	linkStr := resp.Header.Get("Link")
	if linkStr == "" {
		return "", nil
	}

	allLinks := strings.Split(linkStr, ",")

	var link string

	for _, lv := range allLinks {
		pieces := strings.Split(lv, ";")

		// Build the regexp and Check the rel tag
		relNext, err := regexp.Compile("rel=[\"']next[\"']")
		if err != nil {
			return "", err
		}

		if relNext.FindStringIndex(pieces[1]) != nil {
			link = strings.Trim(pieces[0], "<>")
			break
		}
	}

	return link, nil
}