package main

import (
  "fmt"
)

// This functions determines in what sheet is a page (a sheet is formed by two pages)
// ex page 1 and 2 are in sheet 1, pages 3 and 4 are in sheet 2
func getSheet(page int) int {
  //when the given page is even
  if page%2 == 0 {
    return page / 2
  }

  return (page + 1) / 2 // when the given page is odd
}

func countFromTheFront(page int) int {
  if page%2 == 0 {
    return getSheet(page)
  }
  return getSheet(page) - 1
}

func countFromTheBack(pages, page int) int {
  // the total number of page turns to get from the start of the book to the last page
  // minus the total number of pages to get fron the start of the boot to the page of interest
  return countFromTheFront(pages) - countFromTheFront(page)
}

func pageCount(pages, page int) int {
  fromTheFront := countFromTheFront(page)
  fromTheBack := countFromTheBack(pages, page)

  if fromTheFront < fromTheBack {
    return fromTheFront
  } else {
    return fromTheBack
  }
}

func main() {
  var pages, page int
  fmt.Scanf("%d\n%d\n", &pages, &page)
  fmt.Println(pageCount(pages, page))
}
