package builder

import "roboBuilder/domain"

type GundamBuilder struct {
	robot domain.Robot
}

func NewGundamBuilder() *GundamBuilder {
	return &GundamBuilder{
		robot: domain.Robot{},
	}
}

func (d *GundamBuilder) addLeg() {
	d.robot.Leg = "Gundam Leg"
}

func (d *GundamBuilder) addHead() {
	d.robot.Head = "Gundam Head"
}

func (d *GundamBuilder) addArm() {
	d.robot.Arm = "Gundam Arm"
}

func (d *GundamBuilder) addBody() {
	d.robot.Body = "Gundam Body"
}

func (d *GundamBuilder) Build() domain.Robot {
	d.addLeg()
	d.addArm()
	d.addHead()
	d.addBody()
	return d.robot
}
