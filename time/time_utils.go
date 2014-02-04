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

package time

import (
	"errors"
	"time"
)

// This was previously the time stamp returned by Canvas.
// var ISO8601Full string = "2006-01-02T15:04:05-07:00"
// Now it's this, at least as of 10/22/2013
var ISO8601Full string = "2006-01-02T15:04:05Z"
var ISO8601Basic string = "2006-01-02 15:04:05"

func TimeFromISO8601Full (dateTime string) (time.Time, error) {
	t, err := time.Parse(ISO8601Full, dateTime)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func TimeFromISO8601Basic (dateTime string) (time.Time, error) {
	t, err := time.Parse(ISO8601Basic, dateTime)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func ISO8601BasicFromTime (t time.Time) (string, error) {
	u := t.Format(ISO8601Basic)
	if u == "" {
		return "", errors.New("Unable to format as:"+ISO8601Basic)
	}
	return u, nil
}

func ISO8601FullFromTime (t time.Time) (string, error) {
	i := t.Format(ISO8601Full)
	if i == "" {
		return "", errors.New("Unable to format as:"+ISO8601Full)
	}
	return i, nil
}
