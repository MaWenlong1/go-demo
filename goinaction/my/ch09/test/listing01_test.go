// 基础单元测试
package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "http://www.baidu.com"
	statucCode := 200
	t.Log("测试http.Get()")
	{
		t.Logf("\t 检查url：%s 的状态码是否是%d \n", url, statucCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\t Should be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()
			if resp.StatusCode == statucCode {
				t.Logf("\t\tShould receive a  %d status. %v", statucCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a %d  status. %v %v", statucCode, ballotX, resp.StatusCode)
			}
		}
	}
}
