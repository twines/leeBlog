package utils

import (
	"fmt"
	"html/template"
)

func Paginate(path string, totalPage int, currentPage int) template.HTML {
	if totalPage <= 1 {
		return ""
	}
	tpl := `<ul class="pagination">`
	for i := 1; i <= totalPage; i++ {
		if i == 1 {
			if currentPage == 1 {
				tpl += fmt.Sprintf("<li class=\"page-item disabled\"><a class=\"page-link text-dark\" href=\"%s\"  tabindex=\"-1\" aria-disabled=\"true\">%s</a></li>", path, "上一页")
				tpl += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link text-white bg-dark\" href=\"%s?page=%d\">%d</a></li>", path, i, i)
			} else {
				tpl += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link text-white bg-dark\" href=\"%s?page=%d\">%s</a></li>", path, currentPage-1, "上一页")
				tpl += fmt.Sprintf("<li class=\"page-item \"><a class=\"page-link text-dark\" href=\"%s?page=%d\">%d</a></li>", path, i, i)
			}
		} else if i == totalPage {
			if currentPage == totalPage {
				tpl += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link text-white bg-dark\" href=\"%s?page=%d\">%d</a></li>", path, i+1, i)
				tpl += fmt.Sprintf("<li class=\"page-item disabled\"><a class=\"page-link text-dark\" href=\"%s\"  tabindex=\"-1\" aria-disabled=\"true\">%s</a></li>", path, "下一页")
			} else {
				tpl += fmt.Sprintf("<li class=\"page-item \"><a class=\"page-link  text-dark\" href=\"%s?page=%d\">%d</a></li>", path, i, i)
				tpl += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link text-white bg-dark\" href=\"%s?page=%d\">%s</a></li>", path, currentPage+1, "下一页")
			}
		} else {
			if i == currentPage {
				tpl += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link text-white bg-dark\" href=\"%s?page=%d\">%d</a></li>", path, i, i)
			} else {
				tpl += fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link  text-dark\" href=\"%s?page=%d\">%d</a></li>", path, i, i)
			}
		}
	}
	tpl += `</ul>`
	return template.HTML(tpl)
}
