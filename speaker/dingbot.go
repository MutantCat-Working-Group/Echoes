package speaker

import "github.com/blinkbean/dingtalk"

func GetDingBot(ding_token, ding_secret string) *dingtalk.DingTalk {
	return dingtalk.InitDingTalkWithSecret(ding_token, ding_secret)
}
