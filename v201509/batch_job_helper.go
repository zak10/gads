package v201509

import (
	"fmt"
	"reflect"
	"encoding/xml"
	//"net/http"
)

type BatchJobHelper struct {
	Auth
}

func NewBatchJobHelper(auth *Auth) *BatchJobHelper {
	return &BatchJobHelper{Auth: *auth}
}

func (s *BatchJobHelper) UploadBatchJobOperations(jobOperations []interface{}, url TemporaryUrl) (mutateResp SimpleMutateResponse, err error) {

	var operations []Operation
	for _, operation := range jobOperations {
		if operationType, valid := getXsiType(reflect.ValueOf(operation).Type().String()); valid {
			switch reflect.TypeOf(operation).Kind() {
			    case reflect.Map:
			        ops := reflect.ValueOf(operation)

			        keys := ops.MapKeys()

			        for _, action := range keys {
			        	jobs := ops.MapIndex(action)

			        	for i := 0; i < jobs.Len(); i++ {
				            
				            operations = append(operations,
								Operation{
									Operator:   action.String(),
									Operand: 	jobs.Index(i).Interface(),
									Xsi_type: 	operationType,
								},
							)
				        }
					}
	    	}
	    }	
	}

	if len(operations) > 0 {
		mutation := struct {
			XMLName xml.Name
			Ops     []Operation 			`xml:"operations"`
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "mutate",
			},
			Ops: operations,
		}

		fmt.Println(url.Url)

		respBody, err := s.Auth.request(ServiceUrl{url.Url, ""}, "mutate", mutation)
		/*req, err := http.NewRequest("POST", url.Url, "")
		
		req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
		resp, err := s.Auth.Client.Do(req)*/
		fmt.Println(err)
		if err != nil {
			//return mutateResp, err
		}

		fmt.Println(respBody)

		//err = xml.Unmarshal([]byte(respBody), &mutateResp)
	}
	fmt.Println(operations)
	return mutateResp, err
}