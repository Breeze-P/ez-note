package constants

var (
	EtcdAddress     = GetIp("EtcdIp") + ":2379"
	MySQLDefaultDSN = "note:note@tcp(" + GetIp("MysqlIp") + ":3306)/note?charset=utf8&parseTime=True&loc=Local"

	ApiPort  = "8080"
	UserPort = "8888"
	NotePort = "8889"
)

const (
	NoteTableName           = "note"
	UserTableName           = "user"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	ApiServiceName          = "demoapi"
	NoteServiceName         = "demonote"
	UserServiceName         = "demouser"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)
