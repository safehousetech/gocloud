package kinesis

//DeleteStream Delete Stream
func (kinesis *Kinesis) DeleteStream(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var streamName, Region string

	for key, value := range param {
		switch key {
		case "StreamName":
			streamNameV, _ := value.(string)
			streamName = streamNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)
	preparedeletestream(params, Region)

	deletestreamjsonmap := make(map[string]interface{})
	preparedeletestreamdict(deletestreamjsonmap, streamName)

	response := make(map[string]interface{})
	err = kinesis.PrepareSignatureV4query(params, deletestreamjsonmap, response)
	resp = response
	return resp, err
}

//CreateStream Create Stream
func (kinesis *Kinesis) CreateStream(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var streamName, region string
	var shardCount int
	for key, value := range param {
		switch key {
		case "StreamName":
			streamNameV, _ := value.(string)
			streamName = streamNameV

		case "Region":
			regionV, _ := value.(string)
			region = regionV

		case "ShardCount":
			shardCountV, _ := value.(int)
			shardCount = shardCountV
		}
	}

	params := make(map[string]string)
	preparecreatestream(params, region)

	createstreamjsonmap := make(map[string]interface{})
	preparecreatestreamdict(createstreamjsonmap, streamName, shardCount)

	response := make(map[string]interface{})
	err = kinesis.PrepareSignatureV4query(params, createstreamjsonmap, response)
	resp = response
	return resp, err
}

//ListStream List Stream
func (kinesis *Kinesis) ListStream(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var exclusiveStartStreamName, limit, Region string

	for key, value := range param {
		switch key {
		case "ExclusiveStartStreamName":
			exclusiveStartStreamNameV, _ := value.(string)
			exclusiveStartStreamName = exclusiveStartStreamNameV

		case "Limit":
			limitV, _ := value.(string)
			limit = limitV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)
	prepareliststream(params, Region)

	liststreamjsonmap := make(map[string]interface{})
	prepareliststreamdict(liststreamjsonmap, exclusiveStartStreamName, limit)

	response := make(map[string]interface{})
	err = kinesis.PrepareSignatureV4query(params, liststreamjsonmap, response)
	resp = response
	return resp, err
}

//DescribeStream describe stream
func (kinesis *Kinesis) DescribeStream(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var streamName, limit, exclusiveStartShardID, region string

	for key, value := range param {
		switch key {
		case "StreamName":
			streamNameV, _ := value.(string)
			streamName = streamNameV

		case "Limit":
			limitV, _ := value.(string)
			limit = limitV

		case "ExclusiveStartShardId":
			exclusiveStartShardIDV, _ := value.(string)
			exclusiveStartShardID = exclusiveStartShardIDV

		case "Region":
			regionV, _ := value.(string)
			region = regionV
		}
	}

	params := make(map[string]string)
	preparedescribestream(params, region)

	describestreamjsonmap := make(map[string]interface{})
	preparedescribestreamdict(describestreamjsonmap, streamName, limit, exclusiveStartShardID)

	response := make(map[string]interface{})
	err = kinesis.PrepareSignatureV4query(params, describestreamjsonmap, response)
	resp = response
	return resp, err
}
