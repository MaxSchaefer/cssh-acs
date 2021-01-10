# cssh-acs
###### ContainerSSH-AuthConfigServer with database

[Checkout ContainerSSH](https://github.com/ContainerSSH/ContainerSSH)

### Usage

Example usage
```shell
> docker build -t maxschaefer/cssh-acs .
> docker run \
    --name cssh-acs
    --detach
    --publish 8080:80
    --volume "$(pwd)/db:/cssh-acs/db"
    maxschaefer/cssh-acs
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
