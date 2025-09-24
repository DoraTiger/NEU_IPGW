package version

// the build version of NEU_IPGW
var (
	BuildVersion = VersionDefault
	BuildTime    = TimeDefault
	BuildRepo    = RepoDefault
)

// if not use make to build project, use the following default variables for version command
const (
	VersionDefault = "unknown"
	TimeDefault    = "1970-01-01T00:00:00+0000"
	RepoDefault    = "github.com/DoraTiger/NEU_IPGW"
)
