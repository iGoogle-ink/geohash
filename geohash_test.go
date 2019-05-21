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

	geohash := Encode(39.928167, 116.389550, 8)
	fmt.Println(geohash)
}
