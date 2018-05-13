//
// This file generated by rdl 1.5.0
//

package globalip

import (
	"log"

	rdl "github.com/ardielle/ardielle-go/rdl"
)

var schema *rdl.Schema

func init() {
	sb := rdl.NewSchemaBuilder("GlobalIP")
	sb.Version(1)
	sb.Namespace("tk.cty4500")
	sb.Comment("A global IP API in *RDL*")

	tOctet := rdl.NewStringTypeBuilder("Octet")
	tOctet.Pattern("[0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]")
	sb.AddType(tOctet.Build())

	tIPAddress := rdl.NewStringTypeBuilder("IPAddress")
	tIPAddress.Pattern("[0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]\\.[0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]\\.[0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]\\.[0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]")
	sb.AddType(tIPAddress.Build())

	tGlobalIPResponse := rdl.NewStructTypeBuilder("Struct", "GlobalIPResponse")
	tGlobalIPResponse.Field("origin", "IPAddress", false, nil, "")
	sb.AddType(tGlobalIPResponse.Build())

	mGetGlobalIP := rdl.NewResourceBuilder("GlobalIPResponse", "GET", "/ip")
	mGetGlobalIP.Name("GetGlobalIP")
	mGetGlobalIP.Exception("BAD_REQUEST", "ResourceError", "")
	mGetGlobalIP.Exception("NOT_FOUND", "ResourceError", "")
	sb.AddResource(mGetGlobalIP.Build())

	var err error
	schema, err = sb.BuildParanoid()
	if err != nil {
		log.Fatalf("rdl: schema build failed: %s", err)
	}
}

func GlobalIPSchema() *rdl.Schema {
	return schema
}
