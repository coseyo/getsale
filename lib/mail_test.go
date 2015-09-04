package lib

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Test_send(t *testing.T) {

	mail := &Mail{}
	mail.Dial("smtp.ym.163.com", 25, "no-reply@ili.li", "123456aa")
	mail.SetSender("no-reply<no-reply@ili.li>")
	mail.SetReceiver("2268452603@qq.com,396733179@qq.com")
	err := mail.Send("2015-08-27", "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\"><html xmlns=\"http://www.w3.org/1999/xhtml\"><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf8\" /><head><style>tr {font-size:12px;}td {text-align:center; border-right: 1px solid #C1DAD7; border-bottom: 1px solid #C1DAD7; background: #fff; padding: 6px 6px 6px 12px;color: #333;}</style></head><body><table cellspacing=\"0\" cellpadding=\"0\" width=\"100%\" align=\"center\" border=\"0\"><tr><td>保利（珠海）房地产开发有限公司</td><td> </td><td>1803/74969.15</td><td>1263/104502.33</td></tr><tr><td>珠海市广兴昌房地产开发有限公司</td><td>锦绣四季花园</td><td>391/39031.59</td><td>908/42814.37</td></tr><tr><td>珠海市斗门区世荣实业有限公司</td><td> </td><td>698/35260.66</td><td>903/76010.39</td></tr><tr><td>珠海华亿投资有限公司</td><td>华发水岸花园</td><td>130/11049.88</td><td>740/63679.37</td></tr><tr><td>珠海市同裕房地产开发有限公司</td><td>云顶澜山花园</td><td>666/40429.8</td><td>706/82706.89</td></tr><tr><td>保利（珠海）房地产开发有限公司</td><td>保利香槟国际花园</td><td>0/0</td><td>637/59187.21</td></tr><tr><td>珠海华郡房产开发有限公司</td><td>华发水郡花园三期A区二标段</td><td>25/2486.56</td><td>627/60701.95</td></tr><tr><td>珠海市财信投资有限公司</td><td>诚丰水晶座公寓</td><td>131/9478.24</td><td>585/29495.84</td></tr><tr><td>珠海市金地房地产开发有限公司</td><td> </td><td>102/21178.11</td><td>583/61542.61</td></tr><tr><td>珠海华茂房地产投资顾问有限公司</td><td>华发四季名苑</td><td>32/7006.36</td><td>563/77928.29</td></tr><tr><td>珠海市佳誉房地产开发有限公司</td><td>时代成花园</td><td>59/5020.65</td><td>540/52642.68</td></tr><tr><td>珠海中航通用房地产开发有限公司</td><td>中航花园</td><td>315/30319.24</td><td>526/70764.7</td></tr><tr><td>珠海华茂房地产投资顾问有限公司</td><td>华发四季名苑</td><td>16/4416.79</td><td>503/84467.94</td></tr><tr><td>珠海华福商贸发展有限公司</td><td>华发新城六期</td><td>44/3072.87</td><td>486/30256.65</td></tr><tr><td>保利（珠海）房地产开发有限公司</td><td>保利香槟国际花园</td><td>2/258.29</td><td>485/45241.53</td></tr></table></body></html>")

	expect(t, err, nil)
}

//func Test_rawsend(t *testing.T) {
//	m := gomail.NewMessage()
//	m.SetHeader("From", "no-reply<no-reply@ili.li>")
//	m.SetHeader("To", "2268452603@qq.com", "2268452603@qq.com,396733179@qq.com")
//	m.SetHeader("Subject", "Hello!")
//	m.SetBody("text/plain", "Hello!")
//	d := gomail.Dialer{Host: "smtp.ym.163.com", Port: 25, Auth: smtp.PlainAuth("", "no-reply@ili.li", "123456aa", "smtp.ym.163.com")}
//	if err := d.DialAndSend(m); err != nil {
//		panic(err)
//	}
//}
