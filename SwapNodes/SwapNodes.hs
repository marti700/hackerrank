-- Swap  Nodes
--
-- Tree data structure
import Data.Vector.Storable.Mutable (swap)
data Tree a = E | Node a (Tree a) (Tree a) deriving (Show, Read, Eq)

-- In order traversal algorithm
traverseInOrder :: Tree a -> [a]
traverseInOrder E = []
traverseInOrder (Node a l r) = traverseInOrder l ++ [a] ++ traverseInOrder r

swapNodeAt :: Tree a -> Integer -> Tree a
swapNodeAt E _ = E
swapNodeAt (Node a l r) 1 = Node a r l
swapNodeAt (Node a l r) d = Node a (swapNodeAt l (d-1)) (swapNodeAt r (d-1))

makeSingletonTree :: Integer -> Tree Integer
makeSingletonTree (-1) = E
makeSingletonTree x = Node x E E

joinT :: Tree a -> Tree a -> Tree a -> Tree a
joinT E _ _ = E
joinT (Node a left right) leftLeaf rightLeaf = Node a leftLeaf  rightLeaf

buildTree :: Tree Integer -> [[Integer]] -> [Integer] -> Tree Integer
buildTree E xs [] = buildTree (makeSingletonTree 1) xs [1]
-- buildTree E _ (_:_) = E
buildTree (Node a left right) [] _ = Node a left right
buildTree (Node a left right) (x:xs) buff =
  buildTree (appendT (Node a left right) (makeSingletonTree $ head buff) x) xs (tail buff ++ ignoreNull [head x, x !! 1])

-- TODO usit as a where in the buidTree function
ignoreNull :: [Integer] -> [Integer]
ignoreNull = filter (/= -1)


appendT :: Tree Integer -> Tree Integer -> [Integer] -> Tree Integer
appendT E _ _= E
-- appendT (Node _ _ _) E _ = E
appendT (Node a left right) (Node b left1 right1) strN
  | (a == b) && (left == left1) && (right == right1) = joinT (Node a left right) (makeSingletonTree (head strN)) (makeSingletonTree (strN !! 1))
  | otherwise = Node a (appendT left (Node b left1 right1) strN) (appendT right (Node b left1 right1) strN)

convertToIntList :: [String] -> [[Integer]]
convertToIntList = map (map (read :: String -> Integer) . words)

solve :: Tree Integer -> [Integer] -> [[Integer]]
solve t [] = []
solve t (x:xs) =  traverseInOrder (swapT t x) : solve (swapT t x) xs
  where swapT tt h = swapNodesAtHeights tt (heights h)

heights :: Integer -> [Integer]
heights x = [x*j | j <- [1..] ]

swapNodesAtHeights :: (Eq a) => Tree a -> [Integer] -> Tree a
swapNodesAtHeights tree [] = tree
swapNodesAtHeights tree (x:xs)
  | tree == swapNodeAt tree x = tree
  | otherwise = swapNodesAtHeights (swapNodeAt tree x) xs

readLines :: Int -> IO [String]
readLines 0 = return []
readLines x = do
  line <- getLine
  rest <- readLines (x-1)
  return (line : rest)

formatOutput :: [Integer] -> String
formatOutput = foldl (\acc x -> if acc == "" then  acc ++ show x else acc ++ " " ++ show x) ""


main :: IO ()
main = do
  numberOfNodes <-  readLn :: IO Int

  --Handle the tree nodes part of the input
  numberOfTreeNodes <- readLines numberOfNodes
  let treeNodes = convertToIntList numberOfTreeNodes
  numberOfSwapOperations <- readLn :: IO Int

  -- Handle the swap operation part of the input
  swapOpDepth <- readLines numberOfSwapOperations
  let swapOperations = convertToIntList swapOpDepth


  -- Builds the tree
  let tree = buildTree E treeNodes []
  -- Applies the solution
  let x =  solve tree (concat swapOperations)
  -- print to stdout
  --convert to string
  let r = map (unlines .map show) x
  mapM_ (putStrLn . formatOutput) x
  -- print swapOperations