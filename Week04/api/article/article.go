package article

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"week4/internal/service"
)

type GetArticleRequest struct {
	Rid int `form:"rid"`
}
func GetArticle(c *gin.Context) {
	var req GetArticleRequest
	if err := c.ShouldBind(&req);err!=nil{
		log.Printf("get article bind err:%v\n",err)
		c.Status(http.StatusBadRequest)
		return
	}

	article,err := service.Srv.GetArticle(c,req.Rid)
	if err != nil{
		log.Printf("failed to get article: %v\n",err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK,article)
}
