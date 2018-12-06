package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func painc_er(err error)  {

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func main()  {


	//fsq_ak := "PjFtQJWfvKrSLYkSlV-keCKWzmXzSK1Zp3R9S5MV"
	//fsq_sk := "Q48lAPnTPVxq20dUmfux9HVCrtC3h-p3MCTgMyXf"

	iam_ak := "IAM-2zhqSqp62XGcpyBz15u8sJxdmZiE2qyahyE8HXJ3"
	iam_sk := "ss-PgQnaW5GYNi1Jlv9AAM0b5k0b4c_H3Nme0G9_opvb"

	//test_upload(iam_ak,iam_sk,"./f3","f3")

	//test_buckets(iam_ak,iam_sk)

	//test_domainlist(iam_ak,iam_sk)

	//test_stat(iam_ak,iam_sk)

	//test_mkbucket(iam_ak,iam_sk)
	//test_normal()
	//test_normal2()
	test_iam_getfile(iam_ak,iam_sk)


}

func test_normal()  {
	down_url := "http://b1.com:9200/key?e=%d"
	dl := 1642444718+3600
	down_url = fmt.Sprintf(down_url,dl)

	a_key := "PjFtQJWfvKrSLYkSlV-keCKWzmXzSK1Zp3R9S5MV"
	s_key := "Q48lAPnTPVxq20dUmfux9HVCrtC3h-p3MCTgMyXf"
	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(s_key))
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte(down_url))

	sign_url := mac.Sum(nil)

	b64_url := base64.URLEncoding.EncodeToString(sign_url)

	token := a_key+":"+b64_url

	final_url := fmt.Sprintf("http://b1.com:9200/key?e=%d&token=%s",dl,token)

	fmt.Println(final_url)

	req ,_:= http.NewRequest("GET",final_url,nil)
	req.Header.Set("Host","b1.com")

	//client := &http.Client{}
	rs,err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:",err)
	}

	body,_ := ioutil.ReadAll(rs.Body)
	fmt.Println("response:", string(body))
}


func test_normal2()  {
	down_url := "http://b2.com:9200/f2?e=%d"
	dl := 1642444718+3600
	down_url = fmt.Sprintf(down_url,dl)

	a_key := "PjFtQJWfvKrSLYkSlV-keCKWzmXzSK1Zp3R9S5MV"
	s_key := "Q48lAPnTPVxq20dUmfux9HVCrtC3h-p3MCTgMyXf"
	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(s_key))
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte(down_url))

	sign_url := mac.Sum(nil)

	b64_url := base64.URLEncoding.EncodeToString(sign_url)

	token := a_key+":"+b64_url

	final_url := fmt.Sprintf("http://b2.com:9200/f2?e=%d&token=%s",dl,token)

	fmt.Println(final_url)

	req ,_:= http.NewRequest("GET",final_url,nil)
	req.Header.Set("Host","b2.com")

	//client := &http.Client{}
	rs,err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:",err)
	}

	body,_ := ioutil.ReadAll(rs.Body)
	fmt.Println("response:", string(body))
	fmt.Println("req_id:", rs.Header.Get("X-Reqid"))
}



func test_iam_getfile(ak,sk string)  {

	//{
	//	"_id" : ObjectId("5aa8cced43c8ce7bcb83ab3c"),
	//	"created_at" : ISODate("2018-03-14T07:19:09.007Z"),
	//	"access_key" : "IAM-2zhqSqp62XGcpyBz15u8sJxdmZiE2qyahyE8HXJ3",
	//	"secret_key" : "ss-PgQnaW5GYNi1Jlv9AAM0b5k0b4c_H3Nme0G9_opvb",
	//	"user_id" : ObjectId("5aa8ccec43c8ce7bcb83ab3b"),
	//	"enabled" : true
	//}
	//{
	//	"_id" : ObjectId("5aa8ccec43c8ce7bcb83ab3b"),
	//	"root_uid" : 1810657124,
	//	"iuid" : 1517381413,
	//	"alias" : "test4",
	//	"password" : "$2a$10$rc5ZyVYxtCP9MqDEGGgyzucz3IoCDSn4uO6k5JyEKkLrRjALCiueS",
	//	"created_at" : ISODate("2018-03-14T07:19:08.867Z"),
	//	"updated_at" : ISODate("2018-11-21T03:31:18.754Z"),
	//	"last_login_time" : ISODate("0001-01-01T00:00:00Z"),
	//	"enabled" : true,
	//	"group_ids" : [ ],
	//	"policy_ids" : [
	//	ObjectId("5a9fb03443c8ce15db19fb1d"),
	//		ObjectId("5a9fb77e43c8ce15db19fb20")
	//	]
	//}

	down_url := "http://b2.com:9200/f2?e=%d"
	dl := 1642444718+3600
	down_url = fmt.Sprintf(down_url,dl)

	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(sk))
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte(down_url))

	sign_url := mac.Sum(nil)

	b64_url := base64.URLEncoding.EncodeToString(sign_url)

	token := ak+":"+b64_url

	final_url := fmt.Sprintf("http://b2.com:9200/f2?e=%d&token=%s",dl,token)

	fmt.Println(final_url)

	req ,_:= http.NewRequest("GET",final_url,nil)
	req.Header.Set("Host","b2.com")

	//client := &http.Client{}
	rs,err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:",err)
	}

	body,_ := ioutil.ReadAll(rs.Body)
	fmt.Println("response:", string(body))


}

func test_upload(ak,sk ,file_path,file_name string )  {

	up_pl_u :=  base64.URLEncoding.EncodeToString([]byte(`{"scope":"b2","deadline":1642444718}`))

	mac := hmac.New(sha1.New, []byte(sk))
	mac.Write([]byte(up_pl_u))
	b64_pl := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	token := ak+":"+b64_pl+":"+up_pl_u

	URL := fmt.Sprintf("http://b2.com:11200/upload/")

	fields := map[string]string{
		"file": file_name,
		"token":token,
	}

	req, err := multipartUpload(URL, file_path, fields)
	painc_er(err)
	fmt.Println("req",req)
	res ,err :=  http.DefaultClient.Do(req)
	painc_er(err)

	body,_ := ioutil.ReadAll(res.Body)
	fmt.Println("res:",res,"\nbody:", string(body))
	return

}


func multipartUpload(destURL string, file_path string, fields map[string]string) (*http.Request, error) {

	f, err := os.Open(file_path)
	defer f.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", fields["filename"])
	if err != nil {
		return nil, fmt.Errorf("CreateFormFile %v", err)
	}

	_, err = io.Copy(fw, f)
	if err != nil {
		return nil, fmt.Errorf("copying fileWriter %v", err)
	}

	for k, v := range fields {
		_ = writer.WriteField(k, v)
	}

	err = writer.Close() // close writer before POST request
	if err != nil {
		return nil, fmt.Errorf("writerClose: %v", err)
	}

	//resp, err := http.Post(destURL, writer.FormDataContentType(), body)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return resp, nil

	req, err := http.NewRequest("POST", destURL, body)
	if err != nil {
	 return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	return  req,nil
	//if req.Close && req.Body != nil {
	// defer req.Body.Close()
	//}

	// return http.DefaultClient.Do(req)
}

func test_buckets(ak,sk string)  {

	down_url := "http://127.0.0.1:9400/buckets"
	down_url = fmt.Sprintf(down_url)

	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(sk))
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte("/buckets\n"))
	sign_url := mac.Sum(nil)
	b64_url := base64.URLEncoding.EncodeToString(sign_url)
	token := ak+":"+b64_url
	fmt.Println(token)

	req ,_:= http.NewRequest("GET",down_url,nil)
	req.Header.Set("Host","b2.com")
	req.Header.Set("Authorization","QBox "+token)


	rs,err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:",err)
		return
	}

	body,_ := ioutil.ReadAll(rs.Body)
	fmt.Println("response:", string(body))
}

func test_domainlist(ak,sk string)  {

	down_url := "http://127.0.0.1:9400/list?tbl=b2"
	down_url = fmt.Sprintf(down_url)

	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(sk))
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte("/buckets\n"))
	sign_url := mac.Sum(nil)
	b64_url := base64.URLEncoding.EncodeToString(sign_url)
	token := ak+":"+b64_url
	fmt.Println(token)

	req ,_:= http.NewRequest("GET",down_url,nil)
	req.Header.Set("Host","b2.com")
	req.Header.Set("Authorization","QBox "+token)


	rs,err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:",err)
		return
	}

	body,_ := ioutil.ReadAll(rs.Body)
	fmt.Println("response:", string(body))


}

func test_mkbucket(ak, sk string)  {

	bucket := base64.URLEncoding.EncodeToString([]byte("b3"))
	url := "http://127.0.0.1:9400/mkbucketv2/"+bucket+"/region/z0"


	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(sk))
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte("/mkbucketv2/"+bucket+"/region/z0\n"))
	sign_url := mac.Sum(nil)
	b64_url := base64.URLEncoding.EncodeToString(sign_url)
	token := ak+":"+b64_url
	fmt.Println(token)

	req ,err:= http.NewRequest("POST",url,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Host","127.0.0.1:9400")
	req.Header.Set("Authorization","QBox "+token)

	rs,err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:",err)
		return
	}

	body,_ := ioutil.ReadAll(rs.Body)
	fmt.Println("response:", string(body))

}

func test_stat(ak, sk string)  {

	ob := base64.URLEncoding.EncodeToString([]byte("b2:f2"))
	down_url := "http://127.0.0.1:9400/stat/"+ob

	//hmac ,use sha1
	mac := hmac.New(sha1.New, []byte(sk))
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte("/stat/"+ob+"\n"))
	sign_url := mac.Sum(nil)
	b64_url := base64.URLEncoding.EncodeToString(sign_url)
	token := ak+":"+b64_url
	fmt.Println(token)

	req ,_:= http.NewRequest("GET",down_url,nil)
	req.Header.Set("Host","b2.com")
	req.Header.Set("Authorization","QBox "+token)


	rs,err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("err:",err)
		return
	}

	body,_ := ioutil.ReadAll(rs.Body)
	fmt.Println("response:", string(body))
}

func test_drop(ak, sk string)  {

	//url := "http://127.0.0.1:9400/drop/"


}