package main

import "github.com/gin-gonic/gin"
import "os/exec"
import "strconv"
import "bytes"
import "github.com/gin-contrib/cors"

var num [5] int

func main() {

    r := gin.Default()

    //cors
	r.Use(cors.New(cors.Config{
        AllowOrigins:  []string{"*"},
        AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
        AllowHeaders:  []string{"content-type"},
        ExposeHeaders: []string{"X-Total-Count"},
    }))

    //run a container, uplimit is 5
    r.GET("/run", func(c *gin.Context) {

    	var port int
    	i := 0
        for ; i < 5; i++ {
        	if num[i] == 0 {
        		num[i] = 9994+i
        		port = num[i]
        		break
        	}
        }
        if i < 5 {
        	cmd := exec.Command("/bin/bash", "-c", "./docker.sh 0 " + strconv.Itoa(port))
		    cmd.Start()
        }
        c.JSON(200, gin.H{
            "port": port,
        })
    })

    //remove a container which port matchs the parameter [:port]
    r.GET("/rm/:port", func(c *gin.Context) {

    	port,err := strconv.Atoi(c.Param("port"))
    	i := 0
        for ; i < 5; i++ {
        	if num[i] == port {
        		num[i] = 0
        		break
        	}
        }
        var outStr string
        if i < 5 {
        	cmd := exec.Command("/bin/bash", "-c", "./docker.sh 1 " + strconv.Itoa(port))
	        var out bytes.Buffer
		    cmd.Stdout = &out
		    cmd.Run()
		    outStr = out.String()
        }else {
        	outStr = "no this task"
        }
        c.JSON(200, gin.H{
            "port": port,
            "out": outStr,
            "err": err,
        })
    })

    r.Run(":9999") // listen on 0.0.0.0:9999
}
























