package config

const (
	DefaultNEUDir     = ".NEU"
	DefaultGatewayURL = "https://pass.neu.edu.cn/tpass/login"

	DefaultIPGatewayLoginURL    = "https://ipgw.neu.edu.cn/srun_portal_pc?ac_id=1"
	DefaultIPGatewayLoggedinURL = "https://ipgw.neu.edu.cn/srun_portal_success?ac_id=1"

	DefaultIPGatewayAdminLoginURL    = "https://ipgw.neu.edu.cn:8800/sso/neusoft/index"
	DefaultIPGatewayAdminLoggedinURL = "https://ipgw.neu.edu.cn:8800/home"

	DefaultIPGatewayTicketURL = DefaultGatewayURL + "?service=http://ipgw.neu.edu.cn/srun_portal_sso?"
	DefaultIPGatewayAPIURL    = "https://ipgw.neu.edu.cn/v1"
	DefaultIPGatewayLogoutURL = "https://ipgw.neu.edu.cn/cgi-bin/srun_portal?action=logout&username="

	DefaultLogLevel  = "info"
	DefaultLogFormat = "plain"

	DefaultPublicKeyStr = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnjA28DLKXZzxbKmo9/1WkVLf1mr+wtLXLXt6sC4WiBCtsbzF5ewm7ARZeAdS3iZtqlYPn6IcUoOw42H8nAK/tfFcIb6dZ1K0atn0U39oWCGPzYuKtLJeMuNZiDXVuAXtojrckOjLW9B3gUnaNGLuIx0fYe66l0o9WjU2cGLNZQfiIxs2h00z1EA9IdSnVxiVQWSD+lsP3JZXh2TT287la4Y4603SQNKTK/QvXfcmccwTEd1IW6HwGxD6QrkInBiHisKWxmveN7UDSaQRZ/J97G0YC32pD38WT53izXeK0p/kU/X37VP555um1wVWFvPIuc9I7gMP1+hq5a+X6c++tQIDAQAB"
)
