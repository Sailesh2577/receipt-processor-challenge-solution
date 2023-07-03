/*
 * Receipt Processor
 *
 * A simple receipt processor
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Item struct {
	// The Short Product Description for the item.
	ShortDescription string `json:"shortDescription"`
	// The total price payed for this item.
	Price string `json:"price"`
}