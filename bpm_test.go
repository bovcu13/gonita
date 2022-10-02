package main

import (
	"bytes"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestBPMClient_CreateProcessCase(t *testing.T) {

	b := New("isabelle_wu")

	formInput1 := `{
		"account":"james",
		"pm":"44", 
		"tm":["43","42"] 
	}`

	type fields struct {
		serverUri  string
		apiUri     string
		username   string
		password   string
		request    *resty.Request
		token      string
		jSessionId string
	}
	type args struct {
		processId string
		jsonBody  string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCaseId []byte
	}{
		{
			name: "isabelle_wu",
			fields: fields{
				serverUri:  b.serverUri,
				apiUri:     b.apiUri,
				username:   b.username,
				password:   b.password,
				request:    b.request,
				token:      b.token,
				jSessionId: b.jSessionId,
			},
			args: args{
				processId: "7999808492643941641",
				jsonBody:  formInput1,
			},
			wantCaseId: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BPMClient{
				serverUri:  tt.fields.serverUri,
				apiUri:     tt.fields.apiUri,
				username:   tt.fields.username,
				password:   tt.fields.password,
				request:    tt.fields.request,
				token:      tt.fields.token,
				jSessionId: tt.fields.jSessionId,
			}
			if gotCaseId := b.CreateProcessCase(tt.args.processId, tt.args.jsonBody); string(gotCaseId) != string(tt.wantCaseId) {
				t.Errorf("CreateProcessCase() = %v, want %v", gotCaseId, tt.wantCaseId)
			}
		})
	}
}

func TestModelInputStringToRawJson(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `{"pm": 44, "tm": [44]} TO {"modelInput":{"pm":44,"tm":[44]}}`,
			args: args{
				s: `{"pm": 44, "tm": [44]}`,
			},
			want: `{"modelInput":{"pm":44,"tm":[44]}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModelInputStringToRawJson(tt.args.s); got != tt.want {
				t.Errorf("StringToRawJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetStateCaseList(t *testing.T) {

	b := New("isabelle_wu")
	// GetReadyCase("50", "ready", "38")

	type fields struct {
		serverUri  string
		apiUri     string
		username   string
		password   string
		request    *resty.Request
		token      string
		jSessionId string
	}
	type args struct {
		rows   string
		state  string
		userId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: `GetReadyCase("38","50","ready")`,
			fields: fields{
				serverUri:  b.serverUri,
				apiUri:     b.apiUri,
				username:   b.username,
				password:   b.password,
				request:    b.request,
				token:      b.token,
				jSessionId: b.jSessionId,
			},
			args: args{
				rows:   "50",
				state:  "ready",
				userId: "38",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BPMClient{
				serverUri:  tt.fields.serverUri,
				apiUri:     tt.fields.apiUri,
				username:   tt.fields.username,
				password:   tt.fields.password,
				request:    tt.fields.request,
				token:      tt.fields.token,
				jSessionId: tt.fields.jSessionId,
			}
			if got := b.GetStateCaseList(tt.args.rows, tt.args.state, tt.args.userId); !bytes.Equal(got, tt.want) {
				t.Errorf("GetStateCaseList() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestBPMClient_ExecuteTask(t *testing.T) {

	b := New("isabelle_wu")
	type args struct {
		taskId   string
		jsonBody string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `ExecuteTask`,
			args: args{
				taskId:   "100035",
				jsonBody: `{"dstaff":[74,79],"gmapprovalstatus":true}`,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.ExecuteTask(tt.args.taskId, tt.args.jsonBody); got != tt.want {
				t.Errorf("ExecuteTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetCasePendingTaskDetail(t *testing.T) {
	b := New("isabelle_wu")
	type args struct {
		caseId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				caseId: "3209",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetCasePendingTaskDetail(tt.args.caseId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetCasePendingTaskDetail() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestBPMClient_GetCaseArchivedTaskDetail(t *testing.T) {
	b := New("isabelle_wu")
	type args struct {
		caseId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				caseId: "3209",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetCaseArchivedTaskDetail(tt.args.caseId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetCaseArchivedTaskDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetArchivedTaskDetail(t *testing.T) {
	b := New("isabelle_wu")
	type args struct {
		sourceObjectId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				sourceObjectId: "60746",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetArchivedTaskDetail(tt.args.sourceObjectId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetArchivedTaskDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetProcessAllCaseList(t *testing.T) {

	b := New("isabelle_wu")
	type args struct {
		rows      string
		processId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				rows:      "100",
				processId: "7999808492643941641",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetProcessAllCaseList(tt.args.rows, tt.args.processId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetProcessAllCaseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestBPMClient_CreateProcessCaseT(t *testing.T) {
// 	b := New("isabelle_wu")
// 	type args struct {
// 		pm string
// 		tm []string
// 	}
// 	tests := []struct {
// 		name       string
// 		b          *BPMClient
// 		args       args
// 		wantCaseId []byte
// 	}{
// 		{
// 			args: args{
// 				pm: "44",
// 				tm: []string{"38"},
// 			},
// 			wantCaseId: []byte(""),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotCaseId := b.CreateProcessCaseT(tt.args.pm, tt.args.tm); !bytes.Equal(gotCaseId, tt.wantCaseId) {
// 				t.Errorf("BPMClient.CreateProcessCaseT() = %v, want %v", gotCaseId, tt.wantCaseId)
// 			}
// 		})
// 	}
// }

func TestBPMClient_GetProcessInstanceId(t *testing.T) {
	b := New("isabelle_wu")
	type args struct {
		processName string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{},
			want: []byte(""),
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetProcessInstanceId(); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetProcessInstanceId() = %v, want %v", got, tt.want)
			}
		})
	}
}
