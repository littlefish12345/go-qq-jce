package goqqjce

type TroopListRequestV2Simplify struct {
	Uin               int64   `jceId:"0"`
	GetMsfMessageFlag bool    `jceId:"1"`
	Cookie            []byte  `jceId:"2"`
	GroupInfo         []int64 `jceId:"3"`
	GroupFlagExt      byte    `jceId:"4"`
	Version           int32   `jceId:"5"`
	CompanyId         int64   `jceId:"6"`
	VersionNumber     int64   `jceId:"7"`
	GetLongGroupName  bool    `jceId:"8"`
}

type TroopNumStruct struct {
	GroupUin              int64  `jceId:"0"`
	GroupCode             int64  `jceId:"1"`
	Flag                  byte   `jceId:"2"`
	GroupInfoSeqence      int64  `jceId:"3"`
	GroupName             string `jceId:"4"`
	GroupMemo             string `jceId:"5"`
	GroupFlagExt          int64  `jceId:"6"`
	GroupRankSeqence      int64  `jceId:"7"`
	CertificationType     int64  `jceId:"8"`
	ShutUpTimeStamp       int64  `jceId:"9"`
	MyShutUpTimeStamp     int64  `jceId:"10"`
	CmdUinUinFlag         int64  `jceId:"11"`
	AdditionalFlag        int64  `jceId:"12"`
	GroupTypeFlag         int64  `jceId:"13"`
	GroupSecType          int64  `jceId:"14"`
	GroupSecTypeInfo      int64  `jceId:"15"`
	GroupClassExt         int64  `jceId:"16"`
	AppPrivilegeFlag      int64  `jceId:"17"`
	SubscriptionUin       int64  `jceId:"18"`
	MemberNum             int64  `jceId:"19"`
	MemberNumSeqence      int64  `jceId:"20"`
	MemberCardSeqence     int64  `jceId:"21"`
	GroupFlagExt3         int64  `jceId:"22"`
	GroupOwnerUin         int64  `jceId:"23"`
	IsConfGroup           bool   `jceId:"24"`
	IsModifyConfGroupFace bool   `jceId:"25"`
	IsModifyConfGroupName bool   `jceId:"26"`
	CmdUinJoinTime        int64  `jceId:"27"`
	CompanyId             int64  `jceId:"28"`
	MaxGroupMemberNum     int64  `jceId:"29"`
	CmdUinGroupMask       int64  `jceId:"30"`
	HlGuildAppId          int64  `jceId:"31"`
	HlGuildSubType        int64  `jceId:"32"`
	CmdUinRingtoneId      int64  `jceId:"33"`
	CmdUinFlagEx2         int64  `jceId:"34"`
	GroupFlagExt4         int64  `jceId:"35"`
	AppealDeadline        int64  `jceId:"36"`
	GroupFlag             int64  `jceId:"37"`
	GroupRemark           []byte `jceId:"38"`
}
