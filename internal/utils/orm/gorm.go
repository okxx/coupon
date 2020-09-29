package orm

import "gorm.io/gorm"

type Orm struct {
	conn *gorm.DB
}

// connection
func (o *Orm) Connection() error {

	return nil
}