package models

import (
	"time"
)

type CryptoPrice struct {
	Id        string    `json:"id,omitempty,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     float32   `json:"price,omitempty"`
	Currency  string    `json:"currency,omitempty"`
	TimeCheck time.Time `json:"timeCheck,omitempty"`
}

//func (b CryptoPrice) Validate() error {
//	if len([]rune(b.Name)) > 100 {
//		return fmt.Errorf("name len more than 100")
//	} else if len([]rune(b.Description)) > 500 {
//		return fmt.Errorf("description len more than 500")
//	} else if len([]rune(b.AuthorId)) > 100 {
//		return fmt.Errorf("description len more than 500")
//	} else {
//		return nil
//	}
//}
