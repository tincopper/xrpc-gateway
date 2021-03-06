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

// An example resource type from AIP-123 used to test the behavior described in the CreateBookRequest message.  See: https://google.aip.dev/123
type Book struct {
	// Output only. The book's ID.
	Id string `json:"id,omitempty"`
	// Output only. Creation time of the book.
	CreateTime time.Time `json:"createTime,omitempty"`
}
