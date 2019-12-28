package main

import (
  _ "log"
	"fmt"
	"io"
	"golang.org/x/net/html"
	"os"
	"strings"
)
// Anchor is a object that holds an Anchor tag's link, as well as the text
type Anchor struct {
	Href string
	Text string
}
func anchorNode(n *html.Node) Anchor {
	attrText := ""
	linktext := ""
	for _, a := range n.Attr {
		if a.Key == "href" {
			attrText = a.Val
			break
		}
	}
	for nc := n.FirstChild; nc != nil; nc = nc.NextSibling {
		if nc.Type == html.TextNode {
			linktext = linktext + nc.Data
		}
	}
	return Anchor{Href: attrText, Text: strings.TrimSpace(linktext)}
}
func getATagAndText(n *html.Node, a []Anchor) []Anchor {
	if n.Type == html.ElementNode && n.Data == "a" {
		a = append(a, anchorNode(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		a = getATagAndText(c, a)
	}
	return a
}
// ParseAnchors takes an io.Reader and parses for Anchor tags, and text.
// it then returns a list of Anchor instances containing Href, and Text.
func ParseAnchors(r io.Reader) ([]Anchor, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("Error parsing html", err)
		return nil, err
	}
	anchors := getATagAndText(doc, make([]Anchor, 0))
	return anchors, nil
}
func run () error {
	r, err := os.OpenFile("link/ex4.html", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return err
	}
	defer r.Close()
	anchors, err := ParseAnchors(r)
	if err != nil {
		fmt.Println("There was an error parsing HTML", err)
		return err
	}
	for _, v := range anchors {
		fmt.Println(v.Href, "-", v.Text)
	}
	return nil
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		return
	}
}