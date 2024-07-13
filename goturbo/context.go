package goturbo

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Context) String(status int, response string) {
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(response))
}

func (c *Context) JSON(status int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(obj)
}
