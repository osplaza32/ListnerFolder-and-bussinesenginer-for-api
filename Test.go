package main

import (
	"encoding/xml"
	"fmt"
	"osplaza32/Automatitations/Entidades"
)

var XML = []byte(`
<?xml version="1.0" encoding="UTF-8"?>

    <wsp:Policy xmlns:L7p="http://www.layer7tech.com/ws/policy" xmlns:wsp="http://schemas.xmlsoap.org/ws/2002/12/policy">
        <wsp:All wsp:Usage="Required">
            <L7p:AuditDetailAssertion>
                <L7p:Detail stringValue="Policy Fragment: Prueba"/>
            </L7p:AuditDetailAssertion>
        </wsp:All>
    </wsp:Policy>`)

func main() {
	var q Entidades.InitPolicy
	var a Entidades.AuditDetailAssertions

	xml.Unmarshal(XML, &q)
	xml.Unmarshal(XML, &a)

	fmt.Println(q)
	fmt.Println(a)



}
