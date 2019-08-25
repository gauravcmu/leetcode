package main

import (
"fmt",
"gin/gonic"
)

func main() {
	//initialize context 
	ctx := context.New() 

	//return a list of all products
	ctx.R.Get("/products/", func (){
		products, err := ctx.MongoSession.Collection("products").find() 
		if err != nil {
			ctx.R.Error(500, "failed to find in DB")
			return 
		}

		ctx.R.JSON(products)
		} )

	//return a product given an ID
	ctx.R.Get("/products/:productid", func (){
		productid := ctx.R.getparam("productid")
		product, err := ctx.MongoSession.Collection("products").findById(productid)
		if err != nil {
			ctx.R.Error(500, "failed to find in DB")
			return 
		}

		ctx.R.JSON(product)
		})

		//return a list of all products
	ctx.R.Get("/customer/", func (){
		products, err := ctx.MongoSession.Collection("customer").find() 
		if err != nil {
			ctx.R.Error(500, "failed to find in DB")
			return 
		}

		ctx.R.JSON(products)
		} )

	//return a product given an ID
	ctx.R.Get("/customer/:customerid", func (){
		productid := ctx.R.getparam("productid")
		product, err := ctx.MongoSession.Collection("products").findById(productid)
		if err != nil {
			ctx.R.Error(500, "failed to find in DB")
			return 
		}

		ctx.R.JSON(product)
		})

	ctx.R.POST("/customer", func() {
		//read the body and parse into customer object (models.customer)
		customer 
		})

	
	ctx.R.router() 

}

