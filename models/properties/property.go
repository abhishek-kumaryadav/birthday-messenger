package properties

import "birthdaymessenger/pkg/fileutil"

var whatsappTemplate = "whatsapp.template"
var whatsappPhoneNumberId = "whatsapp.phoneNumberId"
var whatsappAuthToken = "whatsapp.authToken"
var gcpTokenPath = "gcp.token.path"
var gcpCredentialsPath = "gcp.credentials.path"
var jobFrequency = "job.frequency"

var props, err = fileutil.ReadPropertiesFile("birthday-messenger.properties")

func GetWhatsappTemplate() string      { return props[whatsappTemplate] }
func GetWhatsappPhoneNumberId() string { return props[whatsappPhoneNumberId] }
func GetWhatsappAuthToken() string     { return props[whatsappAuthToken] }
func GetGcpTokenPath() string          { return props[gcpTokenPath] }
func GetGcpCredentialsPath() string    { return props[gcpCredentialsPath] }
func GetJobFrequency() string          { return props[jobFrequency] }
