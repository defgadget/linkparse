package main

import (
  _ "log"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)
type links struct {
	links map[string]string
}
func getATagAndText(n *html.Node, l *links) {
	attrText := ""
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				attrText = a.Val
				break
			}
		}
		for nc := n.FirstChild; nc != nil; nc = nc.NextSibling {
			if nc.Type == html.TextNode {
				l.links[attrText] = strings.TrimSpace(nc.Data)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getATagAndText(c, l)
	}
}
func main() {
	l := new(links)
	l.links = make(map[string]string)
	r, err := os.OpenFile("link/ex4.html", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer r.Close()
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("Error parsing html", err)
	}
	getATagAndText(doc, l)
	for k, v := range l.links {
		fmt.Println(k, "-", v)
	}
}