package txtypes

type TransactionType struct {
	ID            *int        `json:"id,omitempty"`
	Name          string      `json:"name,required"`
	Description   string      `json:"description"`
	SamplePayload interface{} `json:"sample-payload"`
}
