## Fake GitHub Contributions Generator

### Disclaimer
Please be aware that using this tool to generate fake contributions may violate GitHub's terms of service. It is essential to use this tool responsibly and only for personal or testing purposes. I am not responsible for any misuse or violations.

This is a tool that generates fake GitHub contributions (commits) for a specified date range in order to fill up your GitHub activity graph. It can be useful for personal or testing purposes.


<img src="ksnip.png" alt="screenshot"/>


### Prerequisites for building
Before running this tool, make sure you have the following:

- Go (Golang) installed on your system.
- A GitHub account with scope access token


### Getting Started

- Clone this repository to your local machine.
- Generate a scoped access token for your GitHub account by following this link: GitHub Access Tokens. Make sure to grant the necessary permissions for pushing and cloning private repositories.
- Set up the tool by providing the required information through command-line flags.

### Installing

```shell
$ go install github.com/navicstein/fake-contributions
```

### Running

```shell
$ fake-contributions -commitsPerDay=10 -username="your_username" -accessToken="your_access_token" -emailAddress="your_email@example.com"
```

Usage
Here are the available command-line flags:

- `commitsPerDay`: Number of commits to generate per day (default is 10).
- `workdaysOnly`: If set to true, commits will only be generated on workdays (Monday to Friday).
- `startDate`: Start date in the format "YYYY-MM-DD" (optional, defaults to one year ago from the current date).
- `endDate`: End date in the format "YYYY-MM-DD" (optional, defaults to the current date).
- `username`: Your GitHub username.
- `accessToken`: Your GitHub access token used for pushing and cloning private repositories.
- `emailAddress`: Your commit email address.

### Author
This tool is developed and maintained by me (Navicstein) find me on LinkedIn.

### Acknowledgment
If you find this tool useful, consider giving it a star on GitHub. The author would appreciate your support.