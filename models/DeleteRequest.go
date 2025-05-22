package models

type DeleteRequest struct {
	IdList []int64 `json:"id_list"`
}
