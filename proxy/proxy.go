package proxy

import (
	"bytes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

var client = fasthttp.Client{
	NoDefaultUserAgentHeader: true,
	DisablePathNormalizing:   true,
	Dial:                     fasthttpproxy.FasthttpSocksDialer(os.Getenv("ALL_PROXY")),
}

func Do(c *fiber.Ctx, addr string) error {
	req := c.Request()
	res := c.Response()
	originalURL := utils.CopyString(c.OriginalURL())
	defer req.SetRequestURI(originalURL)
	req.SetRequestURI(addr)
	// NOTE: if req.isTLS is true, SetRequestURI keeps the scheme as https.
	// issue reference:
	// https://github.com/gofiber/fiber/issues/1762
	if scheme := getScheme(utils.UnsafeBytes(addr)); len(scheme) > 0 {
		req.URI().SetSchemeBytes(scheme)
	}

	req.Header.Del(fiber.HeaderConnection)
	if err := client.Do(req, res); err != nil {
		return err
	}
	res.Header.Del(fiber.HeaderConnection)
	return nil
}

func getScheme(uri []byte) []byte {
	i := bytes.IndexByte(uri, '/')
	if i < 1 || uri[i-1] != ':' || i == len(uri)-1 || uri[i+1] != '/' {
		return nil
	}
	return uri[:i-1]
}
