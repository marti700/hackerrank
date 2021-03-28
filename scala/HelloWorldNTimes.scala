object Solution {
  def main(args: Array[String]) {
    val times = scala.io.StdIn.readInt()
    f(times)
    }

    def f(n: Int) : Int = {
      if (n < 0)
        return 0
      println("Hello World")
      f(n-1)
    }
}