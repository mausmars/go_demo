package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func main() {
	var html = template.Must(template.New("https").Parse(`
	<html>
	<head>
	  <title>Https Test</title>
	  <script src="/assets/app.js"></script>
	</head>
	<body>
	  <h1 style="color:red;">Welcome, Ginner! 22</h1>
	</body>
	</html>
	`))

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("https://www.baidu.com/", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}else{
				log.Printf("Success to push!!!")
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	// Listen and Server in https://10.0.106.205:8080/
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}