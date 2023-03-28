package controller

import (
	"app/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type preferResponse struct {
	code int                    //httpステータスコード
	body map[string]interface{} //帰ってくる文字列
}

// controller.CreateUserAction()のテスト
func TestIndex(t *testing.T) {

	tests := []struct {
		book model.Book
		want preferResponse
	}{
		{
			model.Book{
				Title: "title1",
				Body:  "body1",
			},
			preferResponse{
				code: http.StatusOK,
				body: map[string]interface{}{
					"Title": "success: title1",
					"Body":  "success: body1",
				},
			},
		},
	}
	for i, tt := range tests {
		//レスポンス
		//ここに帰ってくる
		response := httptest.NewRecorder()
		//コンテキストを作成
		c, _ := gin.CreateTestContext(response)

		Index(c)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

		//ステータスコードがおかしいもしくは帰ってきたメッセージが想定と違えばダメ
		if response.Code != tt.want.code {
			t.Errorf("%d番目のテストが失敗しました。想定返却コード：%d, 実際の返却コード：%d", i+1, tt.want.code, response.Code)
		} else {
			//実際に帰ってきたレスポンスの中に想定された値が入っているかどうか
			for key := range tt.want.body {
				//値の存在チェック
				if _, exist := responseBody[key]; exist {

					//値の内容チェック
					if responseBody[key] != tt.want.body[key] {
						t.Errorf("%d番目のテストが失敗しました。想定されたキー「%s」の値:%s, 実際に返却された値:%s", i+1, key, tt.want.body[key], responseBody[key])
					} // else{
					//クリアはここだけ
					// }

				} else {
					t.Errorf("%d番目のテストが失敗しました。想定された「%s」がレスポンスボディに入っていません。", i+1, key)
				}
			}
		}
	}
}
