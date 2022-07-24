# Azure Virtual Desktop Client Connectivity Tool

This tool is client side outbound port checker for Azure Virtual Desktop.
Tool checks for required URLs and outbound ports on client side. Not all required/specified URLs combinations can be checked. See the table below for requirements, limitations and how this tool checks the requirements.

This is NOT official Microsoft tool. This tool is based on publicly available documentation available here: <https://docs.microsoft.com/en-us/azure/virtual-desktop/safe-url-list?tabs=azure#remote-desktop-clients>.

## AVD client sites and resources

|Requirement | Outbound Port | Probe DNS Record | Limited Test | Probe URL|
|---|---|---|---|---|
|*.wvd.microsoft.com|443|client.wvd.microsoft.com|Yes|https://client.wvd.microsoft.com/arm/webclient/index.html|
|*.servicebus.windows.net|443|watchdog.servicebus.windows.net|Yes|  |
|go.microsoft.com|443|go.microsoft.com|  |  |
|aka.ms|443|aka.ms|  |  |
|docs.microsoft.com|443|docs.microsoft.com|  |  |
|privacy.microsoft.com|443|privacy.microsoft.com|  |  |
|query.prod.cms.rt.microsoft.com|443|query.prod.cms.rt.microsoft.com|  |  |

## Office 365 URLs and IP address ranges - section 56

|Requirement | Outbound Port | Probe DNS Record | Limited Test | Probe URL|
|---|---|---|---|---|
|*.auth.microsoft.com|443|credentials.auth.microsoft.com|Yes|  |
|*.msftidentity.com|443|msftidentity.com|Yes|  |
|*.msidentity.com|443|msidentity.com|Yes|  |
|account.activedirectory.windowsazure.com|443|account.activedirectory.windowsazure.com|  |  |
|accounts.accesscontrol.windows.net|443|accounts.accesscontrol.windows.net|  |  |
|adminwebservice.microsoftonline.com|443|adminwebservice.microsoftonline.com|  |  |
|api.passwordreset.microsoftonline.com|443|api.passwordreset.microsoftonline.com|  |  |
|autologon.microsoftazuread-sso.com|443|autologon.microsoftazuread-sso.com|  |  |
|becws.microsoftonline.com|443|becws.microsoftonline.com|  |  |
|ccs.login.microsoftonline.com|443|ccs.login.microsoftonline.com|  |  |
|clientconfig.microsoftonline-p.net|443|clientconfig.microsoftonline-p.net|  |  |
|companymanager.microsoftonline.com|443|companymanager.microsoftonline.com|  |  |
|device.login.microsoftonline.com|443|device.login.microsoftonline.com|  |  |
|graph.microsoft.com|443|graph.microsoft.com|  |  |
|graph.windows.net|443|graph.windows.net|  |  |
|login.microsoft.com|443|login.microsoft.com|  |  |
|login.microsoftonline.com|443|login.microsoftonline.com|  |  |
|login.microsoftonline-p.com|443|login.microsoftonline-p.com|  |  |
|login.windows.net|443|login.windows.net|  |  |
|logincert.microsoftonline.com|443|logincert.microsoftonline.com|  |  |
|loginex.microsoftonline.com|443|loginex.microsoftonline.com|  |  |
|login-us.microsoftonline.com|443|login-us.microsoftonline.com|  |  |
|nexus.microsoftonline-p.com|443|nexus.microsoftonline-p.com|  |  |
|nexus.microsoftonline-p.com|443|nexus.microsoftonline-p.com|  |  |
|passwordreset.microsoftonline.com|443|passwordreset.microsoftonline.com|  |  |
|provisioningapi.microsoftonline.com|443|provisioningapi.microsoftonline.com|  |  |
|20.190.128.0/18|443|???|Yes|  |
|40.126.0.0/18|443|???|Yes|  |
|2603:1006:2000::/48|443|???|Yes|  |
|2603:1007:200::/48|443|???|Yes|  |
|2603:1016:1400::/48|443|???|Yes|  |
|2603:1017::/48|443|???|Yes|  |
|2603:1026:3000::/48|443|???|Yes|  |
|2603:1027:1::/48|443|???|Yes|  |
|2603:1036:3000::/48|443|???|Yes|  |
|2603:1037:1::/48|443|???|Yes|  |
|2603:1046:2000::/48|443|???|Yes|  |
|2603:1047:1::/48|443|???|Yes|  |
|2603:1056:2000::/48|443|???|Yes|  |
|2603:1057:2::/48|443|???|Yes|  |
