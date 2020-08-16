package bigtable

//Createbigtabledictnoaryconvert json formation.
func Createbigtabledictnoaryconvert(option Createbigtable, Createbigtablejsonmap map[string]interface{}) {

	if option.tableID != "" {
		Createbigtablejsonmap["tableId"] = option.tableID
	}

	Createbigtablejsonmap["table"] = option.table
}
