package recursion

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"os"
	"strings"

	git "../github"
)

const htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
    body content
    <p>more content</p>
	<a href='2345'>
</body>
</html>`

func Start() {
	resp, _ := git.GetRequest("https://career.habr.com/")
	byte, _ := ioutil.ReadAll(resp.Body)
	doc, err := html.Parse(strings.NewReader(string(byte)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinksl: %v\n", err)
		os.Exit(1)
	}
	visite := visit(nil, doc)
	for _, link := range visite {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for n := n.FirstChild; n != nil; n = n.NextSibling {
		links = visit(links, n)
	}
	return links
}
