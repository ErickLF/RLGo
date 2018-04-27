package main

import (
    "fmt"
    "math/rand"
    "time"
)

const(
	up int = iota
	down
	left
	right
    up_left
    up_right
    down_left
    down_right
	numactions int = 8
	worldheight=7
	worldwidth=10
	)
//----------------------------------------------------------------------
type Sarsa struct{
	Q[][]float64
	startX int
	startY int
	goalX int
	goalY int
	epsilon float64
	alpha float64
	gamma float64
	wind [] int
}
//----------------------------------------------------------------------
func main(){
	rand.Seed(time.Now().Unix())
	fmt.Println(" KING WINDY GRID WORLD")
	sarsa:=Sarsa{}
	sarsa.InitSarsa()
	sarsa.Sarsas(2000)
	sarsa.Print()
}
//----------------------------------------------------------------------
func(q *Sarsa) InitSarsa(){
	//Initialize Q(s,a),for all s belonging to S, a belonging to A(s), arbitrarily, and Q(terminal-state,.)=0
	q.startX,q.startY = 0,3
	q.goalX,q.goalY = 7,3
	q.epsilon,q.gamma,q.alpha=0.1,1.0,0.5
	q.wind = []int{0, 0, 0, 1, 1, 1, 2, 2, 1, 0}
	q.Q = make([][]float64,worldheight*worldwidth)

	for i:=0;i<worldheight*worldwidth;i++{ //initialize the size of the matrix
        q.Q[i] = make([]float64,numactions)
    }
	
    for j:=0;j<worldheight*worldwidth;j++{
    	for i:=0;i<numactions;i++{  
            q.Q[j][i]=rand.Float64()
        }
    }	
}
//----------------------------------------------------------------------
func (s*Sarsa)Sarsas(episodes int){
	for i:=0;i<episodes;i++{ //Repeat (for each episode)
		sx,sy:=s.startX,s.startY //Initialize S
		action:=s.EpsilonGreedy(sx,sy) //Choose A from S using policy derived from Q(e.g.,epsilon-Greedy)
		for sx!=s.goalX || sy!=s.goalY{ //Repeat (for each step of the episode): that is, until you reach the goal
      		reward,sq_sx,sq_sy:=s.Move(sx,sy,action)//Take action A, observe R, S'//sq -> single quote	
			sq_action:=s.EpsilonGreedy(sq_sx,sq_sy)//Choose  A' from S' using policy derived from Q (e.g.,epsisol-Greedy)
			//Q(S,A) <- Q(S,A) + alpha[R+gammaQ(S',A')-Q(S,A)]
			QSA:=s.Q[GetGridPos(sx,sy)][action]
			sqQSA:=s.Q[GetGridPos(sq_sx,sq_sy)][sq_action]
			Q:= QSA + s.alpha*(reward+s.gamma*sqQSA-QSA)
			s.Q[GetGridPos(sx,sy)][action] = Q
			//S <- S'; A <- A'//until S is terminal
			sx,sy=sq_sx,sq_sy
			action=sq_action
		}
	}
}
//----------------------------------------------------------------------
func (q *Sarsa) EpsilonGreedy(x,y int) int{ //epsilon greedy
	if rand.Float64() < 1-q.epsilon { return q.GetAction(x, y) }
	return rand.Intn(numactions)
}
//----------------------------------------------------------------------
func (ql* Sarsa) GetAction(x,y int) int {//both parameters are int and returns int
	gridPos := GetGridPos(x,y) //position in grid
	max := ql.Q[gridPos][0] //taking the first action of position in grid
	action := 0
	for i := 0; i < numactions; i++ {
		if max < ql.Q[gridPos][i] {
			max = ql.Q[gridPos][i]
			action = i
		}
	}
	return action
}
//----------------------------------------------------------------------
func (ql* Sarsa) Move(x,y,action int) (float64,int,int){
  qx,qy:=x,y
  switch action{
      case up:
        if y!=worldheight-1{ qy=y+1 }
      case down:
        if y!=0{ qy=y-1 }
      case left:
        if x!=0{ qx=x-1 }
      case right:
        if x!=worldwidth-1{ qx=x+1 }
      case up_left:
        if x!=0{ qx = x-1 }
        if y!=worldheight-1{ qy=y+1 }
      case up_right:
       if x != worldwidth-1 { qx = x + 1 }
       if y != worldheight-1 { qy = y + 1 }
      case down_left:
       if x != 0 { qx = x - 1 }
       if y != 0 { qy = y - 1 }
      case down_right:
       if x != worldwidth-1 { qx = x + 1 }
       if y != 0 { qy = y - 1 }
   }
   qy+= ql.wind[x]
   if qy > worldheight-1{ qy=worldheight-1 }  
    
   if qy==ql.goalY && qx==ql.goalX{ //If reaches goal
		return 0.0, ql.goalX, ql.goalY 
  }
	
  return -1.0,qx,qy//if doesnt reach goal OR doesnt fall on THE CLIFF just continue on board
}
//----------------------------------------------------------------------
func GetGridPos(x,y int) int{
	return y * worldwidth + x
}
//----------------------------------------------------------------------
func (q *Sarsa) Print() {
	for i := worldheight - 1; i >= 0; i-- {
		for j := 0; j < worldwidth; j++ {
			if i != q.goalY || j != q.goalX {
					switch q.GetAction(j, i) {
						case up:
							fmt.Print("↑  ")
						case down:
							fmt.Print("↓  ")
						case left:
							fmt.Print("←  ")
						case right:
							fmt.Print("→  ")
                        case up_left:
                            fmt.Print("↖  ")
                        case up_right:
                            fmt.Print("↗  ")
                        case down_left:
                            fmt.Print("↙  ")
                        case down_right:
                            fmt.Print("↘  ")
						}
			} else {
				fmt.Print("G  ")
			}
		}
		fmt.Println("")
	}
    
    for i:= 0; i<len(q.wind); i++{
        fmt.Print(q.wind[i],"  ")
    }
    fmt.Println("= wind")
}