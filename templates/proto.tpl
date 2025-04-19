syntax = "proto3";

package {{.PackageName}};
option go_package = "./{{.PackageName}}";

service {{.ServiceName}} {
    rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
    string name = 1;
}

message HelloResponse {
    string message = 1;
}

{{- if .AdditionalMessages}}
{{range .AdditionalMessages}}
message {{.}} {
    string field = 1;
}
{{- end}}
{{- end}}