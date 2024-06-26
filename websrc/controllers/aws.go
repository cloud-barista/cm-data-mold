/*
Copyright 2023 The Cloud-Barista Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package controllers

func GetAWSRegions() []string {
	return []string{
		"us-east-2", "us-east-1", "us-west-1", "us-west-2",
		"af-south-1", "ap-east-1", "ap-south-2", "ap-southeast-3", "ap-southeast-4",
		"ap-south-1", "ap-northeast-3", "ap-northeast-2", "ap-southeast-1", "ap-southeast-2",
		"ap-northeast-1", "ca-central-1", "eu-central-1", "eu-central-2",
		"eu-west-1", "eu-west-2", "eu-west-3", "eu-south-1", "eu-south-2", "eu-north-1",
		"il-central-1", "me-south-1", "me-central-1", "sa-east-1",
	}
}
