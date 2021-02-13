# Samurai
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

`-c | --count` - setting the number of applications for further comparison with the `target`

`-d | --device` - the name of the device to track the category. The default value for the field is "whyred". This field is required for a tool that gets all apps from categories

`-e | --email` - email to connect user to specific `device`. This filed is required for google tool if your want track categories.

`-f | --file` - a keyword file to track the application's target position for those keywords. Keywords must be separated by `\ n`

`--force` - by default, if you start tracking an application, the bot checks if you are tracking the application with the same parameters, and if it finds it, it gets some configuration from the previous launch. This command allows you to create another application with such parameters

`--gsfig` - This is id given to user when he init connection to new `device`.Not required param. This param helps category tracker not creating new connection to device. Must using together with  `token`

`--help` - help message :)

`--img` - start saving images from the application to an external resource. This resource returns new links to images in local storage

`-i | --intenisty` - frequency of one executor tick (saving info from selected store). By default 24h.

`-k | --keywords` - keywords to track the application's target position for those keywords.  `target`.  Keywords must be separated by commas. If `-k` and `-f` are both provided `-f` has more prioritet.


