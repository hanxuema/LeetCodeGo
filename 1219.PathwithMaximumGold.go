package main

func getMaximumGold(grid [][]int) int {
    res := 0
    if grid == nil || len(grid) == 0 {
        return res
    }
    
    dir := make([][]int, 4)
    dir[0] = []int{1,0}
    dir[1] = []int{0,1}
    dir[2] = []int{-1,0}
    dir[3] = []int{0,-1}
    max := 0
    for i := 0; i < len(grid);i++{
        for j:=0; j <len(grid[i]); j++{
            if grid[i][j] > 0{
                backtrack1219(grid,i,j ,dir, 0, &max)
            }  
        }
    }
    return max
}

func backtrack1219(grid [][]int, n int, m int, dir [][]int, res int, max *int) int{
    res += grid[n][m]
    if res>*max{
        *max=res
    }
 
    temp:=grid[n][m]
    grid[n][m] = 0
    for index := 0; index < len(dir); index++{
        nn := n + dir[index][0] //must create new variable, otherwise it will change the value outside loops
        mm := m + dir[index][1]
        if nn < 0 || nn >= len(grid) || mm < 0 || mm >= len(grid[nn]) || grid[nn][mm] == 0 {
            continue
        } 
        
        backtrack1219(grid, nn, mm, dir, res, max)  
    }
    grid[n][m] = temp
 
     return res
   //remove
  
}
