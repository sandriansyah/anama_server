package models

type Product struct {
	Id          int64  `gorm:"primarykey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_Product"`
	Description string `gorm:"type:text" json:"deskripsi"`
}
