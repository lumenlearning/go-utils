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

package io

import (
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	
	str := string(b)
	return str, nil
}

