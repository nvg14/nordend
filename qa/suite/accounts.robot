*** Settings ***
Library    RequestsLibrary
Library    Collections
Resource  ../variables/accounts.robot

*** Test Cases ***
Create a new account
    [Documentation]    API to list all accounts
    [Tags]    Accounts
    ${data}=      create dictionary       amgid=${amgid_1}  name=${name_1}
    ${headers}=     create dictionary       Content-Type=application/json
    ${json}=    evaluate  json.dumps(${data})   json
    ${response}=    Post     ${base_url}${accounts_endpoint}  headers=${headers}    data=${json}
    Status Should Be    201
    
    ${response}    GET    ${base_url}${list_accounts}
    Status Should Be    200
    Log List    ${response.json()}

Get list of accounts
    [Documentation]    API to list all accounts
    [Tags]    Accounts
    ${response}    GET    ${base_url}${list_accounts}
    Status Should Be    200
    Log List    ${response.json()}

Get list of accounts
    [Documentation]    API to list all accounts
    [Tags]    Accounts
    ${response}    GET    ${base_url}${list_accounts}
    Status Should Be    200
    Log List    ${response.json()}