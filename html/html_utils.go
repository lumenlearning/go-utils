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
	"errors"
	"fmt"
	"sync"
	ghtsel "code.google.com/p/go-html-transform/css/selector"
	ghth5 "code.google.com/p/go-html-transform/h5"
	gnhtml "code.google.com/p/go.net/html"
)

func FindNodes(html, cssSelector string) ([]*gnhtml.Node, error) {
	// Create a parse tree of the HTML
	tree, err := ghth5.NewFromString(html)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ghth5.NewFromString => %v", err.Error()))
	}

	// Create a selector chain
	sel, err := ghtsel.Selector(cssSelector)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ghtsel.Selector => %v", err.Error()))
	}

	// Return the nodes that we found, if any.
	return sel.Find(tree.Top()), nil
}

func GetNodeText(n *gnhtml.Node) string {
	nodeTree := ghth5.NewTree(n)
	
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