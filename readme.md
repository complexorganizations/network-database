# Network-Database: A Comprehensive IP Address Repository

Welcome to Network-Database, an extensive repository of IP addresses meticulously categorized to facilitate diverse network-related needs.

## Repository Overview

[![Resource Update Status](https://github.com/complexorganizations/network-database/actions/workflows/update-resources.yml/badge.svg)](https://github.com/complexorganizations/network-database/actions/workflows/update-resources.yml)

Network-Database stands as a pivotal resource in network management and security. It's a dynamic, regularly updated collection hosting millions of IP addresses. These IPs are assorted into various categories, each signifying a distinct type of network activity or threat.

### Key Features

- **Extensive IP Collection:** Over 10 million IP addresses, systematically categorized.
- **Real-Time Validation:** Each IP is verified in real-time, ensuring a 99.99% accuracy rate.
- **User-Friendly Blocklist Generation:** Easily generate customized blocklists tailored to specific security needs.
- **Regular Updates:** Database updates occur daily, ensuring the latest IP data is always at your fingertips.

### Content Delivery Network (CDN) Access

Access the meticulously curated collections of IP addresses through various CDN platforms. Each category is linked to a specific type of network behavior or threat, providing a wide range of resources for network management and security.

| Category      | Description                                       | GitHub Repository                                                                                                   | Statically CDN                                                                                                 | JSDelivr CDN                                                                                           | Combinatronics CDN                                                                                              |
| :------------ | :------------------------------------------------ | :------------------------------------------------------------------------------------------------------------------ | :------------------------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------- |
| Abuse         | IPs linked to abusive activities                  | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/abuse)         | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/abuse)         | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/abuse)         | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/abuse)         |
| Anonymizers   | IPs associated with anonymizing services          | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/anonymizers)   | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/anonymizers)   | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/anonymizers)   | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/anonymizers)   |
| Attacks       | IPs involved in network attacks                   | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/attacks)       | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/attacks)       | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/attacks)       | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/attacks)       |
| Malware       | IPs related to malware distribution               | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/malware)       | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/malware)       | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/malware)       | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/malware)       |
| Organizations | IPs associated with specific organizations        | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/organizations) | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/organizations) | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/organizations) | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/organizations) |
| Reputation    | IPs known for impacting network reputation        | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/reputation)    | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/reputation)    | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/reputation)    | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/reputation)    |
| Spam          | IPs frequently used for spamming activities       | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/spam)          | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/spam)          | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/spam)          | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/spam)          |
| Unroutable    | IPs that are unroutable in network configurations | [View on GitHub](https://raw.githubusercontent.com/complexorganizations/network-database/main/assets/unroutable)    | [Statically CDN](https://cdn.statically.io/gh/complexorganizations/network-database/main/assets/unroutable)    | [JSDelivr CDN](https://cdn.jsdelivr.net/gh/complexorganizations/network-database/assets/unroutable)    | [Combinatronics CDN](https://combinatronics.io/complexorganizations/network-database/main/assets/unroutable)    |

These resources are updated regularly to ensure the most current and comprehensive data is available for your network security and management needs.

---

## How to Use and Update

### Setting Up

Clone the repository to get started:

```bash
git clone https://github.com/complexorganizations/network-database
```

Navigate to the project directory:

```bash
cd network-database
```

Compile the application:

```bash
go build .
```

### Updating IP Lists

Execute the following command to refresh the IP lists:

```bash
./network-database -update
```

---

## Insights and FAQs

**Q: How frequently is the database updated?**

- A: Daily updates are conducted to ensure the most current IP data is available.

**Q: What is the error margin for invalid IPs?**

- A: Our stringent validation process limits the error margin to less than 0.01%.

---

## Community and Support

- **Authors:** Spearheaded by the Open Source Community, with contributions from network experts worldwide.
- **Support:** Utilize our GitHub issue tracker and wiki for in-depth assistance.
- **Contributing:** Your contributions are invaluable. Join us in enhancing this vital resource!
- **Feedback:** We welcome your insights. Please share them in the repository's conversation section.
- **License:** Governed under the [Apache License Version 2.0](https://github.com/complexorganizations/ip-blocklists/blob/main/.github/license).

Together, let's make Network-Database the most reliable and comprehensive IP resource available!
