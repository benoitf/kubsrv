package controller

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
	"os"
	"log"
)

func GetService(c *gin.Context) {
	name := c.Param("name")
	name = strings.Replace(name, ".", string(os.PathSeparator), -1)
	version := c.Param("version")
	cheServiceFile := filepath.Join(cheRegistryRepository, name, version, "CheService.yaml")
	if (gin.IsDebugging()) {
		log.Printf("Requested CheService %s", cheServiceFile)
	}

	c.YAML(200, cheServiceFile)
}
