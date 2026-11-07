package resources

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 内部嵌入
//
//go:embed templates static
var EFS embed.FS

func LoadResourceFromEmbed(r *gin.Engine) {
	tpl, _ := fs.Sub(EFS, "templates")
	//坑点：模板注册名是文件名，不是路径，而用c.HTML()要写注册名，所以用路径会报错；要在gtpl文件中加{{define 注册名}}才能用路径
	r.LoadHTMLFS(http.FS(tpl), "*/*.gtpl") // 不是**/*.gtpl

	/* 坑点1：路径中有"//"不会解析成"/", 资源不同报错
	blog, _ := fs.Sub(EFS, "static/blog")
	r.StaticFS("/blog", http.FS(blog))
	*/

	/* 坑点2：http.Dir要用完整路径
	r.StaticFS("/blog", http.Dir("./static/blog"))
	*/

	static, _ := fs.Sub(EFS, "static")
	timeline, _ := fs.Sub(static, "timeline") //特殊处理

	r.StaticFS("/static", http.FS(static))
	r.StaticFS("/timeline", http.FS(timeline))
}

// 外部加载
func LoadResourceFromLocalFileSystem(r *gin.Engine) {
	r.LoadHTMLGlob("internal/resources/templates/**/*.gtpl")

	r.StaticFS("/static", http.Dir("internal/resources/static/"))
	r.StaticFS("/timeline", http.Dir("internal/resources/static/timeline")) //特殊处理
}

func Init(r *gin.Engine) {
	//二选一
	LoadResourceFromEmbed(r)
}
