package user

import "time"

// JAVA EQUIVALENT -> MODEL
// User object model
type User struct {
	UserID           int       `db:"id" json:"user_id"`
	UserNip          string    `db:"nip" json:"Nip"`
	UserNama         string    `db:"nama_lengkap" json:"Nama_lengkap"`
	UserTanggalLahir time.Time `db:"tanggal_lahir" json:"Tanggal_lahir"`
	UserJabatan      string    `db:"jabatan" json:"Jabatan"`
	UserAlamatEmail  string    `db:"email" json:"Email"`
}
