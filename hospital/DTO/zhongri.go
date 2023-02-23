package DTO

type RegListReq struct {
	RegType      string `json:"regType"`
	Type         string `json:"type"`
	DistID       string `json:"distId"`
	OutpatientID string `json:"outpatientId"`
	DepartmentID string `json:"departmentId"`
	SchDate      string `json:"schDate"`
	MediLevel    string `json:"mediLevel"`
	PageNum      int    `json:"pageNum"`
	PageSize     int    `json:"pageSize"`
	EnumType     string `json:"enumType"`
	Timeout      int    `json:"timeout"`
	IsLoading    bool   `json:"isLoading"`
}

type RegListResp struct {
	Data    Data   `json:"data"`
	NowTime string `json:"nowTime"`
	Result  bool   `json:"result"`
	APIName string `json:"api_name"`
}
type List struct {
	DeptID        int    `json:"deptId"`
	DeptName      string `json:"deptName"`
	DocID         int    `json:"docId"`
	DocName       string `json:"docName"`
	GoodAt        string `json:"goodAt"`
	MediLevel     int    `json:"mediLevel"`
	MediLevelName string `json:"mediLevelName"`
	RegCount      int    `json:"regCount"`
	TitleShown    string `json:"titleShown"`
	VirtualFlag   string `json:"virtualFlag"`
	HeadPhoto     string `json:"headPhoto,omitempty"`
}
type Data struct {
	EndRow            int    `json:"endRow"`
	HasNextPage       bool   `json:"hasNextPage"`
	HasPreviousPage   bool   `json:"hasPreviousPage"`
	IsFirstPage       bool   `json:"isFirstPage"`
	IsLastPage        bool   `json:"isLastPage"`
	List              []List `json:"list"`
	NavigateFirstPage int    `json:"navigateFirstPage"`
	NavigateLastPage  int    `json:"navigateLastPage"`
	NavigatePages     int    `json:"navigatePages"`
	NavigatepageNums  []int  `json:"navigatepageNums"`
	NextPage          int    `json:"nextPage"`
	PageNum           int    `json:"pageNum"`
	PageSize          int    `json:"pageSize"`
	Pages             int    `json:"pages"`
	PrePage           int    `json:"prePage"`
	Size              int    `json:"size"`
	StartRow          int    `json:"startRow"`
	Total             int    `json:"total"`
}

type DoctorDetailResp struct {
	Result bool `json:"result"`
	Data   []struct {
		SchID          int64  `json:"schId"`
		AccessSchID    string `json:"accessSchId"`
		HosDistID      int    `json:"hosDistId"`
		HosDistName    string `json:"hosDistName"`
		DistAddr       string `json:"distAddr"`
		DeptID         int    `json:"deptId"`
		DeptName       string `json:"deptName"`
		DocID          int    `json:"docId"`
		ResNo          int    `json:"resNo"`
		SchDate        string `json:"schDate"`
		WeekType       string `json:"weekType"`
		StartTime      string `json:"startTime"`
		EndTime        string `json:"endTime"`
		DayType        int    `json:"dayType"`
		Cost           int    `json:"cost"`
		SchLevel       string `json:"schLevel"`
		State          int    `json:"state"`
		NumSrcType     int    `json:"numSrcType"`
		StateShown     string `json:"stateShown"`
		RegType        int    `json:"regType"`
		NeedChooseFvSv bool   `json:"needChooseFvSv"`
	} `json:"data"`
	NowTime string `json:"nowTime"`
	APIName string `json:"api_name"`
}

type SubmitReq struct {
	PatID          int    `json:"patId"`
	PatName        string `json:"patName"`
	PcID           int    `json:"pcId"`
	SchID          int64  `json:"schId"`
	IsSendRegSms   string `json:"isSendRegSms"`
	Type           string `json:"type"`
	RegSelectItems struct {
	} `json:"regSelectItems"`
	Vcode        string `json:"vcode"`
	DocID        string `json:"docId"`
	HosDistID    string `json:"hosDistId"`
	DeptID       string `json:"deptId"`
	DistrictID   string `json:"districtId"`
	PatCardType  int    `json:"patCardType"`
	PatCardName  string `json:"patCardName"`
	DeptName     string `json:"deptName"`
	IsHealthCard bool   `json:"isHealthCard"`
	Timeout      int    `json:"timeout"`
	IsLoading    bool   `json:"isLoading"`
}

type SubmitResp struct {
	Result bool `json:"result"`
	Data   struct {
		RegID           int    `json:"regId"`
		AccessRegID     string `json:"accessRegId"`
		PrdCode         string `json:"prdCode"`
		SchID           int64  `json:"schId"`
		DayType         int    `json:"dayType"`
		AccessSchID     string `json:"accessSchId"`
		HosID           int    `json:"hosId"`
		UsID            int64  `json:"usId"`
		PatID           int    `json:"patId"`
		PatName         string `json:"patName"`
		PhoneNo         string `json:"phoneNo"`
		IDCardNo        string `json:"idCardNo"`
		PcID            int    `json:"pcId"`
		PatCardNo       string `json:"patCardNo"`
		PatCardName     string `json:"patCardName"`
		ExpectDate      string `json:"expectDate"`
		ExpectStime     string `json:"expectStime"`
		ExpectEtime     string `json:"expectEtime"`
		ExpectAddr      string `json:"expectAddr"`
		RegStatus       int    `json:"regStatus"`
		PayStatus       int    `json:"payStatus"`
		RegText         string `json:"regText"`
		Cost            int    `json:"cost"`
		ServiceFee      int    `json:"serviceFee"`
		RegFee          int    `json:"regFee"`
		CreateTime      string `json:"createTime"`
		DeptID          int    `json:"deptId"`
		DeptName        string `json:"deptName"`
		AccessDeptID    string `json:"accessDeptId"`
		DeptAddr        string `json:"deptAddr"`
		DocID           int    `json:"docId"`
		AccessDocID     string `json:"accessDocId"`
		DocName         string `json:"docName"`
		TitleShown      string `json:"titleShown"`
		HosDistID       int    `json:"hosDistId"`
		HosDistName     string `json:"hosDistName"`
		ExpiredStatus   int    `json:"expiredStatus"`
		WaitAddr        string `json:"waitAddr"`
		IsExpert        int    `json:"isExpert"`
		RegType         int    `json:"regType"`
		PayCountDown    int    `json:"payCountDown"`
		HosName         string `json:"hosName"`
		HosAlias        string `json:"hosAlias"`
		TakeIndex       int    `json:"takeIndex"`
		TakeIndexStr    string `json:"takeIndexStr"`
		PayExpireTime   string `json:"payExpireTime"`
		HavaComment     string `json:"havaComment"`
		ActualPayAmount int    `json:"actualPayAmount"`
		HisBizAmount    int    `json:"hisBizAmount"`
		HosTel          string `json:"hosTel"`
		DistAddr        string `json:"distAddr"`
		AccessPatID     string `json:"accessPatId"`
		Birthday        string `json:"birthday"`
	} `json:"data"`
	NowTime       string `json:"nowTime"`
	APIName       string `json:"api_name"`
	EnablePayment bool   `json:"enablePayment"`
}
