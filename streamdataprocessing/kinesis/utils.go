package kinesis

func preparedeletestreamdict(deletestreamjsonmap map[string]interface{}, streamName string) {

	if streamName != "" {
		deletestreamjsonmap["StreamName"] = streamName
	}

}

func preparedeletestream(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "Kinesis_20131202.DeleteStream"
}

func preparedescribestreamdict(describeStreamJSONmap map[string]interface{}, streamName string, limit string, exclusiveStartShardID string) {

	if exclusiveStartShardID != "" {
		describeStreamJSONmap["ExclusiveStartShardId"] = exclusiveStartShardID
	}

	if streamName != "" {
		describeStreamJSONmap["StreamName"] = streamName
	}

	if limit != "" {
		describeStreamJSONmap["Limit"] = limit
	}
}

func preparedescribestream(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "Kinesis_20131202.DescribeStream"
}

func preparecreatestreamdict(createstreamjsonmap map[string]interface{}, streamName string, shardCount int) {

	if shardCount > 0 {
		createstreamjsonmap["ShardCount"] = shardCount
	}

	if streamName != "" {
		createstreamjsonmap["StreamName"] = streamName
	}
}

func preparecreatestream(params map[string]string, region string) {

	if region != "" {
		params["Region"] = region
	}
	params["amztarget"] = "Kinesis_20131202.CreateStream"
}

func prepareliststreamdict(liststreamjsonmap map[string]interface{}, exclusiveStartStreamName string, limit string) {

	if exclusiveStartStreamName != "" {
		liststreamjsonmap["ExclusiveStartStreamName"] = exclusiveStartStreamName
	}

	if limit != "" {
		liststreamjsonmap["Limit"] = limit
	}
}

func prepareliststream(params map[string]string, region string) {

	if region != "" {
		params["Region"] = region
	}
	params["amztarget"] = "Kinesis_20131202.ListStreams"
}
