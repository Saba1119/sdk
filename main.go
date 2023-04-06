package main

import ()
type Config struct {
	TableName string `yaml:"tableName"`
	AWS_REGION string `yaml:"AWS_REGION"`
    BUCKET_NAME string `yaml:"BUCKET_NAME"`
    KMS_KEY string `yaml:"KMS_KEY"`
    ImageID            string `yaml:"IMAGE_ID"`
	InstanceType       string `yaml:"INSTANCE_TYPE"`
	GroupName          string `yaml:"GROUP_NAME"`
	AlarmName          string `yaml:"ALARM_NAME"`
	SecurityGroupName  string `yaml:"SECURITYGROUP_NAME"`
	VPCID              string `yaml:"VPC_ID"`
	TargetGroupName1   string `yaml:"TARGETGROUP_NAME1"`
	TargetGroupName2   string `yaml:"TARGETGROUP_NAME2"`
	ALBName            string `yaml:"ALB_NAME"`
	SecurityGroup1FROMPort int64  `yaml:"Sgfrom_PORT1"`
	SecurityGroup1TOPort int64  `yaml:"Sgto_PORT1"`
	SecurityGroup2FROMPort int64  `yaml:"Sgfrom_PORT2"`
	SecurityGroup2TOPort int64  `yaml:"Sgto_PORT2"`
	TargetGroup1PORT     int64  `yaml:"TG1PORT"`
	TargetGroup2PORT     int64  `yaml:"TG2PORT"`
	Listener1PORT        int64  `yaml:"LISTENER1PORT"`
	Listener2PORT        int64  `yaml:"LISTENER2PORT"`

}

func main() {
	CreateS3()
	CreateDb()
	Createec2()
}
