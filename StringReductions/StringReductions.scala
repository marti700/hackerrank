object Solution {
  def strCompress(str: String): String = {
    str.foldLeft("") { (acc, e) =>
      if (acc.indexOf(e) == -1) acc + e else acc + ""
    }
  }

  def main(args: Array[String]) : Unit = {
    val input = io.Source.stdin.getLines().toList.mkString("")
    print(strCompress(input))
  }
}
