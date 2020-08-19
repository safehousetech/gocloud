package aliloadbalancer

import (
	"fmt"
	"github.com/safehousetech/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreateLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	create := map[string]interface{}{
		"RegionId":           "cn-qingdao",
		"LoadBalancerName":   "abc",
		"AddressType":        "internet",
		"InternetChargeType": "paybytraffic",
	}
	_, err := alilb.CreateLoadBalancer(create)
	if err != nil {
		t.Errorf("CreateLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is created successfully.")
}

func TestCreateLoadBalancerBuilder(t *testing.T) {
	var alilb Aliloadbalancer
	create, err := NewCreateLoadBalancerBuilder().
		RegionID("cn-qingdao").
		LoadBalancerName("abc").
		AddressType("internet").
		InternetChargeType("paybytraffic").
		Build()
	if err != nil {
		t.Errorf("CreateLoadBalancer Test Fail: %s", err)
		return
	}
	_, err = alilb.CreateLoadBalancer(create)
	if err != nil {
		t.Errorf("CreateLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is created successfully.")
}

func TestDeleteLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	delete := map[string]interface{}{
		"RegionId":       "cn-qingdao",
		"LoadBalancerId": "lb-m5ehdbs3p10a7kmq344je",
	}
	_, err := alilb.DeleteLoadBalancer(delete)
	if err != nil {
		t.Errorf("DeleteLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is deleted successfully.")
}

func TestDeleteLoadBalancerBuilder(t *testing.T) {
	var alilb Aliloadbalancer
	deleteLoadBalancer, err := NewDeleteLoadBalancerBuilder().
		RegionID("cn-qingdao").
		LoadBalancerID("lb-a2dgmwo53mftn7h34za84").
		Build()
	if err != nil {
		t.Errorf("DeleteLoadBalancer Test Fail: %s", err)
		return
	}
	_, err = alilb.DeleteLoadBalancer(deleteLoadBalancer)
	if err != nil {
		t.Errorf("DeleteLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is deleted successfully.")
}

func TestListLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	list := map[string]interface{}{
		"RegionId": "cn-qingdao",
	}
	resp, err := alilb.ListLoadBalancer(list)
	if err != nil {
		t.Errorf("ListLoadBalancer Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
	t.Logf("Ali LoadBalancer is listed successfully.")
}

func TestListLoadBalancerBuilder(t *testing.T) {
	var alilb Aliloadbalancer
	list, err := NewListLoadBalancerBuilder().
		RegionID("cn-qingdao").
		Build()
	if err != nil {
		t.Errorf("ListLoadBalancer Test Fail: %s", err)
		return
	}
	resp, err := alilb.ListLoadBalancer(list)
	if err != nil {
		t.Errorf("ListLoadBalancer Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
	t.Logf("Ali LoadBalancer is listed successfully.")
}

func TestAttachLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	attach := map[string]interface{}{
		"LoadBalancerId": "lb-m5eemmwmtmt4l6jf73d72",
		"BackendServers": "[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Weight':'100','Type':'ecs'}," +
			"{'ServerId':'i-m5eahbbwqxawpj1opww9','Weight':'100','Type':'ecs'}]",
	}
	_, err := alilb.AttachNodeWithLoadBalancer(attach)
	if err != nil {
		t.Errorf("AttachNodeWithLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is listed successfully.")
}

func TestAttachLoadBalancerBuilder(t *testing.T) {
	var alilb Aliloadbalancer
	attach, err := NewAttachLoadBalancerBuilder().
		LoadBalancerID("lb-m5e7ldi4obgcya3bju3wu").
		BackendServers("[{'ServerId':'i-m5e0cjxe9c7iulv9znvw','Weight':'100','Type':'ecs'}," +
			"{'ServerId':'i-m5efh52hzdkyjoaafwc0','Weight':'100','Type':'ecs'}]").
		Build()
	if err != nil {
		t.Errorf("AttachNodeWithLoadBalancer Test Fail: %s", err)
		return
	}
	_, err = alilb.AttachNodeWithLoadBalancer(attach)
	if err != nil {
		t.Errorf("AttachNodeWithLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is listed successfully.")
}

func TestDettachLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	dettach := map[string]interface{}{
		"RegionId":       "cn-qingdao",
		"LoadBalancerId": "lb-m5eemmwmtmt4l6jf73d72",
		"BackendServers": "[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Type':'ecs'}," +
			"{'ServerId':'i-m5eahbbwqxawpj1opww9','Type':'ecs'}]",
	}
	_, err := alilb.DetachNodeWithLoadBalancer(dettach)
	if err != nil {
		t.Errorf("DetachnodewithLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is detached with node successfully.")
}

func TestDetachLoadBalancerBuilder(t *testing.T) {
	var alilb Aliloadbalancer
	detach, err := NewDetachLoadBalancerBuilder().
		RegionID("cn-qingdao").
		LoadBalancerID("lb-m5eemmwmtmt4l6jf73d72").
		BackendServers("[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Type':'ecs'}," +
			"{'ServerId':'i-m5eahbbwqxawpj1opww9','Type':'ecs'}]").
		Build()
	if err != nil {
		t.Errorf("DetachnodewithLoadBalancer Test Fail: %s", err)
		return
	}
	_, err = alilb.DetachNodeWithLoadBalancer(detach)
	if err != nil {
		t.Errorf("DetachnodewithLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is detached with node successfully.")
}

func TestParseCreateLoadBalancerResp(t *testing.T) {
	var alilb Aliloadbalancer
	create, err := NewCreateLoadBalancerBuilder().
		RegionID("cn-qingdao").
		LoadBalancerName("abc").
		AddressType("internet").
		InternetChargeType("paybytraffic").
		Build()
	if err != nil {
		t.Errorf("CreateLoadBalancer Test Fail: %s", err)
		return
	}
	resp, err := alilb.CreateLoadBalancer(create)
	if err != nil {
		t.Errorf("CreateLoadBalancer Test Fail: %s", err)
		return
	}
	createLoadBalancerResp, err := ParseCreateLoadBalancerResp(resp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", createLoadBalancerResp)
}

func TestParseAttachLoadBalancerResp(t *testing.T) {
	var alilb Aliloadbalancer
	attach, err := NewAttachLoadBalancerBuilder().
		LoadBalancerID("lb-m5eg5vvdshdzix98evkk0").
		BackendServers("[{'ServerId':'i-m5eggpwhjezbnjuorqf5','Weight':'100','Type':'ecs'}," +
			"{'ServerId':'i-m5e5nqgrrrhte615lmkl','Weight':'100','Type':'ecs'}]").
		Build()
	if err != nil {
		t.Errorf("AttachNodeWithLoadBalancer Test Fail: %s", err)
		return
	}
	resp, err := alilb.AttachNodeWithLoadBalancer(attach)
	if err != nil {
		t.Errorf("AttachNodeWithLoadBalancer Test Fail: %s", err)
		return
	}
	attachLoadBalancerResp, err := ParseAttachLoadBalancerResp(resp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", attachLoadBalancerResp)
}

func TestParseDetachLoadBalancerResp(t *testing.T) {
	var alilb Aliloadbalancer
	detach, err := NewDetachLoadBalancerBuilder().
		RegionID("cn-qingdao").
		LoadBalancerID("lb-m5eg5vvdshdzix98evkk0").
		BackendServers("[{'ServerId':'i-m5eggpwhjezbnjuorqf5','Type':'ecs'}," +
			"{'ServerId':'i-m5e5nqgrrrhte615lmkl','Type':'ecs'}]").
		Build()
	if err != nil {
		t.Errorf("DetachnodewithLoadBalancer Test Fail: %s", err)
		return
	}
	resp, err := alilb.DetachNodeWithLoadBalancer(detach)
	if err != nil {
		t.Errorf("DetachnodewithLoadBalancer Test Fail: %s", err)
		return
	}
	detachLoadBalancerResp, err := ParseDetachLoadBalancerResp(resp)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", detachLoadBalancerResp)
}
