import scala.math.abs
import scala.math.pow
import scala.math.sqrt

object Solution {

  def perimeter(pointList: List[(Double, Double)]): Double = {
    mapTuple(distance, pointList, pointList.head).sum

  }

  def distance(point1: (Double, Double), point2: (Double, Double)): Double ={
    sqrt( pow( ( abs( point2._1 - point1._1 ) ), 2 ) + pow( (abs (point2._2 - point1._2 ) ), 2 ) )
  }

  def mapTuple(
      f: ((Double, Double), (Double, Double)) => Double,
      xs: List[(Double,Double)],
      firstPoint: (Double,Double)
  ): List[Double] = {
    xs match {
      case (x::Nil) => f(x, firstPoint) :: Nil
      case (x :: xs) => f(x, xs.head) :: mapTuple(f, xs, firstPoint)
    }
  }

  def toTuple(in: String): (Double, Double) = {
    val tupleString = in.split(' ')
    (tupleString(0).toDouble, tupleString(1).toDouble)q
  }

  def main(args: Array[String]): Unit = {
    val input = io.Source.stdin.getLines().toList.tail.map(toTuple)
    println(perimeter(input))
  }
}
