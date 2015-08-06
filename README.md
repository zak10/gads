# gads

Package gads provides a wrapper for the Google Adwords SOAP API.  Based off of
[emiddleton/gads](https://github.com/emiddleton/gads), this version
was updated to support v201506 and Go 1.5.

Currently this project remains a fork and is a joint effort between
[colinmutter/gads](https://github.com/colinmutter/gads) (working on AWQL,
and Go 1.5 support) and [rfink/gads](https://github.com/rfink/gads)
(working on v201506 compatibility).


## installation

~~~
	go get github.com/colinmutter/gads
~~~

## setup

In order to access the API you will need to sign up for an MMC
account[1], get a developer token[2] and setup authentication[3].
There is a tool in the setup_oauth2 directory that will help you
setup a configuration file.

1. http://www.google.com/adwords/myclientcenter/
2. https://developers.google.com/adwords/api/docs/signingup
3. https://developers.google.com/adwords/api/docs/guides/authentication

Currently, the package will look for these generated oauth json credentials
in a file at ./config.json.  To override, supply the -gads_oauth flag and
specity the file's location:

    go run myproj.go --gads_oauth ~/my_oauth_creds.json

## versions

This project currently supports v201409 and v201506.  To select
the appropriate version, import the specific package:

	  import (
	    gads "github.com/colinmutter/gads/v201506"
	  )

or

	  import (
	    gads "github.com/colinmutter/gads/v201409"
	  )


## usage

The package is comprised of services used to manipulate various
adwords structures.  To access a service you need to create an
gads.Auth and parse it to the service initializer, then can call
the service methods on the service object.

~~~ go
     authConf, err := NewCredentials(oauth2.NoContext)
     campaignService := gads.NewCampaignService(&authConf.Auth)

     campaigns, totalCount, err := cs.Get(
       gads.Selector{
         Fields: []string{
           "Id",
           "Name",
           "Status",
         },
       },
     )
~~~

> Note: This package is a work-in-progress, and may occasionally
> make backwards-incompatible changes.

See godoc for further documentation and examples.

* [godoc.org/github.com/emiddleton/gads](https://godoc.org/github.com/emiddleton/gads)

## about

Gads is developed by [Edward Middleton](https://blog.vortorus.net/)

and supported by:   
 - [Colin Mutter](http://github.com/colinmutter)   
 - [Ryan Fink](http://github.com/rfink)
