package parse

import (
	"fmt"
	"io"

	"github.com/dkvka/gopkg/html"
	xhtml "golang.org/x/net/html"
)

// TODO: Create this with only stdlib

// parse.Links takes in an html document and returns a slice
// of links parsed from it.
func Links(r io.Reader) ([]html.Link, error) {
	doc, err := xhtml.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)

	var links []html.Link

	for _, node := range nodes {
		links = append(links, makeLink(node))
		fmt.Println(node)
	}

	return links, nil
}

func linkNodes(n *xhtml.Node) (nodes []*xhtml.Node) {
	if n.Type == xhtml.ElementNode && n.Data == "a" {
		return []*xhtml.Node{n}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}

	return
}

func makeLink(n *xhtml.Node) (link html.Link) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = "TODO: Parse text"
	return
}
