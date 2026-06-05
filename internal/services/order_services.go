package services

import (
	"log"

	"github.com/crisnet-dev/fastfood/internal/repository"
)

func DeleteAllPendingOrderService() error {
	err := repository.DeleteAllPendingOrders()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
