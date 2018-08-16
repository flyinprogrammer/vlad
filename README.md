# Welcome to Vlad!

This tool implements the Vagrant Cloud API so that you can privately host vagrant boxes. It supports using the [packer vagrant-cloud post-processor](https://www.packer.io/docs/post-processors/vagrant-cloud.html) for publishing boxes as well as the [vagrant](https://www.vagrantup.com/docs/cli/box.html) CLI tool for managing boxes on your hosts.

## Buffalo

Vlad is a [buffalo](https://gobuffalo.io/en/docs/installation) application, so you will need to install that toolchain in order to develop on Vlad.

## Database Setup

Vlad expects you to use postgres to store all the metadata about the boxes.

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a

Note: If you installed postgres with homebrew on os x you might need to run this first:

```bash
$ buffalo db create
ERRO[0000] Error: couldn't create database vlad_development: error creating PostgreSQL database vlad_development: pq: role "postgres" does not exist

$ createuser -s -r postgres
$ buffalo db create
v4.6.4

created database vlad_development
```

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

``` bash
$ buffalo dev
```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
