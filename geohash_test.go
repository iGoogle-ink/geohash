//==================================
//  * Name：Jerry
//  * DateTime：2019/5/21 16:45
//  * Desc：
//==================================
package geohash

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {

	geohash := Encode(31.2851847116, 121.5571761131, 10)
	fmt.Println(geohash)
}
