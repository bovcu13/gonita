package main

import (
	"encoding/json"
	"log"
)

//  GetProcessInstanceId
//  @Description: 取得流程ID
//  @receiver b
//  @return []byte
func (b *BPMClient) GetProcessInstanceId() []byte {
	uri := b.apiUri + "bpm/process?c=100"
	// log.Println("GetProcessInstanceId()- uri", uri)

	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("GetProcessInstanceId() - b.request:\n %+v", b.request)
	// log.Println("GetProcessInstanceId() - Status Code:", resp.StatusCode())
	// log.Printf("GetProcessInstanceId() - resp:\n %+v", string(resp.Body()))

	return resp.Body()
}

// Create a new caseID(process instance)
// ref: https://documentation.bonitasoft.com/bonita/2021.2/api/bpm-api#start-a-process-using-an-instantiation-contract
//
//  CreateProcessCase
//  @Description: Create a new caseID(process instance), Start a new form
//  @receiver b
//  @param processId [表單ID]
//  @param jsonBody 只需要提供"內層"結構(轉成string)
//  @return caseId
//
func (b *BPMClient) CreateProcessCase(processId string, jsonBody ...string) (caseId []byte) {
	// TODO: 傳入改為 interface{} ?
	var s string
	if len(jsonBody) != 0 {
		s = ModelInputStringToRawJson(jsonBody[0])
	} else {
		s = "{}"
	}
	log.Println("CreateProcessCase() - StringToRawJson(): ", s)

	uri := b.apiUri + "bpm/process/" + processId + "/instantiation"
	log.Println("CreateProcessCase()- uri", uri)

	resp, err := b.request.SetBody(s).Post(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CreateProcessCase() - body: ", s)
	log.Printf("CreateProcessCase() - b.request:\n %+v", b.request)
	log.Println("CreateProcessCase() - Status Code:", resp.StatusCode())
	log.Printf("CreateProcessCase() - resp:\n %+v", resp.Body())

	// TODO: 是否只回需回傳 ID?
	return resp.Body() //caseId
}

func (b *BPMClient) CreateProcessCaseTest(processId string) []byte {

	uri := b.apiUri + "bpm/process/" + processId + "/instantiation"
	log.Println("CreateProcessCase()- uri", uri)

	resp, err := b.request.Post(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("CreateProcessCase() - Status Code:", resp.StatusCode())

	return resp.Body() //caseId
}

func ModelInputStringToRawJson(s string) string {
	s1 := []byte(s)
	var s2 *interface{}

	err := json.Unmarshal(s1, &s2)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("StringToRawJson() - json.Unmarshal(s1, &s2) %+v: ", s2)

	s3 := &FormInput{
		ModelInput: s2,
	}

	s4, err := json.Marshal(s3)
	if err != nil {
		log.Fatal(err)
	}

	return string(s4)
}

//  GetStateCaseList
//  @Description: 按簽核狀態取回人員有關的表單
//  @receiver b
//  @param rows 顯示資料量
//  @param state  |ready|
//  @param userId
//  @return string
func (b *BPMClient) GetStateCaseList(rows string, state string, userId string) []byte {

	uri := b.apiUri + "bpm/humanTask?c=" + rows + "&f=state=" + state + "&f=user_id=" + userId
	//log.Println("GetStateCaseList()- uri", uri)

	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("GetStateCaseList() - b.request:\n %+v", b.request)
	log.Println("GetStateCaseList() - Status Code:", resp.StatusCode())
	//log.Printf("CreateProcessCase() - resp:\n %+v", string(resp.Body()))

	return resp.Body()
}

//ExecuteTask
//@Description:審核任務
//@receiver b
//@param taskId
//@param jsonBody 只需要提供"內層"結構(轉成string)
//@return ResponseStatusCode (204為成功)
func (b *BPMClient) ExecuteTask(taskId string, jsonBody ...string) int {

	var s string
	if len(jsonBody) != 0 {
		s = ModelInputStringToRawJson(jsonBody[0])
	} else {
		s = "{}"
	}

	//log.Println("ExecuteTask() - StringToRawJson(): ", s)

	uri := b.apiUri + "bpm/userTask/" + taskId + "/execution?assign=true"
	//log.Println("ExecuteTask()- uri", uri)

	resp, err := b.request.SetBody(s).Post(uri)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("ExecuteTask() - body: ", s)
	//log.Printf("ExecuteTask() - b.request:\n %+v", b.request)
	log.Println("ExecuteTask() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//  GetStateCaseList
//  @Description: 顯示該單待執行任務詳細資料
//  @receiver b
//  @param caseId
//  @return []byte
func (b *BPMClient) GetCasePendingTaskDetail(caseId string) []byte {

	uri := b.apiUri + "bpm/humanTask?f=caseId=" + caseId
	log.Println("GetCasePendingTaskDetail() -uri", uri)
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetCasePendingTaskDetail() - b.request:\n %+v", b.request)
	log.Println("GetCasePendingTaskDetail() - Status Code:", resp.StatusCode())
	return resp.Body()
}

// GetCaseArchivedTaskDetail
// @Description: 顯示該單已完成任務詳細資料
// @receiver b
// @parm caseId
// @return string
func (b *BPMClient) GetCaseArchivedTaskDetail(caseId string) []byte {

	uri := b.apiUri + "bpm/archivedTask?f=caseId=" + caseId
	log.Println("GetCaseArchivedTaskDetail() -uri", uri)
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetCaseArchivedTaskDetail() - b.request:\n %+v", b.request)
	log.Println("GetCaseArchivedTaskDetail() - Status Code:", resp.StatusCode())
	return resp.Body()
}

// GetCaseArchivedTaskDetail
// @Description: 顯示該任務完成後之詳細資料
// @receiver b
// @parm caseId
// @return string
func (b *BPMClient) GetArchivedTaskDetail(sourceObjectId string) []byte {

	uri := b.apiUri + "bpm/archivedHumanTask?f=sourceObjectId=" + sourceObjectId
	log.Println("GetArchivedTaskDetail() -uri", uri)
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetArchivedTaskDetail() - b.request:\n %+v", b.request)
	log.Println("GetArchivedTaskDetail() - Status Code:", resp.StatusCode())
	return resp.Body()
}

//GetProcessAllCaseList
// @Desctiption: 取得該流程所有單況
// @receiver b
// @parm caseId
// @return string
func (b *BPMClient) GetProcessAllCaseList(rows string, processId string) []byte {

	uri := b.apiUri + "bpm/case?c=" + rows + "&f=processDefinitionId=" + processId
	log.Println("GetCaseArchivedTaskDetail() -uri", uri)
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetCaseArchivedTaskDetail() - b.request:\n %+v", b.request)
	log.Println("GetCaseArchivedTaskDetail() - Status Code:", resp.StatusCode())
	return resp.Body()
}

// GetUserId
// @Description: 搜尋人員的userID
// @receiver b
// @parm caseId
// @return string
