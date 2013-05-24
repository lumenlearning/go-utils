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

package html

import (
	"sync"
	"code.google.com/p/go-html-transform/h5"
	gnhtml "code.google.com/p/go.net/html"
)

func GetNodeText(n *gnhtml.Node) string {
	nodeTree := h5.NewTree(n)
	
	texts := make(chan string)
	wg := sync.WaitGroup{}
	finalString := ""
	
	wg.Add(1)
	go func () {
		nodeTree.Walk(func (c *gnhtml.Node) {
			if c.Type == gnhtml.TextNode {
				texts <- c.Data
			}
		})
		
		close(texts)
		wg.Done()
	}()

	wg.Add(1)
	go func () {
		for t := range texts {
			finalString += t
		}

		wg.Done()
	}()
	
	wg.Wait()
	return finalString
}