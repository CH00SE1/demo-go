package inck

type GroupDoc struct {
	placepointname string
	inputmanname   string
}

func (g GroupDoc) TableName() string {
	return "resa_group_req_doc_v"
}
