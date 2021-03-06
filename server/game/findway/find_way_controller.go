package findway

import (
	"net"
	"server/server/game/findway/jps"
	"server/server/game/utils"
)

/**
jps参考 :https://www.cnblogs.com/KillerAery/archive/2020/06/17/12242445.html
		https://zhuanlan.zhihu.com/p/290924212

	就是经过优化的a*算法
	实现思路跟迪杰斯特拉算法一样，都是使用一个openList 和一个 closeList 来刷新节点
*/

/**
A* F值（优先级值）：F = G + H
这条公式意思：F是从起点经过该点再到达终点的预测总耗费值。通过计算F值，我们可以优先选择F值最小的方向来进行搜索。

jps 算法，就是往一个方向一直走，直到撞墙。
*/

type StartController struct {
}

func (c *StartController) Service(param map[string]interface{}, conn net.Conn) error {
	data := param["map"]
	rows := data.([]interface{})
	maps := make([][]int, len(rows))
	for i, row := range rows {
		cols := row.([]interface{})
		for _, v := range cols {
			maps[i] = append(maps[i], int(v.(float64)))
		}

	}
	startI, ok := param["startI"].(float64)
	if !ok {
		// todo 需要把context传进来。 这里需要再校验一下能否直接转为int。
		//log.Errorf("")
	}
	startJ, ok := param["startJ"].(float64)
	endI, ok := param["endI"].(float64)
	endJ, ok := param["endJ"].(float64)
	mapCfg := jps.InitMap(maps)
	jpsController := jps.InitJPS(mapCfg)
	route := jpsController.FindPath(int(startJ), int(startI), int(endJ), int(endI))
	toM := make(map[string]interface{})
	toM["route"] = route
	utils.Send(conn, toM, 1)
	return nil
}
