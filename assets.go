package asset

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9c71a471d1741f25c6a0cf33633b97ca741711ba = "db:\n  development:\n    adapter: mysql\n    encoding: utf8\n    database: blog\n    pool: 4\n    host: localhost\n    port: 3006\n    user: root\n    password: pas\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"config.yml"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546982883, 1546982883643494035),
		Data:     nil,
	}, "/config.yml": &assets.File{
		Path:     "/config.yml",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546982715, 1546982715088302462),
		Data:     []byte(_Assets9c71a471d1741f25c6a0cf33633b97ca741711ba),
	}}, "")
