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
)
