# Available

Available will connect to a WHOIS server to ask for details about the
domains passed to it of course, if no details are returned, we can safely
assume that the domain is available for purchase.
Unfortunately, the WHOIS specification (see http://tools.ietf.org/html/rfc3912)
is very small and contains no information about how a WHOIS server should reply
when you ask for details about a domain. This means programmatically parsing
the response becomes a messy endeavor. To address this issue for now,
we will integrate with only a single WHOIS server, which we can be sure will
have No match somewhere in the response when it has no records for the domain.



The *com.whois-servers.net* WHOIS service supports domain names
for .com and .net, which is why the Domainify program only adds these types of domains.
