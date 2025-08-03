
package models

import "gorm.io/gorm"

type Order struct {
    ID     uint `gorm:"primaryKey"`
    UserID uint
    Items  []Item `gorm:"many2many:order_items"`
}
