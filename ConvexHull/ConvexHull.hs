import Text.Printf ( printf )
import Data.List ( sortBy)

-- Finds the min y-axis coordinate in an array of 2D points, if two y coordinates have the same value
-- this function returns the one with the lower x coordinate value
minY ::  [(Double,Double)] -> (Double,Double)
minY [] = (0,0)
minY (x:xs) = foldl (flip decideBetween) x xs
  where
    decideBetween currentMin currentX
      | snd currentMin == snd currentX = if fst currentMin < fst currentX then currentMin else currentX
      | otherwise = if snd currentMin < snd currentX then currentMin else currentX

-- Calculates the polar angle of all the points in respect to a given point and add the resutl as the third element of a new tupple
getPolarAngle :: (Double, Double) -> [(Double, Double)] -> [(Double, Double, Double)]
getPolarAngle a = map (\x -> (fst x, snd x, abs (atan2 (snd a - snd x) (fst a - fst x))))

-- returns the thrid element of a three element tuple
trd :: (a, b, c) -> c
trd (_,_,x) = x


-- finds the convex hull using the Graham scan alorith
-- uses the cross product (cX function) to determinate if the next point is a left or right turn
-- recieves as argument the point with the lowest Y coordinate, the points and an empty list that will server as the stack
-- https://www.youtube.com/watch?v=B2AJoQSZf4M
convexHull :: (Double,Double) -> [(Double,Double )] -> [(Double,Double)] -> [(Double, Double)]
convexHull _ [] ys = ys
convexHull initialValue xs [] = convexHull initialValue xs [initialValue]
convexHull initialValue (x:xs) [y] = convexHull initialValue xs (x:[y])
convexHull initialValue (x:xs) (b:c:ds)
  | cX b x c >= 0 = convexHull initialValue xs (x:b:c:ds)
  | cX b x c < 0 = convexHull initialValue (x:xs) (c:ds)
  where
    cX (x1,y1) (x2,y2) (x3,y3) = ((x2-x1) * (y3-y1)) - ((y2 - y1) * (x3-x1))


-- removes the points that have the same angle by taking the point which is farther away
-- from the lowest Y coordinate point
removeSameAnglePoints:: (Double, Double) -> [(Double,Double, Double)] -> [(Double, Double, Double)]
removeSameAnglePoints _ [] = []
removeSameAnglePoints a (x:y:[])
  | trd x == trd y = (farthest a x y) : removeSameAnglePoints a []
  | otherwise = [x,y]
removeSameAnglePoints a (x:y:xs)
  | trd x == trd y = removeSameAnglePoints a ((farthest a x y) : xs)
  | otherwise = x : removeSameAnglePoints a (y:xs)

-- Converts a three element tuple to a pair by removien the last touple element
maketouple :: (a, b, c) -> (a, b)
maketouple (a,b,_) = (a,b)

-- finds the point that is farther away from a given reference point
farthest :: (Double, Double) -> (Double, Double, Double) -> (Double, Double, Double) -> (Double, Double, Double)
farthest a b c
  | distance a (maketouple b) > distance a (maketouple c) = b
  | otherwise = c

-- finds the perimeter of a polygon
perimeter :: [(Double, Double)] -> Double
perimeter [] = 0
perimeter [a,b] = distance a b
perimeter (a:b:xs) = distance a b + perimeter (b:xs)

-- finds the disctance bewtween two points
distance :: (Double, Double) -> (Double, Double) -> Double
distance a b = sqrt((fst b- fst a)^2  + (snd b - snd a)^2 )

-- removes the last element of all tuples in a list of Touples and returns a list of pairs
ss :: [(a, b, c)] -> [(a, b)]
ss = map (\(x,y,_) -> (x,y))

-- convert a list of string into a list of pairs
toDouble:: [String] -> [(Double, Double)]
toDouble [] = []
toDouble (a:b:xs) = (read a :: Double, read b :: Double) : toDouble xs

-- sorting function, used with the sortBy function to sort the points by angle
xxx (a,b,c) (d,e,f)
  | c > f = GT
  | otherwise = LT


main :: IO ()
main = do
  (t:input) <- fmap lines getContents
  let inputAsDouble = toDouble $  concatMap words input
  let minP = minY inputAsDouble
  let sortedPoints = ss $ removeSameAnglePoints minP (sortBy xxx (getPolarAngle minP inputAsDouble))
  let convHull = convexHull minP (reverse sortedPoints) []
  printf "%.1f" (perimeter (convHull ++ [head convHull]))
