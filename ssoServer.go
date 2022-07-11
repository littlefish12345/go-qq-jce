package goqqjce

type SsoServerRequestStruct struct {
	Uin     int64  `jceId:"1"`
	Timeout int64  `jceId:"2"`
	C       byte   `jceId:"3"` //一直为0x01
	IMSI    string `jceId:"4"`
	IsWifi  int32  `jceId:"5"` //true=100, false=1
	AppId   int32  `jceId:"6"`
	IMEI    string `jceId:"7"`
	CellId  int64  `jceId:"8"`
	I       int64  `jceId:"9"`
	J       int64  `jceId:"10"`
	K       int64  `jceId:"11"`
	L       bool   `jceId:"12"`
	M       int64  `jceId:"13"`
}

type SsoServerInfoStruct struct {
	Host     string `jceId:"1"`
	Port     int32  `jceId:"2"`
	Location string `jceId:"8"`
}

type SsoServerRequestPacketStruct struct {
	Version     int16             `jceId:"1"`
	PkgType     byte              `jceId:"2"`
	MsgType     int32             `jceId:"3"`
	ReqId       int32             `jceId:"4"`
	ServantName string            `jceId:"5"`
	FuncName    string            `jceId:"6"`
	Buffer      []byte            `jceId:"7"`
	Timeout     int32             `jceId:"8"`
	Context     map[string]string `jceId:"9"`
	Status      map[string]string `jceId:"10"`
}
