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
	"fmt"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type UserAgentClient struct {
	UserAgentString string
	Client *http.Client
}

func (u *UserAgentClient) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", u.UserAgentString)
	return u.Client.Do(req)
}

func (u *UserAgentClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http.NewRequest => %v", err.Error()))
	}
	req.Header.Set("User-Agent", u.UserAgentString)
	return u.Client.Do(req)
}

func (u *UserAgentClient) Head(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http.NewRequest => %v", err.Error()))
	}
	req.Header.Set("User-Agent", u.UserAgentString)
	return u.Client.Do(req)
}

func (u *UserAgentClient) Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("http.NewRequest => %v", err.Error()))
	}
	req.Header.Set("User-Agent", u.UserAgentString)
	req.Header.Set("Content-Type", bodyType)
	return u.Client.Do(req)
}

func (u *UserAgentClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return u.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}