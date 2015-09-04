package models

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"getsale/lib"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
)

type Object struct {
	count        int
	args         string
	TargetUrl    string
	PageLimit    int
	MailTo       string
	MailFrom     string
	MailHost     string
	MailUser     string
	MailPassword string
	MailPort     int
}

func (this *Object) Go() {

	beelog := logs.NewLogger(100)
	beelog.SetLogger("file", fmt.Sprintf("{\"filename\":\"logs/%s.log\"}", time.Now().Format("2006-01-02")))
	beelog.Info("start")

	html, err := this.getHtml()
	if err != nil {
		beelog.Error("getHtml", err)
		return
	}

	i := 1

	for {
		beelog.Info(fmt.Sprintf("get page %d", i))
		if this.PageLimit != 0 && i >= this.PageLimit {
			beelog.Info(fmt.Sprintf("page maxout, page:%d", i))
			break
		}
		phtml, err := this.postHtml()
		if err != nil {
			beelog.Error("postHtml", err)
			break
		}
		i++
		html = html + phtml
	}

	html = this.formatHtml(html)

	err = this.sendMail(html)
	if err != nil {
		beelog.Error("sendMail", err)
	}
	beelog.Info("end")
	return
}

func (this *Object) setPageLimit(doc *goquery.Document) error {
	if this.PageLimit > 0 {
		return nil
	}
	pageHtml := doc.Find("#ctl00_ContentPlaceHolder1_pageInfo").Text()
	// <span id="ctl00_ContentPlaceHolder1_pageInfo">第1页/共30页,共436条记录，每页15条</span>

	start := strings.Index(pageHtml, "共") + 3
	pos := strings.Index(pageHtml[start:], "页")
	s := pageHtml[start:][:pos]
	n, err := strconv.ParseInt(s, 10, 64)
	this.PageLimit = int(n)
	return err

}

func (this *Object) getHtml() (string, error) {

	res, err := request("GET", this.TargetUrl, "")
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("http status not 200, it is %d", res.StatusCode))
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	html, err := this.parse(doc)
	if err != nil {
		return "", err
	}

	this.setArgs(doc)
	if err := this.setPageLimit(doc); err != nil {
		return "", err
	}

	return html, nil
}

func (this *Object) postHtml() (string, error) {
	res, err := request("POST", this.TargetUrl, this.args)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return "", err
	}
	html, err := this.parse(doc)
	if err != nil {
		return "", err
	}

	this.setArgs(doc)

	return html, nil
}

func (this *Object) parse(doc *goquery.Document) (string, error) {
	var html string
	doc.Find(".ListRow1,.ListRow2").Each(func(i int, s *goquery.Selection) {
		html += "<tr>"
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			html += "<td>" + s.Text() + "</td>"
		})
		html += "</tr>"
	})

	return html, nil
}

func (this *Object) setArgs(doc *goquery.Document) {
	v, _ := doc.Find("#__VIEWSTATE").Attr("value")
	e, _ := doc.Find("#__EVENTVALIDATION").Attr("value")
	g, _ := doc.Find("#__VIEWSTATEGENERATOR").Attr("value")
	v = url.QueryEscape(v)
	e = url.QueryEscape(e)
	g = url.QueryEscape(g)
	this.args = "__EVENTTARGET=ctl00%24ContentPlaceHolder1%24lnkBtnNext&__EVENTARGUMENT=&__VIEWSTATE=" + v + "&__EVENTVALIDATION=" + e + "&__VIEWSTATEGENERATOR=" + g
}

func (this *Object) formatHtml(html string) string {
	newHtml := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">"
	newHtml += "<html xmlns=\"http://www.w3.org/1999/xhtml\">"
	newHtml += "<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf8\" />"
	newHtml += "<head><style>"
	newHtml += "tr {font-size:12px;}"
	newHtml += "td {text-align:center; border-right: 1px solid #C1DAD7; border-bottom: 1px solid #C1DAD7; background: #fff; padding: 6px 6px 6px 12px;color: #333;}"
	newHtml += "</style></head>"
	newHtml += "<body>"
	newHtml += "<table cellspacing=\"0\" cellpadding=\"0\" width=\"100%\" align=\"center\" border=\"0\">" + html + "</table></body></html>"
	return newHtml
}

func (this *Object) sendMail(html string) error {
	mail := &lib.Mail{}
	mail.Dial(this.MailHost, this.MailPort, this.MailUser, this.MailPassword)
	mail.SetSender(this.MailFrom)
	mail.SetReceiver(this.MailTo)
	subject := time.Now().Format("2006-01-02")
	return mail.Send(subject, html)
}

func request(method string, url string, params string) (res *http.Response, err error) {
	client := &http.Client{}
	var req *http.Request
	if method == "GET" {
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return
		}
	} else {
		body := bytes.NewBuffer([]byte(params))
		req, err = http.NewRequest("POST", url, body)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Add("Referer", url)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.155 Safari/537.36")
	res, err = client.Do(req)
	return
}
