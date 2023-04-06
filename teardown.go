package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/kms"
    "github.com/aws/aws-sdk-go/service/elbv2"
)

func main() {
    // Initialize a new AWS session
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Delete EC2 instances
    ec2Svc := ec2.New(sess)
    _, err := ec2Svc.TerminateInstances(&ec2.TerminateInstancesInput{
        InstanceIds: []*string{aws.String(os.Getenv("INSTANCE_ID"))},
    })
    if err != nil {
        fmt.Println("Error deleting EC2 instances:", err)
        os.Exit(1)
    }
    fmt.Println("EC2 instances deleted successfully.")
    fmt.Println("Security group deleted successfully.")

    // Delete ALB resources
    elbv2Svc := elbv2.New(sess)
    _, err = elbv2Svc.DeleteLoadBalancer(&elbv2.DeleteLoadBalancerInput{
        LoadBalancerArn: aws.String(os.Getenv("LOAD_BALANCER_ARN")),
    })
    if err != nil {
        fmt.Println("Error deleting ALB resources:", err)
        os.Exit(1)
    }
    fmt.Println("ALB resources deleted successfully.")

    // Delete ALB target group
    _, err = elbv2Svc.DeleteTargetGroup(&elbv2.DeleteTargetGroupInput{
        TargetGroupArn: aws.String(os.Getenv("TARGET_GROUP_ARN")),
    })
    if err != nil {
        fmt.Println("Error deleting ALB target group:", err)
        os.Exit(1)
    }
    fmt.Println("ALB target group deleted successfully.")

    // Delete S3 bucket and objects
    s3Svc := s3.New(sess)
    _, err = s3Svc.DeleteBucket(&s3.DeleteBucketInput{
        Bucket: aws.String(os.Getenv("BUCKET_NAME")),
    })
    if err != nil {
        fmt.Println("Error deleting S3 bucket:", err)
        os.Exit(1)
    }
    fmt.Println("S3 bucket deleted successfully.")


     // Create a KMS service client
    svc := kms.New(sess)

    // Specify the KMS key ID to delete
    keyId := os.Getenv("KMS_ID")

    // Delete the KMS key
    _, err = svc.ScheduleKeyDeletion(&kms.ScheduleKeyDeletionInput{
        KeyId:               aws.String(keyId),
    })
    if err != nil {
        fmt.Println("Error deleting KMS key:", err)
        os.Exit(1)
    }
    fmt.Println("KMS key deleted successfully.")

    // Delete DynamoDB table
    dynamoSvc := dynamodb.New(sess)
    _, err = dynamoSvc.DeleteTable(&dynamodb.DeleteTableInput{
        TableName: aws.String(os.Getenv("TABLE_NAME")),
    })
    if err != nil {
        fmt.Println("Error deleting DynamoDB table:", err)
        os.Exit(1)
    }
    fmt.Println("DynamoDB table deleted successfully.")

}
