package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"net/http"
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
	awsbucket := bucket
	awsKey := jobId
	s3resp, err := s3.New(sess).GetObject(&s3.GetObjectInput{
		Bucket: &awsbucket,
		Key:    &awsKey,
	})
	if err != nil {
		return nil, err
	}
	return apiResponse(http.StatusOK, s3resp.Body)

}
func getListOfJobs(sess *session.Session)(*events.ALBTargetGroupResponse,error) {
	awsbucket := bucket
	awsKey := FilesWithJobList
	s3resp, err := s3.New(sess).GetObject(&s3.GetObjectInput{
		Bucket: &awsbucket,
		Key:    &awsKey,
	})
	if err != nil {
		return nil, err
	}
	return apiResponse(http.StatusOK, s3resp.Body)
}

func main()  {
	lambda.Start(handler)
	
}


func apiResponse(status int, body io.ReadCloser) (*events.ALBTargetGroupResponse, error) {
	resp := events.ALBTargetGroupResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	    resp.StatusCode = status
		stringBody, _ := json.Marshal(body)
		resp.Body = string(stringBody)
		return &resp, nil


}