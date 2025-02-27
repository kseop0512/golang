package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=49545109&recommend_ids=eJxNj8sRw0AIQ6vJHRDic04h7r%2BL2N7Mssc3Qk%2BDN9kpeZXqJ7%2FetMqWq8UWkmmxUw%2BI6o2x0kaWbQwkFLsbqO4%2BsCmPSlY3xCNGVSzUpFrw2SXEk4PtDtt4D4XIpCX64l9lhB4qqmsdxynQfXyrCJ9%2FvYxv9weTbkB2&view_type=search&searchword=python&searchType=search&gz=1&t_ref_content=generic&t_ref=search&relayNonce=cdf55a40c42d3da27b2c&paid_fl=n&search_uuid=6e59ff77-485b-4c89-b1e7-dcd17dc2df48&immediately_apply_layer_open=n#seq=0

type extractedJob struct {
	id       string
	title    string
	location string
}

// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

// https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=python&recruitPage=1&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n
// https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=python&recruitPage=2&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term

	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 1; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}
	// fmt.Println(jobs)
	for i := 1; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)

	pageURL := url + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		// jobs = append(jobs, job)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".area_job>h2.job_tit>a>span").Text())
	location := CleanString(card.Find(".job_condition>span:first-child").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
	}

}

// CleanString cleans a string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	/*
		req, rErr := http.NewRequest("GET", baseURL, nil)

		checkErr(rErr)

		// 프록시로 호출
		purl, err := url.Parse(baseURL)
		client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(purl)}}
		res, err := client.Do(req)
	*/
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	// fmt.Println(doc)
	return pages
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "ID", "Title", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		// id는 링크형태로 변경
		// https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.id, job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
