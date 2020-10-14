package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"net/http"

	"os"
)

const (
	bucket = ""
	FilesWithJobList = ""
	Region = ""
)
type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}
func handler(req events.ALBTargetGroupRequest ) (*events.ALBTargetGroupResponse , error) {
	
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(Region)},
	)
	if req.Path == "" {
     return getListOfJobs(sess)
	} else {
		return getJobs(sess,req.Path)
	}
}

func getJobs(sess *session.Session,jobId string)  (*events.ALBTargetGroupResponse,error){
	file, err := os.Create("tempfile")
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{
			aws.String(err.Error()),
		})
	}
	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(jobId),
		})
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{
			aws.String(err.Error()),
		})
	}
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	result, err := ioutil.ReadAll(file)
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)

}
func getListOfJobs(sess *session.Session)(*events.ALBTargetGroupResponse,error){
	file, err := os.Create("tempfile")
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{
			aws.String(err.Error()),
		})
	}
	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(FilesWithJobList),
		})
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{
			aws.String(err.Error()),
		})
	}
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	result, err := ioutil.ReadAll(file)
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)

}
func main()  {
	lambda.Start(handler)
	f,err := os.Open("mishra")
	if err!= nil {
		fmt.Print("my files list", err)
	}
	fmt.Print(f)
}


func apiResponse(status int, body interface{}) (*events.ALBTargetGroupResponse, error) {
	resp := events.ALBTargetGroupResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status
	switch  body.(type) {
	case  []byte:
		stringBody, _ := json.Marshal(body)
		resp.Body = string(stringBody)
		return &resp, nil
	default:
		resp.StatusCode = http.StatusInternalServerError
		return &resp,nil
	}

}