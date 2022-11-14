package reponse

type ReponseMessge struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (this *ReponseMessge) setCode() {

}

func (this *ReponseMessge) setMessage() {

}

func (this *ReponseMessge) setData() {

}
