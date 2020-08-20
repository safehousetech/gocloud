package vultrcompute

import (
	"errors"

	"github.com/safehousetech/gocloud/vultrauth"
)

// VultrCompute struct
type VultrCompute struct {
}

//CreateNode .
type CreateNode struct {
	DCID                 int      // Location to create this virtual machine in.  See v1/regions/list
	VPSPLANID            int      // Plan to use when creating this virtual machine.  See v1/plans/list
	OSID                 int      // Operating system to use.  See v1/os/list
	ipxeChainURL         string   // (optional) If you've selected the 'custom' operating system, this can be set to chainload the specified URL on bootup, via iPXE
	ISOID                string   // (optional)  If you've selected the 'custom' operating system, this is the ID of a specific ISO to mount during the deployment
	SCRIPTID             int      // (optional) If you've not selected a 'custom' operating system, this can be the SCRIPTID of a startup script to execute on boot.  See v1/startupscript/list
	SNAPSHOTID           string   // (optional) If you've selected the 'snapshot' operating system, this should be the SNAPSHOTID (see v1/snapshot/list) to restore for the initial installation.
	enableIPv6           string   // (optional) 'yes' or 'no'.  If yes, an IPv6 subnet will be assigned to the machine (where available)
	enablePrivateNetwork string   // (optional) 'yes' or 'no'. If yes, private networking support will be added to the new server.
	NETWORKID            []string // (optional) List of private networks to attach to this server.Use either this field or enable_private_network, not both
	label                string   // (optional) This is a text label that will be shown in the control panel
	SSHKEYID             string   // (optional) List of SSH keys to apply to this server on install (only valid for Linux/FreeBSD).See v1/sshkey/list.Separate keys with commas
	autoBackups          string   // (optional) 'yes' or 'no'.  If yes, automatic backups will be enabled for this server (these have an extra charge associated with them)
	APPID                int      // (optional)  If launching an application (OSID 186), this is the APPID to launch.See v1/app/list.
	userdata             string   // (optional)  Base64 encoded user-data
	notifyActivate       string   // (optional, default 'yes') 'yes' or 'no'. If yes, an activation email will be sent when the server is ready.
	ddosProtection       string   // (optional, default 'no') 'yes' or 'no'.  If yes, DDOS protection will be enabled on the subscription (there is an additional charge for this)
	reservedIPv4         string   // (optional) IP address of the floating IP to use as the main IP of this server
	hostname             string   // (optional) The hostname to assign to this server.
	tag                  string   // (optional) The tag to assign to this server.
	FIREWALLGROUPID      string   // (optional) The firewall group to assign to this server. See /v1/firewall/group_list.
}

//StartNode ..
type StartNode struct {
	SUBID int
}

//RebootNode .
type RebootNode struct {
	SUBID int
}

//DeleteNode .
type DeleteNode struct {
	SUBID int // Unique identifier for this subscription.  These can be found using the v1/server/list call.
}

//ListNode .
type ListNode struct {
	SUBID  int    // (optional) Unique identifier of a subscription. Only the subscription object will be returned.
	tag    string // (optional) A tag string. Only subscription objects with this tag will be returned.
	label  string // (optional) A text label string. Only subscription objects with this text label will be returned.
	mainIP string // (optional) An IPv4 address. Only the subscription matching this IPv4 address will be returned.
}

// CreateNodeBuilder pattern code
type CreateNodeBuilder struct {
	createNode *CreateNode
}

//NewCreateNodeBuilder .
func NewCreateNodeBuilder() *CreateNodeBuilder {
	createNode := &CreateNode{}
	b := &CreateNodeBuilder{createNode: createNode}
	return b
}

//DCID .
func (b *CreateNodeBuilder) DCID(dcID int) *CreateNodeBuilder {
	b.createNode.DCID = dcID
	return b
}

//VPSPLANID .
func (b *CreateNodeBuilder) VPSPLANID(vPSPLANID int) *CreateNodeBuilder {
	b.createNode.VPSPLANID = vPSPLANID
	return b
}

//OSID .
func (b *CreateNodeBuilder) OSID(oSID int) *CreateNodeBuilder {
	b.createNode.OSID = oSID
	return b
}

//IpxeChainURL .
func (b *CreateNodeBuilder) IpxeChainURL(ipxeChainURL string) *CreateNodeBuilder {
	b.createNode.ipxeChainURL = ipxeChainURL
	return b
}

//ISOID .
func (b *CreateNodeBuilder) ISOID(iSOID string) *CreateNodeBuilder {
	b.createNode.ISOID = iSOID
	return b
}

//SCRIPTID .
func (b *CreateNodeBuilder) SCRIPTID(sCRIPTID int) *CreateNodeBuilder {
	b.createNode.SCRIPTID = sCRIPTID
	return b
}

//SNAPSHOTID .
func (b *CreateNodeBuilder) SNAPSHOTID(sNAPSHOTID string) *CreateNodeBuilder {
	b.createNode.SNAPSHOTID = sNAPSHOTID
	return b
}

//EnableIPv6 .
func (b *CreateNodeBuilder) EnableIPv6(enableIPv6 string) *CreateNodeBuilder {
	b.createNode.enableIPv6 = enableIPv6
	return b
}

//EnablePrivateNetwork .
func (b *CreateNodeBuilder) EnablePrivateNetwork(enablePrivateNetwork string) *CreateNodeBuilder {
	b.createNode.enablePrivateNetwork = enablePrivateNetwork
	return b
}

//NETWORKID .
func (b *CreateNodeBuilder) NETWORKID(nETWORKID []string) *CreateNodeBuilder {
	b.createNode.NETWORKID = nETWORKID
	return b
}

//Label .
func (b *CreateNodeBuilder) Label(label string) *CreateNodeBuilder {
	b.createNode.label = label
	return b
}

//SSHKEYID .
func (b *CreateNodeBuilder) SSHKEYID(sSHKEYID string) *CreateNodeBuilder {
	b.createNode.SSHKEYID = sSHKEYID
	return b
}

//AutoBackups .
func (b *CreateNodeBuilder) AutoBackups(autoBackups string) *CreateNodeBuilder {
	b.createNode.autoBackups = autoBackups
	return b
}

//APPID .
func (b *CreateNodeBuilder) APPID(aPPID int) *CreateNodeBuilder {
	b.createNode.APPID = aPPID
	return b
}

//UserData .
func (b *CreateNodeBuilder) UserData(userdata string) *CreateNodeBuilder {
	b.createNode.userdata = userdata
	return b
}

//NotifyActivate .
func (b *CreateNodeBuilder) NotifyActivate(notifyActivate string) *CreateNodeBuilder {
	b.createNode.notifyActivate = notifyActivate
	return b
}

//DdosProtection .
func (b *CreateNodeBuilder) DdosProtection(ddosProtection string) *CreateNodeBuilder {
	b.createNode.ddosProtection = ddosProtection
	return b
}

//ReservedIPv4 .
func (b *CreateNodeBuilder) ReservedIPv4(reservedIPv4 string) *CreateNodeBuilder {
	b.createNode.reservedIPv4 = reservedIPv4
	return b
}

//Hostname .Hostname
func (b *CreateNodeBuilder) Hostname(hostname string) *CreateNodeBuilder {
	b.createNode.hostname = hostname
	return b
}

//Tag .
func (b *CreateNodeBuilder) Tag(tag string) *CreateNodeBuilder {
	b.createNode.tag = tag
	return b
}

//FIREWALLGROUPID .
func (b *CreateNodeBuilder) FIREWALLGROUPID(fIREWALLGROUPID string) *CreateNodeBuilder {
	b.createNode.FIREWALLGROUPID = fIREWALLGROUPID
	return b
}

//Build .
func (b *CreateNodeBuilder) Build() (map[string]interface{}, error) {
	if b.createNode.DCID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "DCID")
	}
	if b.createNode.VPSPLANID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "VPSPLANID")
	}
	if b.createNode.OSID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "OSID")
	}

	params := make(map[string]interface{})

	params["DCID"] = b.createNode.DCID
	params["VPSPLANID"] = b.createNode.VPSPLANID
	params["OSID"] = b.createNode.OSID

	if b.createNode.ipxeChainURL != "" {
		params["ipxe_chain_url"] = b.createNode.ipxeChainURL
	}
	if b.createNode.ISOID != "" {
		params["ISOID"] = b.createNode.ISOID
	}
	if b.createNode.SCRIPTID != 0 {
		params["SCRIPTID"] = b.createNode.SCRIPTID
	}
	if b.createNode.SNAPSHOTID != "" {
		params["SNAPSHOTID"] = b.createNode.SNAPSHOTID
	}
	if b.createNode.enableIPv6 != "" {
		params["enable_ipv6"] = b.createNode.enableIPv6
	}
	if b.createNode.enablePrivateNetwork != "" {
		params["enable_private_network"] = b.createNode.enablePrivateNetwork
	}
	if len(b.createNode.NETWORKID) != 0 {
		params["NETWORKID"] = b.createNode.NETWORKID
	}
	if b.createNode.label != "" {
		params["label"] = b.createNode.label
	}
	if b.createNode.SSHKEYID != "" {
		params["SSHKEYID"] = b.createNode.SSHKEYID
	}
	if b.createNode.autoBackups != "" {
		params["auto_backups"] = b.createNode.autoBackups
	}
	if b.createNode.APPID != 0 {
		params["APPID"] = b.createNode.APPID
	}
	if b.createNode.userdata != "" {
		params["userdata"] = b.createNode.userdata
	}
	if b.createNode.notifyActivate != "" {
		params["notify_activate"] = b.createNode.notifyActivate
	}
	if b.createNode.ddosProtection != "" {
		params["ddos_protection"] = b.createNode.ddosProtection
	}
	if b.createNode.reservedIPv4 != "" {
		params["reserved_ip_v4"] = b.createNode.reservedIPv4
	}
	if b.createNode.hostname != "" {
		params["hostname"] = b.createNode.hostname
	}
	if b.createNode.tag != "" {
		params["tag"] = b.createNode.tag
	}
	if b.createNode.FIREWALLGROUPID != "" {
		params["FIREWALLGROUPID"] = b.createNode.FIREWALLGROUPID
	}

	return params, nil
}

// StartNodeBuilder pattern code
type StartNodeBuilder struct {
	startNode *StartNode
}

//NewStartNodeBuilder .
func NewStartNodeBuilder() *StartNodeBuilder {
	startNode := &StartNode{}
	b := &StartNodeBuilder{startNode: startNode}
	return b
}

//SUBID .
func (b *StartNodeBuilder) SUBID(subID int) *StartNodeBuilder {
	b.startNode.SUBID = subID
	return b
}

//Build .
func (b *StartNodeBuilder) Build() (map[string]interface{}, error) {
	if b.startNode.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := map[string]interface{}{
		"SUBID": b.startNode.SUBID,
	}
	return params, nil
}

// RebootNodeBuilder pattern code
type RebootNodeBuilder struct {
	rebootNode *RebootNode
}

//NewRebootNodeBuilder .
func NewRebootNodeBuilder() *RebootNodeBuilder {
	rebootNode := &RebootNode{}
	b := &RebootNodeBuilder{rebootNode: rebootNode}
	return b
}

//SUBID ..
func (b *RebootNodeBuilder) SUBID(subID int) *RebootNodeBuilder {
	b.rebootNode.SUBID = subID
	return b
}

//Build .
func (b *RebootNodeBuilder) Build() (map[string]interface{}, error) {
	if b.rebootNode.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := map[string]interface{}{
		"SUBID": b.rebootNode.SUBID,
	}
	return params, nil
}

// DeleteNodeBuilder pattern code
type DeleteNodeBuilder struct {
	deleteNode *DeleteNode
}

//NewDeleteNodeBuilder .
func NewDeleteNodeBuilder() *DeleteNodeBuilder {
	deleteNode := &DeleteNode{}
	b := &DeleteNodeBuilder{deleteNode: deleteNode}
	return b
}

//SubID .
func (b *DeleteNodeBuilder) SubID(subID int) *DeleteNodeBuilder {
	b.deleteNode.SUBID = subID
	return b
}

//Build ..
func (b *DeleteNodeBuilder) Build() (map[string]interface{}, error) {
	if b.deleteNode.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := map[string]interface{}{
		"SUBID": b.deleteNode.SUBID,
	}
	return params, nil
}

// ListNodeBuilder pattern code
type ListNodeBuilder struct {
	listNode *ListNode
}

//NewListNodeBuilder .
func NewListNodeBuilder() *ListNodeBuilder {
	listNode := &ListNode{}
	b := &ListNodeBuilder{listNode: listNode}
	return b
}

//SUBID ..
func (b *ListNodeBuilder) SUBID(subID int) *ListNodeBuilder {
	b.listNode.SUBID = subID
	return b
}

func (b *ListNodeBuilder) tag(tag string) *ListNodeBuilder {
	b.listNode.tag = tag
	return b
}

func (b *ListNodeBuilder) label(label string) *ListNodeBuilder {
	b.listNode.label = label
	return b
}

//mainIP .
func (b *ListNodeBuilder) mainIP(mainIP string) *ListNodeBuilder {
	b.listNode.mainIP = mainIP
	return b
}

//Build .
func (b *ListNodeBuilder) Build() (map[string]interface{}, error) {
	params := map[string]interface{}{}
	if b.listNode.SUBID != 0 {
		params["SUBID"] = b.listNode.SUBID
	}
	if b.listNode.tag != "" {
		params["tag"] = b.listNode.tag
	}
	if b.listNode.label != "" {
		params["label"] = b.listNode.label
	}
	if b.listNode.mainIP != "" {
		params["main_ip"] = b.listNode.mainIP
	}
	return params, nil
}
