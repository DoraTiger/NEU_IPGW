package config

const (
	DefaultNEUDir     = ".NEU"
	DefaultGatewayURL = "https://pass.neu.edu.cn/tpass/login"

	DefaultConfigSubPath = ".config/doratiger/neu-ipgw"
	CredentialFileName   = "credentials.enc"
	EnvConfigDir         = "NEU_IPGW_CONFIG_DIR"
	EnvMasterKey         = "NEU_IPGW_MASTER_KEY"
	EnvLanguage          = "NEU_IPGW_LANG"
	// DefaultMasterKey is public and meant for baseline obfuscation only.
	DefaultMasterKey = "doratiger-neu-ipgw-default-master-key-v1"

	DefaultIPGatewayLoginURL    = "https://ipgw.neu.edu.cn/srun_portal_pc?ac_id=1"
	DefaultIPGatewayLoggedinURL = "https://ipgw.neu.edu.cn/srun_portal_success?ac_id=1"
	DefaultIPGatewayStatusURL   = "https://ipgw.neu.edu.cn/cgi-bin/rad_user_info"

	DefaultIPGatewayAdminLoginURL    = "https://ipgw.neu.edu.cn:8800/sso/neusoft/index"
	DefaultIPGatewayAdminLoggedinURL = "https://ipgw.neu.edu.cn:8800/home"

	DefaultIPGatewayTicketURL = DefaultGatewayURL + "?service=http://ipgw.neu.edu.cn/srun_portal_sso?"
	DefaultIPGatewayAPIURL    = "https://ipgw.neu.edu.cn/v1"
	DefaultIPGatewayLogoutURL = "https://ipgw.neu.edu.cn/cgi-bin/srun_portal?action=logout&username="

	DefaultSSOBaseURL   = "https://pass.neu.edu.cn/tpass/login"
	DefaultSSOLogoutURL = "https://pass.neu.edu.cn/tpass/logout"
	DefaultPayCASURL    = "https://pay.neu.edu.cn/drCasLogin"
	DefaultPayReferer   = "https://pay.neu.edu.cn/"

	DefaultElectricityQueryURL    = "https://pay.neu.edu.cn/queryDefaultRoominfo"
	DefaultElectricityRoomInfoURL = "https://pay.neu.edu.cn/queryRommInfo"
	DefaultFactoryCode            = "E039"
	DefaultElectricityPrice       = 0.51

	DefaultLogLevel  = "info"
	DefaultLogFormat = "plain"

	DefaultPublicKeyStr = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnjA28DLKXZzxbKmo9/1WkVLf1mr+wtLXLXt6sC4WiBCtsbzF5ewm7ARZeAdS3iZtqlYPn6IcUoOw42H8nAK/tfFcIb6dZ1K0atn0U39oWCGPzYuKtLJeMuNZiDXVuAXtojrckOjLW9B3gUnaNGLuIx0fYe66l0o9WjU2cGLNZQfiIxs2h00z1EA9IdSnVxiVQWSD+lsP3JZXh2TT287la4Y4603SQNKTK/QvXfcmccwTEd1IW6HwGxD6QrkInBiHisKWxmveN7UDSaQRZ/J97G0YC32pD38WT53izXeK0p/kU/X37VP555um1wVWFvPIuc9I7gMP1+hq5a+X6c++tQIDAQAB"
)
