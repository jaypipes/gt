package graphic

// Shamelessly pulled and modified from:
// https://github.com/rivo/tview/blob/f39b95c73dbb30877f4b5145b835333002afb2a8/semigraphics.go

// Block: Box Drawing U+2500-U+257F (http://unicode.org/charts/PDF/U2500.pdf)
const (
	BoxDrawingLightHorizontal                    rune = '\u2500' // ─
	BoxDrawingHeavyHorizontal                    rune = '\u2501' // ━
	BoxDrawingLightVertical                      rune = '\u2502' // │
	BoxDrawingHeavyVertical                      rune = '\u2503' // ┃
	BoxDrawingLightTripleDashHorizontal          rune = '\u2504' // ┄
	BoxDrawingHeavyTripleDashHorizontal          rune = '\u2505' // ┅
	BoxDrawingLightTripleDashVertical            rune = '\u2506' // ┆
	BoxDrawingHeavyTripleDashVertical            rune = '\u2507' // ┇
	BoxDrawingLightQuadrupleDashHorizontal       rune = '\u2508' // ┈
	BoxDrawingHeavyQuadrupleDashHorizontal       rune = '\u2509' // ┉
	BoxDrawingLightQuadrupleDashVertical         rune = '\u250a' // ┊
	BoxDrawingHeavyQuadrupleDashVertical         rune = '\u250b' // ┋
	BoxDrawingLightDownAndRight                  rune = '\u250c' // ┌
	BoxDrawingDownLightAndRightHeavy             rune = '\u250d' // ┍
	BoxDrawingDownHeavyAndRightLight             rune = '\u250e' // ┎
	BoxDrawingHeavyDownAndRight                  rune = '\u250f' // ┏
	BoxDrawingLightDownAndLeft                   rune = '\u2510' // ┐
	BoxDrawingDownLightAndLeftHeavy              rune = '\u2511' // ┑
	BoxDrawingDownHeavyAndLeftLight              rune = '\u2512' // ┒
	BoxDrawingHeavyDownAndLeft                   rune = '\u2513' // ┓
	BoxDrawingLightUpAndRight                    rune = '\u2514' // └
	BoxDrawingUpLightAndRightHeavy               rune = '\u2515' // ┕
	BoxDrawingUpHeavyAndRightLight               rune = '\u2516' // ┖
	BoxDrawingHeavyUpAndRight                    rune = '\u2517' // ┗
	BoxDrawingLightUpAndLeft                     rune = '\u2518' // ┘
	BoxDrawingUpLightAndLeftHeavy                rune = '\u2519' // ┙
	BoxDrawingUpHeavyAndLeftLight                rune = '\u251a' // ┚
	BoxDrawingHeavyUpAndLeft                     rune = '\u251b' // ┛
	BoxDrawingLightVerticalAndRight              rune = '\u251c' // ├
	BoxDrawingVerticalLightAndRightHeavy         rune = '\u251d' // ┝
	BoxDrawingUpHeavyAndRightDownLight           rune = '\u251e' // ┞
	BoxDrawingDownHeavyAndRightUpLight           rune = '\u251f' // ┟
	BoxDrawingVerticalHeavyAndRightLight         rune = '\u2520' // ┠
	BoxDrawingDownLightAndRightUpHeavy           rune = '\u2521' // ┡
	BoxDrawingUpLightAndRightDownHeavy           rune = '\u2522' // ┢
	BoxDrawingHeavyVerticalAndRight              rune = '\u2523' // ┣
	BoxDrawingLightVerticalAndLeft               rune = '\u2524' // ┤
	BoxDrawingVerticalLightAndLeftHeavy          rune = '\u2525' // ┥
	BoxDrawingUpHeavyAndLeftDownLight            rune = '\u2526' // ┦
	BoxDrawingDownHeavyAndLeftUpLight            rune = '\u2527' // ┧
	BoxDrawingVerticalHeavyAndLeftLight          rune = '\u2528' // ┨
	BoxDrawingDownLightAndLeftUpHeavy            rune = '\u2529' // ┨
	BoxDrawingUpLightAndLeftDownHeavy            rune = '\u252a' // ┪
	BoxDrawingHeavyVerticalAndLeft               rune = '\u252b' // ┫
	BoxDrawingLightDownAndHorizontal             rune = '\u252c' // ┬
	BoxDrawingLeftHeavyAndRightDownLight         rune = '\u252d' // ┭
	BoxDrawingRightHeavyAndLeftDownLight         rune = '\u252e' // ┮
	BoxDrawingDownLightAndHorizontalHeavy        rune = '\u252f' // ┯
	BoxDrawingDownHeavyAndHorizontalLight        rune = '\u2530' // ┰
	BoxDrawingRightLightAndLeftDownHeavy         rune = '\u2531' // ┱
	BoxDrawingLeftLightAndRightDownHeavy         rune = '\u2532' // ┲
	BoxDrawingHeavyDownAndHorizontal             rune = '\u2533' // ┳
	BoxDrawingLightUpAndHorizontal               rune = '\u2534' // ┴
	BoxDrawingLeftHeavyAndRightUpLight           rune = '\u2535' // ┵
	BoxDrawingRightHeavyAndLeftUpLight           rune = '\u2536' // ┶
	BoxDrawingUpLightAndHorizontalHeavy          rune = '\u2537' // ┷
	BoxDrawingUpHeavyAndHorizontalLight          rune = '\u2538' // ┸
	BoxDrawingRightLightAndLeftUpHeavy           rune = '\u2539' // ┹
	BoxDrawingLeftLightAndRightUpHeavy           rune = '\u253a' // ┺
	BoxDrawingHeavyUpAndHorizontal               rune = '\u253b' // ┻
	BoxDrawingLightVerticalAndHorizontal         rune = '\u253c' // ┼
	BoxDrawingLeftHeavyAndRightVerticalLight     rune = '\u253d' // ┽
	BoxDrawingRightHeavyAndLeftVerticalLight     rune = '\u253e' // ┾
	BoxDrawingVerticalLightAndHorizontalHeavy    rune = '\u253f' // ┿
	BoxDrawingUpHeavyAndDownHorizontalLight      rune = '\u2540' // ╀
	BoxDrawingDownHeavyAndUpHorizontalLight      rune = '\u2541' // ╁
	BoxDrawingVerticalHeavyAndHorizontalLight    rune = '\u2542' // ╂
	BoxDrawingLeftUpHeavyAndRightDownLight       rune = '\u2543' // ╃
	BoxDrawingRightUpHeavyAndLeftDownLight       rune = '\u2544' // ╄
	BoxDrawingLeftDownHeavyAndRightUpLight       rune = '\u2545' // ╅
	BoxDrawingRightDownHeavyAndLeftUpLight       rune = '\u2546' // ╆
	BoxDrawingDownLightAndUpHorizontalHeavy      rune = '\u2547' // ╇
	BoxDrawingUpLightAndDownHorizontalHeavy      rune = '\u2548' // ╈
	BoxDrawingRightLightAndLeftVerticalHeavy     rune = '\u2549' // ╉
	BoxDrawingLeftLightAndRightVerticalHeavy     rune = '\u254a' // ╊
	BoxDrawingHeavyVerticalAndHorizontal         rune = '\u254b' // ╋
	BoxDrawingLightDoubleDashHorizontal          rune = '\u254c' // ╌
	BoxDrawingHeavyDoubleDashHorizontal          rune = '\u254d' // ╍
	BoxDrawingLightDoubleDashVertical            rune = '\u254e' // ╎
	BoxDrawingHeavyDoubleDashVertical            rune = '\u254f' // ╏
	BoxDrawingDoubleHorizontal                   rune = '\u2550' // ═
	BoxDrawingDoubleVertical                     rune = '\u2551' // ║
	BoxDrawingDownSingleAndRightDouble           rune = '\u2552' // ╒
	BoxDrawingDownDoubleAndRightSingle           rune = '\u2553' // ╓
	BoxDrawingDoubleDownAndRight                 rune = '\u2554' // ╔
	BoxDrawingDownSingleAndLeftDouble            rune = '\u2555' // ╕
	BoxDrawingDownDoubleAndLeftSingle            rune = '\u2556' // ╖
	BoxDrawingDoubleDownAndLeft                  rune = '\u2557' // ╗
	BoxDrawingUpSingleAndRightDouble             rune = '\u2558' // ╘
	BoxDrawingUpDoubleAndRightSingle             rune = '\u2559' // ╙
	BoxDrawingDoubleUpAndRight                   rune = '\u255a' // ╚
	BoxDrawingUpSingleAndLeftDouble              rune = '\u255b' // ╛
	BoxDrawingUpDoubleAndLeftSingle              rune = '\u255c' // ╜
	BoxDrawingDoubleUpAndLeft                    rune = '\u255d' // ╝
	BoxDrawingVerticalSingleAndRightDouble       rune = '\u255e' // ╞
	BoxDrawingVerticalDoubleAndRightSingle       rune = '\u255f' // ╟
	BoxDrawingDoubleVerticalAndRight             rune = '\u2560' // ╠
	BoxDrawingVerticalSingleAndLeftDouble        rune = '\u2561' // ╡
	BoxDrawingVerticalDoubleAndLeftSingle        rune = '\u2562' // ╢
	BoxDrawingDoubleVerticalAndLeft              rune = '\u2563' // ╣
	BoxDrawingDownSingleAndHorizontalDouble      rune = '\u2564' // ╤
	BoxDrawingDownDoubleAndHorizontalSingle      rune = '\u2565' // ╥
	BoxDrawingDoubleDownAndHorizontal            rune = '\u2566' // ╦
	BoxDrawingUpSingleAndHorizontalDouble        rune = '\u2567' // ╧
	BoxDrawingUpDoubleAndHorizontalSingle        rune = '\u2568' // ╨
	BoxDrawingDoubleUpAndHorizontal              rune = '\u2569' // ╩
	BoxDrawingVerticalSingleAndHorizontalDouble  rune = '\u256a' // ╪
	BoxDrawingVerticalDoubleAndHorizontalSingle  rune = '\u256b' // ╫
	BoxDrawingDoubleVerticalAndHorizontal        rune = '\u256c' // ╬
	BoxDrawingLightArcDownAndRight               rune = '\u256d' // ╭
	BoxDrawingLightArcDownAndLeft                rune = '\u256e' // ╮
	BoxDrawingLightArcUpAndLeft                  rune = '\u256f' // ╯
	BoxDrawingLightArcUpAndRight                 rune = '\u2570' // ╰
	BoxDrawingLightDiagonalUpperRightToLowerLeft rune = '\u2571' // ╱
	BoxDrawingLightDiagonalUpperLeftToLowerRight rune = '\u2572' // ╲
	BoxDrawingLightDiagonalCross                 rune = '\u2573' // ╳
	BoxDrawingLightLeft                          rune = '\u2574' // ╴
	BoxDrawingLightUp                            rune = '\u2575' // ╵
	BoxDrawingLightRight                         rune = '\u2576' // ╶
	BoxDrawingLightDown                          rune = '\u2577' // ╷
	BoxDrawingHeavyLeft                          rune = '\u2578' // ╸
	BoxDrawingHeavyUp                            rune = '\u2579' // ╹
	BoxDrawingHeavyRight                         rune = '\u257a' // ╺
	BoxDrawingHeavyDown                          rune = '\u257b' // ╻
	BoxDrawingLightLeftAndHeavyRight             rune = '\u257c' // ╼
	BoxDrawingLightUpAndHeavyDown                rune = '\u257d' // ╽
	BoxDrawingHeavyLeftAndLightRight             rune = '\u257e' // ╾
	BoxDrawingHeavyUpAndLightDown                rune = '\u257f' // ╿
)

// BoxDrawingJoints is a map for joining box drawing (or otherwise) runes.
// The matching will be sorted ascending by rune value, so you don't need to
// provide all rune combinations,
// e.g. (─) + (│) = (┼) will also match (│) + (─) = (┼)
var BoxDrawingJoints = map[string]rune{
	// (─) + (│) = (┼)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightVertical}): BoxDrawingLightVerticalAndHorizontal,
	// (─) + (┌) = (┬)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightDownAndRight}): BoxDrawingLightDownAndHorizontal,
	// (─) + (┐) = (┬)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightDownAndLeft}): BoxDrawingLightDownAndHorizontal,
	// (─) + (└) = (┴)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightUpAndRight}): BoxDrawingLightUpAndHorizontal,
	// (─) + (┘) = (┴)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightUpAndLeft}): BoxDrawingLightUpAndHorizontal,
	// (─) + (├) = (┼)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightVerticalAndRight}): BoxDrawingLightVerticalAndHorizontal,
	// (─) + (┤) = (┼)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightVerticalAndLeft}): BoxDrawingLightVerticalAndHorizontal,
	// (─) + (┬) = (┬)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightDownAndHorizontal,
	// (─) + (┴) = (┴)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightUpAndHorizontal,
	// (─) + (┼) = (┼)
	string([]rune{BoxDrawingLightHorizontal, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (│) + (┌) = (├)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightDownAndRight}): BoxDrawingLightVerticalAndRight,
	// (│) + (┐) = (┤)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightDownAndLeft}): BoxDrawingLightVerticalAndLeft,
	// (│) + (└) = (├)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightUpAndRight}): BoxDrawingLightVerticalAndRight,
	// (│) + (┘) = (┤)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightUpAndLeft}): BoxDrawingLightVerticalAndLeft,
	// (│) + (├) = (├)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightVerticalAndRight}): BoxDrawingLightVerticalAndRight,
	// (│) + (┤) = (┤)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightVerticalAndLeft}): BoxDrawingLightVerticalAndLeft,
	// (│) + (┬) = (┼)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (│) + (┴) = (┼)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (│) + (┼) = (┼)
	string([]rune{BoxDrawingLightVertical, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (┌) + (┐) = (┬)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightDownAndLeft}): BoxDrawingLightDownAndHorizontal,
	// (┌) + (└) = (├)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightUpAndRight}): BoxDrawingLightVerticalAndRight,
	// (┌) + (┘) = (┼)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightUpAndLeft}): BoxDrawingLightVerticalAndHorizontal,
	// (┌) + (├) = (├)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightVerticalAndRight}): BoxDrawingLightVerticalAndRight,
	// (┌) + (┤) = (┼)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightVerticalAndLeft}): BoxDrawingLightVerticalAndHorizontal,
	// (┌) + (┬) = (┬)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightDownAndHorizontal,
	// (┌) + (┴) = (┼)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (┌) + (┴) = (┼)
	string([]rune{BoxDrawingLightDownAndRight, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (┐) + (└) = (┼)
	string([]rune{BoxDrawingLightDownAndLeft, BoxDrawingLightUpAndRight}): BoxDrawingLightVerticalAndHorizontal,
	// (┐) + (┘) = (┤)
	string([]rune{BoxDrawingLightDownAndLeft, BoxDrawingLightUpAndLeft}): BoxDrawingLightVerticalAndLeft,
	// (┐) + (├) = (┼)
	string([]rune{BoxDrawingLightDownAndLeft, BoxDrawingLightVerticalAndRight}): BoxDrawingLightVerticalAndHorizontal,
	// (┐) + (┤) = (┤)
	string([]rune{BoxDrawingLightDownAndLeft, BoxDrawingLightVerticalAndLeft}): BoxDrawingLightVerticalAndLeft,
	// (┐) + (┬) = (┬)
	string([]rune{BoxDrawingLightDownAndLeft, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightDownAndHorizontal,
	// (┐) + (┴) = (┼)
	string([]rune{BoxDrawingLightDownAndLeft, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (┐) + (┼) = (┼)
	string([]rune{BoxDrawingLightDownAndLeft, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (└) + (┘) = (┴)
	string([]rune{BoxDrawingLightUpAndRight, BoxDrawingLightUpAndLeft}): BoxDrawingLightUpAndHorizontal,
	// (└) + (├) = (├)
	string([]rune{BoxDrawingLightUpAndRight, BoxDrawingLightVerticalAndRight}): BoxDrawingLightVerticalAndRight,
	// (└) + (┤) = (┼)
	string([]rune{BoxDrawingLightUpAndRight, BoxDrawingLightVerticalAndLeft}): BoxDrawingLightVerticalAndHorizontal,
	// (└) + (┬) = (┼)
	string([]rune{BoxDrawingLightUpAndRight, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (└) + (┴) = (┴)
	string([]rune{BoxDrawingLightUpAndRight, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightUpAndHorizontal,
	// (└) + (┼) = (┼)
	string([]rune{BoxDrawingLightUpAndRight, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (┘) + (├) = (┼)
	string([]rune{BoxDrawingLightUpAndLeft, BoxDrawingLightVerticalAndRight}): BoxDrawingLightVerticalAndHorizontal,
	// (┘) + (┤) = (┤)
	string([]rune{BoxDrawingLightUpAndLeft, BoxDrawingLightVerticalAndLeft}): BoxDrawingLightVerticalAndLeft,
	// (┘) + (┬) = (┼)
	string([]rune{BoxDrawingLightUpAndLeft, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (┘) + (┴) = (┴)
	string([]rune{BoxDrawingLightUpAndLeft, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightUpAndHorizontal,
	// (┘) + (┼) = (┼)
	string([]rune{BoxDrawingLightUpAndLeft, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (├) + (┤) = (┼)
	string([]rune{BoxDrawingLightVerticalAndRight, BoxDrawingLightVerticalAndLeft}): BoxDrawingLightVerticalAndHorizontal,
	// (├) + (┬) = (┼)
	string([]rune{BoxDrawingLightVerticalAndRight, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (├) + (┴) = (┼)
	string([]rune{BoxDrawingLightVerticalAndRight, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (├) + (┼) = (┼)
	string([]rune{BoxDrawingLightVerticalAndRight, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (┤) + (┬) = (┼)
	string([]rune{BoxDrawingLightVerticalAndLeft, BoxDrawingLightDownAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (┤) + (┴) = (┼)
	string([]rune{BoxDrawingLightVerticalAndLeft, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (┤) + (┼) = (┼)
	string([]rune{BoxDrawingLightVerticalAndLeft, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (┬) + (┴) = (┼)
	string([]rune{BoxDrawingLightDownAndHorizontal, BoxDrawingLightUpAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
	// (┬) + (┼) = (┼)
	string([]rune{BoxDrawingLightDownAndHorizontal, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,

	// (┴) + (┼) = (┼)
	string([]rune{BoxDrawingLightUpAndHorizontal, BoxDrawingLightVerticalAndHorizontal}): BoxDrawingLightVerticalAndHorizontal,
}
