object Solution {
  def strCompress (str: String) : String = {
    group(str.toList).flatMap(x => {
      if (x.length > 1) x.head.toString + x.length.toString
      else x.head.toString
    }).foldLeft("") {(acc, e) => acc + e}
  }

  /**
    * Groups in a list the contiguous equal characters
    *
    * @param str a list of Chars
    * @return a list of list where the inner list contains the grouped equal characters
    */
  def group(str: List[Char]) : List[List[Char]] = {
    str match {
      case Nil => Nil
      case (x::xs) => (x :: xs.takeWhile(y => x == y)) :: group(xs.dropWhile(y => y == x))
    }
  }

  def main(args: Array[String]) : Unit = {
    val input = io.Source.stdin.getLines().toList.head
    print(strCompress(input))
  }
}