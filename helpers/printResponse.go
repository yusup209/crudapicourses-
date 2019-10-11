package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

func JSONify(ctx *fasthttp.RequestCtx, data map[string]interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("err while converting map data to json in JSONify\n%v", err)
		data["err"] = err
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(res)

}
