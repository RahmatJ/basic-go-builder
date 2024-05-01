package builder

import "roboBuilder/domain"

type RobotBuilder interface {
	addLeg()
	addHead()
	addArm()
	addBody()
	Build() domain.Robot
}

func GetBuilder(data string) RobotBuilder {
	switch data {
	case "gundam":
		return NewGundamBuilder()
	default:
		return NewDefaultRobotBuilder()
	}
}
