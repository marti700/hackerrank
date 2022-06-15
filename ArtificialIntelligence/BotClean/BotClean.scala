object Solution {
  /**
   * Converts an array of strings to a cuadratic matrix, the dimensions
   * of the matrix are given by the @param shape
   *
   * @param shape the dimension that the final matrix will have
   * @param arr   a List of String
   * @return an Array[Array[Char]] (a Matrix) wich dimension is the same as the shape parameter
   **/
  def convertToMatrix(shape: Int, arr: Array[String]): Array[Array[Char]] = {
    val newGrid = Array.ofDim[Char](shape, shape)

    for (i <- 0 until arr.length)
      newGrid(i) = arr(i).toArray

    newGrid
  }

  /**
   * given the size of the grid (which is always cuadratic)
   * retuns if a given coordinate is on the grid
   *
   * @param gridSize the size of the grid (the dimension of the matrix)
   * @param pos      an x, y coordinate as a tuple of Ints (Int, Int)
   * @return true if the given coordinate is out of the grid limits, false otherwise
   */
  def outOutOfLimits(gridSize: Int, pos: (Int, Int)): Boolean = {
    if ((pos._1 < 0) || (pos._1 > gridSize - 1)) return true

    if ((pos._2 < 0) || (pos._2 > gridSize - 1)) return true

    false
  }

  /** Returns the location of the bot. Initially the bot is always located in the middle of the grid
   *
   * @param gridSize the size of the map. The grid is always a cuadratic matrix
   * @return a tuple representing the x,y coordinates of the center of the grid
   */
  def locateBot(gridSize: Int): (Int, Int) = {
    val s = gridSize / 2
    (s, s)
  }

  /**
   * perform a BFS search of the grid to find the paths to the nearest dirty zone
   *
   * @param grid    the grid
   * @param pos     the position on which the bot currenlty is
   * @return a List of tuples holding the shortest path that lead to the dirty spot
   */

  def bfs(grid: Array[Array[Char]], pos: (Int, Int)): List[(Int, Int)] = {

    val visited = scala.collection.mutable.Map[String, Boolean]()

    val queue = scala.collection.mutable.Queue.empty[List[(Int, Int)]];
    queue.enqueue( List(pos) )

    while (queue.nonEmpty) {
      val currentPath = queue.dequeue()
      val currentNode = currentPath.head
      if (visited.get(currentNode.toString()).isEmpty) {
        if (grid(currentNode._1)(currentNode._2) == 'd') {
          return (-8,-8)::currentPath // (-8,-8) represents a clean operation
        }
        else {
          visited(currentNode.toString()) = true
          val nei = getNeighbors(grid.length, currentNode)
          nei.foreach(n => {
            queue.enqueue(n::currentPath)
          })
        }
      }
    }
    null
  }

  /**
   * Search recursively the next spot to clean by using BFS to find the nearest dirty spot
   * @param grid an matrix of chars representing the grid
   * @param pos the position on which the bot is
   * @return a tuple list representing the coordinates the bot traveled along
   */
  def botClean(grid: Array[Array[Char]], pos: (Int,Int)) : List[List[(Int,Int)]] = {
    val newGrid = grid.clone()
    val pathToDirt = bfs(newGrid,pos);

    pathToDirt match {
      case null => Nil
      case _ => {
        val dirtCoordinates = pathToDirt.tail.head
        newGrid(dirtCoordinates._1)(dirtCoordinates._2) = '-'
        pathToDirt::botClean(newGrid, dirtCoordinates)
      }
    }
  }

  /**
   * gets the neighbors of this node, since the bot can only move up, down, right or left
   * in the grid just this directions are taken in to consideration
   *
   * @param pos the node (as a (Int, Int) tuple) of which we want to know the neighbors
   * @return the neighbors of the node represented by the pos param
   */
  def getNeighbors(gridSize: Int, pos: (Int, Int)): List[(Int, Int)] = {
    val n = scala.collection.mutable.ListBuffer[(Int, Int)]()

    // if there are neighbors to the right
    if (!outOutOfLimits(gridSize, (pos._1, pos._2 + 1)))
      n.append((pos._1, pos._2 + 1))
    //if there are neighbors up
    if (!outOutOfLimits(gridSize, (pos._1 - 1, pos._2)))
      n.append((pos._1 - 1, pos._2))
    //if there are neighbors to the left
    if (!outOutOfLimits(gridSize, (pos._1, pos._2 - 1)))
      n.append((pos._1, pos._2 - 1))
    // if there are neighbors down
    if (!outOutOfLimits(gridSize, (pos._1 + 1, pos._2)))
      n.append((pos._1 + 1, pos._2))

    n.toList
  }

  /**
   * given the initial position translate a list of path to RIGHT, LEFT, UP, DOWN
   * depending on their coordinates relative to the initial point
   *
   * @param initialPos the reference position from which the directions will be calculated
   * @param positions  a List of paths
   * @return a List of String representing the turns (Right, left, up, down) taken by the bot to arrive to the princess
   */
  def getTurns(initialPos: (Int, Int), positions: List[(Int, Int)]): List[String] = {
    positions match {
      case Nil => Nil
      case x :: xs if x._1 == initialPos._1 && x._2 == initialPos._2 + 1 => "RIGHT" :: getTurns(x, xs)
      case x :: xs if x._1 == initialPos._1 && x._2 == initialPos._2 - 1 => "LEFT" :: getTurns(x, xs)
      case x :: xs if x._1 == initialPos._1 - 1 && x._2 == initialPos._2 => "UP" :: getTurns(x, xs)
      case x :: xs if x._1 == initialPos._1 + 1 && x._2 == initialPos._2 => "DOWN" :: getTurns(x, xs)
      case x :: xs if x._1 == -8 && x._2 == -8 => "CLEAN" :: getTurns(x,xs)
    }
  }


  def main(args: Array[String]): Unit = {
    val x_y = scala.io.StdIn.readLine().split(" ")
    val r = x_y(0).toInt
    val c = x_y(1).toInt
    val grid = new Array[String](5)

    for (i <- 0 until 5) {
      grid.update(i, scala.io.StdIn.readLine())
    }

    val botLocation = (r, c)

    val botMoves = botClean(convertToMatrix(5, grid), botLocation)
      .flatMap(x => getTurns(x.last, x.reverse.tail))

    println(botMoves.head)

  }
}
