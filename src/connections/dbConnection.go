package connections

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"text/template"
)

var database *gorm.DB
var err error

func DatabaseInit() {
	host := "ep-fancy-sunset-61348543-pooler.ap-southeast-1.aws.neon.tech"
	password := "k3cGLHqoe1XD"
	user := "mochamadiqbalaljihad"
	dbName := "upay_db"
	port := "5432"

	sentence := "host={{.hostt}} user={{.usert}} password={{.passwordt}} dbname={{.dbnamet}} port={{.portt}} sslmode=require"
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(sentence)).Execute(b, map[string]interface{}{
		"hostt":     host,
		"usert":     user,
		"passwordt": password,
		"dbnamet":   dbName,
		"portt":     port})

	database, err = gorm.Open(postgres.Open(b.String()), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}
