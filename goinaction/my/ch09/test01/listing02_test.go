// 基本的表组测试
package listing02

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statucCode int
	}{
		{
			"http://www.baidu.com",
			http.StatusOK,
		},
		{
			"https://github.com/MaWenlong1233",
			http.StatusNotFound,
		},
	}
	t.Log("测试下载多个url")
	{
		for _, u := range urls {
			t.Logf("\t When checking %s for status code %d ", u.url, u.statucCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\t Should be able to make the Get call.", ballotX, err)
				}
				t.Log("\t\tShould be able to make the Get call.", checkMark)
				defer resp.Body.Close()
				if resp.StatusCode == u.statucCode {
					t.Logf("\t\tShould receive a  %d status. %v", u.statucCode, checkMark)
				} else {
					t.Errorf("\t\tShould receive a %d  status. %v %v", u.statucCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}
