# Samurai

    /|
                          /'||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||
                         |  ||         __.--._
                         |  ||      /~~   __.-~\ _
                         |  ||  _.-~ / _---._ ~-\/~\
                         |  || // /  /~/  .-  \  /~-\
                         |  ||((( /(/_(.-(-~~~~~-)_/ |
                         |  || ) (( |_.----~~~~~-._\ /
                         |  ||    ) |              \_|
                         |  ||     (| =-_   _.-=-  |~)        ,
                         |  ||      | `~~ |   ~~'  |/~-._-'/'/_,
                         |  ||       \    |        /~-.__---~ , ,
                         |  ||       |   ~-''     || `\_~~~----~
                         |  ||_.ssSS$$\ -====-   / )\_  ~~--~
                 ___.----|~~~|%$$$$$$/ \_    _.-~ /' )$s._
        __---~-~~        |   |%%$$$$/ /  ~~~~   /'  /$$$$$$$s__
      /~       ~\    ============$$/ /        /'  /$$$$$$$$$$$SS-.
    /'      ./\\\\\\_( ~---._(_))$/ /       /'  /$$$$%$$$$$~      \
    (      //////////(~-(..___)/$/ /      /'  /$$%$$%$$$$'         \
     \    |||||||||||(~-(..___)$/ /  /  /'  /$$$%$$$%$$$            |
      `-__ \\\\\\\\\\\(-.(_____) /  / /'  /$$$$%$$$$$%$             |
          ~~""""""""""-\.(____) /   /'  /$$$$$%%$$$$$$\_            /
                        $|===|||  /'  /$$$$$$$%%%$$$$$( ~         ,'|
                    __  $|===|%\/'  /$$$$$$$$$$$%%%%$$|        ,''  |
                   ///\ $|===|/'  /$$$$$$%$$$$$$$%%%%$(            /'
                    \///\|===|  /$$$$$$$$$%%$$$$$$%%%%$\_-._       |
                     `\//|===| /$$$$$$$$$$$%%%$$$$$$-~~~    ~      /
                       `\|-~~(~~-`$$$$$$$$$%%%///////._       ._  |
                       (__--~(     ~\\\\\\\\\\\\\\\\\\\\        \ \
                       (__--~~(       \\\\\\\\\\\\\\\\\\|        \/
                        (__--~(       ||||||||||||||||||/       _/
                         (__.--._____//////////////////__..---~~
                         |   """"'''''           ___,,,,ss$$$%
                        ,%\__      __,,,\sssSS$$$$$$$$$$$$$$%%
                      ,%%%%$$$$$$$$$$\;;;;\$$$$$$$$$$$$$$$$%%%$.
                     ,%%%%%%$$$$$$$$$$%\;;;;\$$$$$$$$$$$$%%%$$$$
                   ,%%%%%%%%$$$$$$$$$%$$$\;;;;\$$$$$$$$$%%$$$$$$,
                  ,%%%%%%%%%$$$$$$$$%$$$$$$\;;;;\$$$$$$%%$$$$$$$$
                 ,%%%%%%%%%%%$$$$$$%$$$$$$$$$\;;;;\$$$%$$$$$$$$$$$
                 %%%%%%%%%%%%$$$$$$$$$$$$$$$$$$\;;;$$$$$$$$$$$$$$$
                   ""==%%%%%%%$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$SV"
                               $$$$$$$$$$$$$$$$$$$$====""""
                                 """""""""~~~~

Bot for scrapping information about application from Google Play or AppStore and storing them to database.

Additionally, the bot can collect information about the position of the application in categories or by given keywords. (keywords gives by user).

Categories for Google Play
```
"apps_topselling_free"  
"apps_topgrossing"  
"apps_movers_shakers"  
"apps_topselling_paid"
```

Categories fro AppStore
```
"newapplications"  
"newfreeapplications"  
"newpaidapplications"  
"topfreeapplications"  
"topgrossingapplications"  
"toppaidapplications"
```

## Configuration

Bot works in different modes GpTool and StoreTool.  For help you can simple type

```sh
$ samurai <type (google|store)> 
```

And then docs appear

Flags for google

```sh

Flags:
  -c, --count int            set count of apps for tracking (default 250)
  -d, --device string        name of user device (default "whyred")
  -e, --email string         email for the device user account
  -f, --file string          file with keywords separated by '
                             '
      --force                force a new tracking instance
      --gsfig int            gsfid instead of user email (must be paired with token)
  -h, --help                 help for google
      --img                  save images to local server collection
  -i, --intensity duration   tracking frequency (default 24h0m0s)
  -k, --keywords string      keywords for tracking separated by commas
  -l, --locale string        Locale for tracking (default "ru_RU")
      --meta                 track only meta information
      --password string      password for the device user account
  -p, --period int           Period of tracking (default 30)
      --proxy string         proxy for external requests from the device
  -t, --target string        Target bundle for tracking
      --token string         token instead of user password (must be paired with gsfid)
```

And flags for store

```sh
Flags:
  -c, --count int            set count of apps for tracking (default 200)
  -f, --file string          file with keywords separated by '
                             '
      --force                force a new tracking instance
  -h, --help                 help for store
      --img                  save images to local server collection
  -i, --intensity duration   tracking frequency (default 24h0m0s)
  -k, --keywords string      keywords for tracking separated by commas
  -l, --locale string        Locale for tracking (default "ru_RU")
      --meta                 track only meta information
  -p, --period int           Period of tracking (default 30)
  -t, --target string        Target bundle for tracking
```

If flags does not contain a default value, it means that they are required.

Let's analyze in more detail

### Goolge tool

`-t | --target (required)` - app bundle for tracking

```sh
$ samurai google -t com.waze
```

`-l | --locale` - where is doing tracking. By default ru_RU.

```sh
$ samurai google -t com.waze -l en_US
```

`-p | --period` - period in days how much time application will be tracked

```sh
$ samurai google -t com.waze -p 5
```

`-i | --intenisty` - frequency of one executor tick (saving info from selected store). By default 24h

```sh
$ samurai google -t com.waze -i 35h
```

`-c | --count (max 250)` - setting the number of applications for further comparison with the `target`

```sh
$ samurai google -t com.waze -c 100
```

`-k | --keywords` - keywords to track the application's target position for those keywords.  `target`.  Keywords must be separated by commas. If `-k` and `-f` are both provided `-f` has more prioritet

```sh
$ samurai google -t com.waze -k "key1, key2, key3"
```

`-f | --file` - a keyword file to track the application's target position for those keywords. Keywords must be separated by `\ n`

```sh
$ samurai google -t com.waze -f path_to_file
```

`--force` - by default, if you start tracking an application, the bot checks if you are tracking the application with the same parameters, and if it finds it, it gets some configuration from the previous launch. This command allows you to create another application with such parameters

```sh
$ samurai google -t com.waze --force
```

`-d | --device (required if --meta not set)` - the name of the device to track the category. The default value for the field is "whyred". This field is required for a tool that gets all apps from categories

```sh
$ samurai google -t com.waze -d your_device
```

`-e | --email (required if --meta not set)` - email to connect user to specific `device`. This filed is required for google tool if your want track categories.

```sh
$ samurai google -t com.waze -e email@something.com
```

`--password (required if --meta not set)` - uses with `email` for connecting to user deivce. Needed for scraping categories

```sh
$ samurai google -t com.waze --password strongpassword
```

`--gsfig` - This is id given to user when he init connection to new `device`.Not required param. This param helps category tracker not creating new connection to device. Must using together with  `token`

```sh
$ samurai google -t com.waze --gsfig 18238898839283912313
```

`--token` - This is token given to user if he successfully connected to device. Not required param. This param helps category tracker not creating new connection to device. Must using together with  `gsfig`

```sh
$ samurai google -t com.waze --token somerandomsymbols123123213lkdla023ajdian92
```

`--proxy` - proxy for http connection

```sh
$ samurai google -t com.waze --proxy http(s)://login:password@ip:port
```

`--meta` - track only meta information (description, rating etc...). Without keywords and categories

```sh
$ samurai google -t com.waze --meta
```

`--img` - start saving images from the application to an external resource. This resource returns new links to images in local storage

```sh
$ samurai google -t com.waze --img
```

`--help` - help message :)

### App Store tool

`-t | --target (required)` - app bundle for tracking.

```sh
$ samurai store -t 512939461
```

`-l | --locale` - where is doing tracking. By default ru_RU.

```sh
$ samurai store -t 512939461 -l en_US
```

`-p | --period` - period in days how much time application will be tracked

```sh
$ samurai store -t 512939461 -p 3
```

`-i | --intenisty` - frequency of one executor tick (saving info from selected store). By default 24h

```sh
$ samurai store -t 512939461 -i 30h
```

`-c | --count (max 200)` - setting the number of applications for further comparison with the `target`

```sh
$ samurai store -t 512939461 -c 50
```

`-k | --keywords` - keywords to track the application's target position for those keywords.  `target`.  Keywords must be separated by commas. If `-k` and `-f` are both provided `-f` has more prioritet

```sh
$ samurai store -t 512939461 -k "key1, key2, key3"
```

`-f | --file` - a keyword file to track the application's target position for those keywords. Keywords must be separated by `\ n`

```sh
$ samurai store -t 512939461 -f path_to_file
```

`--force` - by default, if you start tracking an application, the bot checks if you are tracking the application with the same parameters, and if it finds it, it gets some configuration from the previous launch. This command allows you to create another application with such parameters

```sh
$ samurai store -t 512939461 -p 4 -i 35h --force
```

`--meta` - track only meta information (description, rating etc...). Without keywords and categories

```sh
$ samurai store -t 512939461 --meta
```

`--img` - start saving images from the application to an external resource. This resource returns new links to images in local storage

```sh
$ samurai store -t 512939461 --img
```

`--help` - help message :)

### Config

You may provide config in config dir. Prod.yml uses only for production. Dev.yml for developing. Let's see how it looks like

```yaml
envs: [envs1, envs2, envs3]  
api:  
  url: <url to service where doing scraping>  
  key: <client key>
  grpc_address: <address to category service>  
  grpc_port: <port of category service>  
  image_processing: <address to img processing service>  
database:  
  name: <db name>  
  user: <db user>  
  password: <db password>  
  address: <db address>  
  port: <db password>  
  schema: <where is database schema stored>
``` 

### Envs

Config provided field envs. You can provide next sensitive security data via envs.

```
* api_key
* db_pass
* db_user
* grpc_address
* grpc_port
``` 

### Deploy

For deploing this app you need add config to docker-compose and just ranning

```sh
$ docker-compose up -d <my_compose_name>
```

