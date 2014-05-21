package main 

import (
//	"github.com/astaxie/beego"
    "blog/models"
	"time"
	"fmt"
)

func main(){
		var blog models.Blog
		blog.Title = "title"
		blog.Content = "content"
		blog.Created = time.Now()
		fmt.Println("db operate")
		models.SaveBlog(blog)
}
