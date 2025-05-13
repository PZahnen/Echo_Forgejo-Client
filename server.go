package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

//go:embed client/public/**/*
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
    if useOS {
        log.Print("using live mode")
        return http.FS(os.DirFS("client/public"))
    }

    log.Print("using embed mode")
    fsys, err := fs.Sub(embededFiles, "client/public")
    if err != nil {
        panic(err)
    }

    return http.FS(fsys)
}

func main() {
    e := echo.New()


    // API-Endpunkte
    e.GET("/api/deployment/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, deployments)
    })

	    e.GET("/api/configuration/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, configs)
    })

		    e.GET("/api/settings", func(c echo.Context) error {
        return c.JSON(http.StatusOK, settings)
    })

	    e.GET("/api/configuration/:owner", func(c echo.Context) error {
        owner := c.Param("owner")

        // Filtere die Konfigurationen nach Owner
        var ownerConfigs []ConfigSource
        for _, config := range configs {
            if config.Owner == owner {
                ownerConfigs = append(ownerConfigs, config)
            }
        }

        if len(ownerConfigs) == 0 {
            return c.JSON(http.StatusNotFound, map[string]string{"message": "No configurations found for owner"})
        }

        return c.JSON(http.StatusOK, ownerConfigs)
    })

	    e.GET("/api/template/:id/versions", func(c echo.Context) error {
        id := c.Param("id")

        // Suche nach dem Template
        for _, template := range templateVersions {
            if template.ID == id {
                return c.JSON(http.StatusOK, template)
            }
        }

        return c.JSON(http.StatusNotFound, map[string]string{"message": "Template not found"})
    })

    e.GET("/api/deployment/:owner/:name", func(c echo.Context) error {
        owner := c.Param("owner")
        name := c.Param("name")

        for _, d := range deployments {
            if d.Owner == owner && d.Name == name {
                return c.JSON(http.StatusOK, d)
            }
        }

        return c.JSON(http.StatusNotFound, map[string]string{"message": "Deployment not found"})
    })

e.GET("/api/deployment/:owner", func(c echo.Context) error {
    owner := c.Param("owner")

    var ownerDeployments []Deployment
    for _, d := range deployments {
        if d.Owner == owner {
            ownerDeployments = append(ownerDeployments, d)
        }
    }

    if len(ownerDeployments) == 0 {
        return c.JSON(http.StatusOK, []Deployment{})
    }

    return c.JSON(http.StatusOK, ownerDeployments)
})

    e.POST("/api/deployment", func(c echo.Context) error {
        newDeployment := Deployment{}
        if err := c.Bind(&newDeployment); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
        }

        for _, d := range deployments {
            if d.Owner == newDeployment.Owner && d.Name == newDeployment.Name {
                return c.JSON(http.StatusConflict, map[string]string{"message": "Deployment already exists"})
            }
        }

        deployments = append(deployments, newDeployment)
        return c.JSON(http.StatusCreated, newDeployment)
    })

    e.PUT("/api/deployment/:owner/:name", func(c echo.Context) error {
        owner := c.Param("owner")
        name := c.Param("name")

        updatedDeployment := Deployment{}
        if err := c.Bind(&updatedDeployment); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
        }

        for i, d := range deployments {
            if d.Owner == owner && d.Name == name {
                deployments[i] = updatedDeployment
                return c.JSON(http.StatusOK, updatedDeployment)
            }
        }

        return c.JSON(http.StatusNotFound, map[string]string{"message": "Deployment not found"})
    })

    e.DELETE("/api/deployment/:owner/:name", func(c echo.Context) error {
        owner := c.Param("owner")
        name := c.Param("name")

        for i, d := range deployments {
            if d.Owner == owner && d.Name == name {
                deployments = append(deployments[:i], deployments[i+1:]...)
                return c.JSON(http.StatusOK, map[string]string{"message": "Deployment deleted"})
            }
        }

        return c.JSON(http.StatusNotFound, map[string]string{"message": "Deployment not found"})
    })

    // UI-Integration (muss zuletzt registriert werden)
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	assetHandler := http.FileServer(getFileSystem(useOS))
	e.GET("/*", echo.WrapHandler(assetHandler))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	e.Logger.Fatal(e.Start(":1323"))
}

// starten mit go run server.go mockData.go live