package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

var b *BPMClient
var id string

func main() {

	id = login()
L:
	for {

		var msg string
		fmt.Print("請輸入指令(1:啟單,2:審核任務,3:查看待執行任務,4:登出,5:退出):")

		fmt.Scan(&msg)
		switch msg {
		case "1": //啟單
			getProcessId()
			var pId string
			fmt.Print("請輸入PorcessId:")
			fmt.Scan(&pId)
			SelectProcess(pId)
		case "2": //審核任務
			var taskId string
			var k string
			var v string
			fmt.Print("請輸入要審核任務的Id:")
			fmt.Scan(&taskId)

			fmt.Println("請提供審核內容:")
			fmt.Scan(&k)
			fmt.Scan(&v)

			runTask(taskId, content(k, v))

		case "3": //查看待執行任務
			getTodoTaskTest(id)

		case "4": //登出
			b.GetLogoutService()
			id = login()

		case "5": //退出
			break L

		}
	}

}

//將key中所有的value存入一個切片中
func getWhatValue(resp []byte, key string) []string {

	var s1 interface{}
	err := json.Unmarshal(resp, &s1)
	if err != nil {
		log.Fatal(err)
	}
	objArr, ok := s1.([]interface{})
	value_list := make([]string, len(objArr))
	if !ok {
		log.Fatal("expected an array of objects")
	}
	for i, obj := range objArr {
		obj, ok := obj.(map[string]interface{})
		if !ok {
			log.Fatalf("expected type map[string]interface{}, got %s", reflect.TypeOf(objArr[i]))
		}
		value_list[i] = fmt.Sprintf("%v", obj[key])
	}
	return value_list
}

//判斷任務執行結果
func runTask(taskId string, content string) {

	t := b.ExecuteTask(taskId, content)
	if t == 204 {
		fmt.Println("執行成功!")
	} else {
		fmt.Println("失敗")
	}

}

//顯示可以開啟的任務流程
func getProcessId() {

	d := []string{"name", "id", "activationState"}
	fmt.Println("請選擇啟單流程:")
	//取得流程
	p := b.GetProcessInstanceId()
	v := RebuildJsonTest(p, d)
	if len(p) == 2 {
		fmt.Println("沒有待開啟流程")
	} else {
		tableView(v, d)

	}

}

//檢查輸入的pId是否正確，並執行
func SelectProcess(pId string) {
	var c []byte
	p := b.GetProcessInstanceId()
	v := getWhatValue(p, "id")
	for i := 0; i < len(v); i++ {
		if pId == v[i] {
			c = b.CreateProcessCaseTest(pId)
			fmt.Println("啟單成功！")
			break
		}
	}
	if c == nil {
		fmt.Println("輸入錯誤.")
	}

}

//init
func login() string {
	var username string
	fmt.Print("請輸入Username:")
	fmt.Scan(&username)
	b = New(username)
	return b.GetUserId()
}

func content(key string, value string) string {
	m := make(map[string]string)
	m[key] = value

	c, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
	}
	return string(c)
}

func (b *BPMClient) GetUserId() string {

	uri := b.apiUri + "system/session/unusedId"

	//log.Println("GetUserId() -uri", uri)
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("GetUserId() - b.request:\n %+v", b.request)
	log.Println("GetUserId() - Status Code:", resp.StatusCode())

	s2 := Rebuild(string(resp.Body()), "user_id")

	return s2

}

func Rebuild(jsonBody string, key ...string) string {
	var s1 map[string]string

	err := json.Unmarshal([]byte(jsonBody), &s1)
	if err != nil {
		fmt.Println("失敗", err)
		os.Exit(3)
	}
	s2 := make(map[string]string)
	for _, keyy := range key {
		s2[keyy] = s1[keyy]
	}

	value := s2["user_id"]
	return value
}

func getTodoTaskTest(userId string) {

	todo := []string{"name", "caseId", "assigned_date", "id"}

	fmt.Println("以下為待執行任務:")
	//取得任務
	p := b.GetStateCaseList("100", "ready", userId)
	v := RebuildJsonTest(p, todo)
	if len(p) == 2 {
		fmt.Println("沒有待執行任務")
	} else {
		tableView(v, todo)
	}
}

func tableView(v [][]string, title []string) {

	//初始化tablewriter
	table := tablewriter.NewWriter(os.Stdout)

	//上面的data为表格内容，还需要定義表格頭部
	table.SetHeader(title)

	//将数據添加到table
	table.AppendBulk(v)

	//输出
	table.Render()

}
