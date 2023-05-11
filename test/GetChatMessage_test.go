package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/yacw-team/yacw/routes"
	"github.com/yacw-team/yacw/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetChatMessageCorrectExample(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{"apiKey":"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","chatId":"1"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getMessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"chatId":"1","messages":[{"type":"user","content":"我想知道如何保持健康"},{"type":"assistant","content":"保持身体健康的几个关键因素包括饮食、锻炼和睡眠：\n\n1. 饮食：保持均衡的饮食很重要，包括适量的蛋白质、碳水化合物、脂肪、纤维和维生素。建议遵循健康的饮食原则，如减少加工食品、甜食和盐的摄入。\n\n2. 锻炼：每周至少进行150分钟的中度运动或75分钟的高强度运动。运动可以增强心肺功能、强化肌肉和骨骼，并有助于控制体重和心理压力。\n\n3. 睡眠：保持规律的睡眠模式至关重要，成年人每晚需7-9小时的睡眠。睡眠不足会影响身体的正常功能，增加患疾病的风险。\n\n除了这些因素外，还应该避免吸烟、限制饮酒、保持良好的心理健康、和经常进行身体检查以便及早发现疾病。"}]}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageMissingLength(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{"apiKey":"sk-hISgKGQQ5cZNGHZxbQFXT3BlbkFJ8vyxitPPXM6oqfgTeNl","chatId":"1"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getMessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageExcessiveLength(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{"apiKey":"sk-hISgKGQQ5cZNGHZxbQFXT3BlbkFJ8vyxitPPXM6oqfgTeNlx1","chatId":"1"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getMessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageFormatMixing(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{"apiKey":"sk-hISgKGQQ5cZNGHZxbQFXT3BlbkFJ8vyxitPPXM6oqfgTeNl我","chatId":"1"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getMessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestGetChatMessageApiKeyNull(t *testing.T) {
	utils.InitDBTest()
	var jsonStr = []byte(`{"apiKey":"","chatId":"1"}`)
	req, err := http.NewRequest("POST", "/v1/chat/getMessage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	routes.SetupRouter().ServeHTTP(rr, req)
	expected := `{"errCode":"3004"}`
	assert.Equal(t, expected, rr.Body.String())
}
