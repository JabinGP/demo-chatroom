package pojo

// MessageWithUser 消息包含用户的实体
type MessageWithUser struct {
	Message  Message `xorm:"extends"`
	Sender   User    `xorm:"extends"`
	Receiver User    `xorm:"extends"`
}

// TableName 指定表名
func (MessageWithUser *MessageWithUser) TableName() string {
	return "message"
}
