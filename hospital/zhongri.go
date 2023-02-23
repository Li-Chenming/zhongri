package hospital

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Albert-Zhan/httpc"
	"github.com/Albert-Zhan/httpc/body"

	"register/hospital/DTO"
)

type DoctorConfig struct {
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
}
type PatientConfig struct {
	PatID       int    `json:"patId:"`
	PatName     string `json:"patName:"`
	PcID        int    `json:"pcId:"`
	PatCardType int    `json:"patCardType:"`
}

type Env struct {
	Cookie       string `json:"cookie::"`
	Baseinfo     string `json:"baseinfo::"`
	StartRunTime string `json:"startRunTime::"`
	OffsetMs     int    `json:"offsetMs"`
	GoNum     int    `json:"goNum"`
}

var doctorConfig DoctorConfig
var patientConfig PatientConfig
var StartRunTime time.Time
var BaseEnv Env

func init() {

	absConfFilePath, _ := filepath.Abs("./config/doctor.json")
	fmt.Println("load config from " + absConfFilePath)
	file, err := os.ReadFile(absConfFilePath)
	if err != nil {
		fmt.Printf("reder confgi err %v\n", err)
		return
	}

	err = json.Unmarshal(file, &doctorConfig)
	if err != nil {
		fmt.Printf("reder config Unmarshal  err %v\n", err)
		return
	}


	absConfFilePath, _ = filepath.Abs("./config/BaseEnv.json")
	fmt.Println("load config from " + absConfFilePath)
	file, err = os.ReadFile(absConfFilePath)
	if err != nil {
		fmt.Printf("reder confgi err %v\n", err)
		return
	}

	err = json.Unmarshal(file, &BaseEnv)
	if err != nil {
		fmt.Printf("reder config Unmarshal  err %v\n", err)
		return
	}

	absConfFilePath, _ = filepath.Abs("./config/patient.json")
	fmt.Println("load config from " + absConfFilePath)
	file, err = os.ReadFile(absConfFilePath)
	if err != nil {
		fmt.Printf("reder confgi err %v\n", err)
		return
	}

	err = json.Unmarshal(file, &patientConfig)
	if err != nil {
		fmt.Printf("reder config Unmarshal  err %v\n", err)
		return
	}
	fmt.Printf("就诊人:%v , ID:%v,  就诊日期:%v，医生:%v\n", patientConfig.PatName, patientConfig.PatID, doctorConfig.SchDate, doctorConfig.SchLevel)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	parse, err := time.ParseInLocation("2006-01-02 15:04:05", BaseEnv.StartRunTime, loc)
	if err != nil {
		StartRunTime = time.Now()
	}
	offset := time.Duration(BaseEnv.OffsetMs)
	StartRunTime = parse.Add(offset * time.Millisecond)
}

func RunZhongRi(ctx context.Context) error {

	submitReq := DTO.SubmitReq{
		PatID:          patientConfig.PatID,
		PatName:        patientConfig.PatName,
		PcID:           patientConfig.PcID,
		SchID:          doctorConfig.SchID,
		IsSendRegSms:   "N",
		Type:           "0", // null
		RegSelectItems: struct{}{},
		Vcode:          "", // null
		DocID:          strconv.FormatInt(int64(doctorConfig.DocID), 10),
		HosDistID:      strconv.FormatInt(int64(doctorConfig.HosDistID), 10),
		DeptID:         strconv.FormatInt(int64(doctorConfig.DeptID), 10),
		DistrictID:     strconv.FormatInt(int64(doctorConfig.HosDistID), 10),
		PatCardType:    patientConfig.PatCardType,
		PatCardName:    "社保卡",
		DeptName:       doctorConfig.DeptName,
		IsHealthCard:   false,
		Timeout:        30000,
		IsLoading:      true,
	}
	submitRes, err := Submit(ctx,submitReq)
	if err != nil {
		fmt.Printf("Submit err %v\n", err)
		return err
	}
	fmt.Println(submitRes)
	if submitRes.Result{
		return nil
	}
	return  errors.New("faield")
}

func Submit(ctx context.Context, param DTO.SubmitReq) (*DTO.SubmitResp, error) {

	//新建http客户端
	client := httpc.NewHttpClient()
	//新建一个cookie管理器,后面所有请求的cookie将保存在这
	cookieJar := httpc.NewCookieJar()
	//设置cookie管理器,
	client.SetCookieJar(cookieJar)
	//新建一个请求
	req := httpc.NewRequest(client)
	req.SetMethod("post").SetUrl("https://m3.hsyuntai.com/med/hp/hospitals/100215/registration/submit230")

	//设置头信息,返回byte类型的body
	req.SetHeader("Host", "m3.hsyuntai.com")
	req.SetHeader("yt-baseinfo", "%7B%22openid%22%3A%22ow78d5GSjURPy1O7fOp9_t_2tczU%22%2C%22thirdAppId%22%3A%22wx869a0e459663d82a%22%2C%22access_token%22%3Atrue%2C%22appId%22%3A%22100215%22%7D")
	req.SetHeader("Accept-Language", "zh-cn")
	req.SetHeader("Cookie", BaseEnv.Cookie)
	req.SetHeader("Accept", "application/json;charset=utf-8")
	req.SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.32(0x18002033) NetType/WIFI Language/zh_CN miniProgram/wx869a0e459663d82a")
	req.SetHeader("Referer", "https://m3.hsyuntai.com/")
	req.SetHeader("ngversion", "82.1.0")
	req.SetHeader("x-requested-with", "XMLHttpRequest")

	jsonBody, err2 := setJsonBody(param)
	if err2 != nil {
		return &DTO.SubmitResp{}, err2
	}
	submitResp := DTO.SubmitResp{}

	select {
	case <-ctx.Done():
		submitResp.Result=true
		return &submitResp,nil
	default:

	}
	req.SetBody(jsonBody)
	_, bodyByte, err := req.Send().EndByte()
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	fmt.Println(string(bodyByte))

	err = json.Unmarshal(bodyByte, &submitResp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &submitResp, nil

}

func setJsonBody(param interface{}) (*body.Raw, error) {
	raw := body.NewRawData()

	marshal, err := json.Marshal(param)
	if err != nil {
		return raw, err
	}
	raw.SetData(string(marshal), body.Json)
	return raw, nil
}
