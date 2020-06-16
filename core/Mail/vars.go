package Mail

var (
	MailSetting Setting
)

type Setting struct {
	Sender   string `yaml:"Sender"`
	Nickname string `yaml:"Nickname"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
}
