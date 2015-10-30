package v201509

import (
	"encoding/xml"
)

type BatchJobService struct {
	Auth
}

type BatchJobPage struct {

}

type BatchJobOperations struct {
	BatchJobOperations []BatchJobOperation
}

type BatchJobOperation struct {
	Operator 	string 		`xml:"operator"`
	Operand 	BatchJob 	`xml:"operand"`
}

type BatchJob struct {
	Id 					int64 						`xml:"id,omitempty" json:",string"`
	Status 				string 						`xml:"status,omitempty"`
	ProgressStats 		*ProgressStats 				`xml:"progressStats,omitempty"`
	UploadUrl 			*TemporaryUrl 				`xml:"uploadUrl,omitempty"`
	DownloadUrl 		*TemporaryUrl 				`xml:"downloadUrl,omitempty"`
	ProcessingErrors 	*BatchJobProcessingError 	`xml:"processingErrors,omitempty"`
}

type TemporaryUrl struct {
	Url 		string 	`xml:"url"`
	Expiration 	string 	`xml:"expiration,"`
}

type BatchJobProcessingError struct {
	FieldPath 		string 	`xml:"fieldPath"`
	Trigger 		string 	`xml:"trigger"`
	ErrorString 	string 	`xml:"errorString"`
	Reason 			string 	`xml:"reason"`
}

type ProgressStats struct {
	NumOperationsExecuted 		int64 	`xml:"numOperationsExecuted" json:",string"`
	NumOperationsSucceeded 		int64 	`xml:"numOperationsSucceeded" json:",string"`
	EstimatedPercentExecuted 	int 	`xml:"estimatedPercentExecuted"`
	NumResultsWritten 			int64 	`xml:"numResultsWritten" json:",string"`
}

func NewBatchJobService(auth *Auth) *BatchJobService {
	return &BatchJobService{Auth: *auth}
}

func (s *BatchJobService) Get(selector Selector) (batchJobPage BatchJobPage) {
	return batchJobPage
}

func (s *BatchJobService) Mutate(batchJobOperations BatchJobOperations) (batchJobs []BatchJob, err error) {
	
	mutation := struct {
		XMLName xml.Name
		Ops     []BatchJobOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: batchJobOperations.BatchJobOperations}
	respBody, err := s.Auth.request(batchJobServiceUrl, "mutate", mutation)
	if err != nil {
		return batchJobs, err
	}
	mutateResp := struct {
		BatchJobs []BatchJob `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return batchJobs, err
	}

	return mutateResp.BatchJobs, err
}

func (s *BatchJobService) Query() {
	
}