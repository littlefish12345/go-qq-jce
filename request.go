package goqqjce

type RequestPacketStruct struct {
	Version     int16             `jceId:"1"`
	PackageType byte              `jceId:"2"`
	MessageType int32             `jceId:"3"`
	RequestId   int32             `jceId:"4"`
	ServantName string            `jceId:"5"`
	FuncName    string            `jceId:"6"`
	Buffer      []byte            `jceId:"7"`
	Timeout     int32             `jceId:"8"`
	Context     map[string]string `jceId:"9"`
	Status      map[string]string `jceId:"10"`
}
