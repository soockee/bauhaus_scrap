package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.bauhaus.info/sitemap.xml", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			fmt.Println(string(r.Body))
			// sitemapLinsk := parseSiteMap(string(r.Body))
			// fmt.Println("")
			// for i, sl := range sitemapLinsk {
			// 	link := parseSiteMapLink(sl)
			// 	fmt.Printf("#%d: %v\n", i, link)
			// }
		},
		BrowserEndpoint: "ws://localhost:3000",
	}).Start()
}

func parseSiteMap(s string) []string {
	links := []string{}

	doc, err := xmlquery.Parse(strings.NewReader(s))

	if err != nil {
		panic(err)
	}

	for i, n := range xmlquery.Find(doc, "//sitemapindex/sitemap/loc") {
		fmt.Printf("#%d %s\n", i, n.InnerText())
		links = append(links, n.InnerText())
	}
	return links
}

func parseSiteMapLink(link string) []string {
	links := []string{}
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(link, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			// for each string that is separated by space
			links = strings.Split(string(r.Body), " ")
		},
		BrowserEndpoint: "ws://localhost:3000",
	}).Start()

	return links
}
