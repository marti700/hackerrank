object BotSavesPrincess {
  /**
    * Converts an array of strings to a cuadratic matrix, the dimensions
    * of the matrix are given by the @param shape
    *
    * @param shape the dimension that the final matrix will have
    * @param arr a List of String
    * @return an Array[Array[Char]] (a Matrix) wich dimension is the same as the shape parameter
    */
  def convertToMatrix(shape: Int, arr: Array[String]): Array[Array[Char]] = {
    val newGrid = Array.ofDim[Char](shape, shape)

    for (i <- 0 until arr.length)
      newGrid(i) = arr(i).toArray

    newGrid
  }

  /**
    * since the princess can only be in the corners of the map. Search each corner
    *
    * @param shape the size of the map (the map is always cuadratic)
    * @param grid The map itself
    * @return a tuple of ints the represents the x, y coordinates where the princess is
    */
  def locatePrincess(shape: Int, grid: Array[Array[Char]]): (Int, Int) = {
    if (grid(0)(0) == 'p')  return (0,0)
    if (grid(0)(shape-1) == 'p') return (0, shape-1)
    if (grid(shape-1)(0) == 'p') return (shape-1, 0)
    if (grid(shape-1)(shape-1) == 'p') return (shape-1, shape-1)

    null
  }

  /**
    * Since the bot is always in the middle of the grid and the princess in one
    * of the corners the shortest path is always Down or up and then to the left or right
    * depending of where the princess.
    *
    * This function creats a List with direnctions to the princess using the above logic
    *
    * @param shape the dimension of the map (always cuadratic)
    * @param pLocation (the location of the princess)
    * @return a List containing the shortest path to the princess
    */
  def inferShortestPath(shape: Int, pLocation: (Int, Int)): List[String] = {
    val upperRangeLimit = ((shape/2)) *2

      if (pLocation._1 == 0 && pLocation._2 == 0)
        return Range(0, upperRangeLimit).map(a => if (a < upperRangeLimit/2) "UP" else "LEFT").toList
      if (pLocation._1 == 0 && pLocation._2 ==  shape -1)
        return Range(0, upperRangeLimit).map(a => if (a < upperRangeLimit/2) "UP" else "RIGHT").toList
      if (pLocation._1 == shape-1 && pLocation._2 ==  0)
        return Range(0, upperRangeLimit).map(a => if (a < upperRangeLimit/2) "DOWN" else "LEFT").toList
      if (pLocation._1 == shape-1 && pLocation._2 == shape-1)
        return Range(0, upperRangeLimit).map(a => if (a < upperRangeLimit/2) "DOWN" else "RIGHT").toList

    List.empty
  }

  def main(args: Array[String]): Unit = {
    //read grid size
    val m = scala.io.StdIn.readLine().toInt

    val grid = new Array[String](m)

    // build the grid
    for (i <- 0 until m) {
      grid.update(i, scala.io.StdIn.readLine())
    }

    inferShortestPath(m,
      locatePrincess(m,
        convertToMatrix(m,grid)
      )
    ).foreach(println)
    // val duration = (System.nanoTime - t1)
  }
}
