package types

type Tipset struct {
	Cids    []map[string]string `json:"Cids"`
	Blocks  []interface{}       `json:"Blocks"`
	Heights int                 `json:"Height"`
}
