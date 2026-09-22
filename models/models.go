package models

import "time"

// Tabel: kelas (id, nama_kelas)
type Kelas struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NamaKelas string    `gorm:"not null" json:"nama_kelas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi: 1 Kelas punya banyak Siswa
	Siswas []Siswa `gorm:"foreignKey:IdKelas" json:"siswas,omitempty"`
}

func (Kelas) TableName() string {
	return "kelas"
}

// Tabel: siswas (id, nama, id_kelas)
type Siswa struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"not null" json:"nama"`
	IdKelas   uint      `gorm:"not null" json:"id_kelas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	Kelas        Kelas        `gorm:"foreignKey:IdKelas" json:"kelas,omitempty"`
	KartuPelajar *KartuPelajar `gorm:"foreignKey:IdSiswa" json:"kartu_pelajar,omitempty"`
}

func (Siswa) TableName() string {
	return "siswas"
}

// Tabel: kartu_pelajars (id, nomor_kartu, id_siswa)
type KartuPelajar struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	NomorKartu string    `gorm:"unique;not null" json:"nomor_kartu"`
	IdSiswa    uint      `gorm:"unique;not null" json:"id_siswa"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relasi
	Siswa Siswa `gorm:"foreignKey:IdSiswa" json:"siswa,omitempty"`
}

func (KartuPelajar) TableName() string {
	return "kartu_pelajars"
}
