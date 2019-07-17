package url

import (
	"fmt"
	"net/url"
)

func main() {
	v := make(url.Values)
	v["app_id"] = []string{"appid"}
	query := v.Encode()
	fmt.Println(query)
}
