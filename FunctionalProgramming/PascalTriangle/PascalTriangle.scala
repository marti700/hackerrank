object Solution {
  def pascalTriangle(rows: Int): List[List[scala.math.BigInt]] = {
    (1 to rows.intValue()).toList.map(triangleRow(_))
  }

  def triangleRow (row: scala.math.BigInt) : List[scala.math.BigInt] = {
    for {
      i <- List(row-1)
      j <- (0 to (row-1).intValue()).toList
    } yield nCr(i,j)
  }

  def nCr (n: scala.math.BigInt, k: scala.math.BigInt) : scala.math.BigInt = (fac(n)) / (fac(k) * fac(n-k))

  def fac(n: scala.math.BigInt) : scala.math.BigInt = {
    var res = 1
    for (x <- 1 to n.intValue())
      res = res * x
    return res
  }

  def mapElementsToString(in : List[scala.math.BigInt]) : String = {
    in match {
      case Nil => "\n"
      case (x::xs) => x.toString() + " " + mapElementsToString(xs)
    }
  }

  def buildOutput(in: List[List[scala.math.BigInt]]) : String = {
    in.foldLeft("") {(acc, x) =>  acc + mapElementsToString(x)}
  }

  def main(args: Array[String]) : Unit = {
    val input = io.Source.stdin.getLines().toList.map(_.toInt)
    print(buildOutput(pascalTriangle(input.head)))

  }
}