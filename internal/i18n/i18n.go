package i18n

import "fmt"

const (
	LangChinese     = "zh"
	LangEnglish     = "en"
	DefaultLanguage = LangChinese

	LabelAccount      = "label.account"
	LabelTraffic      = "label.traffic"
	LabelDuration     = "label.duration"
	LabelBalance      = "label.balance"
	LabelOnlineIP     = "label.online_ip"
	MsgNotLoggedIn    = "msg.not_logged_in"
	MsgUsageFetchFail = "msg.usage_fetch_fail"
	MsgLoginUsageFail = "msg.login_usage_fail"
)

var translations = map[string]map[string]string{
	LangChinese: {
		LabelAccount:      "账号",
		LabelTraffic:      "已用流量",
		LabelDuration:     "已用时长",
		LabelBalance:      "账户余额",
		LabelOnlineIP:     "在线 IP",
		MsgNotLoggedIn:    "当前未登录，请先执行 NEU_IPGW login。",
		MsgUsageFetchFail: "获取在线信息失败",
		MsgLoginUsageFail: "登录成功，但获取在线信息失败",
	},
	LangEnglish: {
		LabelAccount:      "Account",
		LabelTraffic:      "Used Traffic",
		LabelDuration:     "Used Duration",
		LabelBalance:      "Balance",
		LabelOnlineIP:     "Online IP",
		MsgNotLoggedIn:    "Not logged in. Run NEU_IPGW login first.",
		MsgUsageFetchFail: "Failed to fetch usage info",
		MsgLoginUsageFail: "Login succeeded, but failed to fetch usage info",
	},
}

func T(lang, key string) string {
	langMap, ok := translations[lang]
	if !ok {
		langMap = translations[DefaultLanguage]
	}
	if text, ok := langMap[key]; ok {
		return text
	}
	return fmt.Sprintf("[%s]", key)
}
