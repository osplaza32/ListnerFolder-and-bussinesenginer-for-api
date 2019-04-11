package Entidades

import "encoding/xml"
type InitPolicy struct {
	XMLName xml.Name `xml:"Policy"`
	Text    string   `xml:",chardata"`
	L7p     string   `xml:"L7p,attr"`
	Wsp     string   `xml:"wsp,attr"`
	All     struct {
		Text  string `xml:",chardata"`
		Usage string `xml:"Usage,attr"`
		AuditDetailAssertions
	} `xml:"All"`
}
type AuditDetailAssertions struct {
	XMLName xml.Name `xml:"AuditDetailAssertion"`
	Text    string   `xml:",chardata"`
	Detail  struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"Detail"`
}


type Policy struct {
	XMLName xml.Name `xml:"Policy"`
	Text    string   `xml:",chardata"`
	L7p     string   `xml:"L7p,attr"`
	Wsp     string   `xml:"wsp,attr"`
	*All
	*HandleErrors
	*Include
	*ApiPortalIntegration
	*Encapsulated
	*Pack
	*OneOrMore
	*ComparisonAssertion
	*AssertionComment
	*HardcodedResponse
	*AuditDetailAssertion
}

type ApiPortalIntegration struct {
	XMLName  xml.Name `xml:"ApiPortalIntegration"`
	Text     string   `xml:",chardata"`
	ApiGroup struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"ApiGroup"`
	ApiId struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"ApiId"`
	PortalManagedApiFlag struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"PortalManagedApiFlag"`
}

type Encapsulated struct {
	XMLName                         xml.Name `xml:"Encapsulated"`
	Text                            string   `xml:",chardata"`
	EncapsulatedAssertionConfigGuid struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"EncapsulatedAssertionConfigGuid"`
	EncapsulatedAssertionConfigName struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"EncapsulatedAssertionConfigName"`
}


type All struct {
	XMLName xml.Name `xml:"All"`
	Text    string   `xml:",chardata"`
	Usage   string   `xml:"Usage,attr"`
	*HandleErrors
	*Include
}
type HandleErrors struct {
	XMLName            xml.Name `xml:"HandleErrors"`
	Text               string   `xml:",chardata"`
	Usage              string   `xml:"Usage,attr"`
	IncludeIOException string   `xml:"includeIOException,attr"`
	VariablePrefix     string   `xml:"variablePrefix,attr"`
	*All
	*Include
}
type Pack struct {
	XMLName xml.Name `xml:"OneOrMore"`
	Text    string   `xml:",chardata"`
	Usage   string   `xml:"Usage,attr"`
	All     struct {
		Text                string `xml:",chardata"`
		Usage               string `xml:"Usage,attr"`
		ComparisonAssertion struct {
			Text          string `xml:",chardata"`
			CaseSensitive struct {
				Text         string `xml:",chardata"`
				BooleanValue string `xml:"booleanValue,attr"`
			} `xml:"CaseSensitive"`
			Expression1 struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"Expression1"`
			Operator struct {
				Text         string `xml:",chardata"`
				OperatorNull string `xml:"operatorNull,attr"`
			} `xml:"Operator"`
			Predicates struct {
				Text       string `xml:",chardata"`
				Predicates string `xml:"predicates,attr"`
				Item       []struct {
					Text    string `xml:",chardata"`
					Binary  string `xml:"binary,attr"`
					Negated struct {
						Text         string `xml:",chardata"`
						BooleanValue string `xml:"booleanValue,attr"`
					} `xml:"Negated"`
					RightValue struct {
						Text        string `xml:",chardata"`
						StringValue string `xml:"stringValue,attr"`
					} `xml:"RightValue"`
				} `xml:"item"`
			} `xml:"Predicates"`
		} `xml:"ComparisonAssertion"`
		SetVariable struct {
			Text             string `xml:",chardata"`
			Base64Expression struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"Base64Expression"`
			VariableToSet struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"VariableToSet"`
		} `xml:"SetVariable"`
		AuditDetailAssertion struct {
			Text               string `xml:",chardata"`
			CustomLoggerSuffix struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"CustomLoggerSuffix"`
			Detail struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"Detail"`
			LoggingOnly struct {
				Text         string `xml:",chardata"`
				BooleanValue string `xml:"booleanValue,attr"`
			} `xml:"LoggingOnly"`
		} `xml:"AuditDetailAssertion"`
		HardcodedResponse struct {
			Text               string `xml:",chardata"`
			Base64ResponseBody struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"Base64ResponseBody"`
			ResponseContentType struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"ResponseContentType"`
			ResponseStatus struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"ResponseStatus"`
		} `xml:"HardcodedResponse"`
	} `xml:"All"`
	TrueAssertion    string `xml:"TrueAssertion"`
	AssertionComment struct {
		Text       string `xml:",chardata"`
		Properties struct {
			Text     string `xml:",chardata"`
			MapValue string `xml:"mapValue,attr"`
			Entry    struct {
				Text string `xml:",chardata"`
				Key  struct {
					Text        string `xml:",chardata"`
					StringValue string `xml:"stringValue,attr"`
				} `xml:"key"`
				Value struct {
					Text        string `xml:",chardata"`
					StringValue string `xml:"stringValue,attr"`
				} `xml:"value"`
			} `xml:"entry"`
		} `xml:"Properties"`
	} `xml:"assertionComment"`
}



type Include struct {
	XMLName    xml.Name `xml:"Include"`
	Text       string   `xml:",chardata"`
	PolicyGuid struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"PolicyGuid"`
}
type OneOrMore struct {
	XMLName xml.Name `xml:"OneOrMore"`
	Text    string   `xml:",chardata"`
	Usage   string   `xml:"Usage,attr"`
}
type ComparisonAssertion struct {
	XMLName       xml.Name `xml:"ComparisonAssertion"`
	Text          string   `xml:",chardata"`
	CaseSensitive struct {
		Text         string `xml:",chardata"`
		BooleanValue string `xml:"booleanValue,attr"`
	} `xml:"CaseSensitive"`
	Expression1 struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"Expression1"`
	Operator struct {
		Text         string `xml:",chardata"`
		OperatorNull string `xml:"operatorNull,attr"`
	} `xml:"Operator"`
	Predicates struct {
		Text       string `xml:",chardata"`
		Predicates string `xml:"predicates,attr"`
		Item       []struct {
			Text    string `xml:",chardata"`
			Binary  string `xml:"binary,attr"`
			Negated struct {
				Text         string `xml:",chardata"`
				BooleanValue string `xml:"booleanValue,attr"`
			} `xml:"Negated"`
			RightValue struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"RightValue"`
		} `xml:"item"`
	} `xml:"Predicates"`
}
type SetVariable struct {
	XMLName          xml.Name `xml:"SetVariable"`
	Text             string   `xml:",chardata"`
	Base64Expression struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"Base64Expression"`
	VariableToSet struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"VariableToSet"`
}
type AuditDetailAssertion struct {
	XMLName            xml.Name `xml:"AuditDetailAssertion"`
	Text               string   `xml:",chardata"`
	CustomLoggerSuffix struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"CustomLoggerSuffix"`
	Detail struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"Detail"`
	LoggingOnly struct {
		Text         string `xml:",chardata"`
		BooleanValue string `xml:"booleanValue,attr"`
	} `xml:"LoggingOnly"`
}
type HardcodedResponse struct {
	XMLName            xml.Name `xml:"HardcodedResponse"`
	Text               string   `xml:",chardata"`
	Base64ResponseBody struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"Base64ResponseBody"`
	ResponseContentType struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"ResponseContentType"`
	ResponseStatus struct {
		Text        string `xml:",chardata"`
		StringValue string `xml:"stringValue,attr"`
	} `xml:"ResponseStatus"`
}
type AssertionComment struct {
	XMLName    xml.Name `xml:"assertionComment"`
	Text       string   `xml:",chardata"`
	Properties struct {
		Text     string `xml:",chardata"`
		MapValue string `xml:"mapValue,attr"`
		Entry    struct {
			Text string `xml:",chardata"`
			Key  struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"key"`
			Value struct {
				Text        string `xml:",chardata"`
				StringValue string `xml:"stringValue,attr"`
			} `xml:"value"`
		} `xml:"entry"`
	} `xml:"Properties"`
}

