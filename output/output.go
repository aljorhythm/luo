package output

type Data interface {
}

type outputV1 struct {
}

type outputV1data struct {
	Version int `json:"version"`
}

func (o outputV1) print(data Data) []byte {
	outputData := outputV1data{Version: 1}
	return toJson(outputData)
}
