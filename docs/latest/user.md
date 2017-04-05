### User Management Commands

The `amp user` command manages all user related operations for AMP.

```
    $ amp user

    Usage:	amp user [OPTIONS] COMMAND [ARGS...]

    User management operations

    Options:
          --help   Print usage

    Commands:
      forgot-login Retrieve account name
      get          Get user
      ls           List users
      rm           Remove user
      signup       Signup for a new account
      verify       Verify account

    Run 'amp user COMMAND --help' for more information on a command.
```

### Examples

* To create a user -
```
    $ amp user signup
```
    [_an email with a verification token will be sent to the given email address_]

* To verify newly created account -
```
    $ amp user verify
```

* To retrieve account name -
```
    $ amp user forgot-login
```
    [_an email with the username will be sent to the registered email address_]

* To retrieve a list of users -
```
    $ amp user ls
```

* To retrieve details of a specific user -
```
    $ amp user get
```

* To remove of a user -
```
    $ amp user rm
    [or]
    $ amp user del
```

* To login to AMP -
```
    $ amp login
```

* To switch between accounts (user or org):
```
    $ amp switch
```

* To see who's currently logged in (user or org):
```
    $ amp whoami
```

* To logout of an account:
```
    $ amp logout
```