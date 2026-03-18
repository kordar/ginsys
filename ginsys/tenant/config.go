package tenant

var (
	Open   bool
	DB     map[string]string // key => 租户id+业务数据库name
	InitDB string
)
