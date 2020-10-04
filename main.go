/*
 * API Specification for Edge AppManager
 *
 * This is a API specification for Edge AppManager module of Aricent Edge Compute solution.
 *
 * API version: 1.0.0
 * Contact: srishti.srivastava@aricent.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "usermanagementservice/routes"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}