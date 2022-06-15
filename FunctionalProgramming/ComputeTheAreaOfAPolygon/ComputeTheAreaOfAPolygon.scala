import scala.math.abs
import scala.math.pow
import scala.math.sqrt

object Solution {

  /**
    * Calculate the perimeter of a poligon given it's vertices
    *
    * @param pointList a list of tuples
    * @return the perimeter of the poligon
    */
  def perimeter(pointList: List[(Double, Double)]): Double = {
    mapTuple(cross, pointList, pointList.head).sum
  }

  /**
    * applys the formula described here https://www.mathopenref.com/coordpolygonarea.html
    * to two given coordinates.
    *
    * @param point1 a tuple representing the x,y coordinates of the first point
    * @param point2 a tuple representing the x,y coordinates of the second point
    * @return the distance between two points
    */
  def cross(point1: (Double, Double), point2: (Double, Double)): Double = {
    ( (point1._1 * point2._2) - (point1._2 * point2._1) ) / 2
  }

  /**
    * Applys a (Double, Double), (Double,Double) => Double function to a list of points
    *
    * @param f a (Double, Double), (Double,Double) => Double
    * @param xs a list of (Double, Double)
    * @param firstPoint the firt point of the polygon
    * @return a list of (Double, Double)
    */
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

  /**
    * Transform a string in the format "2 2" to a tuple (2.0, 2.0)
    *
    * @param in a String in consisting of two numbers separeted by an space
    * @return a touble of (Double, Double)
    */
  def toTuple(in: String): (Double, Double) = {
    val tupleString = in.split(' ')
    (tupleString(0).toDouble, tupleString(1).toDouble)
  }

  def main(args: Array[String]): Unit = {
    val input = io.Source.stdin.getLines().toList.tail.map(toTuple)
    println(perimeter(input))
  }
}