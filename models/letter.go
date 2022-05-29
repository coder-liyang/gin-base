package models

type Letter struct {
	Model
	FromUserid string `json:"from_userid" gorm:"index"`
	ToUserid string `json:"to_userid" gorm:"index"`
	Content string `json:"content"`
	SendTime string `json:"send_time"`
	DeliveryTime string `json:"delivery_time"` //送达时间（送达时间是在发信的时候就确定的，因此此字段有值并不代表送达）
	WriteStatus int `json:"write_status"` //写信状态 0:草稿 1:完成
	DeliveryStatus int `json:"delivery_status"` //寄信状态 0:未寄信 1:已寄信未收到 2:已寄信已收到
	ReadStatus int `json:"read_status"` //读信状态 0:未读 1:已读
}