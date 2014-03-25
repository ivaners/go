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
	// "strings"
	"database/sql"
	"github.com/hailiang/socks"
	_ "github.com/mattn/go-sqlite3"
)

func read(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开出错，可能找不到文件", path)
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

func setProxy() {
	dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "61.136.68.76:1080")
	tr := &http.Transport{Dial: dialSocksProxy}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", "http://ip.cn/", nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("代理错误...")
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
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

func sqlite() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// ipList := strings.Split(read("t.txt"), "\n")
	// fmt.Println(ipList[0])
	// txt := read("t.txt")

	// if txt == "" {
	// 	return
	// }
	// fmt.Println("...")
	// sqlite()
	setProxy()
}
