/**
*
最小化编译
采用：go build -ldflags "-s -w" 这种方式编译。

解释一下参数的意思：

-ldflags： 表示将后面的参数传给连接器（5/6/8l）
-s：去掉符号信息
-w：去掉DWARF调试信息
注意：
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	// "reflect"
	"regexp"
	"strings"
)

func read(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	var sList string
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			return sList
		}
		if 0 == n {
			break
		}
		sList = string(buf[:n])
	}
	return sList
}

func getRequest() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.topcashback.co.uk/join", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:27.0) Gecko/20100101 Firefox/27.0")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("网络链接失败。。")
	} else {
		if resp.StatusCode == 200 {
			// req, _ := http.NewRequest("GET", "https://www.topcashback.co.uk/join", nil)
			// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:27.0) Gecko/20100101 Firefox/27.0")
			// for i := range resp.Cookies() {
			// 	req.AddCookie(resp.Cookies()[i])
			// }

			// client := new(http.Client)
			// resp, err := client.Do(req)
			// if err != nil {
			// 	fmt.Println("获取表单信息失败!")
			// } else {
			// 	if resp.StatusCode == 200 {
			defer resp.Body.Close()
			restr, _ := ioutil.ReadAll(resp.Body)
			result := string(restr)
			str := regexp.MustCompile(`name="__VIEWSTATE"\s.*value="(.*)"`)
			fmt.Println(str.FindStringSubmatch(result)[1])
			// }
			// }

		}
	}
}

func main() {
	ipList := strings.Split(read("t.txt"), "\n")
	fmt.Println(ipList[0])

}
