package cr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateRepoAuthorization invokes the cr.UpdateRepoAuthorization API synchronously
// api document: https://help.aliyun.com/api/cr/updaterepoauthorization.html
func (client *Client) UpdateRepoAuthorization(request *UpdateRepoAuthorizationRequest) (response *UpdateRepoAuthorizationResponse, err error) {
	response = CreateUpdateRepoAuthorizationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRepoAuthorizationWithChan invokes the cr.UpdateRepoAuthorization API asynchronously
// api document: https://help.aliyun.com/api/cr/updaterepoauthorization.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRepoAuthorizationWithChan(request *UpdateRepoAuthorizationRequest) (<-chan *UpdateRepoAuthorizationResponse, <-chan error) {
	responseChan := make(chan *UpdateRepoAuthorizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRepoAuthorization(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateRepoAuthorizationWithCallback invokes the cr.UpdateRepoAuthorization API asynchronously
// api document: https://help.aliyun.com/api/cr/updaterepoauthorization.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateRepoAuthorizationWithCallback(request *UpdateRepoAuthorizationRequest, callback func(response *UpdateRepoAuthorizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRepoAuthorizationResponse
		var err error
		defer close(result)
		response, err = client.UpdateRepoAuthorization(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateRepoAuthorizationRequest is the request struct for api UpdateRepoAuthorization
type UpdateRepoAuthorizationRequest struct {
	*requests.RoaRequest
	RepoNamespace string           `position:"Path" name:"RepoNamespace"`
	RepoName      string           `position:"Path" name:"RepoName"`
	AuthorizeId   requests.Integer `position:"Path" name:"AuthorizeId"`
}

// UpdateRepoAuthorizationResponse is the response struct for api UpdateRepoAuthorization
type UpdateRepoAuthorizationResponse struct {
	*responses.BaseResponse
}

// CreateUpdateRepoAuthorizationRequest creates a request to invoke UpdateRepoAuthorization API
func CreateUpdateRepoAuthorizationRequest() (request *UpdateRepoAuthorizationRequest) {
	request = &UpdateRepoAuthorizationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "UpdateRepoAuthorization", "/repos/[RepoNamespace]/[RepoName]/authorizations/[AuthorizeId]", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRepoAuthorizationResponse creates a response to parse from UpdateRepoAuthorization response
func CreateUpdateRepoAuthorizationResponse() (response *UpdateRepoAuthorizationResponse) {
	response = &UpdateRepoAuthorizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}