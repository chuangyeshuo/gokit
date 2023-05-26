// Code generated by log-gen. DO NOT EDIT.
package models

import "time"

type ShopOrder struct {
	id           uint64    `json:"id"`// bigint(20) UNSIGNED2
	uuid         string    `json:"uuid"`// varchar(50) CHARACTER SET utf8mb43
	user_id      uint64    `json:"user_id"`// bigint(20) UNSIGNED3
	payment      uint8     `json:"payment"`// tinyint(4) UNSIGNED3
	fee_amount   uint      `json:"fee_amount"`// int(11) UNSIGNED3
	prepare_time uint      `json:"prepare_time"`// int(11) UNSIGNED3
	paid_time    uint      `json:"paid_time"`// int(11) UNSIGNED3
	goods_detail string    `json:"goods_detail"`// text CHARACTER SET utf8mb43
	status       uint8     `json:"status"`// tinyint(4) UNSIGNED3
	created_ts   time.Time `json:"created_ts"`// timestamp2
	updated_ts   time.Time `json:"updated_ts"`// timestamp3
	logistics_id uint64    `json:"logistics_id"`// bigint(20) UNSIGNED3
	flag         uint      `json:"flag"`// int(11) UNSIGNED3
}

func (s *ShopOrder) Pk() string {
	return "id"
}
func (s *ShopOrder) TableName() string {
	return "shop_order"
}
