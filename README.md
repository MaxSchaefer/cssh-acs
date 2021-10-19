# cssh-acs
###### ContainerSSH-AuthConfigServer with database

[Checkout ContainerSSH](https://github.com/ContainerSSH/ContainerSSH)

### Intention

We at [one medialis](https://medialis.one) would like to give our customers SSH/SFTP access to their websites hosted by us.
For this we use ContainerSSH to dynamically launch containers with their mounted resources.
Because ContainerSSH relies on an external auth, cssh-acs provides an easy way to do this.

###### ***NOTE: This is not an official one medialis product***

### Usage

Example usage
```shell
> docker run \
    --name cssh-acs
    --detach
    --publish 8080:80
    --volume "$(pwd)/db:/cssh-acs/db"
    ghcr.io/maxschaefer/cssh-acs
```

### Database

Database structure
```
db/
    users/
        user.json    
```

User model
```
{
  "username": String,
  "password": String, // SHA512 Hash
  "config": Object // https://github.com/ContainerSSH/configuration/blob/f1696ce58c9d317bba1eb8afa250f678efdbb487/appconfig.go#L19
}
```
Checkout /example (Password: user)
