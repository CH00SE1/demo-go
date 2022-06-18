package main

import (
	"fmt"
	"math/rand"
	"os"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var intLetters = []rune("0123456789")

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// 创建测试数据
func createFile(text *string) {

	create, err := os.Create("C:\\Users\\Administrator\\Desktop\\模拟数据.json")

	if err != nil {
		fmt.Println(err)
	}

	defer create.Close()

	writeString, err := create.WriteString("[" + *text + "]")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(writeString)
}

func main() {

	jsonStr := ""

	var text *string

	text = &jsonStr

	num := 20000

	for i := 0; i < num; i++ {

		str := "{\"setlId\":\"36284113" + RandomString(5, defaultLetters) + "\",\"mdtrtId\":\"56741021\",\"psnName\":\"陈雪英\",\"psnCertType\":\"01\",\"psnCertTypeName\":\"居民身份证（户口簿）\",\"certno\":\"430121193310307923\",\"psnNo\":\"43000020100030375675\",\"gend\":\"2\",\"gendName\":\"女\",\"naty\":\"01\",\"natyName\":\"汉族\",\"insutype\":\"390\",\"insutypeName\":\"城乡居民基本医疗保险\",\"psnType\":\"16\",\"psnTypeName\":\"居民(老)\",\"begndate\":\"2022-01-01\",\"enddate\":\"2022-01-01\",\"setlTime\":\"2022-01-01 08:09:49\",\"medfeeSumamt\":576.07,\"hiAgreSumfee\":576.07,\"fundPaySumamt\":210.00,\"psnPay\":366.07,\"acctPay\":0.00,\"cashPayamt\":366.07,\"fulamtOwnpayAmt\":0.00,\"preselfpayAmt\":0.00,\"inscpAmt\":576.07,\"dedcStd\":0.00,\"crtDedc\":0.00,\"actPayDedc\":0.00,\"hifpPay\":210.00,\"poolPropSelfpay\":0.7000,\"cvlservPay\":0.00,\"hifesPay\":0.00,\"hifmiPay\":0.00,\"hifobPay\":0.00,\"hifdmPay\":0.00,\"mafPay\":0.00,\"othfundPay\":0.00,\"cvlservFlag\":\"0\",\"cvlservFlagName\":\"否\",\"cvlservLv\":null,\"cvlservLvName\":null,\"spPsnType\":null,\"spPsnTypeName\":null,\"spPsnTypeLv\":null,\"spPsnTypeLvName\":null,\"clctGrde\":null,\"clctGrdeName\":null,\"flxempeFlag\":\"0\",\"flxempeFlagName\":\"否\",\"nwbFlag\":null,\"nwbFlagName\":null,\"insuAdmdvs\":\"430105\",\"insuAdmdvsName\":\"开福区\",\"empNo\":\"430105040300\",\"empName\":\"迎宾路社区\",\"empType\":\"01\",\"empTypeName\":\"01\",\"empMgtType\":null,\"empMgtTypeName\":null,\"payLoc\":\"2\",\"payLocName\":\"医疗机构\",\"fixmedinsCode\":\"P43010201038\",\"fixmedinsName\":\"湖南达嘉维康医药产业股份有限公司五一路分店\",\"hospLv\":\"0\",\"hospLvName\":\"0\",\"mdtrtCertType\":\"02\",\"mdtrtCertTypeName\":\"居民身份证\",\"medType\":\"14\",\"medTypeName\":\"门诊慢特病\",\"setlType\":\"2\",\"setlTypeName\":\"联网结算\",\"clrType\":\"11\",\"clrTypeName\":\"门诊\",\"clrWay\":\"1\",\"clrWayName\":\"按项目\",\"clrOptins\":\"430199\",\"clrOptinsName\":\"长沙市市本级\",\"refdSetlFlag\":\"0\",\"refdSetlFlagName\":\"否\",\"mdtrtCertNo\":\"430121193310307923\",\"diseNo\":\"M01603\",\"diseName\":\"糖尿病并发症\"}"

		if i == num-1 {
			jsonStr += str
		} else {
			jsonStr += (str + ",")
		}

	}

	createFile(text)

}
