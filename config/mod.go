package config

import (
	"fmt"
	"library/parameter"
)

const Version string = "1.0.0"

const Help string = `
Usage:

	version			print config version

	server-title [.]
	server-addr [.]
	server-password [.]
	server-protocol [.]
	server-cdn [.]
	server-key [.]

	mysql-user [.]
	mysql-password [.]
	mysql-host [.]
	mysql-dbname [.]
	mysql-charset [.]
	mysql-parse-time [.]
	mysql-loc [.]

	cert-cert [.]
	cert-key  [.]

Default:
  MySQL:
	User:      "root"
	Password:  "password"
	Host:      "localhost"
	DBName:    "book"
	Charset:   "utf8mb4"
	ParseTime: "true"
	Loc:       "Local"
  Server:
	Title:    "Library"
	Addr:     ":8080"
	Password: "pass"
	Protocol: "https"
	CDN:      "/"
	Key:      "12345678123456781234567812345678"
  Cert
	Cert: "./cert/cert.pem"
	Key:  "./cert/key.pem"
`

var Args = parameter.GetArgs()
type Arg = parameter.Arg

func DeleteParseModule() {
	parameter.DeleteFromBaseArgs("config")
}

func AddParseModule() {
	parameter.AddToBaseArgs("config", Arg{
		Block:	true,
		Executor: config,
	})
	Args["version"] = Arg{
		Size: 0,
		Block: false,
		Executor: printVersion,
	}
	Args["help"] = Arg{
		Size: 0,
		Block: false,
		Executor: printHelp,
	}
}
func AddParseMySQL() {
	Args["mysql-user"] = Arg{
		Size:     1,
		Block:    false,
		Executor: setMySqlUser,
	}
	Args["mysql-password"] = Arg{
		Size: 1,
		Block: false,
		Executor: setMySqlPassword,
	}
	Args["mysql-host"] = Arg{
		Size: 1,
		Block: false,
		Executor: setMySqlHost,
	}
	Args["mysql-dbname"] = Arg{
		Size: 1,
		Block: false,
		Executor: setMySqlDBName,
	}
	Args["mysql-charset"] = Arg{
		Size: 1,
		Block: false,
		Executor: setMySqlCharset,
	}
	Args["mysql-parse-time"] = Arg{
		Size: 1,
		Block: false,
		Executor: setMySqlParseTime,
	}
	Args["mysql-loc"] = Arg{
		Size: 1,
		Block: false,
		Executor: setMySqlLoc,
	}
}
func AddParseServer() {
	Args["server-title"] = Arg{
		Size: 1,
		Block: false,
		Executor: setServerTitle,
	}
	Args["server-addr"] = Arg{
		Size: 1,
		Block: false,
		Executor: setServerAddr,
	}
	Args["server-password"] = Arg{
		Size: 1,
		Block: false,
		Executor: setServerPassword,
	}
	Args["server-protocol"] = Arg{
		Size: 1,
		Block: false,
		Executor: setServerProtocol,
	}
	Args["server-cdn"] = Arg{
		Size: 1,
		Block: false,
		Executor: setServerCDN,
	}
	Args["server-key"] = Arg{
		Size: 1,
		Block: false,
		Executor: setServerKey,
	}
}

func AddParseCert() {
	Args["cert-cert"] = Arg{
		Size: 1,
		Block: false,
		Executor: setCertCert,
	}
	Args["cert-key"] = Arg{
		Size: 1,
		Block: false,
		Executor: setCertKey,
	}
}

func AddParsePath() {
	Args["path-theme"] = Arg{
		Size: 1,
		Block: false,
		Executor: setPathTheme,
	}
}

func config(args []string) {
	parameter.GenericParseArgs(&Args, args[1:])
}

func printHelp([]string) {
	fmt.Print(Help)
}

func printVersion([]string) {
	fmt.Println(Version)
}

func setServerTitle(arg []string) {
	Server.Title = arg[1]
}
func setServerAddr(arg []string) {
	Server.Addr = arg[1]
}
func setServerPassword(arg []string) {
	Server.Password = arg[1]
}
func setServerProtocol(arg []string) {
	Server.Protocol = arg[1]
}
func setServerCDN(arg []string) {
	Server.CDN = arg[1]
}
func setServerKey(arg []string) {
	Server.Key = arg[1]
}
func setMySqlUser(arg []string) {
	MySQL.User = arg[1]
}
func setMySqlPassword(arg []string) {
	MySQL.Password = arg[1]
}
func setMySqlHost(arg []string) {
	MySQL.Host = arg[1]
}
func setMySqlDBName(arg []string) {
	MySQL.DBName = arg[1]
}
func setMySqlCharset(arg []string) {
	MySQL.Charset = arg[1]
}
func setMySqlParseTime(arg []string) {
	MySQL.ParseTime = arg[1]
}
func setMySqlLoc(arg []string) {
	MySQL.Loc = arg[1]
}
func setCertCert(arg []string) {
	Cert.Cert = arg[1]
}
func setCertKey(arg []string) {
	Cert.Key = arg[1]
}
func setPathTheme(arg []string) {
	Path.Theme = arg[1]
}