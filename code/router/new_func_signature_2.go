import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luci/luci-go/server/router"
	"golang.org/x/net/context"
)

// START OMIT
func helloPage(c *router.Context)

// END OMIT
type Context struct {
	Context context.Context
	Writer  http.ResponseWriter
	Request *http.Request
	Params  httprouter.Params
}

