
package models

import "gorm.io/gorm"

type Cart struct {
    ID      uint `gorm:"primaryKey"`
    UserID  uint
    Items   []Item `gorm:"many2many:cart_items"`
}
