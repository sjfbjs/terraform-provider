package cbn

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

// DeleteCenBandwidthPackage invokes the cbn.DeleteCenBandwidthPackage API synchronously
// api document: https://help.aliyun.com/api/cbn/deletecenbandwidthpackage.html
func (client *Client) DeleteCenBandwidthPackage(request *DeleteCenBandwidthPackageRequest) (response *DeleteCenBandwidthPackageResponse, err error) {
	response = CreateDeleteCenBandwidthPackageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCenBandwidthPackageWithChan invokes the cbn.DeleteCenBandwidthPackage API asynchronously
// api document: https://help.aliyun.com/api/cbn/deletecenbandwidthpackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCenBandwidthPackageWithChan(request *DeleteCenBandwidthPackageRequest) (<-chan *DeleteCenBandwidthPackageResponse, <-chan error) {
	responseChan := make(chan *DeleteCenBandwidthPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCenBandwidthPackage(request)
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

// DeleteCenBandwidthPackageWithCallback invokes the cbn.DeleteCenBandwidthPackage API asynchronously
// api document: https://help.aliyun.com/api/cbn/deletecenbandwidthpackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCenBandwidthPackageWithCallback(request *DeleteCenBandwidthPackageRequest, callback func(response *DeleteCenBandwidthPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCenBandwidthPackageResponse
		var err error
		defer close(result)
		response, err = client.DeleteCenBandwidthPackage(request)
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

// DeleteCenBandwidthPackageRequest is the request struct for api DeleteCenBandwidthPackage
type DeleteCenBandwidthPackageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	CenBandwidthPackageId string           `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteCenBandwidthPackageResponse is the response struct for api DeleteCenBandwidthPackage
type DeleteCenBandwidthPackageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCenBandwidthPackageRequest creates a request to invoke DeleteCenBandwidthPackage API
func CreateDeleteCenBandwidthPackageRequest() (request *DeleteCenBandwidthPackageRequest) {
	request = &DeleteCenBandwidthPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteCenBandwidthPackage", "cbn", "openAPI")
	return
}

// CreateDeleteCenBandwidthPackageResponse creates a response to parse from DeleteCenBandwidthPackage response
func CreateDeleteCenBandwidthPackageResponse() (response *DeleteCenBandwidthPackageResponse) {
	response = &DeleteCenBandwidthPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
