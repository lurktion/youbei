package static

import "embed"

//go:embed index.html favicon.ico css img js fonts
var Static embed.FS
