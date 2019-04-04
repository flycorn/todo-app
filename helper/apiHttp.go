package helper

import (
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"reflect"
	"sort"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

//HTTP请求工具
type ApiHttp struct {
	StatusCode int
	Header http.Header
	Body string
}

/**
	GET请求
	params:
		apiUrl 接口地址
		arg
          第一个参数 请求的参数键值对【map[string]interface{}】
          第二个参数 请求返回的json转换结构体
 */
func (h *ApiHttp) Get(apiUrl string, arg ...interface{}) error{
	var resModel interface{}

	l := len(arg)
	if l > 0 {
		if reflect.TypeOf(arg[0]).Kind() == reflect.Map {
			param := arg[0].(map[string]interface{})
			for k, v := range param{
				str := v.(string)
				if strings.Contains(apiUrl, "?") {
					//包含
					apiUrl += "&"
				} else {
					//不包含
					apiUrl += "?"
				}
				apiUrl += k+"="+str
			}
		}
		if l > 1 {
			resModel = arg[1]
		}
	}

	//发送请求
	resp, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //断开请求

	//请求结果
	h.StatusCode = resp.StatusCode
	h.Header = resp.Header
	if err != nil {
		return err
	}

	//读取接口返回的数据
	resData, _ := ioutil.ReadAll(resp.Body)
	h.Body = string(resData)

	if l > 1 {
		//json转结构体数据
		err = json.Unmarshal(resData, &resModel)
	}
	return err
}

/**
	POST请求
	params:
		apiUrl 接口地址
		arg
          第一个参数 请求的参数键值对【map[string]interface{}】
          第二个参数 请求返回的json转换结构体
 */
func (h *ApiHttp) POST(apiUrl string, arg ...interface{}) error{
	var resModel interface{}
	var r http.Request

	l := len(arg)
	if l > 0 {
		r.ParseForm()
		if reflect.TypeOf(arg[0]).Kind() == reflect.Map {
			param := arg[0].(map[string]interface{})
			for k, v := range param{
				str := v.(string)
				r.Form.Add(k, str)
			}
		}
		if l > 1 {
			resModel = arg[1]
		}
	}

	bodystr := strings.TrimSpace(r.Form.Encode())
	request, err := http.NewRequest("POST", apiUrl, strings.NewReader(bodystr))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")

	//发送请求
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	defer resp.Body.Close() //断开请求
	//请求结果
	h.StatusCode = resp.StatusCode
	h.Header = resp.Header
	if err != nil {
		return err
	}

	//读取接口返回的数据
	resData, _ := ioutil.ReadAll(resp.Body)
	h.Body = string(resData)

	if l > 1 {
		//json转结构体数据
		err = json.Unmarshal(resData, &resModel)
	}
	return err
}

/**
   获取签名
 */
func (h *ApiHttp) GetSign(params interface{}, key string, arg ...interface{}) string{
	var sign string
	l := len(arg)
	if(l > 0){
		t := arg[0].(int)
		//其它生成签名规则
		switch t {
		case 2:
			sign = signType2(params, key)
		default:
			sign = signType1(params, key)
		}
	} else {
		sign = signType1(params, key)
	}
	return sign
}

//签名规则1【默认】
func signType1(paramData interface{}, key string) string{
	var sign string

	var keys []string
	var params map[string]interface{}

	if reflect.TypeOf(paramData).Kind() == reflect.Map {
		params = paramData.(map[string]interface{})
		for k, _ := range params{
			keys = append(keys, k)
		}
		sort.Strings(keys)
	}

	builder := strings.Builder{}
	l := len(keys)
	for k, v := range keys {
		builder.WriteString(v)
		builder.WriteString("=")
		builder.WriteString(fmt.Sprint(params[v]))
		if l < k+1{
			builder.WriteString("&")
		}
	}
	builder.WriteString(key)

	h := md5.New()
	h.Write([]byte(builder.String()))
	sign = hex.EncodeToString(h.Sum(nil))
	return sign
}

//签名规则2
func signType2(paramData interface{}, key string) string{
	var sign string

	var keys []string
	var params map[string]interface{}

	if reflect.TypeOf(paramData).Kind() == reflect.Map {
		params = paramData.(map[string]interface{})
		for k, _ := range params{
			keys = append(keys, k)
		}
		sort.Strings(keys)
	}

	builder := strings.Builder{}
	l := len(keys)
	for k, v := range keys {
		builder.WriteString(v)
		builder.WriteString("=")
		builder.WriteString(fmt.Sprint(params[v]))
		if l < k+1{
			builder.WriteString("&")
		}
	}
	builder.WriteString("&key="+key)

	h := md5.New()
	h.Write([]byte(builder.String()))
	sign = hex.EncodeToString(h.Sum(nil))
	return sign
}


