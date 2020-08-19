# gocloud compute - ECS

## Configure Ali-cloud credentials

Create `alicloudconfig.json` as follows,
```
{
  "AliAccessKeyID": "xxxxxxxxxxxx",
  "AliAccessKeySecret": "xxxxxxxxxxxx"
}
```

also You can setup environment variables as

```
export AliAccessKeyID =  "xxxxxxxxxxxx"
export AliAccessKeySecret = "xxxxxxxxxxxx"
```

## Initialize library

```
import "github.com/safehousetech/gocloud/gocloud"

alicloud, _ := gocloud.CloudProvider(gocloud.Aliprovider)
```

### Create instance

```
    create := map[string]interface{}{
        "RegionId":        "cn-qingdao",
        "ImageId":         "centos_7_04_64_20G_alibase_201701015.vhd",
        "InstanceType":    "ecs.xn4.small",
        "SecurityGroupId": "sg-m5egbo9s5xb21kpu6nk2",
    }

    resp, err := alicloud.CreateNode(create)
    createNodeResp, err := ecs.ParseCreateNodeResp(resp)
    fmt.Println(createNodeResp.StatusCode)
    fmt.Println(createNodeResp.InstanceId)
```

or

```
    createNode, err := ecs.NewCreateNodeBuilder().
        RegionID("cn-qingdao").
        ImageID("centos_7_04_64_20G_alibase_201701015.vhd").
        InstanceType("ecs.xn4.small").
        SecurityGroupID("sg-m5egbo9s5xb21kpu6nk2").
        Build()
    if err != nil {
        //...
    }
    resp, err := alicloud.CreateNode(createNode)
    createNodeResp, err := ecs.ParseCreateNodeResp(resp)
    fmt.Println(createNodeResp.StatusCode)
    fmt.Println(createNodeResp.InstanceId)
```

### Start instance

```
    start := map[string]interface{}{
        "InstanceId": "i-m5e3ee3z8wdy8ktdq591",
    }

    resp, err := alicloud.StartNode(start)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    startNode, err := ecs.NewStartNodeBuilder().
        InstanceID("i-m5e3xm2ph7d501uauqhy").
        Build()
    if err != nil {
        //...
    }
    resp, err := alicloud.StartNode(startNode)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Stop instance

```
    stop := map[string]interface{}{
        "InstanceId": "i-m5e3ee3z8wdy8ktdq591",
        "ForceStop":  false,
    }

    resp, err := alicloud.StopNode(stop)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    stopNode, err := ecs.NewStopNodeBuilder().
        InstanceID("i-m5e3xm2ph7d501uauqhy").
        ForceStop(false).
        Build()
    if err != nil {
        //...
    }
    resp, err := alicloud.StopNode(stopNode)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Reboot instance

```
    reboot := map[string]interface{}{
        "InstanceId": "i-m5e3ee3z8wdy8ktdq591",
        "ForceStop":  false,
    }

    resp, err := alicloud.RebootNode(reboot)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    rebootNode, err := ecs.NewRebootNodeBuilder().
        InstanceID("i-m5e3xm2ph7d501uauqhy").
        ForceStop(false).
        Build()
    if err != nil {
        //...
    }
    resp, err := alicloud.RebootNode(rebootNode)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

### Delete instance

```
    dlete := map[string]interface{}{
        "InstanceId": "i-m5e3ee3z8wdy8ktdq591",
    }

    resp, err := alicloud.DeleteNode(delete)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```

or

```
    deleteNode, err := ecs.NewDeleteNodeBuilder().
        InstanceID("i-m5e3xm2ph7d501uauqhy").
        Build()
    if err != nil {
        //...
    }
    resp, err := alicloud.DeleteNode(deleteNode)
    response := resp.(map[string]interface{})
    fmt.Println(response["body"])
```