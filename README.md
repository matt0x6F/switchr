switchr
===

**Introduction**

Switchr is a command line application for publicly shared pairing stations.

At Capstone we have team pair-programming stations that have all of our team utilities loaded. One of the issues we ran into was quickly switching between team members that were in control. We sought to solve Git first and then be able to add on AWS/GCP credential switching as well. 

This project reflects an early and unmatured release of switchr, which is really our first GoLang app.

**Using Switchr**

Switchr comes with a full help menu courtesy of Cobra.
Create a file in your home directory called `.switchr.yaml` with the following format:

```
profiles:
    - name: Profile One
      email: you@email.com
      key: key_name
```

Switchr always looks for keys in `~/.ssh`
