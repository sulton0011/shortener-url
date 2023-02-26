package v1

import (
	"github.com/skip2/go-qrcode"
	"github.com/teris-io/shortid"
)

func (g *GetUrlResponse) GetShortUrl(HTTPScheme, host, port string) string {
	if g == nil {
		return ""
	}
	g.ShortUrl = HTTPScheme + "://" + host + port + "/" + g.ShortUrl
	return g.ShortUrl
}

func (g *GetUrlResponse) GetQrCode(shortUrl string, size int) []byte {
	if g == nil || shortUrl == "" {
		return []byte{}
	}

	png, _ := qrcode.Encode(shortUrl, qrcode.Medium, size)
	g.QrCode = png
	return g.QrCode
}

func (g *CreateUrlRequest) GetShortUrl() string {
	if g == nil {
		return ""
	}
	if g.ShortUrl == "" {
		urlCode, _ := shortid.Generate()
		return urlCode
	}
	return g.ShortUrl
}
