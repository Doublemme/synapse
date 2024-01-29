package types

import "gorm.io/gorm"

type DatabaseConnFunc func() gorm.Dialector
