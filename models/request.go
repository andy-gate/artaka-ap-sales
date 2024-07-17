package models

type APRequest struct {
	Store			[]APRequestStore	`json:"store"`
}

type APRequestStore struct {
	Store_id			string			`json:"store_id"`
	Transactions		[]Trx			`json:"transactions"`
}

type Trx struct {
	Invoice_no					string			`json:"invoice_no"`
	Trans_date					string			`json:"trans_date"`
	Trans_time					string			`json:"trans_time"`
	Sequence_unique				string			`json:"sequence_unique"`
	Item_name					string			`json:"item_name"`
	Item_code					string			`json:"item_code"`
	Item_barcode				string			`json:"item_barcode"`
	Item_cat_name				string			`json:"item_cat_name"`
	Item_cat_code				string			`json:"item_cat_code"`
	Item_qty					string			`json:"item_qty"`
	Item_unit					string			`json:"item_unit"`
	Item_price_per_unit			string			`json:"item_price_per_unit"`
	Item_discount				string			`json:"item_discount"`
	Item_price_amount			string			`json:"item_price_amount"`
	Item_vat					string			`json:"item_vat"`
	Item_tax					string			`json:"item_tax"`
	Item_total_discount			string			`json:"item_total_discount"`
	Item_total_price_amount		string			`json:"item_total_price_amount"`
	Item_total_vat				string			`json:"item_total_vat"`
	Item_total_tax				string			`json:"item_total_tax"`
	Item_total_service_charge	string			`json:"item_total_service_charge"`
	Invoice_tax					string			`json:"invoice_tax"`
	Invoice_discount			string			`json:"invoice_discount"`
	Transaction_amount			string			`json:"transaction_amount"`
	Currency					string			`json:"currency"`
	Rate						string			`json:"rate"`
	Payment_type				string			`json:"payment_type"`
	Payment_by					string			`json:"payment_by"`
	Username					string			`json:"username"`
	Buyer_barcode				string			`json:"buyer_barcode"`
	Buyer_name					string			`json:"buyer_name"`
	Buyer_flight_no				string			`json:"buyer_flight_no"`
	Buyer_destination			string			`json:"buyer_destination"`
	Buyer_nationality			string			`json:"buyer_nationality"`
	Remark						string			`json:"remark"`
	Tax_id						string			`json:"tax_id"`
	Payment_name				string			`json:"payment_name"`
	Payment_time				string			`json:"payment_time"`
	Distance					string			`json:"distance"`
	Journey_time				string			`json:"journey_time"`
}

type APResponse struct {
	Status				bool				`json:"status"`
	Success_insert		int					`json:"success_insert"`
	Failed_insert		int					`json:"failed_insert"`
	Success_data		[]Trx				`json:"success_data"`
	Failed_data			[]Trx				`json:"failed_data"`
	Failed_response		[]string			`json:"failed_response"`
}