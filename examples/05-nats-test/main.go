package main

import (
	"github.com/joek/dkom/ysfbot"
	"gobot.io/x/gobot"
)

func main() {
	var order *ysfbot.Order

	ysf := createNewYsf()
	master := gobot.NewMaster()

	ysf.On(ysfbot.OrderCreated, func(m interface{}) {
		o := m.(ysfbot.Order)
		order = &o
		ysf.Publish(ysfbot.UpdateOrderStatus, ysfbot.OrderStatus{
			ID:     o.ID,
			Status: "In Preperation",
		})
	})

	master.AddRobot(ysf)
	master.Start()
}

func createNewYsf() ysfbot.YSFRobot {
	natsURL := "127.0.0.1:32240"
	natsClientID := 1234
	natsUser := "natsuser"
	natsPassword := "xqwwVp9C&Hn6jXcux4r)vq"
	orderUpdateURL := "http://beershop.local/orderStatusUpdate"
	return ysfbot.NewYSFRobot(natsURL, natsClientID, natsUser, natsPassword, orderUpdateURL)
}
