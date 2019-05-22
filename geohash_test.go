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

	geohash := Encode(31.2791193622, 121.5543168783, 8)
	fmt.Println(geohash)
}
