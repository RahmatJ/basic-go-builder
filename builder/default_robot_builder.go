package builder

import "roboBuilder/domain"

type DefaultRobotBuilder struct {
	robot domain.Robot
}

func NewDefaultRobotBuilder() *DefaultRobotBuilder {
	return &DefaultRobotBuilder{
		robot: domain.Robot{},
	}
}

func (d *DefaultRobotBuilder) addLeg() {
	d.robot.Leg = "Default Leg"
}

func (d *DefaultRobotBuilder) addHead() {
	d.robot.Head = "Default Head"
}

func (d *DefaultRobotBuilder) addArm() {
	d.robot.Arm = "Default Arm"
}

func (d *DefaultRobotBuilder) addBody() {
	d.robot.Body = "Default Body"
}

func (d *DefaultRobotBuilder) Build() domain.Robot {
	d.addLeg()
	d.addArm()
	d.addHead()
	d.addBody()
	return d.robot
}
