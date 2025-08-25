

# Tokeninfo


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**audience** | **String** | Who is the intended audience for this token. In general the same as issued_to. |  [optional] |
|**email** | **String** | The email address of the user. Present only if the email scope is present in the request. |  [optional] |
|**expiresIn** | **Integer** | The expiry time of the token, as number of seconds left until expiry. |  [optional] |
|**issuedTo** | **String** | To whom was the token issued to. In general the same as audience. |  [optional] |
|**scope** | **String** | The space separated list of scopes granted to this token. |  [optional] |
|**userId** | **String** | The obfuscated user id. |  [optional] |
|**verifiedEmail** | **Boolean** | Boolean flag which is true if the email address is verified. Present only if the email scope is present in the request. |  [optional] |



