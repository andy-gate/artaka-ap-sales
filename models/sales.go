package models

import (
	"encoding/json"
	"time"
)

type Sales struct {
	Id 					string			`json:"id"`
	Create_dtm 			time.Time 		`json:"create_dtm"`
	Sales_id 			string 			`json:"sales_id"`
	User_id 			string 			`json:"user_id"`
	Outlet_id 			string 			`json:"outlet_id"`
	Sales_type 			string			`json:"sales_type"`
	Customer_id 		string 			`json:"customer_id"`
	Products 			json.RawMessage `json:"products"`
	Subtotal 			int 			`json:"subtotal"`
	Total_diskon 		int 			`json:"total_diskon"`
	Total_tax 			json.RawMessage `json:"total_tax"`
	Total_bill 			int 			`json:"total_bill"`
	Payment_method 		string 			`json:"payment_method"`
	Payment_due_date 	string 			`json:"payment_due_date"`
	Total_payment 		int 			`json:"total_payment"`
	Exchange 			int 			`json:"exchange"`
	Notes 				string 			`json:"notes"`
	Total_buy_cost 		int 			`json:"total_buy_cost"`
	Payment_date 		string 			`json:"payment_date"`
	Reward_id 			string 			`json:"Reward_id"`
	Points_redeem 		int 			`json:"points_redeem"`
}

type Product struct {
	Tax					string			`json:"tax"`
	Name				string			`json:"name"`
	Notes				string			`json:"notes"`
	Units				string			`json:"units"`
	Diskon				int				`json:"diskon"`
	Sku_id				string			`json:"sku_id"`
	Variant				string			`json:"variant"`
	Buy_cost			int				`json:"buy_cost"`
	Category			string			`json:"category"`
	Sell_cost			int				`json:"sell_cost"`
	Salestype_up		int				`json:"salestype_up"`
	Number_orders		int				`json:"number_orders"`
	Modifier_price		int				`json:"modifier_price"`
	Modifier_option		string			`json:"modifier_option"`
}

type SalesQuery struct {
	Id 					int			`json:"id"`
}