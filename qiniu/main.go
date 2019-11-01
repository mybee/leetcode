package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	AccessKey  = "IGkXhvImV0mONZ7DQx46uJSz0YB8SUiqtp7m1nMz"
	SecretKKey = "yF6mZKaAr9X-gHi2b4dcvIArzO6iSx5DQl-WTUIv"
)

func main() {
	/*官网文档  https://developer.qiniu.com/kodo/manual/1208/upload-token
	用户根据业务需求，确定上传策略要素，构造出具体的上传策略。例如用户要向空间 my-bucket 上传一个名为 sunflower.jpg 的图片，授权有效期截止到 2015-12-31 00:00:00（该有效期指上传完成后在七牛生成文件的时间，而非上传的开始时间），并且希望得到图片的名称、大小、宽高和校验值。那么相应的上传策略各字段分别为：
	scope = 'my-bucket:sunflower.jpg'
	deadline = 1451491200
	returnBody = '{
				"name": $(fname),
				"size": $(fsize),
				"w": $(imageInfo.width),
				"h": $(imageInfo.height),
				"hash": $(etag)
	}'
	*/
	putPolicy := map[string]interface{}{}
	putPolicy["scope"] = "app-space-1144894155"
	putPolicy["deadline"] = time.Now().Unix() + int64(time.Hour/time.Second) //截至时间！！！一开始写了有效时长上去，而且1s是1e9 。。。
	putPolicy["insertOnly"] = 0
	putPolicy["returnBody"] = `{
            "name": "$(fname)",
            "size": "$(fsize)",
            "w": "$(imageInfo.width)",
            "h": "$(imageInfo.height)",
            "hash": "$(etag)"
        }`
	putPolicyBytes, err := json.Marshal(putPolicy)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("putPolicy:", string(putPolicyBytes))
	encodedPutPolicy := base64.URLEncoding.EncodeToString(putPolicyBytes)
	fmt.Println("encodePutPolicy:", encodedPutPolicy)

	/*官网文档
	sign = hmac_sha1(encodedPutPolicy, "<SecretKKey>")
	#假设 SecretKKey 为 MY_SECRET_KEY，实际签名为：
	sign = "c10e287f2b1e7f547b20a9ebce2aada26ab20ef2"
	注意：签名结果是二进制数据，此处输出的是每个字节的十六进制表示，以便核对检查。
	*/
	mac := hmac.New(sha1.New, []byte(SecretKKey)) //坑点,顺序跟说的不一样？？？看了SDK才知道这样弄，不然一直bad token
	mac.Write([]byte(encodedPutPolicy))
	//sign := fmt.Sprintf("%x\n", mac.Sum(nil))
	//encodeSign := base64.URLEncoding.EncodeToString([]byte(sign))
	digest := mac.Sum(nil) //坑点

	/*官网文档
	encodedSign = urlsafe_base64_encode(sign)
	#最终签名值为：
	encodedSign = "wQ4ofysef1R7IKnrziqtomqyDvI="
	*/
	encodeSign := base64.URLEncoding.EncodeToString(digest)
	fmt.Println("encodeSign:", encodeSign)

	/*官网文档
	uploadToken = AccessKey + ':' + encodedSign + ':' + encodedPutPolicy
	#假设用户的 AccessKey 为 MY_ACCESS_KEY ，则最后得到的上传凭证应为：
	uploadToken = "MY_ACCESS_KEY:wQ4ofysef1R7IKnrziqtomqyDvI=:eyJzY29wZSI6Im15LWJ1Y2tldDpzdW5mbG93ZXIuanBnIiwiZGVhZGxpbmUiOjE0NTE0OTEyMDAsInJldHVybkJvZHkiOiJ7XCJuYW1lXCI6JChmbmFtZSksXCJzaXplXCI6JChmc2l6ZSksXCJ3XCI6JChpbWFnZUluZm8ud2lkdGgpLFwiaFwiOiQoaW1hZ2VJbmZvLmhlaWdodCksXCJoYXNoXCI6JChldGFnKX0ifQ=="
	注意：为确保客户端、业务服务器和七牛服务器对于授权截止时间的理解保持一致，需要同步校准各自的时钟。频繁返回 401 状态码时请先检查时钟同步性与生成 deadline 值的代码逻辑。
	*/
	uploadToken := AccessKey + ":" + encodeSign + ":" + encodedPutPolicy

	form := map[string]string{"token": uploadToken, "key": "test.go"}
	newfileUploadRequest("http://upload.qiniup.com", form, "file", "./test.go")
}
func newfileUploadRequest(uri string, form map[string]string, formFileName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range form {
		err = writer.WriteField(key, val)
		if err != nil {
			return err
		}
	}
	part, err := writer.CreateFormFile(formFileName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Host", "upload.qiniu.com")
	req.Header.Set("Content-Length", fmt.Sprint(body.Len()))
	fmt.Println("reqHeader:", req.Header)
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	Body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("code:", resp.StatusCode, "\nheader:", resp.Header, "\n", string(Body))
	return nil
}
