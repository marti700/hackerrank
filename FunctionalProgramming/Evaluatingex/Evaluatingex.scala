object Solution {
  def e (n: Double) : Double = {
    def fac (f: Double) : Double = {
      f match {
        case 1 => 1
        case _ => f * fac(f-1)
      }
    }

    (1 to 9).toList.foldLeft(1.0) {(acc, elem) => acc + (scala.math.pow(n,elem) /fac(elem)) }
  }

  def main (args: Array[String]) : Unit = {
    val stdin = io.Source.stdin.getLines().toList.map(_.toDouble)
    stdin.tail.foreach(element => println(e(element)))
  }
}

