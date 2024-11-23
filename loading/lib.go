package loading

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

/// a loading bar

func LoadingBar() {
	fmt.Println("\u001b[33m✨ Loading... Please wait! ✨")

	// length is 50
	totalSteps := 50
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 100) // 模拟延迟
		// 计算已经完成的进度
		progress := (i * totalSteps) / 100
		// 构建进度条
		progressBar := "\u001b[36m["

		// 填充进度条的完成部分"»"
		for j := 0; j < progress; j++ {
			progressBar += "\u001b[32m»" // 
		}

		// 填充进度条的空白部分
		for j := progress; j < totalSteps; j++ {
			progressBar += "\u001b[37m»" // 空白字符
		}

		progressBar += "\u001b[36m]"
		writer := bufio.NewWriter(os.Stdout)
		fmt.Fprintf(writer, "\u001b[1000D\u001b[36m%s %d%% 🐢", progressBar, i) 
		writer.Flush()
	}

	// 显示完成信息
	fmt.Println("\n\u001b[32m🎉 Loading Complete! 🎉")
}
