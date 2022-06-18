package topen

type TOpen struct {
	SetlID            string
	MdtrtID           string
	PsnName           string
	PsnCertType       string
	PsnCertTypeName   string
	Certno            string
	PsnNo             string
	Gend              string
	GendName          string
	Naty              string
	NatyName          string
	Insutype          string
	InsutypeName      string
	PsnType           string
	PsnTypeName       string
	Begndate          string
	Enddate           string
	SetlTime          string
	MedfeeSumamt      float64
	HiAgreSumfee      float64
	FundPaySumamt     float64
	PsnPay            float64
	AcctPay           float64
	CashPayamt        float64
	FulamtOwnpayAmt   float64
	PreselfpayAmt     float64
	InscpAmt          float64
	DedcStd           float64
	CrtDedc           float64
	ActPayDedc        float64
	HifpPay           float64
	PoolPropSelfpay   float64
	CvlservPay        float64
	HifesPay          float64
	HifmiPay          float64
	HifobPay          float64
	HifdmPay          float64
	MafPay            float64
	OthfundPay        float64
	CvlservFlag       string
	CvlservFlagName   string
	CvlservLv         string
	CvlservLvName     string
	SpPsnType         string
	SpPsnTypeName     string
	SpPsnTypeLv       string
	SpPsnTypeLvName   string
	ClctGrde          string
	ClctGrdeName      string
	FlxempeFlag       string
	FlxempeFlagName   string
	NwbFlag           string
	NwbFlagName       string
	InsuAdmdvs        string
	InsuAdmdvsName    string
	EmpNo             string
	EmpName           string
	EmpType           string
	EmpTypeName       string
	EmpMgtType        string
	EmpMgtTypeName    string
	PayLoc            string
	PayLocName        string
	FixmedinsCode     string
	FixmedinsName     string
	HospLv            string
	HospLvName        string
	MdtrtCertType     string
	MdtrtCertTypeName string
	MedType           string
	MedTypeName       string
	SetlType          string
	SetlTypeName      string
	ClrType           string
	ClrTypeName       string
	ClrWay            string
	ClrWayName        string
	ClrOptins         string
	ClrOptinsName     string
	RefdSetlFlag      string
	RefdSetlFlagName  string
	MdtrtCertNo       string
	DiseNo            string
	DiseName          string
}

type Seltid struct {
	SetlID string
}

// 自定义表明
func (TOpen) TableName() string {
	return "t_open_copy"
}

func (Seltid) TableName() string {
	return "t_open_copy"
}
