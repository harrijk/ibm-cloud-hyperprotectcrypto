/*Package examples provides grep11 function call examples.
Regular function calls have four steps. The following is one example to generate a key.

	1 cryptoClient := pb.NewCryptoClient(conn)      //create crypto client
	2 Template := util.NewAttributeMap              //create template
	3 keygenmsg, err := &pb.GenerateKeyRequest()    //create RPC request parameters
	4 if err != nill {...}                          //check err

*/
package examples
