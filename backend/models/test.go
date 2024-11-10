package models

type Test struct {
    ID   uint   `gorm:"primaryKey"`
    User string `gorm:"type:varchar(100)"`
}

// TableName sets the insert table name for this struct type
func (Test) TableName() string {
    return "test"
}
