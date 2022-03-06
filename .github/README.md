# network-database

A collection of IP addresses linked to specific types of material.

[![Updating the resources](https://github.com/complexorganizations/network-database/actions/workflows/update-resources.yml/badge.svg)](https://github.com/complexorganizations/network-database/actions/workflows/update-resources.yml)

## Features

- IP address information depending on activity and other variables.
- Blocklists may be easily created using the data given.

### CDN
| Name | Description | GitHub | GitLab | Statically | JSDelivr | Combinatronics |
| :--- | :---------- | :----- | :----- | :--------- | :------- | :------------- |
| Abuse | A collection of abuse-related IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/abuse` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/abuse` | `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/abuse` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/abuse` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/abuse` |
| Anonymizers | A collection of anonymizers-related IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/anonymizers` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/anonymizers` | `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/anonymizers` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/anonymizers` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/anonymizers` |
| Attacks | A collection of attacks-related IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/attacks` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/attacks` | `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/attacks` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/attacks` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/attacks` |
| Malware | A collection of malware-related IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/malware` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/malware` | `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/malware` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/malware` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/malware` |
| Organizations | A collection of organizations-related IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/organizations` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/organizations` | `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/organizations` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/organizations` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/organizations` |
| Reputation | A collection of reputation-related IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/reputation` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/reputation` | `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/reputation` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/reputation` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/reputation` |
| Spam | A collection of spam-related IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/spam` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/spam` |  `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/spam` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/spam` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/spam` |
| Unroutable | A collection of unroutable IP addresses. | `https://raw.githubusercontent.com/complexorganizations/ip-blocklists/main/assets/unroutable` | `https://gitlab.com/prajwal-koirala/ip-blocklists/-/raw/main/assets/unroutable` |  `https://cdn.statically.io/gh/complexorganizations/ip-blocklists/main/assets/unroutable` | `https://cdn.jsdelivr.net/gh/complexorganizations/ip-blocklists/assets/unroutable` | `https://combinatronics.io/complexorganizations/ip-blocklists/main/assets/unroutable` |


---
## Updating

Clone the project

```bash
git clone https://github.com/complexorganizations/ip-blocklists
```

Go to the project directory

```bash
cd ip-blocklists
```

Build the application

```bash
go build .
```

Update the lists.

```bash
./ip-blocklists -update
```

---
## FAQ

#### What proportion of IP addresses are invalid?

- Since each and every IP address is verified for validity, the percentage is generally 0%.

---
## Authors

- [@prajwal-koirala](https://github.com/prajwal-koirala)

## Credits

Open Source Community

## Support

Please utilize the github repo issue and wiki for help.

## Contributing

Contributions are always welcome!

## Feedback

Please utilize the github repo conversations to offer feedback.

## License

[Apache License Version 2.0](https://github.com/complexorganizations/ip-blocklists/blob/main/.github/license)
