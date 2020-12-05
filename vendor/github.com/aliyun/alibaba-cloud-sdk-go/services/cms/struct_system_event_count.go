package cms

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

// SystemEventCount is a nested struct in cms response
type SystemEventCount struct {
	Content      string `json:"Content" xml:"Content"`
	Product      string `json:"Product" xml:"Product"`
	Name         string `json:"Name" xml:"Name"`
	GroupId      string `json:"GroupId" xml:"GroupId"`
	Num          int64  `json:"Num" xml:"Num"`
	Level        string `json:"Level" xml:"Level"`
	Status       string `json:"Status" xml:"Status"`
	ResourceId   string `json:"ResourceId" xml:"ResourceId"`
	RegionId     string `json:"RegionId" xml:"RegionId"`
	InstanceName string `json:"InstanceName" xml:"InstanceName"`
	Time         int64  `json:"Time" xml:"Time"`
}
