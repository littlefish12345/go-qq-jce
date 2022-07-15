package goqqjce

type ClientRegisterPackStruct struct {
	Uin                int64  `jceId:"0"`
	Bid                int64  `jceId:"1"`
	ConnType           byte   `jceId:"2"`
	Other              string `jceId:"3"`
	Status             int32  `jceId:"4"`
	OnlinePush         bool   `jceId:"5"`
	IsOnline           bool   `jceId:"6"`
	IsShowOnline       bool   `jceId:"7"`
	KickPC             bool   `jceId:"8"`
	KickWeak           bool   `jceId:"9"`
	TimeStamp          int64  `jceId:"10"`
	IOSVersion         int64  `jceId:"11"`
	NetType            byte   `jceId:"12"`
	BuildVersion       string `jceId:"13"`
	RegType            byte   `jceId:"14"`
	DeviceParam        []byte `jceId:"15"`
	Guid               []byte `jceId:"16"`
	LocaleId           int32  `jceId:"17"`
	SilentPush         byte   `jceId:"18"`
	DeviceName         string `jceId:"19"`
	DeviceType         string `jceId:"20"`
	OSVersion          string `jceId:"21"`
	OpenPush           bool   `jceId:"22"`
	LargeSeq           int64  `jceId:"23"`
	LastWatchStartTime int64  `jceId:"24"`
	OldSsoIp           int64  `jceId:"26"`
	NewSsoIp           int64  `jceId:"27"`
	ChannelNum         string `jceId:"28"`
	CpId               int64  `jceId:"29"`
	VendorName         string `jceId:"30"`
	VendorOSName       string `jceId:"31"`
	IOSIdfa            string `jceId:"32"`
	B769Request        []byte `jceId:"33"`
	IsSetStatus        bool   `jceId:"34"`
	ServerBuf          []byte `jceId:"35"`
	SetMute            bool   `jceId:"36"`
	ExtOnlineStatus    int64  `jceId:"38"`
	BatteryStatus      int32  `jceId:"39"`
}

type ClientRegisterResponsePackStruct struct {
	Uin                      int64  `jceId:"0"`
	Bid                      int64  `jceId:"1"`
	ReplyCode                byte   `jceId:"2"`
	Result                   string `jceId:"3"`
	ServerTime               int64  `jceId:"4"`
	LogQQ                    byte   `jceId:"5"`
	NeedKik                  bool   `jceId:"6"`
	UpdateFlag               byte   `jceId:"7"`
	TimeStamp                int64  `jceId:"8"`
	CrashFlag                byte   `jceId:"9"`
	ClientIp                 string `jceId:"10"`
	ClientPort               int32  `jceId:"11"`
	HelloInterval            int32  `jceId:"12"`
	LargeSeq                 int32  `jceId:"13"`
	LargeSeqUpdate           byte   `jceId:"14"`
	D769RspBody              []byte `jceId:"15"`
	Status                   int32  `jceId:"16"`
	ExtOnlineStatus          int64  `jceId:"17"`
	ClientBatteryGetInterval int64  `jceId:"18"`
	ClientAutoStatusInterval int64  `jceId:"19"`
}
