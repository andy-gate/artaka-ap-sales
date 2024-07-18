package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"fmt"
	"io"
	"net/http"

	"github.com/andy-gate/artaka-ap-sales/models"
	"github.com/gin-gonic/gin"
)

func SendSales(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.SalesQuery
	c.BindJSON(&query)

	var sales []models.Sales

	if err := models.MPosGORM.Raw("select * from sales where id = ?", query.Id).Scan(&sales).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if (sales != nil) {
		result := submitSalesAP1(sales[0], "foy shele mart")
		c.JSON(http.StatusOK, result)
	  } else {
		c.JSON(http.StatusOK, json.RawMessage(`[]`))
	  }
}

func submitSalesAP1(sales models.Sales, outlet_name string) models.APResponse {
	apiUrl := "https://api-ecsys.angkasapura2.co.id/api/auth/login"

	bodyReq := models.UserLogin{
		Username: "api.artakapos.ho.4546!",
		Password: "api.artakapos.ho.4546!",
	}

	bodyBytes, err := json.Marshal(&bodyReq)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST",apiUrl,bytes.NewBuffer(bodyBytes))
	if err != nil {   
		fmt.Printf("Request Failed: %s", err)
		return models.APResponse{}
	}

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

	defer resp.Body.Close()

	res := models.LoginResponse{}
  	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)

	var storeData models.StoreDetail

	for i := range res.User.Store {
		if strings.EqualFold(res.User.Store[i].Store_name, outlet_name) {
			storeData = res.User.Store[i]
		}
	}

	var reqAPStore models.APRequestStore
	reqAPStore.Store_id = storeData.Store_id
	var trxAP []models.Trx
	var listProduct []models.Product
	_ = json.Unmarshal(sales.Products, &listProduct)

	for in := range listProduct {		
		var temp models.Trx
		temp.Invoice_no = sales.Sales_id
		temp.Trans_date = sales.Create_dtm.Format("2006-01-02")
		temp.Trans_time = sales.Create_dtm.Format("2006-01-02 15:04:05")
		temp.Sequence_unique = strconv.Itoa(in+1)
		temp.Item_name = listProduct[in].Name
		temp.Item_code = listProduct[in].Sku_id
		temp.Item_barcode = listProduct[in].Sku_id
		temp.Item_cat_name = listProduct[in].Category
		temp.Item_cat_code = listProduct[in].Category
		temp.Item_qty = strconv.Itoa(listProduct[in].Number_orders)
		temp.Item_unit = listProduct[in].Units
		temp.Item_price_per_unit = strconv.Itoa(listProduct[in].Sell_cost)
		temp.Item_discount = strconv.Itoa(listProduct[in].Diskon)
		temp.Item_price_amount = strconv.Itoa(listProduct[in].Sell_cost - listProduct[in].Diskon)
		temp.Item_vat = "0"
		temp.Item_tax = "0"
		temp.Item_total_discount = strconv.Itoa(listProduct[in].Diskon * listProduct[in].Number_orders)
		temp.Item_total_price_amount = strconv.Itoa((listProduct[in].Sell_cost - listProduct[in].Diskon) * listProduct[in].Number_orders)
		temp.Item_total_vat = "0"
		temp.Item_total_tax = "0"
		temp.Item_total_service_charge = "0"
		temp.Invoice_tax = "0"
		temp.Invoice_discount = strconv.Itoa(sales.Total_diskon)
		temp.Transaction_amount = strconv.Itoa(listProduct[in].Sell_cost * listProduct[in].Number_orders)
		temp.Currency = "IDR"
		temp.Rate = "1"
		temp.Payment_type = sales.Payment_method
		temp.Payment_by = ""
		temp.Username = ""
		temp.Buyer_barcode = ""
		temp.Buyer_name = ""
		temp.Buyer_flight_no = ""
		temp.Buyer_destination = ""
		temp.Buyer_nationality = ""
		temp.Remark = "Success"
		temp.Tax_id = ""
		temp.Payment_name = sales.Payment_method
		temp.Payment_time = sales.Payment_date
		temp.Distance = "0"
		temp.Journey_time = "0"
		trxAP = append(trxAP, temp)
	}
	reqAPStore.Transactions = trxAP
	var reqAP models.APRequest
	reqAP.Store = append(reqAP.Store, reqAPStore)
	apiUrlAP := "https://api-ecsys.angkasapura2.co.id/api/v1/transaction"

	bodyBytesReq, err := json.Marshal(&reqAP)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", string(bodyBytesReq))

	request, _ := http.NewRequest("POST", apiUrlAP, bytes.NewBuffer(bodyBytesReq))
	request.Header.Set("Authorization", res.Token)
	request.Header.Set("Content-Type", "application/json")

	if err != nil {   
		fmt.Printf("Request Failed: %s", err)
		return models.APResponse{}
	}

    respAP, err := client.Do(request)
    if err != nil {
        panic(err)
    }

	defer respAP.Body.Close()

	resAP := models.APResponse{}
	bodyAP, _ := io.ReadAll(respAP.Body)
	_ = json.Unmarshal(bodyAP, &resAP)
	return resAP
}