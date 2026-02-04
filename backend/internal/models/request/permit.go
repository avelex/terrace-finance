package request

type SignedPermit struct {
	Networks map[uint32]string `json:"networks"`
}
