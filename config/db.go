package config

import (
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/whatsauth"
)

var IteungIPAddress string = os.Getenv("ITEUNGBEV1")

var MongoString string = os.Getenv("MONGOSTRING")

var MariaStringAkademik string = os.Getenv("MARIASTRINGAKADEMIK")

var Ulbimariainfo = atdb.DBInfo{
	DBString: MariaStringAkademik,
	DBName:   "DB_ULBI",
}

var Ulbimariaconn = atdb.MariaConnect(Ulbimariainfo)

var Usertables = [4]whatsauth.LoginInfo{mhs, dosen, user, user1}

var mhs = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "Password",
	Phone:    "Phone",
	Username: "Login",
	Uuid:     "simak_mst_mahasiswa",
	Login:    "2md5",
}

var dosen = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "Password",
	Phone:    "Phone",
	Username: "Login",
	Uuid:     "simak_mst_dosen",
	Login:    "2md5",
}

var user = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "Phone",
	Username: "user_name",
	Uuid:     "simak_besan_users",
	Login:    "2md5",
}

var user1 = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "Phone",
	Username: "user_name",
	Uuid:     "besan_users",
	Login:    "2md5",
}
