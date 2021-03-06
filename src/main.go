package main

import (
	restapi "github.com/cloud-barista/cb-ladybug/src/rest-api"
	"github.com/cloud-barista/cb-ladybug/src/utils/config"
)

// @title CB-Laydbug REST API
// @version 1.0
// @description CB-Laydbug REST API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://cloud-barista.github.io
// @contact.email contact-to-cloud-barista@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /ladybug
func main() {

	config.Setup()
	restapi.Server()

}
