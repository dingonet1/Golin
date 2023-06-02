//go:build windows

package windows

import (
	"fmt"
	"golin/global"
	"regexp"
	"strconv"
	"strings"
)

func disk() {
	value := global.ExecCommands(`wmic logicaldisk get DeviceID, FileSystem, Size, FreeSpace`)
	value = strings.ReplaceAll(value, "\r\r\n", "\n")
	echo := ""
	for _, v := range strings.Split(value, "\n") {
		if strings.Count(v, "DeviceID") > 0 {
			continue
		}
		re := regexp.MustCompile(`\s{2,}`)
		v = re.ReplaceAllString(v, " ")
		data := strings.Split(v, " ")
		if len(data) == 5 {
			free, _ := strconv.Atoi(data[2])
			freeg := fmt.Sprintf("%d/G", free/1024.0/1024.0/1024.0)
			all, _ := strconv.Atoi(data[3])
			allg := fmt.Sprintf("%d/G", all/1024.0/1024.0/1024.0)
			if freeg == "0/G" || allg == "0/G" {
				freeg = fmt.Sprintf("%d/M", free/1024.0/1024.0)
				allg = fmt.Sprintf("%d/M", all/1024.0/1024.0)
			}
			percentage := (float64(free) / float64(all)) * 100
			percentageStr := fmt.Sprintf("%.2f%%", percentage)

			echo += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", data[0], data[1], freeg, allg, percentageStr)
		}
	}
	html = strings.ReplaceAll(html, "磁盘信息结果", echo)
}