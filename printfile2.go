package main
import(
	"fmt"
	"io/ioutil" // 入出力関連のユーティリティ関数が定義されているパッケージ
	/*io/ioutil
		func ReadFile(), ReadDir(), WriteFile()などが定義されている。
		参照： https://waman.hatenablog.com/entry/2017/10/05/132727
	*/
)
func main(){
	PrintFile("test.txt")
}

// ファイルを読み込んで内容を表示する
func PrintFile(filename string){
	// ファイルの内容を全部読む --- (*1)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// 繰り返し表示する --- (*2)
	line := ""
	aline := ""
	for i := 0; i<len(bytes); i++ {
		b := bytes[i]
		// アスキー文字の表示用
		c := string(b)
		if b < 32 || b > 126 {c="_"}
		aline += c
		// 画面に表示する --- (*3)
		m := i % 16
		if m == 0 {
			line += fmt.Sprintf("%04d", i) // %04x ---
		}
		line += fmt.Sprintf("%02x", b)
		switch m {
		case 7: // 見やすくするための区切り
			line += "|" 
		case 15: //区切りとアスキー文字
			fmt.Println(line + "|" + aline)
			line = ""
			aline = ""
		default:
			line += " "
		}
	}
	// 表示残しを出力 --- (*4)
	if line != "" {
		fmt.Printf("%-53s|%s\n", line, aline)
	}
}