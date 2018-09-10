package _rsa

import (
	"fmt"
	"encoding/base64"
	"testing"
)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCGelAapImDcWTfOU3xFQedpRq8
EoLcoPRa0Xu4GsmtaqS34SgmLd+NwA9VfzB0vLdcwLFEV9lsnr2aScyTVr0feac6
xVIWI7SHmQLGlCw4Akb5kYpSdpHxuxfiTLYgWeAfrDw4idwvKopCmq+XQlpbzUVc
+Af0egG/GyV1bQMy3wIDAQAB
-----END PUBLIC KEY-----
`)


var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCGelAapImDcWTfOU3xFQedpRq8EoLcoPRa0Xu4GsmtaqS34Sgm
Ld+NwA9VfzB0vLdcwLFEV9lsnr2aScyTVr0feac6xVIWI7SHmQLGlCw4Akb5kYpS
dpHxuxfiTLYgWeAfrDw4idwvKopCmq+XQlpbzUVc+Af0egG/GyV1bQMy3wIDAQAB
AoGAd5OdVQOHr5SOEJvg/YUon3onWqLPMCYDAhqR+3P+bzCLxWH1lHVn+qCUQf71
9Ge2WEiTl+TR1e8VQ3Np8H3JbjmH9hHaeZN0k3fQ3LfCqB9c/J9LvSlbScWPoqZU
+fTdrVG62s6330D94oeUOJdnkI7vwA5kIddAgtFhOEXP8oECQQCXGv9TcdmtNZby
8D6Zd6XGqRLNHIRKuS4GjC5Txiy0XGGibUGtS/R5VUEgORsryJJrE3Rq8Li5KCDX
h16di85BAkEA49Rn7cXPhIWdACTvcs6u+vgualGgeslM+Q59Sxd+9hir+z/ZEyN7
RzaHILzVORoVOvSk5zvKZ4SBFQes73j5HwJAdE05O/ai/igDGNVEuUZX+AHmEKzk
Pcct36hBeAVOHzwDgcrqBAI0FrdBuxV5DkgAOh3tOuowo4J5VKpCbqxOgQJBAInQ
I2nN/UBBfn3m2b0NzwTa+WCwra14dQo2vI2e0drVg7rAnXOZQ+oIzuZ7s1MONwdp
kxKYtU+29EofUXmuKKsCQHAbpFRFH8sOB/F3UDFjo/yakRfiEBggikcRFOJ1jY6k
C2aEY5hDv0f6EXWuyg7iw/9320wV7OPhKCW6J9FANvA=
-----END RSA PRIVATE KEY-----
`)


func Test_RSA(t *testing.T){
	raw := "ea4d35ed8fe19e11b17917ad273021_11_22_x1dd"
	encryptBytes, err := RsaEncrypt([]byte(raw), publicKey)
	if err!=nil{
		t.Error("RsaEncrypt error", err)
		return
	}
	fmt.Println("RsaEncrypt success", base64.StdEncoding.EncodeToString(encryptBytes))
	decryptBytes, err:= RsaDecrypt(encryptBytes, privateKey)
	if err!=nil{
		t.Error("RsaDecrypt error", err)
		return
	}
	if raw == string(decryptBytes){
		t.Log("RsaDecrypt success", raw, string(decryptBytes))
	}else{
		t.Error("RsaDecrypt failed", raw, string(decryptBytes))
	}

}
