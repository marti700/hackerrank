object ConvexHull {

  /** Finds the point with the minimum Y coordinate if multiple points with the same min Y coordinate are found
    * the point with the lowest X coordinate is taken
    * @param p a list of point represented as a list of pairs of Double
    * @return a pair of Double representing the point with the minimum Y coordinate
    */
  def minY(p: List[(Double, Double)]): (Double, Double) = {

    /** Decides which is have the lowest Y coordinate between two points
      * @param currentMin a pair consisting of the current Y coordinate point minimum found until now
      * @param currentX the current point to be compared with the current minimum
      * @return a pair representing the point with the lowest Y coordinate between the two evaluated points
      */
    def decideBetween(
        currentMin: (Double, Double),
        currentX: (Double, Double)
    ): (Double, Double) = {
      if (currentMin._2 < currentX._2) currentMin
      else currentX
    }

    p match {
      case Nil       => (0, 0)
      case (x :: xs) => xs.foldLeft(x) { (acc, x) => decideBetween(acc, x) }
    }
  }

  /** Find the polar angle of a list of point with respect to a given reference point
    * @param l The reference point (in our case is the point with the lowest Y coordinate)
    * @param p a list of pairs
    * @return a list of tuples (Double,Double,Double) in which the first two are the coordinates of a point in the plane
    *         and the last one is the angle (measured in radians) it makes with the reference point.
    */
  def getPolarAngle(
      l: (Double, Double),
      p: List[(Double, Double)]
  ): List[(Double, Double, Double)] = {
    p.map(x => (x._1, x._2, math.abs(math.atan2(l._2 - x._2, l._1 - x._1))))
  }

  /** Removes the points with the same angle from a list of three element tuple (Double, Double, Double) by taking the
    * point that is farther away from the reference point
    * @param lp The reference point (in our case is the point with the lowest Y coordinate)
    * @param pts A list of points of three element tuples
    * @return a list of three element tuples with the points that have the same angle removed
    */
  def removeSameAnglePoints(
      lp: (Double, Double),
      pts: List[(Double, Double, Double)]
  ): List[(Double, Double, Double)] = {

    /** determins which of the evaluated points is farther from the reference point
      * @param lp the reference point
      * @param a a pair representing a point int he plane
      * @param b a pair representing a point in the plane
      * @return a pair representing the point that is farther away from the reference point
      */
    def farthest(
        lp: (Double, Double),
        a: (Double, Double, Double),
        b: (Double, Double, Double)
    ) = {
      if (distance(lp, (a._1, a._2)) > distance(lp, (b._1, b._2))) a
      else b
    }

    (lp, pts) match {
      case (_, Nil) => Nil
      case (a, (x :: y :: Nil)) if x._3 == y._3 =>
        (farthest(a, x, y)) :: removeSameAnglePoints(a, Nil)
      case (a, (x :: y :: xs)) if x._3 == y._3 =>
        removeSameAnglePoints(a, farthest(a, x, y) :: xs)
      case (a, x :: xs) => x :: removeSameAnglePoints(a, xs)
    }
  }

  /** Find the distance between two points using the euclidean's distance formula
    * @param a a pair of Double representing a point in the plane
    * @param b a pair of Double representing a point in the plane
    * @return the distance between the two given points
    */
  def distance(a: (Double, Double), b: (Double, Double)): Double = {
    math.sqrt(math.pow(b._2 - a._2, 2) + math.pow((b._1 - a._1), 2))
  }

  /** Removes the last element of a three element tuple
    * @param p a List of three element tuples
    * @return a list of pairs
    */
  def removeAngle(p: List[(Double, Double, Double)]): List[(Double, Double)] =
    p.map(x => (x._1, x._2))

  /** Calculates the cross product between tree points in the plane
    * @param a a point in the plane
    * @param b a point in the plane
    * @param c a point in the plane
    * @return -1 if the are formed by the paralelogram of the vectors formed by a b c is negative
    *         0 if the points are collinear
    *         1 if the are formed by the paralelogram of the vectors formed by a b c is positive
    */
  def cX(
      a: (Double, Double),
      b: (Double, Double),
      c: (Double, Double)
  ): Double = {
    (a, b, c) match {
      case ((x1, y1), (x2, y2), (x3, y3)) =>
        ((x2 - x1) * (y3 - y1)) - ((y2 - y1) * (x3 - x1))
    }
  }

  /** Finds the convexHull using the Graham Scan algorithm
    * @param p a list of points
    * @param stack a List to be used as a stack
    * @return a list of pairs representings the points that are part of the convex hull
    */
  def convexHull(
      p: List[(Double, Double)],
      stack: List[(Double, Double)]
  ): List[(Double, Double)] = {
    (p, stack) match {
      case (Nil, stack)       => stack
      case (x :: xs, Nil)     => convexHull(xs, List(x))
      case (x :: xs, List(s)) => convexHull(xs, x :: List(s))
      case (x :: xs, b :: c :: ds) if cX(c, x, b) >= 0 =>
        convexHull(xs, x :: b :: c :: ds)
      case (x :: xs, b :: c :: ds) if cX(c, x, b) < 0 =>
        convexHull(x :: xs, c :: ds)
    }
  }

  /** Finds the perimeter of a polygon
    * @param p a list of pairs representing the points that are part of the polygon
    * @return a Double representing the area of the polygon
    */
  def perimeter(p: List[(Double, Double)]): Double = {
    p match {
      case Nil          => 0.0
      case List(a, b)   => distance(a, b)
      case a :: b :: xs => distance(a, b) + perimeter(b :: xs)
    }
  }

  /** Converts a string of the format  'Number Number' to a pair of doubles
    * @param s String in the form 'Number Number'
    * @return a pair of double
    */
  def toTupleDouble(s: String): (Double, Double) = {
    val tuple = s.split(" ")
    (tuple(0).toDouble, tuple(1).toDouble)
  }

  def main(args: Array[String]): Unit = {
    val l = io.Source.stdin.getLines().toList.tail.map(x => toTupleDouble(x))
    val min = minY(l)
    val sortedPoints = removeAngle(
      removeSameAnglePoints(min, getPolarAngle(min, l).sortBy(_._3))
    )
    val ch = convexHull(sortedPoints, Nil)
    println(perimeter(ch ++ List(ch.head)))
  }
}
