object ListReplication {
  def main(args: Array[String]) {
    val stdin = io.Source.stdin.getLines().toList
    val l = stdin.map(_.toInt)
    f(l.head, l.tail)
  }

  def f(num: Int, xs: List[Int]): List[Int] = {
    val buffer = scala.collection.mutable.ListBuffer[Int]()

    def go(i: Int, list: List[Int]): List[Int] = {
      list match {
        case x :: xs => { go1(num, x); go(i, xs) }
        case Nil     => buffer.toList
      }
    }

    def go1(j: Int, k: Int): Int = {
      if (j < 1) return 0
      buffer += k
      go1(j - 1, k)
    }

    go(num, xs)
  }
}
