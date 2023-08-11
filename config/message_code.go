package config

/*** Code message type:
* Validate: MSG_Vxxxx => validate general
* Error: MSG_Exxxx => error general
* Info: MSG_Ixxxx => info general
* ------------------------------------
* Validate Create Form: MSG_CVxxxx => validate general
* Error Create Form: MSG_CExxxx => error general
* Info Create Form: MSG_CIxxxx => info general
* ------------------------------------
* Validate Update Form: MSG_UVxxxx => validate general
* Error Update Form: MSG_UExxxx => error general
* Info Update Form: MSG_UIxxxx => info general
* ------------------------------------
* Validate Delete Form: MSG_DVxxxx => validate general
* Error Delete Form: MSG_DExxxx => error general
* Info Delete Form: MSG_DIxxxx => info general
***/
const (
	// Param error
	PARAM_ERROR string = "MSG_V0000"
	// Validate
	VALIDATE string = "VALIDATE"
	// Param require
	REQUIRED string = "MSG_V1000"
	// Min length
	MIN_LENGTH string = "MSG_V1001"
	// Max length
	MAX_LENGTH string = "MSG_V1002"
	// Key error not found
	KEY_NOT_FOUND string = "MSG_S0000"
	// System error
	SYSTEM_ERROR string = "MSG_S0001"
	// Token invalid
	TOKEN_INCORRECT string = "MSG_S0002"
	// Unauthorized
	UNAUTHORIZED string = "MSG_S0003"
)
