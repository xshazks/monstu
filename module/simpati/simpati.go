package simpati

import (
	"fmt"

	"iteung/helper"
	"iteung/helper/wacipher"
	"iteung/model"
)

func GetLoginInfofromPhoneNumber(phonenumber string) (response model.Simpati) {
	fmt.Println("phonenumber : " + phonenumber)
	if phonenumber != "" {
		response.Username = GetUsernamefromPhonenumber(phonenumber)
		fmt.Println("username : " + response.Username)
		if response.Username != "" {
			response.Password = UpdatePasswordfromUsername(response.Username)
			fmt.Println("password : " + response.Password)
			if response.Password != "" {
				response.Login = "Login"
				response.Userid = GetUserIdfromUsername(response.Username)
			}

		}
	}
	return response
}

func GetUsernamefromPhonenumber(phone_number string) (username string) {
	username = GetUsernamefromPhonenumberInTable(phone_number, "simak_mst_mahasiswa")
	fmt.Println(username)
	if username == "" {
		username = GetUsernamefromPhonenumberInTable(phone_number, "simak_mst_dosen")

	}
	return username
}

func GetUsernamefromPhonenumberInTable(phone_number string, tabel string) (username string) {
	db := helper.MariaConnect("db_ulbi")
	err := db.QueryRow("select Login from "+tabel+" where Handphone = ?", phone_number).Scan(&username)
	if err != nil {
		fmt.Printf("GetUsernamefromPhonenumberInTable %v: %v\n", tabel, err)
	}
	return username
}

func GetHashPasswordfromUsername(username string) (hashpassword string) {
	db := helper.MariaConnect("db_ulbi")
	err := db.QueryRow("select user_password from simak_besan_users where user_name = ?", username).Scan(&hashpassword)
	if err != nil {
		fmt.Printf("GetHashPasswordfromUsername: %v\n", err)
	}
	return hashpassword
}

func UpdatePasswordfromUsername(username string) (newPassword string) {
	newPassword = wacipher.RandomString(10)
	var temp interface{}
	var err error
	db := helper.MariaConnect("db_ulbi")

	err = db.QueryRow("update simak_mst_mahasiswa set Password = MD5(MD5(?)) where Login = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	err = db.QueryRow("update simak_mst_dosen set Password = MD5(MD5(?)) where Login = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	err = db.QueryRow("update simak_besan_users  set user_password = MD5(MD5(?)) where user_name = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	err = db.QueryRow("update besan_users  set user_password = MD5(MD5(?)) where user_name = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	return newPassword
}

func GetUserIdfromUsername(username string) (userid string) {
	db := helper.MariaConnect("db_ulbi")
	err := db.QueryRow("select user_id from simak_besan_users where user_name = ?", username).Scan(&userid)
	if err != nil {
		fmt.Printf("GetHashPasswordfromUsername: %v\n", err)
	}
	return userid
}

func GetJadwalKuliah(kodedosen string) (jadwal []model.JadwalKuliah) {
	db := helper.MariaConnect("db_ulbi")
	rows, _ := db.Query("select JadwalID, Nama, NamaKelas, HariID, JamMulai, JamSelesai, RuangID, Kehadiran from simak_trn_jadwal where DosenID = ? and TahunID = ?", kodedosen, getTahunID())
	jadwal = []model.JadwalKuliah{}
	for rows.Next() {
		var r model.JadwalKuliah
		rows.Scan(&r.JadwalID, &r.Nama, &r.NamaKelas, &r.HariID, &r.JamMulai, &r.JamSelesai, &r.RuangID, &r.Kehadiran)
		jadwal = append(jadwal, r)
	}
	return jadwal
}

func getTahunID() (TahunID model.TahunID) {
	db := helper.MariaConnect("db_ulbi")
	db.QueryRow("SELECT TahunID FROM simak_mst_tahun where NA = 'N' group by TahunID order by TahunID DESC limit 1").Scan(&TahunID)
	return TahunID

}
