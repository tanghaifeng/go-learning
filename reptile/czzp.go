package reptile

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
)

// 常州招聘
// 页面规则  https://www.zhipin.com/c101191100/b_%E6%96%B0%E5%8C%97%E5%8C%BA/?query=Java&page=2&ka=page-2
func ZhangzhouZhaoPing() {
	// 省份
	content := ""
	for i := 1; i < 100; i++ { // 默认100页
		//p := strconv.Itoa(i)

		// http://rili.offcn.com/jszp/s25-2021/2.html  页面规则
		res, err := http.Get("http://www.cfang.com/loupan/all")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		//
		//// 获取页数
		//pages := doc.Find(".oPubPages font")
		//
		//fmt.Println(pages.Text())
		//
		//page := pages.Text()
		//if page == "" {
		//	break
		//}
		//re := regexp.MustCompile("[0-9]+")
		//fmt.Println(re.FindAllString(page, -1))
		//pageArr := re.FindAllString(page, -1)[0]
		//fmt.Println(pageArr)
		//pageSize, err := strconv.Atoi(pageArr)
		//if pageSize < i {
		//	break
		//}

		if err != nil {
			log.Fatal(err)
		}
		doc.Find(".result-item flex").Each(func(i int, s *goquery.Selection) {

			t1 := s.Find(".title").Text()

			//t2 := s.Find(".job-area").Text()
			//
			//t3 := s.Find(".job-pub-time").Text()
			//
			//t4 := s.Find(".company-text h3 a").Text()
			//
			//t5 := s.Find(".company-text p em ").Text()

			//t6 := s.Find("td:nth-child(6)").Text()
			//
			//t7 := s.Find("td:nth-child(7)").Text()
			//
			//t8, _ := s.Find("a").Attr("href")
			if t1 != "" {
				content += t1 + "\n"
				fmt.Println(content)
			}
		})
	}
	fmt.Println(content)
	ioutil.WriteFile("java 新北"+".txt", []byte(content), 0644)

}
