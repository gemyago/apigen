package resources

import "embed"

//go:generate cp -r ../../../../../../generators/go-apigen-server/pom.xml ./go-apigen-server.xml
//go:generate cp -r ../../../../../../.versions ./.versions

//go:embed go-apigen-server.xml
var goApigenServer embed.FS

//go:embed .versions
var versions embed.FS
