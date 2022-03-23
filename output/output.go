package output

type Data interface {
}

type outputV1 struct {
}

func (o outputV1) print(data Data) []byte {
	return toJson(data)
}
