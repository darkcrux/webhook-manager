package repository

// import (
// 	"github.com/jinzhu/gorm"

// 	"github.com/darkcrux/notification-manager/internal/component/notification"
// )

// type GormNotificationRepository struct {
// 	db *gorm.DB
// }

// func NewGormNotificationRepository(db *gorm.DB) notification.Notification {
// 	return &GormNotificationRepository{db}
// }

// func (repo *GormNotificationRepository) Save(tx *notification) (id int, err error) {
// 	if err = repo.db.Save(tx).Error; err != nil {
// 		// what went wrong?
// 		return
// 	}
// 	id = *tx.ID
// 	return
// }

// func (repo *GormNotificationRepository) GetByTxIDAndCustomerID(txID, customerID int) (wh *notification.notification, err error) {

// 	return
// }
