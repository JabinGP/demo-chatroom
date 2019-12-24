package pojo

import "time"

// Message 消息实体，对应表message
type Message struct {
	ID         int64
	SenderID   int64
	ReceiverID int64
	Content    string
	SendTime   int64
	CreatedAt  time.Time `xorm:"created"` // 这个Field将在Insert时自动赋值为当前时间
	UpdatedAt  time.Time `xorm:"updated"` // 这个Field将在Insert或Update时自动值为当前时间
	DeletedAt  time.Time `xorm:"deleted"` // 如果带DeletedAt这个字段和标签，xorm删除时自动软删除
}
