
package models

import "gorm.io/gorm"

type Item struct {
    ID       uint    `gorm:"primaryKey"`
    Name     string
    Price    float64
}
