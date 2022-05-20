// 学習教材: https://news.mynavi.jp/techplus/article/gogogo-8/
package main
import(
	"fmt"
	"os"
)
// ファイルを読み込んで内容を表示する
func PrintFile(filename string){
	// ファイルを開く --- (*1)
	fp, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	// 読み込みバッファを用意する(一時記憶領域) --- (*2)
	/*make関数
		make([]T, len, cap)
		第一引数:[]T --- スライスの要素の型を指定
		第二引数:len --- スライスの長さを指定
		第三引数:cap --- スライスの容量を指定
	*/
	/*スライス []スライス名
		要素数の指定を必要としない配列
		⇒要素の追加や削除を自由に行える
		var subjects[3] string = [3]string { "国語", "数学", "英語" } //配列
		subjects := []string{ "国語", "数学", "英語" } // スライス
	*/
	/*スライスの長さと容量
		
	*/
	buf := make([]byte, 1)
	// 繰り返し表示する --- (*3)
	for {
		// 1バイト読み込む --- (*4)
		cnt, _ := fp.Read(buf)
		if cnt == 0{
			break
		}
		// 画面に表示する --- (*5)
		fmt.Printf("[%d]", buf[0])
	}
}
func main(){
	PrintFile("test.txt")
	fmt.Printf("\n")
}