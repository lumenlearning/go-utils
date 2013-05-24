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

package web

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	gnpublicsuffix "code.google.com/p/go.net/publicsuffix"
	sel "code.google.com/p/go-html-transform/css/selector"
	h5 "code.google.com/p/go-html-transform/h5"
	lmnhttp "lumenlearning.com/util/http"
)

func CanvasWebLogin(username, password, login1, login2, useragent string) (*lmnhttp.UserAgentClient, error) {
	// Create the client that we'll be using for this session. Use a CookieJar 
	//   object to handle all of the cookie business.
	c := http.Client{}
	jarOptions := cookiejar.Options{}
	jarOptions.PublicSuffixList = gnpublicsuffix.List 
	jar, err := cookiejar.New(&jarOptions)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cookiejar.New => %v", err.Error()))
	}
	c.Jar = jar

	// Wrap the Client object in a UserAgentClient so that we can 
	client := lmnhttp.UserAgentClient{}
	client.UserAgentString = useragent
	client.Client = &c

	// login1 is the first login page, the one that gives us the 
	//   authenticity_token value required for login on Canvas.
	resp, err := client.Get(login1)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("client.Get => %v", err.Error()))
	}

	// Get the page containing the  authenticity_token that we need to
	//   successfully complete the login.
	respBytes, err := lmnhttp.ReadResponseBody(resp)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("lmnhttp.ReadResponseBody => %v", err.Error()))
	}
	respStr := string(*respBytes)

	// Find the <input> element that contains the value we're looking for.
	chn, err := sel.Selector("#login_form input[name=authenticity_token]")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("sel.Selector => %v", err.Error()))
	}
	respTree, err := h5.NewFromString(respStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("h5.NewFromString => %v", err.Error()))
	}
	authTokenNode := chn.Find(respTree.Top())[0]
	authToken := ""
	for _, a := range authTokenNode.Attr {
		if a.Key == "value" {
			authToken = a.Val
			break
		}
	}
	if authToken == "" {
		return nil, errors.New("Could not find a value for authenticity token.")
	}

	// With the given credentials and our newly found authenticity_token, try
	//   to log in to Canvas.
	val := url.Values{}
	val.Set("authenticity_token", authToken)
	val.Set("redirect_to_ssl", "1")
	val.Set("pseudonym_session[unique_id]", username)
	val.Set("pseudonym_session[password]", password)
	val.Set("pseudonym_session[remember_me]", "0")

	resp, err = client.PostForm(login2, val)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("client.PostForm => %v", err.Error()))
	}

	return &client, nil
}
