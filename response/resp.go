package response

type Resp struct {
	Status int16  `json:"status"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
	Data   any    `json:"data"`
}

func RespOk(d any) Resp {
	r := Resp{
		Status: 200,
		Msg:    "OK",
		Error:  "",
		Data:   d,
	}
	return r
}
func RespErr(code int16, msg, err string) Resp {
	r := Resp{
		Status: code,
		Msg:    msg,
		Error:  err,
		Data:   nil,
	}
	return r
}
