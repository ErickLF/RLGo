package main
import (
    "fmt"
    "math/rand"
    "time"
)
//----------------------------------------------------------------------
const (
      up int = iota
      down
      left
      right
      worldheight int = 4 //world height
      worldwidth int = 12 //world width
      numactions int = 4 //number of actions
)
//----------------------------------------------------------------------
type QLearning struct{
    Q [][]float64
    alpha float64  //exange rate
    gamma float64  //how much does the past matter?
    epsilon float64 
    startX int
    startY int
    goalX int 
    goalY int 
}
//----------------------------------------------------------------------
func main() {
    rand.Seed(time.Now().Unix())
    ql:=QLearning{}
	ql.Initialize()
	ql.QLearning(1000)
	ql.Print()
}
//----------------------------------------------------------------------
func (ql* QLearning) Initialize(){
	ql.startX,ql.startY = 0,0
    ql.goalX,ql.goalY = 11,0
    ql.alpha,ql.gamma,ql.epsilon=0.5,1.0,0.1
    ql.Q = make([][]float64,worldheight*worldwidth) //Nuestra Q es representada con el tablero 
	
    for i:=0;i<worldheight*worldwidth;i++{ //initialize the size of the matrix
        ql.Q[i] = make([]float64,numactions)
    }
	
    for j:=0;j<worldheight*worldwidth;j++{
    	for i:=0;i<numactions;i++{  
            ql.Q[j][i]=rand.Float64()
        }
    }	
}
//----------------------------------------------------------------------
func (ql* QLearning) QLearning(episode int){
    for i:=0;i<episode;i++{//Repeart (for each step of episode )
        sx,sy:=ql.startX,ql.startY //Initialize S
        for sx!=ql.goalX || sy!=ql.goalY{ //loop until reaches goal
            action:=ql.EpsilonGreedy(sx,sy)  //choose A from S using policy derived from Q
            //take action A,observe R,S' //sq -> single quote
            R,sqsx,sqsy:=ql.Move(sx,sy,action)
            QSA:=ql.Q[GetGridPos(sx,sy)][action]
            maxAction:=ql.GetAction(sx,sy)
            sqQSA:=ql.Q[GetGridPos(sqsx,sqsy)][maxAction]
            //Q(S,A)<-Q(S,A)+alpha*[R+gammaMax v(a) Q(S',a)-Q(S,A)]
            Q:=QSA+ql.alpha*(R+ql.gamma*(sqQSA-QSA))
			ql.Q[GetGridPos(sx,sy)][action]=Q
            //S<-S' //until S is terminal
            sx,sy=sqsx,sqsy
        }
    }
}
//----------------------------------------------------------------------
func (q *QLearning) EpsilonGreedy(x,y int) int{ //epsilon greedy
	if rand.Float64() < 1-q.epsilon { return q.GetAction(x, y) }
	return rand.Intn(numactions)
}
//----------------------------------------------------------------------
func (ql* QLearning) GetAction(x,y int) int {//both parameters are int and returns int
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
func (ql* QLearning) Move(x,y,action int) (float64,int,int){
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
  }
  if qy == 0 && qx >= 1 && qx < worldwidth-1 { //If falls in the cliff
		return -100.0, 0, 0
  }else if qy==ql.goalY && qx==ql.goalX{ //If reaches goal
		return 0.0, ql.goalX, ql.goalY 
  }
	
  return -1.0,qx,qy//if doesnt reach goal OR doesnt fall on THE CLIFF just continue on board
}
//----------------------------------------------------------------------
func GetGridPos(x,y int) int{
	return y * worldwidth + x
}
//----------------------------------------------------------------------
func (q *QLearning) Print() {
	for i := worldheight - 1; i >= 0; i-- {
		for j := 0; j < worldwidth; j++ {
			if i != q.goalY || j != q.goalX {
				if j > 0 && j < worldwidth-1 && i == 0 {
					fmt.Print("←  ")
				} else {
					switch q.GetAction(j, i) {
						case up:
							fmt.Print("↑  ")
						case down:
							fmt.Print("↓  ")
						case left:
							fmt.Print("←  ")
						case right:
							fmt.Print("→  ")
						}
				}
			} else {
				fmt.Print("G ")
			}
		}
		fmt.Println("")
	}
}