object ArrayOfNElements {
  def main(args: List[String]) : Unit = {
    val n = scala.io.StdIn.readInt()
    println(f(n))
  }
  def f(n: Int): List[Int] = {
    n match {
      case 0 => Nil
      case _ => n::f(n-1)
    }
  }
}