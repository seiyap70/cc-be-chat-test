package cmd_handler

import (
	"reflect"
	"testing"
)

func TestParseRequest(t *testing.T) {
	testRequest := []struct {
		jsonStr    string
		targetType reflect.Type
	}{
		{jsonStr: `{"cmd":"join_room", "id": "1", "data":{"userName":"seiya", "roomName":"chatroom1"}}`, targetType: reflect.TypeOf(&joinRoom{})},
		{jsonStr: `{"cmd":"left_room", "id": "1", "data":{}}`, targetType: reflect.TypeOf(&leftRoom{})},
		{jsonStr: `{"cmd":"send_chat", "id": "1", "data":{"text":"xxxxxxx"}}`, targetType: reflect.TypeOf(&sendChatMessage{})},
		{jsonStr: `{"cmd":"popular_word", "id": "1", "data":{}}`, targetType: reflect.TypeOf(&popularWord{})},
		{jsonStr: `{"cmd":"stats_user", "id": "1", "data":{"userName":"seiya"}}`, targetType: reflect.TypeOf(&statsUser{})},
	}
	for _, test := range testRequest {
		request, err := ParseRequest([]byte(test.jsonStr))
		if err != nil {
			t.Fatalf("parse %s failed, error: %s", test.jsonStr, err)
		}
		if reflect.TypeOf(request.Data) != test.targetType {
			t.Fatalf("test %s, target %v, get %v", test.jsonStr, test.targetType, reflect.TypeOf(request.Data))
		}
	}
}
