package main
import (
  "fmt"

  //"log"
  "github.com/gocolly/colly"
  //"github.com/gocolly/colly/proxy"
)

const LINK = "https://protonmail.com/blog/"

func main() {
  scrapPage(LINK)
}

func scrapPage(LINK string) {
	// Instantiate default collector
	c := colly.NewCollector(
    colly.AllowedDomains("protonmail.com"),
	)
//   if p, err := proxy.RoundRobinProxySwitcher(
//     "http://35.182.149.39:3129",
//     "https://35.182.149.39:3129",
//     "socks://88.198.50.103:1080",
//   ); err == nil {
//   c.SetProxyFunc(p)
// }

  //"http://35.182.149.39:3129", "https://35.182.149.39:3129", "socks://88.198.50.103:1080"
	c.OnHTML("h3", func(e *colly.HTMLElement) {

		println(e.Text)

	})


	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
  c.Visit(LINK)

}
