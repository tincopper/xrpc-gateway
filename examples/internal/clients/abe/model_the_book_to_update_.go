/*
 * A Bit of Everything
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Contact: none@example.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package abe

import (
	"time"
)

// The book's `name` field is used to identify the book to be updated. Format: publishers/{publisher}/books/{book}
type TheBookToUpdate_ struct {
	// Output only. The book's ID.
	Id string `json:"id,omitempty"`
	// Output only. Creation time of the book.
	CreateTime time.Time `json:"createTime,omitempty"`
}
