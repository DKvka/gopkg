package html

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link in an html document
// <a href="...">Text</a>
type Link struct {
	Href string
	Text string
}

// TODO: Create this with only stdlib

// ParseLinks takes in an html document and returns a slice
// of links parsed from it.
func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)

	var links []Link

	for _, node := range nodes {
		links = append(links, makeLink(node))
		fmt.Println(node)
	}

	return links, nil
}

func linkNodes(n *html.Node) (nodes []*html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
w
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}

	return
}

func makeLink(n *html.Node) (link Link) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = "TODO: Parse text"
	return
}
