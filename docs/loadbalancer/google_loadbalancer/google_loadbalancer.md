# gocloud loadbalancer - gce

## Configure Google Cloud credentials.

Download your service account credentials file from Google Cloud and save it as `googlecloudinfo.json` in your <b>HOME/.gocloud</b> directory.

You can also set the credentials as environment variables:
```js
export PrivateKey =  "xxxxxxxxxxxx"
export Type =  "xxxxxxxxxxxx"
export ProjectID = "xxxxxxxxxxxx"
export PrivateKeyID = "xxxxxxxxxxxx"
export ClientEmail = "xxxxxxxxxxxx"
export ClientID = "xxxxxxxxxxxx"
export AuthURI = "xxxxxxxxxxxx"
export TokenURI = "xxxxxxxxxxxx"
export AuthProviderX509CertURL = "xxxxxxxxxxxx"
export ClientX509CertURL =  "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/safehousetech/gocloud/gocloud"

googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
```

### Create loadbalancer

```js
 creatloadbalancer := map[string]interface{}{
   "Project":  "sheltermap-1493101612061",
   "Name"   :  "google-loadbalancer",
   "Region":   "us-central1",
   "Instances": []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-1"},
	}

  resp, err := googlecloud.CreateLoadBalancer(creatloadbalancer)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])

  ```

### List loadbalancer

```js
 listloadbalancer := map[string]string{
     "Project": "sheltermap-1493101612061",
     "Region":  "us-central1",
 }

 resp, err := googlecloud.ListLoadBalancer(listloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Attach node with loadbalancer

```js

  attachnodewithloadbalancer := map[string]interface{}{
	"Project":    "sheltermap-1493101612061",
	"Region":     "us-central1",
	"TargetPool": "google-loadbalancer",
	"Instances":  []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-2"},
	}

 resp, err := googlecloud.AttachNodeWithLoadBalancer(attachnodewithloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Detach node with loadbalancer

```js
  detachnodewithloadbalancer := map[string]interface{}{
      "Project":    "sheltermap-1493101612061",
      "Region":     "us-central1",
      "TargetPool": "google-loadbalancer",
      "Instances":  []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-2"},
	}

  resp, err := googlecloud.DetachNodeWithLoadBalancer(detachnodewithloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete loadbalancer
```js
  deleteloadbalancer := map[string]string{
		"Project":    "sheltermap-1493101612061",
		"Region":     "us-central1",
		"TargetPool": "google-loadbalancer",
	}

 resp, err := googlecloud.DeleteLoadBalancer(deleteloadbalancer)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```