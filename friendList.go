package goqqjce

type FriendListRequest struct {
	RequestType     int32   `jceId:"0"`
	IfReflush       bool    `jceId:"1"`
	Uin             int64   `jceId:"2"`
	StartIndex      int16   `jceId:"3"`
	FriendCount     int16   `jceId:"4"`
	GroupId         byte    `jceId:"5"`
	IfGetGroupInfo  bool    `jceId:"6"`
	GroupStartIndex byte    `jceId:"7"`
	GroupCount      byte    `jceId:"8"`
	IfGetMsfGroup   bool    `jceId:"9"`
	IfShowTermType  bool    `jceId:"10"`
	Version         int64   `jceId:"11"`
	UinList         []int64 `jceId:"12"`
	AppType         int32   `jceId:"13"`
	IfGetDovId      bool    `jceId:"14"`
	IfGetBothFlag   bool    `jceId:"15"`
	D50Request      []byte  `jceId:"16"`
	D6BRequest      []byte  `jceId:"17"`
	SnsTypeList     []int64 `jceId:"18"`
}

type FriendInfoStruct struct {
	FriendUin   int64  `jceId:"0"`
	GroupId     byte   `jceId:"1"`
	FaceId      int16  `jceId:"2"`
	Remark      string `jceId:"3"`
	Status      byte   `jceId:"5"`
	Level       byte   `jceId:"6"`
	NickName    string `jceId:"14"`
	Network     byte   `jceId:"20"`
	NetworkType int32  `jceId:"24"`
	CardId      []byte `jceId:"41"`
}
