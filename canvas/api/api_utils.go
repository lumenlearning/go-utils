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

package api

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	lmnHttp "github.com/lumenlearning/go-utils/http"
)

func CallAPI(srv, api, auth string, cln *http.Client) (*http.Response, error) {
	// Put together the full URL for the request
	url := "https://"+srv+"/api/v1/"+api

	return AuthorizedCall(url, auth, cln)
}

func AuthorizedCall(url, auth string, cln lmnHttp.Clientish) (*http.Response, error) {
	// Get an http.Client if one was not provided
	if cln == nil {
		cln = &http.Client{}
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
		return resp, err
	}

	// Return the http.Response object.
	return resp, nil
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