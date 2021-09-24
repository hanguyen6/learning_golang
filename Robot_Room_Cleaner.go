// https://leetcode.com/problems/robot-room-cleaner/
/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 * 
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

type Point struct {
    row int
    col int
}


func goBack(robot *Robot) {
    robot.TurnRight()
    robot.TurnRight()
    robot.Move()
    robot.TurnRight()
    robot.TurnRight()
}

func backTrack(robot *Robot, row int, col int, d int, visited *map[Point]bool) {
    
    directions := [4][2]int{  {-1, 0},    // up 
                                {0, 1},     // right
                                {1, 0},     // down
                              {0, -1} }   // left
                

    
    // Mark cell as visited and clean the cell 
    p := Point{row, col}
    (*visited)[p] = true 
    robot.Clean()
    
    // Keep moving in clockwise direction 
    for i:=0; i < 4 ; i++ {
        newD := (d + i) % 4
        newRow := row + directions[newD][0]
        newCol := col + directions[newD][1]
        newPoint := Point{newRow, newCol}
        
        _, found := (*visited)[newPoint]
        if (!found && robot.Move()) {
            backTrack(robot, newRow, newCol, newD, visited)
            goBack(robot)
        }
        
        robot.TurnRight()          
    }
}

func cleanRoom(robot *Robot) {
    visited := make(map[Point]bool)
    backTrack(robot, 0, 0, 0, &visited)
}
