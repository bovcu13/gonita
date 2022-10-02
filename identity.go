package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

//AddGroup
//  @Description: 新增部門
//  @receiver b
//  @param bodyInput
//  @return []byte
func (b *BPMClient) AddGroup(bodyInput string) []byte {

	uri := b.apiUri + "identity/group"
	log.Println("AddGroup() -uri", uri)
	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("AddGroup() - body: ", bodyInput)
	log.Printf("AddGroup() - b.request:\n %+v", b.request)
	log.Println("AddGroup() - Status Code:", resp.StatusCode())
	return resp.Body()
}

//EditGroup
//  @Description: 更新部門
//  @receiver b
//  @param bodyInput
//  @param groupId
//  @return int (200為成功)
func (b *BPMClient) EditGroup(bodyInput string, groupId string) int {

	//s := StringToRawJson(bodyInput)
	uri := b.apiUri + "identity/group/" + groupId
	log.Println("EditGroup() -uri", uri)
	resp, err := b.request.SetBody(bodyInput).Put(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("EditGroup() - body: ", bodyInput)
	log.Printf("EditGroup() - b.request:\n %+v", b.request)
	log.Println("EditGroup() - Status Code:", resp.StatusCode())
	return resp.StatusCode()
}

//DeleteGroup
//  @Description: 刪除部門
//  @receiver b
//  @param groupId
//  @return int (200為成功)
func (b *BPMClient) DeleteGroup(groupId string) int {

	uri := b.apiUri + "identity/group/" + groupId
	log.Println("DeleteGroup() -uri", uri)
	resp, err := b.request.Delete(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DeleteGroup() - b.request:\n %+v", b.request)
	log.Println("DeleteGroup() - Status Code:", resp.StatusCode())
	return resp.StatusCode()
}

//AddUser
//  @Description: 新增人員
//  @receiver b
//  @param userName
//  @param manager_id
//  @param enabled
//  @return bool
func (b *BPMClient) AddUser(userName string, enabled bool, manager_id ...string) (int, bool) {
	Str_enabled := fmt.Sprint(enabled)
	var bodyInput string
	if len(manager_id) != 0 {
		bodyInput = `{"userName":"` + userName + `", "password":"bpm", "manager_id":"` + manager_id[0] + `", "enabled":"` + Str_enabled + `"}`
	} else {
		bodyInput = `{"userName":"` + userName + `", "password":"bpm", "enabled":"` + Str_enabled + `"}`
	}

	uri := b.apiUri + "identity/user"
	log.Println("AddUser() -uri", uri)
	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("AddUser() - body: ", bodyInput)
	log.Printf("AddUser() - b.request:\n %+v", b.request)
	log.Println("AddUser() - Status Code:", resp.StatusCode())
	resp.Body()
	var s2 map[string]interface{}
	err2 := json.Unmarshal(resp.Body(), &s2)
	if err2 != nil {
		log.Fatal(err2)
	}

	id := fmt.Sprintf("%v", s2["id"])
	f, err3 := strconv.Atoi(id)
	if err3 != nil {
		log.Fatal(err3)
	}
	//str2 := fmt.Sprintf("%v", s2["userName"])
	if resp.StatusCode() == 200 {
		errProfileMember := b.AddProfileMember("1", id)
		errProfileManager := b.AddProfileMember("2", id)
		errProfessionalContactData := b.AddProfessionalContactData(id, userName)
		if errProfileMember == 200 && errProfileManager == 200 && errProfessionalContactData == 200 {
			return f, true
		}
	}
	return 0, false
}

//AddProfileMember
//  @Description: 設定為後台管理員或使用者
//  @receiver b
//  @param bodyInput
//  @return []byte
func (b *BPMClient) AddProfileMember(profile_id string, id string) int {
	//s := StringToRawJson(bodyInput)
	bodyInput := `{"profile_id":"` + profile_id + `", "member_type":"USER", "user_id":"` + id + `"}`
	uri := b.apiUri + "portal/profileMember"
	log.Println("AddProfileMember() - uri", uri)

	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("AddProfileMember() - body: ", bodyInput)
	log.Printf("AddProfileMember() - b.request:\n %+v", b.request)
	log.Println("AddProfileMember() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//AddMembership
//  @Description: 設定隸屬部門
//  @receiver b
//  @param bodyInput
//  @return []byte
func (b *BPMClient) AddMembership(bodyInput string) []byte {

	uri := b.apiUri + "identity/membership"
	log.Println("AddMembership() - uri", uri)

	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("AddMembership() - body: ", bodyInput)
	log.Printf("AddMembership() - b.request:\n %+v", b.request)
	log.Println("AddMembership() - Status Code:", resp.StatusCode())

	return resp.Body()
}

//AddProfessionalContactData
//  @Description: 新增人員聯繫資訊
//  @receiver b
//  @param id
//  @param username
//  @return int
func (b *BPMClient) AddProfessionalContactData(id string, username string) int {

	bodyInput := `{"id":"` + id + `", "email":"` + username + `@hta.com.tw"}`

	uri := b.apiUri + "identity/professionalcontactdata"
	log.Println("AddProfessionalContactData() - uri", uri)

	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("AddProfessionalContactData() - body: ", bodyInput)
	log.Printf("AddProfessionalContactData() - b.request:\n %+v", b.request)
	log.Println("AddProfessionalContactData() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//EditUser
//  @Description: 編輯人員
//  @receiver b
//  @param bodyInput
//  @return int
func (b *BPMClient) EditUser(userId string, bodyInput string) int {

	uri := b.apiUri + "identity/user/" + userId
	log.Println("EditGroup() - uri", uri)

	resp, err := b.request.SetBody(bodyInput).Put(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("EditUser() - body: ", bodyInput)
	log.Printf("EditUser() - b.request:\n %+v", b.request)
	log.Println("EditUser() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//DeleteMembership
//  @Description: 刪除人員隸屬部門
//  @receiver b
//  @param userId
//  @param groupId
//  @param roleId
//  @return int
func (b *BPMClient) DeleteMembership(userId string, groupId string, roleId string) int {
	uri := b.apiUri + "identity/membership/" + userId + "/" + groupId + "/" + roleId
	log.Println("DeleteMembership() - uri", uri)

	resp, err := b.request.Delete(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DeleteMembership() - b.request:\n %+v", b.request)
	log.Println("DeleteMembership() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//DeleteUser
//  @Description: 刪除人員
//  @receiver b
//  @param bodyInput
//  @return int
func (b *BPMClient) DeleteUser(userId string) int {
	uri := b.apiUri + "identity/user/" + userId
	log.Println("DeleteUser() - uri", uri)

	resp, err := b.request.Delete(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DeleteUser() - b.request:\n %+v", b.request)
	log.Println("DeleteUser() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

func RebuildJson(jsonBody string, key ...string) string {
	var s1 map[string]interface{}

	err := json.Unmarshal([]byte(jsonBody), &s1)
	if err != nil {
		log.Println("反序列化失败", err)
	}
	s2 := make(map[string]interface{})
	for _, keyy := range key {
		s2[keyy] = s1[keyy]
	}

	s3, err := json.Marshal(s2)
	if err != nil {
		log.Fatal(err)
	}
	return string(s3)
}

func RebuildJsonTest(jsonBody []byte, key []string) [][]string {

	//jsonBodyㄉ反序列化..json>>>object
	var s1 interface{}
	err := json.Unmarshal(jsonBody, &s1)
	if err != nil {
		log.Fatal(err)
	}

	//object array分割
	objArr, ok := s1.([]interface{})
	if !ok {
		log.Fatal("expected an array of objects")
	}

	//宣告s2重建json
	s2 := make([][]string, len(objArr))
	for i := range s2 {
		s2[i] = make([]string, len(key))
	}

	//迭代
	for i, obj := range objArr {
		obj, ok := obj.(map[string]interface{})
		if !ok {
			log.Fatalf("expected type map[string]interface{}, got %s", reflect.TypeOf(objArr[i]))
		}
		for j := 0; j < len(key); j++ {
			s2[i][j] = fmt.Sprintf("%v", obj[key[j]])
		}
	}
	return s2
}

//GetUserMembership
//  @Description: 查看人員隸屬部門角色
func (b *BPMClient) GetUserMembership(userId string) []byte {
	uri := b.apiUri + "identity/membership?c=25&f=user_id=" + userId
	log.Println("GetUserMembership() -uri", uri)
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GetUserMembership() - b.request:\n %+v", b.request)
	log.Println("GetUserMembership() - Status Code:", resp.StatusCode())
	return resp.Body()
}

//EditUserMembership
//  @Description: 編輯人員隸屬部門(for 單一角色)
//  @reciver b
//  @param jsonBody
//  @return bool
func (b *BPMClient) EditUserMembership(group_id, role_id, jsonBody string) bool {
	var s1 map[string]interface{}
	err := json.Unmarshal([]byte(jsonBody), &s1)
	if err != nil {
		log.Fatal(err)
	}
	userId := fmt.Sprintf("%v", s1["user_id"])
	//i := b.GetUserMembership(userId)
	//result := string(i)[1 : len(i)-1]
	//log.Print(result)
	//var s2 map[string]interface{}
	//err2 := json.Unmarshal([]byte(result), &s2)
	//if err2 != nil {
	//	log.Fatal(err2)
	//}
	//groupId := fmt.Sprintf("%v", s2["group_id"])
	//roleId := fmt.Sprintf("%v", s2["role_id"])

	err3 := b.DeleteMembership(userId, group_id, role_id)
	if err3 != 200 {
		return false
	}
	err4 := b.AddMembership(jsonBody)
	if err4 == nil {
		log.Fatal(err4)
	}
	return true
}
