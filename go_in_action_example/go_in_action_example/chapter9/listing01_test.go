// 这个示例程序展示如何写基础单元测试
package chapter9

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload确认http包的Get函数可以下载内容
func TestDownload(t *testing.T)  {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200
	t.Log("Given the need to test download content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)
			defer resp.Body.Close()
			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}

		}
	}
}

