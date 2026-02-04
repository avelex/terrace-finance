package response

type PermitPayload struct {
	Networks map[uint32]string `json:"networks"`
}
