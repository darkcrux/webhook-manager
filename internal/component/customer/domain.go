package customer

type Customer struct {
	ID                 *int   `json:"id,omitempty"`
	CustomerExternalID string `json:"customer-external-id,required"`
	UniqueKey          string `json:"unique-key,required"`
}
