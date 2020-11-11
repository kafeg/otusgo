package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestMainHW04(t *testing.T) {
	hw04()
}

func TestWordCount(t *testing.T) {
	testString := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean dictum dignissim pharetra. Nullam nec eros laoreet, tempus sapien at, luctus justo. Fusce orci erat, ac lorem rutrum, lobortis dignissim metus. Maecenas sagittis tellus commodo orci interdum feugiat. Vestibulum aliquam dolor vel gravida viverra. Quisque porta fermentum aliquam. Nam molestie in metus ut viverra.
Ut vulputate elementum. Curabitur facilisis consequat elit, vitae eros elementum sit amet. Proin faucibus nulla posuere, interdum augue quis, consectetur eros eros eros eros eros. Cras egestas diam vel vestibulum luctus. Sed ut augue finibus, efficitur augue at, dapibus turpis. In congue feugiat arcu sit amet rhoncus. Interdum et malesuada fames ac ante ipsum primis in faucibus. Vestibulum venenatis varius sem, id pulvinar mauris ultrices sed. Mauris quis rutrum massa.
Vestibulum in tellus non nulla volutpat rutrum et ac dolor. Etiam tempor volutpat orci vitae euismod. Praesent ut massa dictum, enim id, euismod nunc. Integer nunc lorem, fermentum id laoreet a, tincidunt a leo. Sed pulvinar est ac odio feugiat, id suscipit metus cursus. Maecenas eget condimentum est. Ut lobortis lobortis nisl sit amet consequat.
Donec ullamcorper congue varius. Duis sagittis tellus nisl, sed eleifend nulla lacinia non. Nunc sed libero ac sem sodales maximus. Pellentesque pellentesque pellentesque enim. Vestibulum sem mauris, pharetra et eros nec, dictum dignissim. Nulla facilisi. Nullam iaculis lobortis nibh, nec vestibulum sem sollicitudin at. Sed commodo arcu eu elit tempus, eget pulvinar est posuere. Praesent ac leo rhoncus, rhoncus felis eu, posuere magna. Curabitur aliquet, nibh sit amet vestibulum pellentesque, lectus quam ullamcorper urna, sed facilisis tellus nunc eget libero. Nunc et tempor dolor. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.
Fusce maximus eros in lacus vestibulum, eu accumsan ligula aliquet. Nullam semper vestibulum nibh, in mattis risus laoreet ac. Aliquam erat volutpat. Vestibulum auctor lectus vel lobortis egestas. Nunc ullamcorper tincidunt massa, eu sagittis nisi consectetur at. Vestibulum leo ipsum, pharetra a porttitor ac, venenatis vitae lectus. Integer elementum risus a bibendum placerat. Vestibulum a blandit metus. Nunc et varius mi, vel dictum sem. Mauris non ante ornare, convallis leo id, semper enim. Sed posuere scelerisque lorem quis finibus. Vestibulum nulla ligula, imperdiet ut felis eget, auctor maximus dolor.
Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vivamus molestie lorem et convallis vestibulum. Nunc cursus iaculis neque, sodales cursus nulla blandit ut. Integer sit amet aliquam erat, non tempor lacus. Quisque ligula quam, finibus consectetur quam eget, congue blandit ligula. Etiam mauris velit, cursus eu nunc ac, tincidunt cursus neque. Pellentesque tincidunt leo. Aenean luctus libero nec magna gravida, pharetra congue nunc elementum. Quisque lobortis mattis sapien vel rhoncus. Vestibulum quis sapien vel enim tempor scelerisque quis ut tellus. Curabitur leo metus, aliquet lobortis sagittis sit amet, blandit eu risus. Aenean eu tellus risus. Suspendisse at pharetra odio, in porttitor enim.
Nulla a aliquet leo. Proin sollicitudin sit amet nibh eget pellentesque. Nullam vitae faucibus felis. Duis non tortor efficitur, venenatis sapien feugiat, finibus arcu. Duis a nulla pharetra pharetra pharetra.`

	testResult := map[string]int {
		"amet":6,
		"cursus":5,
		"eros":9,
		"lobortis":7,
		"nulla":7,
		"nunc":10,
		"pellentesque":6,
		"pharetra":8,
		"tellus":6,
		"vestibulum":13,
	}

	wordsCnt := wordsCount(testString)
	fmt.Printf("%v\n", wordsCnt)

	require.Equal(t, testResult, wordsCnt, "wordsCount")
}
