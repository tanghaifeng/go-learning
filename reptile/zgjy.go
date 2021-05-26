package reptile

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// * 也算不上正在意义上的爬虫

// 各个省份教师招聘 2019-2021的数据
// 每个省份数据写到txt文件中，并用&隔开

// 页面规则  http://rili.offcn.com/jszp/s25-2021/2.html
func ExampleScrape() {
	// 省份
	for k := 1; k < 32; k++ {
		// Find the review items
		content := ""
		k := strconv.Itoa(k)
		for y := 2019; y < 2022; y++ {
			yy := strconv.Itoa(y)
			for i := 1; i < 100; i++ { // 默认100页
				p := strconv.Itoa(i)

				// http://rili.offcn.com/jszp/s25-2021/2.html  页面规则
				res, err := http.Get("http://rili.offcn.com/jszp/s" + k + "-" + yy + "/" + p + ".html")
				if err != nil {
					log.Fatal(err)
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
				}

				// Load the HTML document
				doc, err := goquery.NewDocumentFromReader(res.Body)

				// 获取页数
				pages := doc.Find(".oPubPages font")

				fmt.Println(pages.Text())

				page := pages.Text()
				if page == "" {
					break
				}
				re := regexp.MustCompile("[0-9]+")
				fmt.Println(re.FindAllString(page, -1))
				pageArr := re.FindAllString(page, -1)[0]
				fmt.Println(pageArr)
				pageSize, err := strconv.Atoi(pageArr)
				if pageSize < i {
					break
				}

				if err != nil {
					log.Fatal(err)
				}
				doc.Find(".ksrl_rlxq_tab tr").Each(func(i int, s *goquery.Selection) {

					t1 := s.Find("td:nth-child(1)").Text()

					t2 := s.Find("td:nth-child(2)").Text()

					t3 := s.Find("td:nth-child(3)").Text()

					t4 := s.Find("td:nth-child(4)").Text()

					t5 := s.Find("td:nth-child(5)").Text()

					t6 := s.Find("td:nth-child(6)").Text()

					t7 := s.Find("td:nth-child(7)").Text()

					t8, _ := s.Find("a").Attr("href")
					if t1 != "" {
						content += t1 + "&" + t2 + "&" + t3 + "&" + t4 + "&" + t5 + "&" + t6 + "&" + t7 + "&" + t8 + "\n"
						fmt.Println(content)
					}
				})
			}
		}
		fmt.Println(content)
		ioutil.WriteFile(k+".txt", []byte(content), 0644)
	}
}
