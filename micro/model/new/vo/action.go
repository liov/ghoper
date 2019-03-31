package vo

type Like struct {
	ID     uint64 `gorm:"primary_key" json:"id"`
	RefID  uint64 `json:"ref_id"`
	Kind   string `gorm:"type:varchar(10)" json:"kind"`
	User   User   `json:"user"`
	UserID uint64 `json:"user_id"`
}

type Collection struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	RefID       uint64 `json:"ref_id"`
	Kind        string `gorm:"type:varchar(10)" json:"kind"`
	FavoritesID uint64 `json:"favorites_id"`
	UserID      uint64 `json:"user_id"`
}

type Favorites struct {
	ID               uint64 `gorm:"primary_key" json:"id"`
	Name             string `gorm:"type:varchar(20)" json:"name"`
	UserID           uint64 `json:"user_id"`
	Count            uint64 `json:"count"`
	FollowUsersCount uint64 `json:"follow_users_count"`
}

type FavoritesOwn struct {
	FollowUsers []User       `json:"follow_users"`
	Collections []Collection `gorm:"many2many:collection_favorites" json:"collections"`
}
